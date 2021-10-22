package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcuslf/sucellus/definitions"
	"github.com/joaomarcuslf/sucellus/models"
	"github.com/joaomarcuslf/sucellus/repositories"
)

type ServicePage struct {
	Title   string
	Service models.Service
}

type ServicePagesHandler struct {
	repository definitions.Repository
	connection definitions.DatabaseClient
}

func NewServicePagesHandler(connection definitions.DatabaseClient) *ServicePagesHandler {
	return &ServicePagesHandler{
		repository: repositories.NewServiceRepository(connection),
		connection: connection,
	}
}

func (s *ServicePagesHandler) ListServices(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"list_services.html",
		ServicePage{
			Title: "List Services",
		},
	)
}

func (s *ServicePagesHandler) AddService(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"add_service.html",
		ServicePage{
			Title: "Add Service",
		},
	)
}

func (s *ServicePagesHandler) ShowService(c *gin.Context) {
	id := c.Param("id")

	data, _ := s.repository.Get(c, id)

	c.HTML(
		http.StatusOK,
		"show_service.html",
		ServicePage{
			Title:   "Show Service",
			Service: data.(models.Service),
		},
	)
}
