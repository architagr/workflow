package Middlewares

import (
	util "github.com/architagr/workflow/internal/util"
	gin "github.com/gin-gonic/gin"
	"net/http"
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
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Api key is not presend in header"})
			return
		}
		keys := map[string]string{
			"abc": "Archit",
		}
		companyName, found := keys[apikey]
		if !found {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid Api key, restricted endpoint"})
			return
		}
		if len(c.Keys) == 0 {
			c.Keys = make(map[string]interface{})
		}
		c.Keys[GetContextKeysName(ContextCmpnayDetails)] = companyName
		c.Keys[GetContextKeysName(ContextApiKey)] = apikey
		// add session verification here, like checking if the user and authType
		// combination actually exists if necessary. Try adding caching this (redis)
		// since this middleware might be called a lot
		c.Next()
	}
}
