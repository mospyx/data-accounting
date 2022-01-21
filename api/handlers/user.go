package handlers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/mospyx/data-accounting/pkg/models"
	"net/http"
)

func GetCurrentUser(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userIdstr, _ := claims["id"]
	userId := uint(userIdstr.(float64))
	user, err := models.GetUser(userId)
	if err != nil {
		CaptureException(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
