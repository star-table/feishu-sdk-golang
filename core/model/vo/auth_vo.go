package vo

type TenantAccessTokenRespVo struct {
	CommonVo
	TenantAccessToken string `json:"tenant_access_token"`
	Expire int64 `json:"expire"`
}

