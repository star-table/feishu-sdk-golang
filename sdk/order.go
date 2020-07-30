package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//查询用户是否在应用开通范围 https://open.feishu.cn/document/ukTMukTMukTM/uATNwUjLwUDM14CM1ATN
func (t Tenant) CheckUser(req vo.CheckUserReq) (*vo.CheckUserResp, error) {
	queryParams := map[string]interface{}{}
	if req.UserId != "" {
		queryParams["user_id"] = req.UserId
	}
	if req.OpenId != "" {
		queryParams["open_id"] = req.OpenId
	}
	respBody, err := http.Get(consts.ApiCheckUser, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CheckUserResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//查询租户购买的付费方案 https://open.feishu.cn/document/ukTMukTMukTM/uETNwUjLxUDM14SM1ATN
func (t Tenant) GetOrderList(req vo.GetOrderListReq) (*vo.GetOrderListResp, error) {
	queryParams := map[string]interface{}{}
	queryParams["page_size"] = req.PageSize
	if req.Status != nil {
		queryParams["status"] = req.Status
	}
	if req.PageToken != nil {
		queryParams["page_token"] = req.PageToken
	}
	if req.TenantKey != nil {
		queryParams["tenant_key"] = req.TenantKey
	}
	respBody, err := http.Get(consts.ApiGetOrderList, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetOrderListResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//查询订单详情 https://open.feishu.cn/document/ukTMukTMukTM/uITNwUjLyUDM14iM1ATN
func (t Tenant) GetOrderInfo(orderId string) (*vo.OrderInfoResp, error) {
	queryParams := map[string]interface{}{}
	queryParams["order_id"] = orderId
	respBody, err := http.Get(consts.ApiGetOrderInfo, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.OrderInfoResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
