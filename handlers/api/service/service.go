package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcuslf/sucellus/definitions"
	"github.com/joaomarcuslf/sucellus/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

type ServiceHandler struct {
	repository *repositories.Service
}

func NewServiceHandler(connection definitions.DatabaseClient) *ServiceHandler {
	return &ServiceHandler{
		repository: repositories.NewServiceRepository(connection),
	}
}

func (s *ServiceHandler) List(c *gin.Context) {
	data, err := s.repository.Query(c, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (s *ServiceHandler) Create(c *gin.Context) {
	err := s.repository.Create(c, c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
	return
}

func (s *ServiceHandler) Get(c *gin.Context) {
	id := c.Param("id")

	data, err := s.repository.Get(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (s *ServiceHandler) Update(c *gin.Context) {
	id := c.Param("id")

	err := s.repository.Update(c, id, c.Request.Body)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := s.repository.Get(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated",
		"data":    data,
	})
}

func (s *ServiceHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	_, err := s.repository.Get(c, id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = s.repository.Delete(c, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted",
	})
}
