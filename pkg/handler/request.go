package handler

import (
	"task/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(g *gin.Engine) {
	println("routel vannee")
	g.POST("/", controllers.Incoming)
}
