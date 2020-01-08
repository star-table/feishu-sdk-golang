package vo

type IsUserAdminResp struct {
	CommonVo
	Data IsUserAdminRespData `json:"data"`
}

type IsUserAdminRespData struct {
	IsAppAdmin bool `json:"is_app_admin"`
}
