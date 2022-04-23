package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	resp := &Response{
		Code: 200,
		Msg:  "成功",
		Data: data,
	}
	c.JSON(http.StatusOK, resp)
}

func ResponseError(c *gin.Context, msg string) {
	resp := &Response{
		Code: 500,
		Msg:  msg,
	}
	c.JSON(http.StatusInternalServerError, resp)
}
