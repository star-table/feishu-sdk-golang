package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//获取 app_access_token（企业自建应用）https://open.feishu.cn/document/ukTMukTMukTM/uADN14CM0UjLwQTN
func GetAppAccessTokenInternal(appId, appSecret string) (*vo.AppAccessTokenInternalRespVo, error){
	reqBody := map[string]interface{}{
		"app_id": appId,
		"app_secret": appSecret,
	}
	respBody, err := http.Post(consts.ApiAppAccessTokenInternal, nil, json.ToJsonIgnoreError(reqBody))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.AppAccessTokenInternalRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取 app_access_token（应用商店应用）https://open.feishu.cn/document/ukTMukTMukTM/uEjNz4SM2MjLxYzM
func GetAppAccessToken(appId, appSecret, appTicket string) (*vo.AppAccessTokenRespVo, error){
	reqBody := map[string]interface{}{
		"app_id": appId,
		"app_secret": appSecret,
		"app_ticket": appTicket,
	}
	respBody, err := http.Post(consts.ApiAppAccessToken, nil, json.ToJsonIgnoreError(reqBody))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.AppAccessTokenRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取 tenant_access_token（企业自建应用）https://open.feishu.cn/document/ukTMukTMukTM/uIjNz4iM2MjLyYzM
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

//获取 tenant_access_token（应用商店应用）https://open.feishu.cn/document/ukTMukTMukTM/uMjNz4yM2MjLzYzM
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

//重新推送 app_ticket https://open.feishu.cn/document/ukTMukTMukTM/uQjNz4CN2MjL0YzM
func AppTicketResend(appId, appSecret string) (*vo.CommonVo, error){
	reqBody := map[string]interface{}{
		"app_id": appId,
		"app_secret": appSecret,
	}
	respBody, err := http.Post(consts.ApiAppTicketResend, nil, json.ToJsonIgnoreError(reqBody))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取登录用户身份 https://open.feishu.cn/document/ukTMukTMukTM/ukTNz4SO1MjL5UzM
func GetOauth2AccessToken(req vo.OAuth2AccessTokenReqVo) (*vo.OAuth2AccessTokenRespVo, error){
	respBody, err := http.Post(consts.ApiOAuth2AccessToken, nil, json.ToJsonIgnoreError(req))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.OAuth2AccessTokenRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取用户信息 https://open.feishu.cn/document/ukTMukTMukTM/uAjNz4CM2MjLwYzM
func GetOAuth2UserInfo(userAccessToken string) (*vo.OAuth2UserInfoRespVo, error){
	respBody, err := http.Get(consts.ApiOAuth2AccessToken, nil, http.BuildTokenHeaderOptions(userAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.OAuth2UserInfoRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
