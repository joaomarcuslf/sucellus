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
	router.LoadHTMLGlob("templates/**/*")

	service_pages := handlers.NewServicePagesHandler(a.connection)

	router.GET("/", service_pages.ListServices)

	router.GET("/services", service_pages.ListServices)
	router.GET("/services/new", service_pages.AddService)
	router.GET("/services/:id", service_pages.ShowService)
	router.GET("/services/:id/edit", service_pages.ListServices)

	api_routes := router.Group("/api")
	{
		api_routes.Use(middlewares.JSONContent())
		api_routes.GET("/health", api_handlers.Health)

		service_routes := api_routes.Group("/services")
		{
			service_api := service_handlers.NewServiceHandler(a.connection)

			service_routes.GET("/", service_api.List)
			service_routes.POST("/", service_api.Create)
			service_routes.GET("/:id", service_api.Get)
			service_routes.PUT("/:id", service_api.Update)
			service_routes.DELETE("/:id", service_api.Delete)

			service_routes.POST("/:id/start", service_api.StartService)
			service_routes.POST("/:id/stop", service_api.StopService)

			service_routes.POST("/start", service_api.StartServices)
			service_routes.POST("/stop", service_api.StopServices)
		}
	}

	router.Run(":" + a.Port)
}
