package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"
	"testing"
)

var calendarIdV4 = "feishu.cn_8brvNcf3OwwQOrux8ixWhc@group.calendar.feishu.cn"
var eventIdV4 = "a7f6ab60-a239-45e0-87a7-a19c550e37ef_0"

func TestTenant_CreateCalendarV4(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2e99b3ab0b0f1654")
	t.Log(e)

	resp, err := tenant.CreateCalendarV4(vo.CreateCalendarV4Req{
		Summary:     "BBB",
		Description: "AAAtest",
		Permissions: "public",
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 200000)
}

func TestTenant_DeleteCalendarV4(t *testing.T) {
	tenantKey := "130736cdb40f9758"
	tenant := GetTenant(tenantKey)
	resp, err := tenant.DeleteCalendarV4("")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestTenant_GetCalendarV4(t *testing.T) {
	tenantKey := "2e99b3ab0b0f1654"
	tenant := GetTenant(tenantKey)
	resp, err := tenant.GetCalendarV4("feishu.cn_Y8iPaCjvuBeG6WEdJ326qg@group.calendar.feishu.cn")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestTenant_GetCalendarListV4(t *testing.T) {
	tenantKey := "2e99b3ab0b0f1654"
	tenant := GetTenant(tenantKey)
	resp, err := tenant.GetCalendarListV4(10, "", "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestTenant_UpdateCalendarV4(t *testing.T) {
	tenantKey := "2e99b3ab0b0f1654"
	calendarId := "feishu.cn_Y8iPaCjvuBeG6WEdJ326qg@group.calendar.feishu.cn"
	tenant := GetTenant(tenantKey)
	resp, err := tenant.UpdateCalendarV4(calendarId, vo.UpdateCalendarV4Req{
		Summary:      "测试001",
		Description:  "测试001 desc",
		Permissions:  "public",
		Color:        0,
		SummaryAlias: "",
		CalendarId:   calendarId,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestTenant_SearchCalendarV4(t *testing.T) {
	tenantKey := "2e99b3ab0b0f1654"
	tenant := GetTenant(tenantKey)
	resp, err := tenant.SearchCalendarV4(vo.SearchCalendarV4Req{
		Query: "测试",
	}, "", 10)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestTenant_SubscribeCalendarV4(t *testing.T) {
	tenantKey := "2e99b3ab0b0f1654"
	calendarId := "feishu.cn_Y8iPaCjvuBeG6WEdJ326qg@group.calendar.feishu.cn"
	tenant := GetTenant(tenantKey)
	resp, err := tenant.SubscribeCalendarV4(calendarId)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestTenant_UnsubscribeCalendarV4(t *testing.T) {
	tenantKey := "2e99b3ab0b0f1654"
	calendarId := "feishu.cn_Y8iPaCjvuBeG6WEdJ326qg@group.calendar.feishu.cn"
	tenant := GetTenant(tenantKey)
	resp, err := tenant.UnsubscribeCalendarV4(calendarId)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestTenant_CreateCalendarEventV4(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2e99b3ab0b0f1654")
	t.Log(e)

	resp, err := tenant.CreateCalendarEventV4(calendarIdV4, vo.CreateCalendarEventV4Req{
		Summary:     "a",
		Description: "a",
		StartTime: vo.CalendarEventTime{
			Timestamp: "1617176660",
		},
		EndTime: vo.CalendarEventTime{
			Timestamp: "1617177660",
		},
		Vchat:           vo.Vchat{},
		Visibility:      "public",
		AttendeeAbility: "",
		FreeBusyStatus:  "",
		Location:        vo.Location{},
		Color:           0,
		Reminders:       nil,
		Recurrence:      "",
		Schemas:         nil,
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_DeleteCalendarEventV4(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2e99b3ab0b0f1654")
	t.Log(e)

	resp, err := tenant.DeleteCalendarEventV4(calendarIdV4, eventIdV4)

	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_GetCalendarEventV4(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2e99b3ab0b0f1654")
	t.Log(e)

	resp, err := tenant.GetCalendarEventV4(calendarIdV4, eventIdV4)

	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)

	resp1, err1 := tenant.GetCalendarEventListV4(calendarIdV4, eventIdV4, 50, nil, nil)

	log.Info(json.ToJsonIgnoreError(resp1), err1)
	assert.Equal(t, err1, nil)
	assert.Equal(t, resp1.Code, 0)
}

func TestTenant_UpdateCalendarEventV4(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2e99b3ab0b0f1654")
	t.Log(e)

	summary := "aaa"
	description := "des"
	resp, err := tenant.UpdateCalendarEventV4(calendarIdV4, eventIdV4, vo.UpdateCalendarEventV4Req{
		Summary:     &summary,
		Description: &description,
	})

	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_AddCalendarEventAttendeesV4(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2e99b3ab0b0f1654")
	t.Log(e)

	resp, err := tenant.AddCalendarEventAttendeesV4(calendarIdV4, eventIdV4, "open_id", vo.AddCalendarEventAttendeesV4Req{
		Attendees: []vo.AttendeesV4{
			vo.AttendeesV4{
				Type:       "user",
				IsOptional: false,
				UserId:     "ou_3018c5fbbb152f1b1b8f4c1c547df2b9",
			},
		},
	})

	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_GetCalendarEventAttendeesV4(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2e99b3ab0b0f1654")
	t.Log(e)

	resp, err := tenant.GetCalendarEventAttendeesV4(calendarIdV4, eventIdV4, 50, nil, "open_id")

	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_DeleteCalendarEventAttendeesV4(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2e99b3ab0b0f1654")
	t.Log(e)

	resp, err := tenant.DeleteCalendarEventAttendeesV4(calendarIdV4, eventIdV4, "open_id", vo.DeleteCalendarEventAttendeesV4Req{
		AttendeeIds: []string{"ou_3018c5fbbb152f1b1b8f4c1c547df2b9"},
	})

	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}
