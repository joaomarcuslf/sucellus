package server

import (
	"github.com/gin-gonic/gin"
	"github.com/joaomarcuslf/sucellus/definitions"
	handlers "github.com/joaomarcuslf/sucellus/handlers"
	api_handlers "github.com/joaomarcuslf/sucellus/handlers/api"
	service_handlers "github.com/joaomarcuslf/sucellus/handlers/api/service"
	middlewares "github.com/joaomarcuslf/sucellus/middlewares"
)

type Server struct {
	Port       string
	connection definitions.DatabaseClient
}

func NewServer(port string, connection definitions.DatabaseClient) *Server {
	return &Server{
		Port:       port,
		connection: connection,
	}
}

func (a *Server) Run() {
	router := gin.New()

	router.Use(gin.Logger())

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", handlers.Index)

	api_routes := router.Group("/api")
	{
		api_routes.Use(middlewares.JSONContent())
		api_routes.GET("/health", api_handlers.Health)

		service_routes := api_routes.Group("/services")
		{
			service := service_handlers.NewServiceHandler(a.connection)

			service_routes.GET("/", service.List)
			service_routes.POST("/", service.Create)
			service_routes.GET("/:id", service.Get)
			service_routes.PUT("/:id", service.Update)
			service_routes.DELETE("/:id", service.Delete)

			service_routes.POST("/start", service.Start)
			service_routes.POST("/stop", service.Stop)
		}
	}

	router.Run(":" + a.Port)
}
