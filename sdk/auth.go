package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

// 获取 app_access_token（企业自建应用）https://open.feishu.cn/document/ukTMukTMukTM/uADN14CM0UjLwQTN
func GetAppAccessTokenInternal(appId, appSecret string) (*vo.AppAccessTokenInternalRespVo, error) {
	reqBody := map[string]interface{}{
		"app_id":     appId,
		"app_secret": appSecret,
	}
	respBody, err := http.Post(consts.ApiAppAccessTokenInternal, nil, json.ToJsonIgnoreError(reqBody))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.AppAccessTokenInternalRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// 获取 app_access_token（应用商店应用）https://open.feishu.cn/document/ukTMukTMukTM/uEjNz4SM2MjLxYzM
func GetAppAccessToken(appId, appSecret, appTicket string) (*vo.AppAccessTokenRespVo, error) {
	reqBody := map[string]interface{}{
		"app_id":     appId,
		"app_secret": appSecret,
		"app_ticket": appTicket,
	}
	respBody, err := http.Post(consts.ApiAppAccessToken, nil, json.ToJsonIgnoreError(reqBody))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.AppAccessTokenRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// 获取 tenant_access_token（企业自建应用）https://open.feishu.cn/document/ukTMukTMukTM/uIjNz4iM2MjLyYzM
func GetTenantAccessTokenInternal(appId string, appSecret string) (*vo.TenantAccessTokenRespVo, error) {
	reqBody := map[string]interface{}{
		"app_id":     appId,
		"app_secret": appSecret,
	}
	respBody, err := http.Post(consts.ApiTenantAccessTokenInternal, nil, json.ToJsonIgnoreError(reqBody))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.TenantAccessTokenRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// 获取 tenant_access_token（应用商店应用）https://open.feishu.cn/document/ukTMukTMukTM/uMjNz4yM2MjLzYzM
func GetTenantAccessToken(appAccessToken string, tenantKey string) (*vo.TenantAccessTokenRespVo, error) {
	reqBody := map[string]interface{}{
		"app_access_token": appAccessToken,
		"tenant_key":       tenantKey,
	}
	respBody, err := http.Post(consts.ApiTenantAccessToken, nil, json.ToJsonIgnoreError(reqBody))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.TenantAccessTokenRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// 重新推送 app_ticket https://open.feishu.cn/document/ukTMukTMukTM/uQjNz4CN2MjL0YzM
func AppTicketResend(appId, appSecret string) (*vo.CommonVo, error) {
	reqBody := map[string]interface{}{
		"app_id":     appId,
		"app_secret": appSecret,
	}
	respBody, err := http.Post(consts.ApiAppTicketResend, nil, json.ToJsonIgnoreError(reqBody))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// 获取登录用户身份 https://open.feishu.cn/document/ukTMukTMukTM/ukTNz4SO1MjL5UzM
func GetOauth2AccessToken(req vo.OAuth2AccessTokenReqVo) (*vo.OAuth2AccessTokenRespVo, error) {
	respBody, err := http.Post(consts.ApiOAuth2AccessToken, nil, json.ToJsonIgnoreError(req))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.OAuth2AccessTokenRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// GetOauth2AccessTokenV1 https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/authen-v1/authen/access_token
func GetOauth2AccessTokenV1(req vo.OAuth2AccessTokenReqVo) (*vo.OAuth2AccessTokenRespVoV1, error) {
	respBody, err := http.Post(consts.ApiOAuth2AccessTokenV1, nil, json.ToJsonIgnoreError(req))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.OAuth2AccessTokenRespVoV1{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

func AuthenAccessToken(appAccessToken string, grantType string, code string) (*vo.AuthenAccessTokenResp, error) {
	req := map[string]string{
		"code":       code,
		"grant_type": grantType,
	}
	respBody, err := http.Post(consts.ApiAuthenAccessToken, nil, json.ToJsonIgnoreError(req), http.BuildTokenHeaderOptions(appAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.AuthenAccessTokenResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// 获取用户信息 https://open.feishu.cn/document/ukTMukTMukTM/uAjNz4CM2MjLwYzM
func GetOAuth2UserInfo(userAccessToken string) (*vo.OAuth2UserInfoRespVo, error) {
	respBody, err := http.Get(consts.ApiOAuth2GetUserInfoByAccessToken, nil, http.BuildTokenHeaderOptions(userAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.OAuth2UserInfoRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// code2session https://open.feishu.cn/document/ukTMukTMukTM/ukjM04SOyQjL5IDN?lang=zh-CN
func TokenLoginValidate(appAccessToken string, code string) (*vo.TokenLoginValidateResp, error) {
	req := map[string]string{
		"code": code,
	}
	respBody, err := http.Post(consts.ApiTokenLoginValidate, nil, json.ToJsonIgnoreError(req), http.BuildTokenHeaderOptions(appAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.TokenLoginValidateResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// 刷新access_token https://open.feishu.cn/document/ukTMukTMukTM/uQDO4UjL0gDO14CN4gTN
func RefreshUserAccessToken(appId, appSecret, appTicket, refreshToken string) (*vo.RefreshAccessTokenResp, error) {
	appAccessToken, err := GetAppAccessToken(appId, appSecret, appTicket)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	req := map[string]string{
		"app_access_token": appAccessToken.AppAccessToken,
		"grant_type":       "refresh_token",
		"refresh_token":    refreshToken,
	}
	respBody, err := http.Post(consts.ApiRefreshAccessToken, nil, json.ToJsonIgnoreError(req))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.RefreshAccessTokenResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
