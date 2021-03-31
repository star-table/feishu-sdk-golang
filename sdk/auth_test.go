package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"
	"testing"
)

func TestGetTenantAccessTokenInternal(t *testing.T) {
	resp, err := GetTenantAccessTokenInternal(consts.TestAppId, consts.TestAppSecret)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestGetAppAccessTokenInternal(t *testing.T) {
	resp, err := GetAppAccessToken(consts.TestAppId, consts.TestAppSecret, "")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestAppTicketResend(t *testing.T) {
	resp, err := AppTicketResend("cli_9d5e49aae9ae9101", "HDzPYfWmf8rmhsF2hHSvmhTffojOYCdI")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
	//
	//resp, err = AppTicketResend("cli_9d40f5bf08f95108", "Apx5vdWeIxVzDBQ6ARte6grZgOCgbhgP")
	//log.Info(json.ToJsonIgnoreError(resp), err)
	//assert.Equal(t, err, nil)
	//assert.Equal(t, resp.Code, 0)
	//
	//resp, err = AppTicketResend("cli_9d511af667bd1101", "GESwFvoks5xQEr1t7dC4uhKR3cm1bg3S")
	//log.Info(json.ToJsonIgnoreError(resp), err)
	//assert.Equal(t, err, nil)
	//assert.Equal(t, resp.Code, 0)

	//resp, err = AppTicketResend("cli_9d3ae99f30eb9102", "O7Z43hPSFF1PHhRk8gVBOd3l36g8L5UG")
	//log.Info(json.ToJsonIgnoreError(resp), err)
	//assert.Equal(t, err, nil)
	//assert.Equal(t, resp.Code, 0)
}

func TestTokenLoginValidate(t *testing.T) {
	resp, err := TokenLoginValidate("a-566311d2cf88d054a4fcfc23233a448f2fccba11", "1c3a78be18ac815a")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestRefreshUserAccessToken(t *testing.T) {
	t.Log(RefreshUserAccessToken(consts.TestAppId, consts.TestAppSecret, consts.TestTicket, "aaaaa"))
}

func TestGetOauth2AccessToken(t *testing.T) {
	t.Log(GetOauth2AccessToken(vo.OAuth2AccessTokenReqVo{
		AppId:     consts.TestAppId,
		AppSecret: consts.TestAppSecret,
		//AppAccessToken: "a-4d0bea08bf46580b9cc9bf8edc0f5736fdfa7673",
		GrantType: "authorization_code",
		Code:      "Ea3OYm95NV8Qn4IB2HxF5g",
	}))

}

func TestGetOAuth2UserInfo(t *testing.T) {
	t.Log(GetOAuth2UserInfo("u-e3PphRj4a4LiklDaeWjnre"))
}
