package vo

type AppAccessTokenInternalRespVo struct {
	CommonVo
	TenantAccessToken string `json:"tenant_access_token"`
	AppAccessToken string `json:"app_access_token"`
	Expire int64 `json:"expire"`
}

type AppAccessTokenRespVo struct {
	CommonVo
	AppAccessToken string `json:"app_access_token"`
	Expire int64 `json:"expire"`
}

type TenantAccessTokenRespVo struct {
	CommonVo
	TenantAccessToken string `json:"tenant_access_token"`
	Expire int64 `json:"expire"`
}

type OAuth2AccessTokenReqVo struct {
	AppId string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	AppAccessToken string `json:"app_access_token"`
	GrantType string `json:"grant_type"`
	Code string `json:"code"`
	RefreshToken string `json:"refresh_token"`
}

type OAuth2AccessTokenRespVo struct {
	AccessToken string `json:"access_token"`
	AvatarUrl string `json:"avatar_url"`
	ExpiresIn int64 `json:"expires_in"`
	Name string `json:"name"`
	EnName string `json:"en_name"`
	OpenId string `json:"open_id"`
	TenantKey string `json:"tenant_key"`
	RefreshToken string `json:"refresh_token"`
	TokenType string `json:"token_type"`
}

type OAuth2UserInfoRespVo struct {
	AvatarUrl string `json:"AvatarUrl"`
	Name string `json:"Name"`
	Email string `json:"Email"`
	Status int `json:"Status"`
	EmployeeId string `json:"EmployeeId"`
	Mobile string `json:"Mobile"`
}

type TokenLoginValidateResp struct {
	Data TokenLoginValidateRespData `json:"data"`

	CommonVo
}

type TokenLoginValidateRespData struct {
	Uid string `json:"uid"`
	OpenId string `json:"open_id"`
	UnionId string `json:"union_id"`
	SessionKey string `json:"session_key"`
	TenantKey string `json:"tenant_key"`
	EmployeeId string `json:"employee_id"`
	TokenType string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn int64 `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}