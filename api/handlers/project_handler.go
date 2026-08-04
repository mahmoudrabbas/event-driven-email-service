package handlers

import (
	"net/http"

	"example.com/models"
	"example.com/rabbitmq"
	"example.com/shared"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProject(c *gin.Context) {

	var req models.CreateProjectRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	pid := uuid.New().String()

	event := shared.ProjectCreatedEvent{

		ReferenceID: pid,

		ClientName: req.ClientName,
		Email:      req.Email,
		Company:    req.Company,

		ProjectName: req.ProjectName,
		Service:     req.Service,
		Budget:      req.Budget,

		Description: req.Description,

		MeetingLink: "https://meet.valtq.dev/discovery/" + pid,
	}

	err := rabbitmq.Publish(event)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{

		"message": "Discovery submitted successfully",

		"reference_id": pid,
	})
}

func GETHome(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"msg": "zai elfol ya 3m el7ag",
	})
}
