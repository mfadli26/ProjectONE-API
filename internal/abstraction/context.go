package abstraction

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Context struct {
	echo.Context
	Auth *AuthContext
	Trx  *TrxContext
}

type AuthContext struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthContextClaim struct {
	ID    []byte `json:"id"`
	Name  []byte `json:"name"`
	Email []byte `json:"email"`
	jwt.StandardClaims
}

type TrxContext struct {
	Db *gorm.DB
}
