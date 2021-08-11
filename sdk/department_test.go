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

func TestTenant_GetUserBatchGetV2(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetUserBatchGetV2(nil, []string{"ou_14e6b0d968a56bdbbdd511f896400e3c", "ou_fb0c3442f0ec403cdd6971234088d975", "ou_0e40d4035ceb951467beb62bb1f3e5ba", "ou_dcff2de6a28060eff9f0d9665d2d28be", "ou_e1b43c426e884c586d52751853896688"})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}

func TestTenant_GetScopeV2(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "138337e75b0b575d")
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
	tenant, e := BuildTenant(app.AppAccessToken, "2ee83c762dcf1657")
	t.Log(e)

	resp, err := tenant.GetDepartmentInfo("g9gged32cca9e9fe")
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

func TestTenant_GetDepartmentInfoBatch(t *testing.T) {
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

	deps := make([]string, 0)
	for _, dep := range resp.Data.DepartmentInfos {
		deps = append(deps, dep.Id)
	}
	resp1, err := tenant.GetDepartmentInfoBatch(deps)
	log.Info(json.ToJsonIgnoreError(resp1), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp1.Code, 0)
	t.Log(json.ToJsonIgnoreError(resp1))

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
	tenant, e := BuildTenant(app.AppAccessToken, "2ee83c762dcf1657")
	t.Log(e)

	resp1, err := tenant.GetDepartmentUserDetailListV2("od-938643fa45292b9906968f49857ebe18", "763bd1e74d05ea7c", 100, false)
	log.Info(json.ToJsonIgnoreError(resp1), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp1.Code, 0)
	t.Log(json.ToJsonIgnoreError(resp1))
	//resp, err := tenant.GetDepartmentSimpleList("0", 0, 100, true)
	//log.Info(json.ToJsonIgnoreError(resp), err)
	//assert.Equal(t, err, nil)
	//assert.Equal(t, resp.Code, 0)
	//
	//for _, dep := range resp.Data.DepartmentInfos {
	//	resp1, err := tenant.GetDepartmentUserDetailListV2(dep.Id, "", 100, true)
	//	log.Info(json.ToJsonIgnoreError(resp1), err)
	//	assert.Equal(t, err, nil)
	//	assert.Equal(t, resp1.Code, 0)
	//	t.Log(json.ToJsonIgnoreError(resp1))
	//}

}

func TestTenant_GetUsersV3(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	if e != nil {
		t.Fatal(e)
	}
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	if e != nil {
		t.Fatal(e)
	}
	t.Log(e)
	resp, e := tenant.GetUsersV3("", "", "", "", 100)
	if e != nil {
		t.Fatal(e)
	}
	for _, user := range resp.Data.Items {
		t.Log(json.ToJsonIgnoreError(user.DepartmentIds))
	}
	t.Log(json.ToJsonIgnoreError(resp))
}
