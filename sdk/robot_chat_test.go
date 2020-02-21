package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"
	"testing"
)

func TestTenant_CreateChat(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.CreateChat(vo.CreateChatReqVo{
		Name:        "test",
		Description: "hhhh",
		OpenIds: []string{
			"ou_e1b43c426e884c586d52751853896688",
			"ou_22f5529e4a1bc5c3a1271ca609b24ed5",
		},
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_ChatList(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.ChatList(0, "")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_ChatInfo(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.ChatInfo("")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_UpdateChat(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.UpdateChat(vo.UpdateChatReqVo{
		ChatId:      "",
		OwnerUserId: "",
		OwnerOpenId: "",
		Name:        "",
		I18nNames:   nil,
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_AddChatUser(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.AddChatUser(vo.UpdateChatMemberReqVo{
		ChatId:  "",
		UserIds: nil,
		OpenIds: nil,
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)

	resp1, err1 := tenant.RemoveChatUser(vo.UpdateChatMemberReqVo{
		ChatId:  "",
		UserIds: nil,
		OpenIds: nil,
	})
	log.Info(json.ToJsonIgnoreError(resp1), err1)
	assert.Equal(t, err1, nil)
	assert.Equal(t, resp1.Code, 0)
}

func TestTenant_DisbandChat(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.DisbandChat(vo.UpdateChatData{ChatId: ""})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}
