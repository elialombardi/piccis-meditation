package service

import (
    "errors"
    "time"

    "github.com/elialombardi/piccis-meditation/backend/pkg/auth"
    "github.com/elialombardi/piccis-meditation/backend/pkg/dto"
    "github.com/elialombardi/piccis-meditation/backend/pkg/models"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type AuthService struct {
    db        *gorm.DB
    jwtSecret string
    jwtTTL    time.Duration
}

func NewAuthService(db *gorm.DB, secret string, ttl time.Duration) *AuthService {
    return &AuthService{db: db, jwtSecret: secret, jwtTTL: ttl}
}

func (s *AuthService) Register(req dto.RegisterRequest) (*dto.AuthResponse, error) {
    var existing models.User
    if err := s.db.Where("email = ?", req.Email).First(&existing).Error; err == nil {
        return nil, errors.New("email already registered")
    }
    hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    u := models.User{Email: req.Email, PasswordHash: string(hash)}
    if err := s.db.Create(&u).Error; err != nil {
        return nil, err
    }
    return s.issue(u)
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.AuthResponse, error) {
    var u models.User
    if err := s.db.Where("email = ?", req.Email).First(&u).Error; err != nil {
        return nil, errors.New("invalid credentials")
    }
    if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)); err != nil {
        return nil, errors.New("invalid credentials")
    }
    return s.issue(u)
}

func (s *AuthService) issue(u models.User) (*dto.AuthResponse, error) {
    token, err := auth.Generate(s.jwtSecret, u.ID, s.jwtTTL)
    if err != nil {
        return nil, err
    }
    return &dto.AuthResponse{
        Token: token,
        User:  dto.UserDTO{ID: u.ID, Email: u.Email},
    }, nil
}