package dto

type RegisterRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
    Token string `json:"token"`
    User  UserDTO `json:"user"`
}

type UserDTO struct {
    ID    uint   `json:"id"`
    Email string `json:"email"`
}