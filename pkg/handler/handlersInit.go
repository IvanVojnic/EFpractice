package handler

import (
	"EFpractic2/pkg/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

/*func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}*/

func (h *Handler) InitRoutes(router *echo.Echo) *echo.Echo {

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	rAct := router.Group("/act")
	rAct.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339} ${host} ${method}]`,
	}))

	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello worl")
	})

	rAct.POST("/create", h.createUser)
	rAct.GET("/get", h.getUser)
	rAct.POST("/update", h.updateUser)
	rAct.GET("/delete", h.deleteUser)
	rAct.GET("/getAllUser", h.getAllUsers)
	router.Start(":3000")
	return router
}
