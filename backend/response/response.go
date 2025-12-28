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

func Success(c *gin.Context, status int, msg string) {
	c.JSON(http.StatusOK, Response{
		Status: 200,
		Msg:    msg,
	})
}

func FailBadRequest(c *gin.Context, status int, msg string) {
	c.JSON(http.StatusBadRequest, Response{
		Status: 400,
		Msg:    msg,
	})
}
