package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Status int         `json:"status"`
}

func SuccessWithData(c *gin.Context, msg string, data interface{}, status int) {
	c.JSON(http.StatusOK, Response{
		Msg:    msg,
		Data:   data,
		Status: 200,
	})
}

func SuccessWithNoData(c *gin.Context, msg string, status int) {
	c.JSON(http.StatusOK, Response{
		Msg:    msg,
		Data:   nil,
		Status: 200,
	})
}

func FailBadRequest(c *gin.Context, status int, msg string) {
	c.JSON(http.StatusBadRequest, Response{
		Status: 400,
		Msg:    msg,
	})
}
