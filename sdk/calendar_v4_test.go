package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"
	"testing"
)

func TestTenant_CreateCalendarV4(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2e99b3ab0b0f1654")
	t.Log(e)

	resp, err := tenant.CreateCalendarV4(vo.CreateCalendarV4Req{
		Summary:     "test日历002",
		Description: "测试使用2",
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
