package handlers

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CaptureException(c *gin.Context, err error) {
	if hub := sentrygin.GetHubFromContext(c); hub != nil {
		claims := jwt.ExtractClaims(c)
		userID, ok := claims["id"]
		hub.ConfigureScope(func(scope *sentry.Scope) {
			usr := sentry.User{
				IPAddress: c.ClientIP(),
			}
			if ok {
				usr.ID = strconv.Itoa(int(userID.(float64)))
			}
			scope.SetUser(usr)
			hub.CaptureException(err)
		})
	}
	fmt.Println(err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
}
