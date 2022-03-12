package global

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//type errorMsgs []*Error
//
//// 错误处理的结构体
//type Error struct {
//	StatusCode int    `json:"-"`
//	Code       int    `json:"code"`
//	Msg        string `json:"msg"`
//}
//
//var (
//	Success     = NewError(http.StatusOK, 0, "success")
//	ServerError = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
//	NotFound    = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
//)
//
//func OtherError(message string) *Error {
//	return NewError(http.StatusForbidden, 100403, message)
//}
//
//func (e *Error) Error() string {
//	return e.Msg
//}
//
//func NewError(statusCode, Code int, msg string) *Error {
//	return &Error{
//		StatusCode: statusCode,
//		Code:       Code,
//		Msg:        msg,
//	}
//}
//
//// 404处理
//func HandleNotFound(c *gin.Context) {
//	err := NotFound
//	c.JSON(err.StatusCode, err)
//	return
//}
//
//func ErrHandler() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		c.Next()
//		if length := len(c.Errors); length > 0 {
//			e := c.Errors[length-1]
//			err := e.Err
//			if err != nil {
//				var Err *Error
//				if e, ok := err.(*Error); ok {
//					Err = e
//				} else if e, ok := err.(error); ok {
//					Err = OtherError(e.Error())
//				} else {
//					Err = ServerError
//				}
//				// 记录一个错误的日志
//				c.JSON(Err.StatusCode, Err)
//				return
//			}
//		}
//
//	}
//}

// 错误处理的结构体
type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

var (
	Success       = NewError(http.StatusOK, 0, "success")
	ServerError   = NewError(http.StatusInternalServerError, 200500, "The system is abnormal. Please try again later")
	NotFound      = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	DatabaseError = NewError(http.StatusOK, 10001, http.StatusText(http.StatusServiceUnavailable))
)

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}

// 404处理
func HandleNotFound(c *gin.Context) {
	err := NotFound
	c.JSON(err.StatusCode, err)
	return
}

func ErrHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var Err *Error
				if e, ok := err.(*Error); ok {
					Err = e
				} else if e, ok := err.(error); ok {
					Err = OtherError(e.Error())
				} else {
					Err = ServerError
				}
				//reflectType := reflect.TypeOf(err)
				//reflectValue := reflect.ValueOf(err)
				//switch reflectType {
				//case int:
				//	response.RayWithoutData(reflectValue, c)
				//
				//
				//}
				c.JSON(Err.StatusCode, Err)
				return
			}
		}()
		c.Next()
	}
}
