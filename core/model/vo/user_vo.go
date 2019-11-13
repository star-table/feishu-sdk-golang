package vo

type GetDepartmentUserListRespVo struct {
	CommonVo
	Data *GetDepartmentUserListRespVoData `json:"data"`
}

type GetDepartmentUserDetailListRespVo struct {
	CommonVo
	Data *GetDepartmentUserDetailListRespVoData `json:"data"`
}

type GetDepartmentUserDetailListRespVoData struct {
	HasMore bool `json:"has_more"`
	UserInfos []UserDetailInfo `json:"user_infos"`
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

type GetUserBatchGetRespVo struct {
	CommonVo
	Data *GetUserBatchGetRespVoData `json:"data"`
}

type GetUserBatchGetRespVoData struct {
    UserInfos []UserDetailInfo `json:"user_infos"`
}

type UserDetailInfo struct {
	Name string `json:"name"`
	NamePy string `json:"name_py"`
	EnName string `json:"en_name"`
	EmployeeId string `json:"employee_id"`
	EmployeeNo string `json:"employee_no"`
	OpenId string `json:"open_id"`
	Status string `json:"status"`
	EmployeeType string `json:"employee_type"`
	Avatar71 string `json:"avatar_71"`
	Avatar240 string `json:"avatar_240"`
	Avatar640 string `json:"avatar_640"`
	AvatarUrl string `json:"avatar_url"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	Description string `json:"description"`
	Country string `json:"country"`
	City string `json:"city"`
	WorkStation string `json:"work_station"`
	IsTenantManager bool `json:"is_tenant_manager"`
	JoinTime int64 `json:"join_time"`
	UpdateTime int64 `json:"update_time"`
	LeaderEmployeeId string `json:"leader_employee_id"`
	LeaderOpenId string `json:"leader_open_id"`
	Departments []string `json:"departments"`
	CustomAttr map[string]Entry `json:"custom_attr"`
}

type Entry struct {
	Value string `json:"value"`
}
