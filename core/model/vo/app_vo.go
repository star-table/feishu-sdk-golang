package vo

type IsUserAdminResp struct {
	CommonVo
	Data IsUserAdminRespData `json:"data"`
}

type IsUserAdminRespData struct {
	IsAppAdmin bool `json:"is_app_admin"`
}

type AdminUserListResp struct {
	CommonVo
	Data AdminUserListRespData `json:"data"`
}

type AdminUserListRespData struct {
	UserList []AdminUserData `json:"user_list"`
}

type AdminUserData struct {
	OpenId  string `json:"open_id"`
	UserId  string `json:"user_id"`
	UnionId string `json:"union_id"`
}
