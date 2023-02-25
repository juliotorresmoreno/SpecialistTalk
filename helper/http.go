package helper

import (
	"bufio"
	"bytes"
	"net"
	"net/http"
	"reflect"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/juliotorresmoreno/SpecialistTalk/model"
	"github.com/labstack/echo/v4"
)

func Paginate(c echo.Context) (int, int) {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	skip, _ := strconv.Atoi(c.QueryParam("skip"))
	if limit == 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	return limit, skip
}

func ValidateSession(c echo.Context) (*model.User, error) {
	session := c.Get("session")
	if session == nil {
		return nil, HTTPErrorUnauthorized
	}
	return session.(*model.User), nil
}

func GetPayload(c echo.Context, payload interface{}) (*model.User, error) {
	session, err := ValidateSession(c)
	if err != nil {
		return session, err
	}

	kindOfJ := reflect.ValueOf(payload).Kind()
	if kindOfJ != reflect.Ptr {
		return session, echo.NewHTTPError(http.StatusInternalServerError, "payload must a pointer")
	}

	err = c.Bind(payload)
	if err != nil {
		return session, HTTPErrorBadRequest
	}

	_, err = govalidator.ValidateStruct(payload)
	if err != nil {
		return session, MakeHTTPError(http.StatusBadRequest, err)
	}

	return session, nil
}

func NewRegisterHTTPResponseWriter(w http.ResponseWriter) *RegisterHTTPResponseWriter {
	return &RegisterHTTPResponseWriter{
		w:          w,
		header:     &http.Header{},
		StatusCode: 0,
		Buffer:     bytes.NewBuffer([]byte{}),
	}
}

type RegisterHTTPResponseWriter struct {
	w          http.ResponseWriter
	StatusCode int
	Buffer     *bytes.Buffer
	header     *http.Header
}

func (u *RegisterHTTPResponseWriter) Header() http.Header {
	return *u.header
}

func (u *RegisterHTTPResponseWriter) Write(b []byte) (int, error) {
	return u.Buffer.Write(b)
}

func (u *RegisterHTTPResponseWriter) WriteHeader(statusCode int) {
	u.StatusCode = statusCode
}

func (u *RegisterHTTPResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hj, _ := u.w.(http.Hijacker)
	return hj.Hijack()
}

func (u *RegisterHTTPResponseWriter) Push() {
	for key := range *u.header {
		value := u.header.Get(key)
		u.w.Header().Set(key, value)
	}
	u.w.WriteHeader(u.StatusCode)
	u.w.Write(u.Buffer.Bytes())
}
