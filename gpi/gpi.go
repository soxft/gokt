package gpi

import "github.com/gin-gonic/gin"

func New(ctx *gin.Context) *Gpi {
	return &Gpi{
		Ctx: ctx,
	}
}

func (c *Gpi) Out(success bool, msg string, data interface{}) {
	c.Ctx.JSON(200, gin.H{
		"success": success,
		"message": msg,
		"data":    data,
	})
}

func (c *Gpi) Success(msg string) {
	c.Out(true, msg, gin.H{})
}

func (c *Gpi) SuccessWithData(msg string, data interface{}) {
	c.Out(true, msg, data)
}

func (c *Gpi) Fail(msg string) {
	c.Out(false, msg, gin.H{})
}

func (c *Gpi) FailWithData(msg string, data interface{}) {
	c.Out(false, msg, data)
}

func (c *Gpi) Abort(httpCode int, msg string, errorCode int) {
	c.Ctx.AbortWithStatusJSON(httpCode, gin.H{
		"success": false,
		"message": msg,
		"data": gin.H{
			"errorCode": errorCode,
		},
	})
}

func (c *Gpi) Abort401(msg string, errorCode int) {
	c.Abort(401, msg, errorCode)
}
