package middleware

import (
	"fmt"
	"global-auth/api/response"
	"global-auth/config"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var secretKey = []byte(config.Env("JWT_SECRET_KEY"))
var adminSecretKey = []byte(config.Env("JWT_ADMIN_SECRET_KEY"))

type Middleware struct {
}

func (m Middleware) JwtSign(id primitive.ObjectID, isAdmin bool) (string, error) {
	ttl := 1440 * time.Minute
	var claims = jwt.MapClaims{}
	claims["id"] = id
	claims["exp"] = time.Now().UTC().Add(ttl).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if !isAdmin {
		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			return tokenString, err
		}
		return tokenString, nil
	} else {
		tokenString, err := token.SignedString(adminSecretKey)
		if err != nil {
			return tokenString, err
		}
		return tokenString, nil
	}
}

func (m Middleware) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("authorization")
		if len(authorization) < 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ResErr(http.StatusUnauthorized, "unauthorize"))
			return
		}
		tokenString := strings.Split(authorization, " ")
		if len(tokenString) < 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ResErr(http.StatusUnauthorized, "unauthorize"))
			return
		}
		token, err := jwt.Parse(tokenString[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return secretKey, nil
		})

		if tokenClaims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("user", tokenClaims)
			ctx.Next()
		} else {
			fmt.Println(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ResErr(http.StatusUnauthorized, err.Error()))
		}
	}
}

func (m Middleware) AuthenticationAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("authorization")
		if len(authorization) < 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ResErr(http.StatusUnauthorized, "unauthorize"))
			return
		}
		tokenString := strings.Split(authorization, " ")
		if len(tokenString) < 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ResErr(http.StatusUnauthorized, "unauthorize"))
			return
		}
		token, err := jwt.Parse(tokenString[1], func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}
			return adminSecretKey, nil
		})

		if tokenClaims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("user", tokenClaims)
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.ResErr(http.StatusUnauthorized, err.Error()))
		}
	}
}
