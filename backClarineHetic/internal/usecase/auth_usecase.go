package usecase

import (
    "backClarineHetic/internal/domain"
    "backClarineHetic/pkg/jwt"
    "errors"
    "golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
    Signup(user *domain.User) (string, error)
    Login(email, password string) (string, error)
}

type authUsecase struct {
    userRepo domain.UserRepository
}

func NewAuthUsecase(userRepo domain.UserRepository) AuthUsecase {
    return &authUsecase{userRepo: userRepo}
}

func (a *authUsecase) Signup(user *domain.User) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    user.Password = string(hashedPassword)

    if err := a.userRepo.Create(user); err != nil {
        return "", err
    }

    token, err := jwt.GenerateToken(user.Email, user.UUID.String())
    if err != nil {
        return "", err
    }
    return token, nil
}

func (a *authUsecase) Login(email, password string) (string, error) {
    user, err := a.userRepo.FindByEmail(email)
    if err != nil {
        return "", err
    }
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("identifiants invalides")
    }
    token, err := jwt.GenerateToken(user.Email, user.UUID.String())
    if err != nil {
        return "", err
    }
    return token, nil
}
