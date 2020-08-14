package workflow

import (
	"flag"
	"fmt"
	"github.com/architagr/workflow/models"
	"github.com/architagr/workflow/controller"
	"github.com/gin-gonic/gin"
)

var configuration models.Configuration

func main() {
	env := flag.String("env", "dev", "environment to be used")
	flag.Parse()

	fmt.Printf("Envirnment used is : %s\n", *env)

	configuration, err := configuration.Init(*env)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("config file issue for environment %s\n", *env)
	} else {
		fmt.Printf("config file read for env %s\n", *env)
	}
	r := gin.Default()
	configureRoutes(r, &configuration)
	r.Run(configuration.Ip + ":" + configuration.Port)
}

func configureRoutes(r *gin.Engine, configuration *models.Configuration) {
	r.GET("/ping", func(c *gin.Context) {
		controller.CreateWorkFlow(c, configuration)
	})
	r.GET("/ping1", func(c *gin.Context) {
		controller.UpdateWorkFlow(c, configuration)
	})


	workflowVersion1:=r.Group("v1")
	{
		flow:= workflowVersion1.Group("workflow")
		{
			flow.GET("",func(c *gin.Context) {
				controller.CreateWorkFlow(c, configuration)
			})
			flow.POST("",func(c *gin.Context) {
				controller.UpdateWorkFlow(c, configuration)
			})
		}
	}
}
