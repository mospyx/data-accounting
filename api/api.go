package api

import (
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/mospyx/data-accounting/api/handlers"
)

const (
	serverPort = ":5050"
)

func Start() error {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Length", "Content-Type"}
	config.AllowAllOrigins = true
	config.ExposeHeaders = []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin"}
	config.AllowCredentials = true

	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	r.Use(cors.New(config))
	r.Use(gin.Recovery())
	authMiddleware, err := jwt.New(handlers.JWT)
	if err != nil {
		return err
	}

	r.LoadHTMLGlob("./public/*.html")
	r.Use(static.Serve("/", static.LocalFile("public", true)))
	r.Static("/public", "public")
	frontend := r.Group("/")
	{
		frontend.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{})
		})
		frontend.GET("/register", func(c *gin.Context) {
			c.HTML(http.StatusOK, "register.html", gin.H{})
		})
	}

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authMiddleware.LoginHandler)
			auth.POST("/register", handlers.Register)
			auth.GET("/refresh", authMiddleware.RefreshHandler)
		}

		user := api.Group("/user")
		{
			user.GET("current_user", authMiddleware.MiddlewareFunc(), handlers.GetCurrentUser)
		}

		companies := api.Group("/companies")
		{
			companies.GET("", handlers.GetCompanies)
			companies.POST("", handlers.CreateCompany)
			company := companies.Group("/:company_id")
			{
				company.GET("", handlers.GetCompany)
				company.DELETE("", handlers.DeleteCompany)
				companyProfile := company.Group("/company_profile")
				{
					companyProfile.GET("", handlers.GetCompanyProfile)
					companyProfile.PUT("", handlers.UpdateCompanyProfile)
				}
				employees := company.Group("/employees")
				employees.GET("", handlers.GetEmployees)
				employees.POST("", handlers.CreateEmployee)
				{
					employee := employees.Group("/:employee_id")
					{
						employee.GET("", handlers.GetEmployee)
						employee.PUT("", handlers.UpdateEmployee)
						employee.DELETE("", handlers.DeleteEmployee)
					}
				}
			}
		}
	}

	log.Println("Starting server on port", serverPort)
	err = r.Run(serverPort)
	if err != nil {
		return err
	}

	return nil
}
