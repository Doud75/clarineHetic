package domain

import "github.com/google/uuid"

type User struct {
    UUID     uuid.UUID `json:"uuid"`
    Username string    `json:"username"`
    Email    string    `json:"email"`
    Password string    `json:"-"`
}

type UserRepository interface {
    Create(user *User) error
    FindByEmail(email string) (*User, error)
}
