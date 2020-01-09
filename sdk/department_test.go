package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"

	"testing"
)

func TestTenant_GetScope(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetScope()
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)


	//openIds := resp.Data.AuthedOpenIds
	openIds := []string{"ou_b7d861c94cea4d316a6bfc5e8421994c"}
	resp1, err := tenant.GetUserBatchGet(nil, openIds)
	log.Info(json.ToJsonIgnoreError(resp1), err)
	//assert.Equal(t, err, nil)
	//	//assert.Equal(t, resp1.Code, 0)

	t.Log(json.ToJsonIgnoreError(resp1))

	resp2, err := tenant.GetUserBatchGetV2(nil, openIds)
	log.Info(json.ToJsonIgnoreError(resp2), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp2.Code, 0)

	t.Log(json.ToJsonIgnoreError(resp2))
}

func TestTenant_GetScopeV2(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetScopeV2()
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestGetDepartmentSimpleList(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetDepartmentSimpleList("0", 0, 100, true)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_GetDepartmentSimpleListV2(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetDepartmentSimpleListV2("0", "", 100, true)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestTenant_GetDepartmentInfo(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetDepartmentInfo("0")
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_GetDepartmentUserList(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetDepartmentUserList("0", 0, 100, true)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_GetDepartmentUserV2List(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetDepartmentUserListV2("0", "", 100, true)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_GetDepartmentUserDetailList(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetDepartmentSimpleList("0", 0, 100, true)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)

	for _, dep := range resp.Data.DepartmentInfos {
		resp1, err := tenant.GetDepartmentUserDetailList(dep.Id, 0, 100, true)
		log.Info(json.ToJsonIgnoreError(resp1), err)
		assert.Equal(t, err, nil)
		assert.Equal(t, resp1.Code, 0)
	}

}


func TestTenant_GetDepartmentUserDetailListV2(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetDepartmentSimpleList("0", 0, 100, true)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)

	for _, dep := range resp.Data.DepartmentInfos {
		resp1, err := tenant.GetDepartmentUserDetailListV2(dep.Id, "", 100, true)
		log.Info(json.ToJsonIgnoreError(resp1), err)
		assert.Equal(t, err, nil)
		assert.Equal(t, resp1.Code, 0)
		t.Log(json.ToJsonIgnoreError(resp1))
	}

}