package utils

import (
	"time"
	"cloud9/config"
	"github.com/dgrijalva/jwt-go"
)

func TokenCreate(account string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = jwt.MapClaims{
		"token_owner": account,
		"exp":         time.Now().Unix() + 3600*24*7,
	}
	tokenStr, err := token.SignedString([]byte(config.JwtTokenKey))
	return tokenStr, err
}

func TokenParse() {

}
