package sdk

import (
	"fmt"
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//创建日历v4版本 https://open.feishu.cn/document/ukTMukTMukTM/uUDM3YjL1AzN24SNwcjN
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

//创建日程v4 https://open.feishu.cn/document/ukTMukTMukTM/uYTM3YjL2EzN24iNxcjN
func (t Tenant) CreateCalendarEventV4(calendarId string, bodyParams vo.CreateCalendarEventV4Req) (*vo.CalendarEventInfoV4Resp, error) {
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarEventCreateV4, calendarId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CalendarEventInfoV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//删除日程v4 https://open.feishu.cn/document/ukTMukTMukTM/ucTM3YjL3EzN24yNxcjN
func (t Tenant) DeleteCalendarEventV4(calendarId string, eventId string) (*vo.CommonVo, error) {
	respBody, err := http.Delete(fmt.Sprintf(consts.ApiCalendarEventDeleteV4, calendarId, eventId), nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取日程v4 https://open.feishu.cn/document/ukTMukTMukTM/ugTM3YjL4EzN24COxcjN
func (t Tenant) GetCalendarEventV4(calendarId string, eventId string) (*vo.CalendarEventInfoV4Resp, error) {
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarEventGetV4, calendarId, eventId), nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CalendarEventInfoV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取日程列表v4 https://open.feishu.cn/document/ukTMukTMukTM/ukTM3YjL5EzN24SOxcjN
func (t Tenant) GetCalendarEventListV4(calendarId string, eventId string, pageSize int, pageToken *string, syncToken *string) (*vo.GetCalendarEventListV4Resp, error) {
	queryParams := map[string]interface{}{
		"page_size":  pageSize,
		"page_token": pageToken,
		"sync_token": syncToken,
	}
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarEventBatchGetV4, calendarId), queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarEventListV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//更新日程v4 https://open.feishu.cn/document/ukTMukTMukTM/uAjM3YjLwIzN24CMycjN
func (t Tenant) UpdateCalendarEventV4(calendarId string, eventId string, bodyParams vo.UpdateCalendarEventV4Req) (*vo.CalendarEventInfoV4Resp, error) {
	respBody, err := http.Patch(fmt.Sprintf(consts.ApiCalendarEventUpdateV4, calendarId, eventId), nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CalendarEventInfoV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//创建日程参与人v4 https://open.feishu.cn/document/ukTMukTMukTM/uAjM3YjLwIzN24CMycjN
func (t Tenant) AddCalendarEventAttendeesV4(calendarId string, eventId string, userIdType string, bodyParams vo.AddCalendarEventAttendeesV4Req) (*vo.AddCalendarEventAttendeesV4Resp, error) {
	params := map[string]interface{}{
		"user_id_type": userIdType,
	}
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarEventAttendeesAddV4, calendarId, eventId), params, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.AddCalendarEventAttendeesV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//删除日程参与人v4 https://open.feishu.cn/document/ukTMukTMukTM/uAzM3YjLwMzN24CMzcjN
func (t Tenant) DeleteCalendarEventAttendeesV4(calendarId string, eventId string, userIdType string, bodyParams vo.DeleteCalendarEventAttendeesV4Req) (*vo.CommonVo, error) {
	params := map[string]interface{}{
		"user_id_type": userIdType,
	}
	respBody, err := http.Post(fmt.Sprintf(consts.ApiCalendarEventAttendeesDeleteV4, calendarId, eventId), params, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取日程参与人列表v4 https://open.feishu.cn/document/ukTMukTMukTM/uEzM3YjLxMzN24SMzcjN
func (t Tenant) GetCalendarEventAttendeesV4(calendarId string, eventId string, pageSize int, pageToken *string, userIdType string) (*vo.GetCalendarEventAttendeesV4Resp, error) {
	queryParams := map[string]interface{}{
		"page_size":    pageSize,
		"page_token":   pageToken,
		"user_id_type": userIdType,
	}
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarEventAttendeesGetV4, calendarId, eventId), queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarEventAttendeesV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取参与人群成员列表v4 https://open.feishu.cn/document/ukTMukTMukTM/uATN3YjLwUzN24CM1cjN
func (t Tenant) GetCalendarEventAttendeesChatMembersV4(calendarId string, eventId string, attendeeId string, pageSize int, pageToken *string, userIdType string) (*vo.GetCalendarEventAttendeesV4Resp, error) {
	queryParams := map[string]interface{}{
		"page_size":    pageSize,
		"page_token":   pageToken,
		"user_id_type": userIdType,
	}
	respBody, err := http.Get(fmt.Sprintf(consts.ApiCalendarEventAttendeesChatMembersGetV4, calendarId, eventId, attendeeId), queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetCalendarEventAttendeesV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
