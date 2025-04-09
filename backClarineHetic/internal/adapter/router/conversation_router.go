package router

import (
    "backClarineHetic/internal/adapter/controller"
    "github.com/gin-gonic/gin"
)

func NewConversationRouter(r *gin.Engine, ConversationController controller.ConversationController) {
    conversationGroup := r.Group("/conversation")
    {
        conversationGroup.GET("/", ConversationController.GetConversation)
        conversationGroup.POST("/:uuid", ConversationController.GetConversationByUuid)
    }
}
