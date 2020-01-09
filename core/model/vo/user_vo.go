package vo

type GetDepartmentUserListRespVo struct {
	CommonVo
	Data *GetDepartmentUserListRespVoData `json:"data"`
}

type GetDepartmentUserListV2RespVo struct {
	CommonVo
	Data *GetDepartmentUserListV2RespVoData `json:"data"`
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

type GetDepartmentUserListV2RespVoData struct {
	HasMore bool `json:"has_more"`
	PageToken string `json:"page_token"`
	UserList []UserRestInfoVo `json:"users"`
}

type UserRestInfoVo struct {
	EmployeeId string `json:"employee_id"`
	OpenId string `json:"open_id"`
	UserId string `json:"user_id"`
	Name string `json:"name"`
	EmployeeNo string `json:"employee_no"`
}

type GetDepartmentSimpleListRespVo struct {
	CommonVo
	Data *GetDepartmentSimpleListRespVoData `json:"data"`
}

type GetDepartmentSimpleListV2RespVo struct {
	CommonVo
	Data *GetDepartmentSimpleListV2RespVoData `json:"data"`
}

type GetScopeRespVo struct {
	CommonVo
	Data *GetScopeRespData `json:"data"`
}

type GetScopeRespV2Vo struct {
	CommonVo
	Data *GetScopeRespV2Data `json:"data"`
}

type GetScopeRespData struct {
	AuthedDepartments []string `json:"authed_departments"`
	AuthedEmployeeIds []string `json:"authed_employee_ids"`
	AuthedOpenIds []string `json:"authed_open_ids"`
}

type GetScopeRespV2Data struct {
	AuthedDepartments []string `json:"authed_departments"`
	AuthedUsers []ScopeUser `json:"authed_users"`
}

type ScopeUser struct {
	OpenId string `json:"open_id"`
	UserId string `json:"user_id"`
}

type GetDepartmentSimpleListRespVoData struct {
	HasMore bool `json:"has_more"`
	PageToken string `json:"page_token"`
	DepartmentInfos []DepartmentRestInfoVo `json:"department_infos"`
}

type GetDepartmentSimpleListV2RespVoData struct {
	HasMore bool `json:"has_more"`
	PageToken string `json:"page_token"`
	DepartmentInfos []DepartmentRestInfoVo `json:"departments"`
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

type GetDepartmentInfoBatchRespVo struct {
	CommonVo
	Data GetDepartmentInfoBatchRespData `json:"data"`
}

type GetDepartmentInfoBatchRespData struct {
	Departments []DepartmentDetailInfo `json:"departments"`
	Errors []GetUserBatchError `json:"errors"`
}

type DepartmentDetailInfo struct {
	ChatId string `json:"chat_id"`
	HasChild bool `json:"has_child"`
	Id string `json:"id"`
	Leader Leader `json:"leader"`
	MemberCount int `json:"member_count"`
	Name string `json:"name"`
	ParentId string `json:"parent_id"`
	Status int `json:"status"`
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
	Status int `json:"status"`
	EmployeeType int `json:"employee_type"`
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

type UserDetailInfoV2 struct {
	Name string `json:"name"`
	NamePy string `json:"name_py"`
	EnName string `json:"en_name"`
	EmployeeId string `json:"employee_id"`
	EmployeeNo string `json:"employee_no"`
	OpenId string `json:"open_id"`
	Status UserStatus `json:"status"`
	EmployeeType int `json:"employee_type"`
	Avatar UserAvatar `json:"avatar"`
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
	Leader Leader `json:"leader"`
	Departments []string `json:"departments"`
	Positions []Position `json:"positions"`
	Orders []Order `json:"orders"`
	CustomAttrs []CustomAttr `json:"custom_attrs"`
}

type CustomAttr struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Value CustomAttrValue `json:"value"`
}

type CustomAttrValue struct {
	Url string `json:"url"`
	PcUrl string `json:"pc_url"`
}

type Order struct {
	DepartmentId string `json:"department_id"`
	UserOrder int `json:"user_order"`
	DepartmentOrder int `json:"department_order"`
}

type Position struct {
	PositionCode string `json:"position_code"`
	PositionName string `json:"position_name"`
	DepartmentId string `json:"department_id"`
	Leader Leader `json:"leader"`
	IsMajor bool `json:"is_major"`
}

type Leader struct {
	OpenId string `json:"open_id"`
	UserId string `json:"user_id"`
	PositionCode string `json:"position_code"`
}

type UserStatus struct {
	IsFrozen bool `json:"is_frozen"`
	IsResigned bool `json:"is_resigned"`
	IsActivated bool `json:"is_activated"`
}

type Entry struct {
	Value string `json:"value"`
}

type SearchUserResp struct {
	CommonVo
	Data *SearchUserRespData `json:"data"`
}

type SearchUserRespData struct {
	HasMore bool `json:"has_more"`
	PageToken string `json:"page_token"`
	Users []SearchUserInfo `json:"users"`
}

type SearchUserInfo struct {
	Avatar UserAvatar `json:"avatar"`
	DepartmentIds []string `json:"department_ids"`
	Name string `json:"name"`
	OpenId string `json:"open_id"`
	UserId string `json:"user_id"`
}

type UserAvatar struct {
	Avatar72 string `json:"avatar_72"`
	Avatar240 string `json:"avatar_240"`
	Avatar640 string `json:"avatar_640"`
	AvatarOrigin string `json:"avatar_origin"`
}

type GetUserBatchGetV2Resp struct {
	CommonVo
	Data GetUserBatchGetV2RespData `json:"data"`
}

type GetUserBatchGetV2RespData struct {
	HasMore bool `json:"has_more"`
	PageToken string `json:"page_token"`
	Users []UserDetailInfoV2 `json:"users"`
	Errors []GetUserBatchError `json:"errors"`
}

type GetUserBatchError struct {
	Id string `json:"id"`
	Code int `json:"code"`
	Msg string `json:"msg"`
}