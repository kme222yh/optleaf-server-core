/*  Usage

r.Use(exception.HandlerMiddleware())

exception.Raise(c, message, code, detail, status)

*/


package exception


import (
    "github.com/gin-gonic/gin"
)


func Raise(c *gin.Context, message string, code string, detail interface{}, status int) {
    e := new(exc)
    e.Body.Message = message
    e.Body.Code = code
    e.Body.Detail = detail
    e.HttpStatus = status
    c.Set("CustomException", e)
}


func HandlerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()
        val, ok := c.Get("CustomException")
        if ok == false {
            return
        }
        response := val.(*exc)
        if response != nil {
            c.AbortWithStatusJSON(response.HttpStatus, response.Body)
        }
    }
}


type body struct {
    Message string      `json:"message"`
    Code string         `json:"code"`
    Detail interface{}  `json:"detail"`
}


type exc struct {
    Body body
    HttpStatus int
}
