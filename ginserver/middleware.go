package ginserver

import (
	"cloud9/config"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
	//"net/http"
)

var err = errors.New("token expired")

func loginMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, cookieErr := c.Request.Cookie("cloud")
		if cookieErr != nil {
			return
		}
		tokenObj, parseErr := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			mapClaims := token.Claims.(jwt.MapClaims)
			if int64(mapClaims["exp"].(float64)) < time.Now().Unix() {
				return nil, err
			}
			return []byte(config.JwtTokenKey), nil
		})
		if parseErr == nil && tokenObj.Valid {
			c.Set("name", tokenObj.Claims.(jwt.MapClaims)["token_owner"])
		}
	}
}
