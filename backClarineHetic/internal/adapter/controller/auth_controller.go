package controller

import (
    "backClarineHetic/internal/adapter/dto"
    "backClarineHetic/internal/domain"
    "backClarineHetic/internal/usecase"
    "net/http"

    "github.com/gin-gonic/gin"
)

type AuthController interface {
    Signup(c *gin.Context)
    Login(c *gin.Context)
}

type authController struct {
    authUsecase usecase.AuthUsecase
}

func NewAuthController(authUsecase usecase.AuthUsecase) AuthController {
    return &authController{authUsecase: authUsecase}
}

func (a *authController) Signup(c *gin.Context) {
    var req dto.SignupRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user := domain.User{
        Username: req.Username,
        Email:    req.Email,
        Password: req.Password,
    }

    // Récupération du token après inscription
    token, err := a.authUsecase.Signup(&user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "utilisateur créé",
        "token":   token,
    })
}

func (a *authController) Login(c *gin.Context) {
    var req dto.LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := a.authUsecase.Login(req.Email, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
