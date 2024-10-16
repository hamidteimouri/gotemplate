package common

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
}

// SetResponse is a general method to generate an HTTP request
func SetResponse(data interface{}, err error) *Response {
	resp := &Response{}
	if err != nil {
		resp.Error = err.Error()
		return resp
	}

	resp.Data = data
	return resp
}

// Resp returns a response with given status code
func Resp(c echo.Context, code int, data interface{}, err error) error {
	return c.JSON(code, SetResponse(data, err))
}

// RespBadRequest returns a response with status code 400
func RespBadRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, SetResponse(nil, err))
}

// RespNotFound return a response with status code 404
func RespNotFound(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusNotFound, SetResponse(data, nil))
}

// RespUnprocessableEntity returns a response with status code 422
func RespUnprocessableEntity(c echo.Context, err error) error {
	return c.JSON(http.StatusUnprocessableEntity, SetResponse(nil, err))
}

// RespOK returns a response with status code 200
func RespOK(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, SetResponse(data, nil))
}

// RespOKCreated returns a response with status code 201
func RespOKCreated(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusCreated, SetResponse(data, nil))
}

// RespInternalError returns a response with status code 500
func RespInternalError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, SetResponse(nil, err))
}
