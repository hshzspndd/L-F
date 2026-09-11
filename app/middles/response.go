package middles

import (
	"github.com/gin-gonic/gin"
)

type ResponseForm struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *ResponseForm) Error() string {
	return e.Message
}

func GetError(code int, message string) *ResponseForm {
	return &ResponseForm{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	var response_success ResponseForm
	response_success = ResponseForm{
		Code:    0,
		Message: "success",
		Data:    data,
	}
	c.JSON(200, response_success)
}

func GlobalResponseError(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		if err, ok := c.Errors.Last().Err.(*ResponseForm); ok {
			c.JSON(400, ResponseForm{
				Code:    err.Code,
				Message: err.Message,
				Data:    err.Data,
			})
		}
		c.Abort()
	}

}
