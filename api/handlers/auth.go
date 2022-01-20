package handlers

import (
	"errors"
	"github.com/mospyx/data-accounting/pkg/models"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var JWT = initJWT()

func initJWT() *jwt.GinJWTMiddleware {
	jwtInit := jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour * 24,
		MaxRefresh:  time.Hour * 24,
		IdentityKey: "id",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					"id":   v.ID,
					"role": v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			u := &models.User{
				ID:   uint(claims["id"].(float64)),
				Role: models.UserRole(claims["role"].(string)),
			}
			return u
		},
		Authenticator: Login,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*models.User); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"Error": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}

	return &jwtInit
}

type LoginHandlerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) (interface{}, error) {
	var user LoginHandlerRequest

	err := c.BindJSON(&user)
	if err != nil {
		return nil, err
	}

	userDict, err := models.Login(user.Email)
	if err != nil {
		return nil, err
	}

	if !CheckPasswordHash(user.Password, userDict.Password) {
		return nil, errors.New("wrong email or password")
	}
	if !userDict.Active {
		return nil, errors.New("account deactivated")
	}

	return userDict, nil
}

type RegisterHandlerRequest struct {
	Email     string `json:"email"`
	Password1 string `json:"password_1"`
	Password2 string `json:"password_2"`
}

func Register(c *gin.Context) {
	var req RegisterHandlerRequest

	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		CaptureException(c, err)
		return
	}

	req.Email = strings.ToLower(req.Email)

	regexpEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !regexpEmail.MatchString(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Oops! It doesn't look like correct email! "})
		CaptureException(c, err)
		return
	}

	email := models.CheckEmail(req.Email)
	if email != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already in use. "})
		return
	}

	pass := models.Password(req.Password1)
	if !pass {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Your password must be 8+ characters and contain at least one uppercase, one symbol and one number digit. "})
		return
	}

	if req.Password1 != req.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords doesn't match. "})
		return
	}

	hash, err := HashPassword(req.Password1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		CaptureException(c, err)
		return
	}

	usr := models.User{
		Email:    req.Email,
		Password: hash,
		Role:     "general",
		Active:   false,
	}

	hostName := os.Getenv("HOST_NAME")
	if hostName == "local" {
		usr.Active = true
	}

	if err = usr.Create(); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		CaptureException(c, err)
		return
	}

	c.JSON(http.StatusCreated, usr)
	return
}

func role(claims jwt.MapClaims) models.UserRole {
	return models.UserRole(claims["role"].(string))
}

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		if role(claims).IsAdmin() {
			c.Next()
		} else {
			c.Status(http.StatusForbidden)
			return
		}
	}
}
