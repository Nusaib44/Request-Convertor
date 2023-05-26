package common

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	Status  bool
	Message string
	Errors  interface{}
	Data    interface{}
}

func ErrorMessage(g *gin.Context, msg string, errors interface{}, code int) {
	responce := response{
		Status:  false,
		Message: msg,
		Errors:  errors,
	}
	g.JSON(code, responce)
}
func SurcessMessage(g *gin.Context, msg string, data interface{}) {
	responce := response{
		Status:  true,
		Message: msg,
		Errors:  nil,
		Data:    data,
	}
	g.JSON(200, responce)
}
