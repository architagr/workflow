package version1

import (
	"fmt"
	http "net/http"

	middleware "github.com/architagr/workflow/Middlewares"
	routeHandler "github.com/architagr/workflow/RouteHandler/Version1"
	models "github.com/architagr/workflow/models"
	gin "github.com/gin-gonic/gin"
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
				c.JSON(http.StatusOK,
					models.ListResponse{
						TotalCount:     len(allWorkflow),
						StartPageIndex: 1,
						EndPageIndex:   len(allWorkflow),
						Data:           allWorkflow,
						Status:         http.StatusOK,
						Message:        fmt.Sprintf("total count of records are %v", len(allWorkflow)),
					},
				)

			})
			flow.POST("", func(c *gin.Context) {
				var objA models.Flow
				if errA := c.ShouldBind(&objA); errA == nil {

					apiKey := fmt.Sprintf("%v", c.Keys[middleware.GetContextKeysName(middleware.ContextApiKey)])
					var flowhadler routeHandler.FlowHandler
					flowhadler.FlowDetails = objA
					newWorkflow, err := flowhadler.AddNewFlow(apiKey)
					if err == nil {
						c.JSON(http.StatusOK,
							models.Response{
								Data:    newWorkflow,
								Status:  http.StatusOK,
								Message: fmt.Sprintf("Workflow Added with flow Id %v", newWorkflow.Id),
							},
						)
					} else {
						c.JSON(http.StatusBadRequest,
							models.Response{
								Error: []models.ErrorDetail{
									models.ErrorDetail{
										ErrorType:    models.GetErrorTypeName(models.ErrorTypeError),
										ErrorMessage: fmt.Sprintf("%v", err),
									},
								},
								Status:  http.StatusBadRequest,
								Message: fmt.Sprintf("Workflow not added"),
							},
						)
					}
				} else {
					c.JSON(http.StatusBadRequest,
						models.Response{
							Error: []models.ErrorDetail{
								models.ErrorDetail{
									ErrorType:    models.GetErrorTypeName(models.ErrorTypeError),
									ErrorMessage: fmt.Sprintf("%v", errA),
								},
							},
							Status:  http.StatusBadRequest,
							Message: fmt.Sprintf("Workflow not added"),
						},
					)
				}
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
