package routes

import (
	"archengine/internal/app/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
}
