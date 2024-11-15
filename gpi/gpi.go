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
func (c *Gpi) Out(code int, success bool, msg string, data interface{}) {
	c.Ctx.JSON(code, gin.H{
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
	c.Out(200, true, msg, gin.H{})
}

// SuccessWithData sends a JSON response indicating a successful operation with the given message and data.
//
// Parameters:
// - msg: A string containing the success message.
// - data: An interface{} containing the data to be sent.
func (c *Gpi) SuccessWithData(msg string, data interface{}) {
	c.Out(200, true, msg, data)
}

// Fail sends a JSON response indicating a failed operation with the given message.
//
// Parameters:
// - msg: A string containing the failure message.
func (c *Gpi) Fail(msg string) {
	c.Out(200, false, msg, gin.H{})
}

// FailWithData sends a JSON response indicating a failed operation with the given message and data.
//
// Parameters:
// - msg: A string containing the failure message.
// - data: An interface{} containing the data to be sent.
func (c *Gpi) FailWithData(msg string, data interface{}) {
	c.Out(200, false, msg, data)
}

func (c *Gpi) FailWithHttpCode(code int, msg string) {
	c.Out(code, false, msg, gin.H{})
}

// Abort sends a JSON response with the given HTTP status code, message, and error code, and aborts the request.
//
// Parameters:
// - httpCode: An integer representing the HTTP status code.
// - msg: A string containing the abort message.
// - errorCode: An integer representing the error code.
func (c *Gpi) Abort(code int, msg string, error string) {
	c.Ctx.AbortWithStatusJSON(code, gin.H{
		"success": false,
		"message": msg,
		"data": gin.H{
			"error": error,
		},
	})
}

// Abort401 sends a JSON response with a 401 HTTP status code, the given message, and error code, and aborts the request.
//
// Parameters:
// - msg: A string containing the abort message.
// - errorCode: An integer representing the error code.
func (c *Gpi) Abort401(msg string, errors string) {
	c.Abort(401, msg, errors)
}

func (c *Gpi) Abort200(msg string, errors string) {
	c.Abort(200, msg, errors)
}

func (c *Gpi) Abort204(msg string, errors string) {
	c.Abort(200, msg, errors)
}
