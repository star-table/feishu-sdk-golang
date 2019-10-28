package sdk

import (
	"github.com/polaris-team/feishu-sdk-golang/core/consts"
	"github.com/polaris-team/feishu-sdk-golang/core/model/vo"
	"github.com/polaris-team/feishu-sdk-golang/core/util/http"
	"github.com/polaris-team/feishu-sdk-golang/core/util/json"
	"github.com/polaris-team/feishu-sdk-golang/core/util/log"
)

func GetTenantAccessTokenInternal(appId string, appSecret string) (*vo.TenantAccessTokenRespVo, error){
	reqBody := map[string]interface{}{
		"app_id": appId,
		"app_secret": appSecret,
	}
	respBody, err := http.Post(consts.ApiTenantAccessTokenInternal, nil, json.ToJsonIgnoreError(reqBody))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.TenantAccessTokenRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

func GetTenantAccessToken(appAccessToken string, tenantKey string) (*vo.TenantAccessTokenRespVo, error){
	reqBody := map[string]interface{}{
		"app_access_token": appAccessToken,
		"tenant_key": tenantKey,
	}
	respBody, err := http.Post(consts.ApiTenantAccessToken, nil, json.ToJsonIgnoreError(reqBody))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.TenantAccessTokenRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}