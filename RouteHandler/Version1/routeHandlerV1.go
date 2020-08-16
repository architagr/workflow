package Version1

import (
	"github.com/architagr/workflow/models"
)

type FlowHandler struct {
	FlowId      int64
	TaskId      int64
	FlowDetails models.Flow
}

// @Summary this method gets all the flows for the company whose company id is been passed in token
// @Description 登录
// @Produce json
// @Param body body controllers.LoginParams true "body参数"
// @Success 200 {string} string "ok" "返回用户信息"
// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
// @Failure 401 {string} string "err_code：10001 登录失败"
// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
// @Router /user/person/login [post]
func (k *FlowHandler) GetAll(apiKey string) ([]models.Flow, int) {
	flows := []models.Flow{
		models.Flow{Id: 1},
	}
	return flows, len(flows)
}

func (k *FlowHandler) AddNewFlow(apiKey string) (models.Flow, error) {
	k.FlowDetails.Id = 1
	k.FlowDetails.Version = "1"
	return k.FlowDetails, nil
}
