package router

import (
    "backClarineHetic/internal/adapter/controller"
    "github.com/gin-gonic/gin"
)

func NewEventRouter(r *gin.Engine, eventController controller.EventController) {
    eventGroup := r.Group("/event")
    {
        eventGroup.GET("/", eventController.GetEvent)
        /*eventGroup.GET("/:uuid", eventController.GetEventByUuid)*/
        eventGroup.POST("/", eventController.CreateEvent)
    }
}
