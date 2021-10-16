package server

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/joaomarcuslf/sucellus/handlers"
	api "github.com/joaomarcuslf/sucellus/handlers/api"
	service "github.com/joaomarcuslf/sucellus/handlers/api/service"
)

type Server struct {
	Port string
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

func (a *Server) Run() {
	router := gin.New()

	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		handlers.Index(c)
	})

	api_routes := router.Group("/api")
	{
		api_routes.GET("/health", func(c *gin.Context) {
			api.Health(c)
		})

		service_routes := api_routes.Group("/services")
		{
			service_routes.GET("/", func(c *gin.Context) {
				service.List(c)
			})

			service_routes.POST("/", func(c *gin.Context) {
				service.Create(c)
			})

			service_routes.GET("/:id", func(c *gin.Context) {
				service.Get(c)
			})

			service_routes.PUT("/:id", func(c *gin.Context) {
				service.Update(c)
			})

			service_routes.DELETE("/:id", func(c *gin.Context) {
				service.Delete(c)
			})
		}
	}

	router.Run(":" + a.Port)
}
