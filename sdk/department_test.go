package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"
	"testing"
)

func TestGetDepartmentSimpleList(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, "751a7ac21dfc5391c65dcf8bfc2cb8504c61903b")
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetDepartmentSimpleList( "0", 0,100, true)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}


func TestTenant_GetDepartmentInfo(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, "751a7ac21dfc5391c65dcf8bfc2cb8504c61903b")
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetDepartmentInfo("0")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}
