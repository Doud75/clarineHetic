package router

import (
    "backClarineHetic/internal/adapter/controller"
    "github.com/gin-gonic/gin"
)

func NewAuthRouter(r *gin.Engine, authController controller.AuthController) {
    authGroup := r.Group("/auth")
    {
        authGroup.POST("/signup", authController.Signup)
        authGroup.POST("/login", authController.Login)
    }
}
