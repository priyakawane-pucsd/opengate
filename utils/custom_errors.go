package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("code:%d error:%s", c.Code, c.Message)
}

func NewCustomError(code int, message string) error {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func NewBadRequestError(message string) error {
	return NewCustomError(http.StatusBadRequest, message)
}

func NewInternalServerError(message string) error {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func WriteError(ctx *gin.Context, err error) {
	if cErr, ok := err.(*CustomError); ok {
		ctx.JSON(cErr.Code, cErr)
		return
	}
	ctx.JSON(http.StatusInternalServerError, &CustomError{Code: http.StatusInternalServerError, Message: err.Error()})
}
