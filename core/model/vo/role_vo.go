package vo

type RoleListResp struct {
	CommonVo
	Data RoleListRespData `json:"data"`
}

type RoleListRespData struct {
	RoleList []Role `json:"role_list"`
}

type Role struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type RoleMemberListResp struct {
	CommonVo
	Data RoleMemberListRespData `json:"data"`
}

type RoleMemberListRespData struct {
	HasMore bool `json:"has_more"`
	PageToken string `json:"page_token"`
	UserList []UserRestInfoVo `json:"user_list"`
}