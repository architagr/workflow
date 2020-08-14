package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/architagr/workflow/models"
)

func CreateWorkFlow(c *gin.Context,configuration *models.Configuration){
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"port":configuration.Port,
		"Connection": configuration.DataBaseConnection,
	})
}

func UpdateWorkFlow(c *gin.Context, configuration *models.Configuration){
	var x map[string]interface{}
	if c.ShouldBind(&x) == nil {

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "hi from routehandler ",
		"port":configuration.Port,
		"Connection": configuration.DataBaseConnection,
		"body":x,
	})
}