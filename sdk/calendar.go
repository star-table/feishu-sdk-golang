package sdk

import (
	"fmt"
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//获取日历 https://open.feishu.cn/document/ukTMukTMukTM/uMDN04yM0QjLzQDN?lang=zh-CN
func (t Tenant) GetCalendar(calendarId string) (*vo.CommonCalendarResp, error) {
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarGet, calendarId), nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonCalendarResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取日历列表
func (t Tenant) GetCalendarList(maxResults *int64, pageToken *string, syncToken *string) (*vo.CalendarListResp, error) {
	queryParams := map[string]interface{}{
		"max_results": maxResults,
		"page_token":  pageToken,
		"sync_token":  syncToken,
	}
	respBody, err := http.Get(consts.ApiCalendarListGet, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CalendarListResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//创建日历 https://open.feishu.cn/document/ukTMukTMukTM/uQTM14CNxUjL0ETN?lang=zh-CN
func (t Tenant) CreateCalendar(bodyParams vo.CreateCalendarReq) (*vo.CommonCalendarResp, error) {
	respBody, err := http.Post(consts.ApiCalendarCreate, nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonCalendarResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//更新日历 https://open.feishu.cn/document/ukTMukTMukTM/uYTM14iNxUjL2ETN?lang=zh-CN
func (t Tenant) UpdateCalendar(calendarId string, bodyParams vo.UpdateCalendarReq) (*vo.CommonCalendarResp, error) {
	respBody, err := http.Patch(fmt.Sprintf(consts.ApiCalendarUpdate, calendarId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonCalendarResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//删除日历 https://open.feishu.cn/document/ukTMukTMukTM/uUTM14SNxUjL1ETN?lang=zh-CN
func (t Tenant) DeleteCalendar(calendarId string) (*vo.CommonVo, error) {
	respBody, err := http.Delete(fmt.Sprintf(consts.ApiCalendarUpdate, calendarId), nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取日程列表 https://open.feishu.cn/document/ukTMukTMukTM/ukTM14SOxUjL5ETN
func (t Tenant) GetCalendarEventList(calendarId string, maxResults *int64, pageToken *string, syncToken *string) (*vo.CalendarEventListResp, error) {
	queryParams := map[string]interface{}{
		"max_results": maxResults,
		"page_token":  pageToken,
		"sync_token":  syncToken,
	}
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarEventCreate, calendarId), queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CalendarEventListResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//创建日程 https://open.feishu.cn/document/ukTMukTMukTM/ugTM14COxUjL4ETN?lang=zh-CN
func (t Tenant) CreateCalendarEvent(calendarId string, bodyParams vo.CreateCalendarEventReq) (*vo.CommonCalendarEventResp, error) {
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarEventCreate, calendarId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonCalendarEventResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//删除日程 https://open.feishu.cn/document/ukTMukTMukTM/uAjM14CMyUjLwITN?lang=zh-CN
func (t Tenant) DeleteCalendarEvent(calendarId string, eventId string) (*vo.CommonVo, error) {
	respBody, err := http.Delete(fmt.Sprintf(consts.ApiCalendarEventDelete, calendarId, eventId), nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//更新日程 https://open.feishu.cn/document/ukTMukTMukTM/uEjM14SMyUjLxITN
func (t Tenant) UpdateCalendarEvent(calendarId string, eventId string, bodyParams vo.CreateCalendarEventReq) (*vo.CommonCalendarEventResp, error) {
	respBody, err := http.Patch(fmt.Sprintf(consts.ApiCalendarEventDelete, calendarId, eventId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonCalendarEventResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//邀请/移除日程参与者 https://open.feishu.cn/document/ukTMukTMukTM/uIjM14iMyUjLyITN
func (t Tenant) UpdateCalendarEventAttendees(calendarId string, eventId string, bodyParams vo.UpdateCalendarEventAtendeesReq) (*vo.UpdateCalendarEventAtendeesResp, error) {
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarEventAttendeesUpdate, calendarId, eventId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.UpdateCalendarEventAtendeesResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取访问控制列表 https://open.feishu.cn/document/ukTMukTMukTM/uMjM14yMyUjLzITN
func (t Tenant) GetCalendarAttendeesAcl(calendarId string) (*vo.GetCalendarAttendeesResp, error) {
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarAttendeesGet, calendarId), nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarAttendeesResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//创建访问控制 https://open.feishu.cn/document/ukTMukTMukTM/uQjM14CNyUjL0ITN
func (t Tenant) AddCalendarAttendeesAcl(calendarId string, bodyParams vo.AddCalendarAttendeesAclReq) (*vo.GetCalendarAttendeesResp, error) {
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarAttendeesGet, calendarId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarAttendeesResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//删除访问控制 https://open.feishu.cn/document/ukTMukTMukTM/uUjM14SNyUjL1ITN
func (t Tenant) DeleteCalendarAttendeesAcl(calendarId string, ruleId string) (*vo.CommonVo, error) {
	respBody, err := http.Delete(fmt.Sprintf(consts.ApiCalendarAttendeesDelete, calendarId, ruleId), nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
