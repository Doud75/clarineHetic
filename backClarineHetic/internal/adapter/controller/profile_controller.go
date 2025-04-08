package controller

import (
    "backClarineHetic/internal/usecase"
    "github.com/gin-gonic/gin"
    "net/http"
)

type ProfileController interface {
    SearchProfile(ctx *gin.Context)
}

type profileController struct {
    profileUseCase usecase.ProfileUseCase
}

func NewProfileController(profileUseCase usecase.ProfileUseCase) ProfileController {
    return &profileController{profileUseCase: profileUseCase}
}

func (p *profileController) SearchProfile(c *gin.Context) {
    searchTerm := c.Query("search_term")
    if searchTerm == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Le paramètre search_term est requis"})
        return
    }

    users, err := p.profileUseCase.SearchProfile(searchTerm)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": users})
}
