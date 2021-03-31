package sdk

import (
	"fmt"
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//CreateCalendarV4 创建日历v4版本 https://open.feishu.cn/document/ukTMukTMukTM/uUDM3YjL1AzN24SNwcjN
func (t Tenant) CreateCalendarV4(bodyParams vo.CreateCalendarV4Req) (*vo.CommonCalendarV4Resp, error) {
	respBody, err := http.Post(consts.ApiCalendarCreateV4, nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonCalendarV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//DeleteCalendarV4 删除日历 https://open.feishu.cn/document/ukTMukTMukTM/uYDM3YjL2AzN24iNwcjN
func (t Tenant) DeleteCalendarV4(calendarId string) (*vo.VoidV4Resp, error) {
	respBody, err := http.Delete(fmt.Sprintf(consts.ApiCalendarDeleteV4, calendarId), nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.VoidV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//GetCalendarV4 获取日历 https://open.feishu.cn/document/ukTMukTMukTM/ucDM3YjL3AzN24yNwcjN
func (t Tenant) GetCalendarV4(calendarId string) (*vo.GetCalendarV4Resp, error) {
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarGetV4, calendarId), nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// 获取日历列表 https://open.feishu.cn/document/ukTMukTMukTM/ugDM3YjL4AzN24COwcjN  todo
func (t Tenant) GetCalendarListV4(pageSize int, pageToken, syncToken string) (*vo.GetCalendarListV4Resp, error) {
	queryParams := map[string]interface{}{
		"page_size":  pageSize,
		"page_token": pageToken,
		"sync_token": syncToken,
	}
	respBody, err := http.Post(consts.ApiCalendarGetListV4, queryParams, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarListV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//UpdateCalendarV4 更新日历 https://open.feishu.cn/document/ukTMukTMukTM/ukDM3YjL5AzN24SOwcjN
func (t Tenant) UpdateCalendarV4(calendarId string, bodyParams vo.UpdateCalendarV4Req) (*vo.CommonCalendarV4Resp, error) {
	respBody, err := http.Patch(fmt.Sprintf(consts.ApiCalendarUpdateV4, calendarId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonCalendarV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//SearchCalendarV4 搜索日历 https://open.feishu.cn/document/ukTMukTMukTM/uATM3YjLwEzN24CMxcjN
func (t Tenant) SearchCalendarV4(bodyParams vo.SearchCalendarV4Req, pageToken string, pageSize int) (*vo.SearchCalendarV4Resp, error) {
	queryParams := map[string]interface{}{
		"page_token": pageToken,
		"page_size":  pageSize,
	}
	respBody, err := http.Post(consts.ApiCalendarSearchV4, queryParams, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.SearchCalendarV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//UnsubscribeCalendarV4 取消订阅日历 https://open.feishu.cn/document/ukTMukTMukTM/ugDO3YjL4gzN24CO4cjN
func (t Tenant) UnsubscribeCalendarV4(calendarId string) (*vo.VoidV4Resp, error) {
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarUnsubscribeV4, calendarId), nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.VoidV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
//SubscribeCalendarV4 订阅日历 https://open.feishu.cn/document/ukTMukTMukTM/ucDO3YjL3gzN24yN4cjN
func (t Tenant) SubscribeCalendarV4(calendarId string) (*vo.SubscribeCalendarV4Resp, error) {
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarSubscribeV4, calendarId), nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.SubscribeCalendarV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//SubscriptionCalendarV4 订阅日历变更事件 https://open.feishu.cn/document/ukTMukTMukTM/ugTO2YjL4kjN24CO5YjN/subscribe%20calendar%20changed%20event
func (t Tenant) SubscriptionCalendarV4() (*vo.VoidV4Resp, error) {
	respBody, err := http.Post(consts.ApiCalendarSubscriptionV4, nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.VoidV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
