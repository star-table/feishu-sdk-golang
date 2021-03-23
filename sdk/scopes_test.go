package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
	"testing"
)

func TestTenant_GetScopes(t *testing.T) {
	tenant := GetTenant("2c0c7abea54f9758")
	resp, err := tenant.GetScopes()
	if err != nil {
		t.Error(err)
		return
	}
	log.Info(json.ToJsonIgnoreError(resp))
}

func TestTenant_ApplyScopes(t *testing.T) {
	tenant := GetTenant("2c0c7abea54f9758")
	resp, err := tenant.ApplyScopes()
	if err != nil {
		t.Error(err)
		return
	}
	log.Info(json.ToJsonIgnoreError(resp))
}

// 测试时，公用方法，获取 tenant。
func GetTenant(outOrgId string) *Tenant {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	if e != nil {
		log.Error(e)
		return nil
	}
	tenant, e := BuildTenant(app.AppAccessToken, outOrgId)
	if e != nil {
		log.Error(e)
		return nil
	}
	return tenant
}
