package vo

type GetScopesResp struct {
	CommonVo
	Data GetScopesRespData `json:"data"`
}

type GetScopesRespData struct {
	Scopes []GetScopesRespDataScopes `json:"scopes"`
}

type GetScopesRespDataScopes struct {
	ScopeName string `json:"scope_name"`
	GrantStatus int `json:"grant_status"`
}

type ApplyScopesResp struct {
	CommonVo
	Data ApplyScopesRespData `json:"data"`
}

type ApplyScopesRespData struct {

}
