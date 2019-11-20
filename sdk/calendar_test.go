package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"
	"testing"
)

var calendarId = "feishu.cn_lNRyD5571uIiNGqLl1KZwe@group.calendar.feishu.cn"
var calendarEventId = "9e1e3aed-3fa9-486c-99bf-4b42ffdbe384"

func TestTenant_GetCalendar(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetCalendar(calendarId)
	log.Info("日历:", json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 200000)

	resp2, err2 := tenant.GetCalendarList(nil, nil, nil)
	log.Info("日历列表：", json.ToJsonIgnoreError(resp2), err2)
	assert.Equal(t, err2, nil)
	assert.Equal(t, resp2.Code, 200000)
}

func TestTenant_CreateCalendar(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.CreateCalendar(vo.CreateCalendarReq{
		Summary:     "日历就是日历",
		Description: "测试使用",
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 200000)
}

func TestTenant_UpdateCalendar(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.UpdateCalendar(calendarId, vo.UpdateCalendarReq{
		Summary:           "日历就是日历",
		Description:       "测试使用哈",
		DefaultAccessRole: "free_busy_reader",
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 200000)
}

func TestTenant_CreateCalendarEvent(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	attend := []vo.Attendees{}
	//attend = append(attend, vo.Attendees{
	//	OpenId:      "ou_87f1b2210acad10a90cc3690802626d7",
	//	EmployeeId:  "",
	//	DisplayName: "fanfan",
	//	Optional:    true,
	//})
	resp, err := tenant.CreateCalendarEvent(calendarId, vo.CreateCalendarEventReq{
		Summary:     "测试测试不加用户怕",
		Description: "日程二号",
		Start: vo.TimeFormat{
			Date: "2019-11-23",
		},
		End: vo.TimeFormat{
			Date: "2019-11-25",
		},
		Attendees: &attend,
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 200000)
}

func TestTenant_GetCalendarEventList(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetCalendarEventList("feishu.cn_NEb1HqWfXyjSKu4zQ4fe3b@group.calendar.feishu.cn", nil, nil, nil)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 200000)
}

func TestTenant_UpdateCalendarEvent(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.UpdateCalendarEvent(calendarId, calendarEventId, vo.CreateCalendarEventReq{
		Summary:     "好一个日程",
		Description: "hhh",
		Start: vo.TimeFormat{
			Date: "2019-11-19",
		},
		End: vo.TimeFormat{
			Date: "2019-11-21",
		},
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 200000)
}

func TestTenant_DeleteCalendarEvent(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.DeleteCalendarEvent("feishu.cn_NEb1HqWfXyjSKu4zQ4fe3b@group.calendar.feishu.cn", "3cca07ad-206e-41c5-92d9-29be0118b043")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 200000)
}

func TestTenant_UpdateCalendarEventAttendees(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	//{"employee_id":"","open_id":"ou_87f1b2210acad10a90cc3690802626d7","user_id":"","name":"樊宇","employee_no":""}
	attend := []vo.AttendeesResp{}
	attend = append(attend, vo.AttendeesResp{
		Status: "invite",
		Attendees: vo.Attendees{
			OpenId:      "ou_87f1b2210acad10a90cc3690802626d7",
			EmployeeId:  "",
			DisplayName: "fanfan",
			Optional:    true,
		},
	})
	resp, err := tenant.UpdateCalendarEventAttendees("feishu.cn_NEb1HqWfXyjSKu4zQ4fe3b@group.calendar.feishu.cn", "3cca07ad-206e-41c5-92d9-29be0118b043", vo.UpdateCalendarEventAtendeesReq{
		Attendees: attend,
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 200000)
}

func TestTenant_GetCalendarAttendeesAcl(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetCalendarAttendeesAcl(calendarId)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 200000)
}

func TestTenant_AddCalendarAttendeesAcl(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.AddCalendarAttendeesAcl(calendarId, vo.AddCalendarAttendeesAclReq{
		Role: consts.AccessRoleReader,
		Scope: vo.CalendarScope{
			Type:   "user",
			OpenId: "ou_87f1b2210acad10a90cc3690802626d7",
		},
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 200000)
}

func TestTenant_DeleteCalendarAttendeesAcl(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)
	resp, err := tenant.DeleteCalendarAttendeesAcl(calendarId, "ou_87f1b2210acad10a90cc3690802626d7")

	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 200000)
}
