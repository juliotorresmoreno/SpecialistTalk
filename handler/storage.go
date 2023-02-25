package handler

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/juliotorresmoreno/SpecialistTalk/helper"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"github.com/juliotorresmoreno/SpecialistTalk/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StorageHandler struct{}

func AttachStorage(g *echo.Group) {
	u := &StorageHandler{}

	g.POST("", u.add)
	g.GET("/:filename", u.getFromHTTP)
}

type POSTAddPayload struct {
	Description string
}

func (u *StorageHandler) add(c echo.Context) error {
	session, err := helper.ValidateSession(c)
	if err != nil {
		return err
	}
	conf := configs.GetConfig()
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

	mongoCli, err := services.GetPoolMongo()
	if err != nil {
		return helper.HTTPErrorInternalServerError
	}

	now := time.Now()
	id := primitive.NewObjectID()
	f := &model.File{
		ID:        &id,
		Name:      file.Filename,
		Owner:     session.Username,
		CreatedAt: &now,
	}
	db := mongoCli.Database(conf.Mongo.StorageDB)
	collection := db.Collection(f.TableName())
	ctx := context.Background()
	_, err = collection.InsertOne(ctx, f)
	if err != nil {
		return helper.HTTPErrorInternalServerError
	}

	ctx = context.Background()
	bucket := conf.Minio.Bucket
	opts := minio.PutObjectOptions{}
	info, err := minioCli.PutObject(ctx, bucket, f.ID.Hex(), src, file.Size, opts)
	if err != nil {
		return helper.HTTPErrorNotFound
	}

	return c.JSON(200, info)
}

// Una alternativa mas
func (u *StorageHandler) Reverse() echo.MiddlewareFunc {
	conf := configs.GetConfig().Minio
	url, _ := url.Parse(conf.Url + conf.Bucket)
	targets := []*middleware.ProxyTarget{{URL: url}}
	return middleware.Proxy(middleware.NewRoundRobinBalancer(targets))
}

// Una alternativa mas
func (u *StorageHandler) Get(c echo.Context) error {
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

	if c.QueryParam("download") == "true" {
		response.Header().Set("Content-Type", "application/octet-stream")
		response.WriteHeader(resp.StatusCode)

		_, err = io.Copy(response, resp.Body)
		if err != nil {
			return helper.HTTPErrorInternalServerError
		}
		return nil
	}

	b := make([]byte, 512)
	_, _ = resp.Body.Read(b)

	contentType := http.DetectContentType(b)
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
