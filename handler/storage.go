package handler

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/juliotorresmoreno/SpecialistTalk/helper"
	"github.com/juliotorresmoreno/SpecialistTalk/services"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
)

type StorageHandler struct{}

func AttachStorage(g *echo.Group) {
	//conf := configs.GetConfig().Minio
	u := &StorageHandler{}

	//g.GET("/:filename", u.get)
	g.POST("", u.add)
	g.GET("/:filename", u.getFromHTTP)

	/*
		url, _ := url.Parse("http://localhost:8000/" + conf.Bucket)
		targets := []*middleware.ProxyTarget{{URL: url}}
		g.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))*/
}

type POSTAddPayload struct {
	Description string
}

func (u *StorageHandler) add(c echo.Context) error {
	conf := configs.GetConfig().Minio
	minioCli, err := services.NewMinioClient()
	if err != nil {
		return helper.HTTPErrorInternalServerError
	}

	p := &POSTAddPayload{}
	c.Bind(p)
	file, err := c.FormFile("file")
	if err != nil {
		helper.MakeHTTPError(500, err)
	}

	src, err := file.Open()
	if err != nil {
		helper.MakeHTTPError(500, err)
	}

	ctx := context.Background()
	bucket := conf.Bucket
	opts := minio.PutObjectOptions{}
	info, err := minioCli.PutObject(ctx, bucket, file.Filename, src, file.Size, opts)
	if err != nil {
		return helper.HTTPErrorNotFound
	}

	return c.JSON(200, info)
}

func (u *StorageHandler) get(c echo.Context) error {
	conf := configs.GetConfig().Minio
	minioCli, err := services.NewMinioClient()
	if err != nil {
		return helper.HTTPErrorInternalServerError
	}
	filename := c.Param("filename")

	ctx := context.Background()
	bucket := conf.Bucket
	opts := minio.GetObjectOptions{}
	o, err := minioCli.GetObject(ctx, bucket, filename, opts)
	if err != nil {
		return helper.HTTPErrorNotFound
	}
	defer o.Close()

	buff := bytes.NewBuffer([]byte{})
	_, err = io.Copy(buff, o)

	if err != nil {
		return helper.HTTPErrorInternalServerError
	}

	b := buff.Bytes()
	contentType := http.DetectContentType(b[:512])

	return c.Blob(http.StatusOK, contentType, b)
}

var responseHeaders = []string{
	"Accept-Ranges",
	"Content-Encoding",
	"Content-Security-Policy",
	"Date",
	"Etag",
	"Last-Modified",
	"Server",
	"Strict-Transport-Security",
	"Vary",
	"Connection",
	"Transfer-Encoding",
}

// getFromHTTP Eficiente e insegura forma de consultar los datos desde el almacen de archivos
func (u *StorageHandler) getFromHTTP(c echo.Context) error {
	conf := configs.GetConfig().Minio
	filename := c.Param("filename")
	url := conf.Url + "/" + conf.Bucket + "/" + filename

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return helper.HTTPErrorInternalServerError
	}
	for key := range c.Request().Header {
		value := c.Request().Header.Get(key)
		req.Header.Add(key, value)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return helper.MakeHTTPError(500, err)
	}
	response := c.Response()
	headers := resp.Header
	for _, key := range responseHeaders {
		value := headers.Get(key)
		if value != "" {
			response.Header().Set(key, value)
		}
	}

	b := make([]byte, 512)
	resp.Body.Read(b)

	contentType := http.DetectContentType(b[:512])

	response.Header().Set("Content-Type", contentType)
	response.WriteHeader(resp.StatusCode)

	_, err = response.Write(b)
	if err != nil {
		return helper.HTTPErrorInternalServerError
	}

	_, err = io.Copy(response, resp.Body)
	if err != nil {
		return helper.HTTPErrorInternalServerError
	}

	return nil
}