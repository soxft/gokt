package gpi

import "github.com/gin-gonic/gin"

// New creates a new Gpi instance with the provided gin context.
//
// Parameters:
// - ctx: The gin context to be used.
//
// Returns:
// - A pointer to the newly created Gpi instance.
func New(ctx *gin.Context) *Gpi {
	return &Gpi{
		Ctx: ctx,
	}
}

// Out sends a JSON response with the given success status, message, and data.
//
// Parameters:
// - success: A boolean indicating the success status.
// - msg: A string containing the message to be sent.
// - data: An interface{} containing the data to be sent.
func (c *Gpi) Out(success bool, msg string, data interface{}) {
	c.Ctx.JSON(200, gin.H{
		"success": success,
		"message": msg,
		"data":    data,
	})
}

// Success sends a JSON response indicating a successful operation with the given message.
//
// Parameters:
// - msg: A string containing the success message.
func (c *Gpi) Success(msg string) {
	c.Out(true, msg, gin.H{})
}

// SuccessWithData sends a JSON response indicating a successful operation with the given message and data.
//
// Parameters:
// - msg: A string containing the success message.
// - data: An interface{} containing the data to be sent.
func (c *Gpi) SuccessWithData(msg string, data interface{}) {
	c.Out(true, msg, data)
}

// Fail sends a JSON response indicating a failed operation with the given message.
//
// Parameters:
// - msg: A string containing the failure message.
func (c *Gpi) Fail(msg string) {
	c.Out(false, msg, gin.H{})
}

// FailWithData sends a JSON response indicating a failed operation with the given message and data.
//
// Parameters:
// - msg: A string containing the failure message.
// - data: An interface{} containing the data to be sent.
func (c *Gpi) FailWithData(msg string, data interface{}) {
	c.Out(false, msg, data)
}

// Abort sends a JSON response with the given HTTP status code, message, and error code, and aborts the request.
//
// Parameters:
// - httpCode: An integer representing the HTTP status code.
// - msg: A string containing the abort message.
// - errorCode: An integer representing the error code.
func (c *Gpi) Abort(httpCode int, msg string, errorCode int) {
	c.Ctx.AbortWithStatusJSON(httpCode, gin.H{
		"success": false,
		"message": msg,
		"data": gin.H{
			"errno": errorCode,
		},
	})
}

// Abort401 sends a JSON response with a 401 HTTP status code, the given message, and error code, and aborts the request.
//
// Parameters:
// - msg: A string containing the abort message.
// - errorCode: An integer representing the error code.
func (c *Gpi) Abort401(msg string, errorCode int) {
	c.Abort(401, msg, errorCode)
}
