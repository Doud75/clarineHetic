package router

import (
    "backClarineHetic/internal/adapter/controller"
    "github.com/gin-gonic/gin"
)

func NewProfileRouter(r *gin.Engine, profileController controller.ProfileController) {
    profileGroup := r.Group("/profile")
    {
        profileGroup.GET("/search-term", profileController.SearchProfile)
    }
}
