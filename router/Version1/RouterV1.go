package version1

import (
	"fmt"
	middleware "github.com/architagr/workflow/Middlewares"
	routeHandler "github.com/architagr/workflow/RouteHandler/Version1"
	gin "github.com/gin-gonic/gin"
	http "net/http"
)

//route paramenters indexing
const (
	FlowId = 1
	TaskId = 2
)

//route paramented name
var routeParameterName = map[int]string{
	FlowId: "flowId",
	TaskId: "taskId",
}

// Get route parameter string according to the index number of the route parameter

func getRouteParameter(code int) string {
	return routeParameterName[code]
}

func RouterVersion1(r *gin.Engine) {

	workflow := r.Group("v1")
	workflow.Use(middleware.AuthenticationRequired())
	{
		flow := workflow.Group("flow")
		{
			flow.GET("", func(c *gin.Context) {
				apiKey := fmt.Sprintf("%v", c.Keys[middleware.GetContextKeysName(middleware.ContextApiKey)])
				var flowhadler routeHandler.FlowHandler
				allWorkflow, _ := flowhadler.GetAll(apiKey)
				c.JSON(http.StatusOK, gin.H{
					"Data":        allWorkflow,
					"CompanyName": c.Keys[middleware.GetContextKeysName(middleware.ContextCmpnayDetails)],
					"apikey":      apiKey,
				})

			})
			flow.POST("", func(c *gin.Context) {
				fmt.Println("v1/flow -> post")
				c.String(http.StatusOK, c.FullPath())
			})
		}

		flowById := workflow.Group("flow/:" + getRouteParameter(FlowId))
		{
			flowById.GET("", func(c *gin.Context) {
				fmt.Println("v1/flow/:flowId -> get")
				c.String(http.StatusOK, c.FullPath())

			})
			flowById.PUT("", func(c *gin.Context) {
				fmt.Println("v1/flow/:flowId -> put")
				c.String(http.StatusOK, c.FullPath())

			})
			flowById.DELETE("", func(c *gin.Context) {
				fmt.Println("v1/flow/:flowId -> delete")
				c.String(http.StatusOK, c.FullPath())
			})

		}

	}

}
