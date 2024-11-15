package gpi

import "github.com/gin-gonic/gin"

type Gpier interface {
	Out(success bool, msg string, data interface{})
	Success(msg string, data interface{})
	Fail(msg string)
	FailWithMsg(msg string, data interface{})
	Abort(httpCode int, msg string, errorCode int)
	Abort401(msg string, errCode int)
}

type Gpi struct {
	Ctx *gin.Context
}
