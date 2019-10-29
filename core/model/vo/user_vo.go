package vo

type GetDepartmentUserListRespVo struct {
	CommonVo
	Data *GetDepartmentUserListRespVoData `json:"data"`
}

type GetDepartmentUserListRespVoData struct {
	HasMore bool `json:"has_more"`
	UserList []UserRestInfoVo `json:"user_list"`
}

type UserRestInfoVo struct {
	EmployeeId string `json:"employee_id"`
	OpenId string `json:"open_id"`
	Name string `json:"name"`
	EmployeeNo string `json:"employee_no"`
}

type GetDepartmentSimpleListRespVo struct {
	CommonVo
	Data *GetDepartmentSimpleListRespVoData `json:"data"`
}

type GetDepartmentSimpleListRespVoData struct {
	HasMore bool `json:"has_more"`
	DepartmentInfos []DepartmentRestInfoVo `json:"department_infos"`
}

type DepartmentRestInfoVo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ParentId string `json:"parent_id"`
}

type GetDepartmentInfoRespVo struct {
	CommonVo
	Data *GetDepartmentInfoRespVoData `json:"data"`
}

type GetDepartmentInfoRespVoData struct {
	DepartmentInfo *DepartmentDetailInfoVo `json:"department_info"`
}

type DepartmentDetailInfoVo struct {
	Id string `json:"id"`
	LeaderEmployeeId string `json:"leader_employee_id"`
	LeaderOpenId string `json:"leader_open_id"`
	ChatId string `json:"chat_id"`
	MemberCount int `json:"member_count"`
	Name string `json:"name"`
	ParentId string `json:"parent_id"`
	Status int `json:"status"`
}