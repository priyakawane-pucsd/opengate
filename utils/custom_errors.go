package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("code:%d error:%s", c.StatusCode, c.Message)
}

func NewCustomError(code int, message string) error {
	return &CustomError{
		StatusCode: code,
		Message:    message,
	}
}

func NewBadRequestError(message string) error {
	return NewCustomError(http.StatusBadRequest, message)
}

func NewInternalServerError(message string) error {
	return &CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

func WriteError(ctx *gin.Context, err error) {
	if cErr, ok := err.(*CustomError); ok {
		ctx.JSON(cErr.StatusCode, cErr)
		return
	}
	ctx.JSON(http.StatusInternalServerError, &CustomError{StatusCode: http.StatusInternalServerError, Message: err.Error()})
}

func WriteResponse(ctx *gin.Context, res any) {
	ctx.JSON(http.StatusOK, res)
}
