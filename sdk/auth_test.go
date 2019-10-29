package sdk

import (
	"github.com/polaris-team/feishu-sdk-golang/core/consts"
	"github.com/polaris-team/feishu-sdk-golang/core/util/json"
	"github.com/polaris-team/feishu-sdk-golang/core/util/log"
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