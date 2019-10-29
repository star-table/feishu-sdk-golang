package sdk

import (
	"github.com/polaris-team/feishu-sdk-golang/core/consts"
	"github.com/polaris-team/feishu-sdk-golang/core/util/json"
	"github.com/polaris-team/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"
	"testing"
)

func TestGetDepartmentSimpleList(t *testing.T) {
	tenant, _ := BuildTenantInternal(consts.TestAppId, consts.TestAppSecret)

	resp, err := tenant.GetDepartmentSimpleList( "0", 0,100, true)
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}