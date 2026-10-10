package auth

import (
    "errors"
    "time"

    "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID uint `json:"uid"`
    jwt.RegisteredClaims
}

func Generate(secret string, userID uint, ttl time.Duration) (string, error) {
    c := Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }
    return jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString([]byte(secret))
}

func Parse(secret, token string) (*Claims, error) {
    t, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (any, error) {
        return []byte(secret), nil
    })
    if err != nil {
        return nil, err
    }
    if c, ok := t.Claims.(*Claims); ok && t.Valid {
        return c, nil
    }
    return nil, errors.New("invalid token")
}