package Middlewares

import (
	"net/http"

	util "github.com/architagr/workflow/internal/util"
	models "github.com/architagr/workflow/models"
	gin "github.com/gin-gonic/gin"
)

const (
	ContextApiKey        = 1
	ContextCmpnayDetails = 2
)

var ContextKeysName = map[int]string{
	ContextApiKey:        "apikey",
	ContextCmpnayDetails: "companyDetails",
}

func GetContextKeysName(code int) string {
	return ContextKeysName[code]
}

// @Summary 登录
// @Description 登录
// @Produce json
// @Param body body controllers.LoginParams true "body参数"
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router /user/person/login [post]
func AuthenticationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		apikey := util.GetApiKey(c.Request)
		if apikey == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				models.Response{
					Status: http.StatusUnauthorized,
					Error: []models.ErrorDetail{
						{
							ErrorType:    models.GetErrorTypeName(models.ErrorTypeFatal),
							ErrorMessage: "Api key is not presend in header",
						},
					},
					Message: "One or more error occured",
				},
			)
			return
		}
		keys := map[string]string{
			"abc": "Archit",
		}
		companyName, found := keys[apikey]
		if !found {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				models.Response{
					Status: http.StatusUnauthorized,
					Error: []models.ErrorDetail{
						models.ErrorDetail{
							ErrorType:    models.GetErrorTypeName(models.ErrorTypeFatal),
							ErrorMessage: "invalid Api key, restricted endpoint",
						},
					},
					Message: "One or more error occured",
				},
			)
			return
		}
		if len(c.Keys) == 0 {
			c.Keys = make(map[string]interface{})
		}
		c.Keys[GetContextKeysName(ContextCmpnayDetails)] = companyName
		c.Keys[GetContextKeysName(ContextApiKey)] = apikey

		c.Next()
	}
}
