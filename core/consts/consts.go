package consts

//API Host const， v3
const (
	//获取 app_access_token（企业自建应用）
	ApiAppAccessTokenInternal = "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal/"
	//获取 app_access_token（应用商店应用）
	ApiAppAccessToken = "https://open.feishu.cn/open-apis/auth/v3/app_access_token/"
	//获取 tenant_access_token（应用商店应用）
	ApiTenantAccessToken = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/"
	//获取 tenant_access_token（企业自建应用）
	ApiTenantAccessTokenInternal = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/"
	//重新推送 app_ticket
	ApiAppTicketResend = "https://open.feishu.cn/open-apis/auth/v3/app_ticket/resend/"
	//获取登录用户身份
	ApiOAuth2AccessToken = "https://open.feishu.cn/connect/qrconnect/oauth2/access_token/"
	//code2session
	ApiTokenLoginValidate = "https://open.feishu.cn/open-apis/mina/v2/tokenLoginValidate"

	//////////////////部门和用户
	//获取通讯录授权范围
	ApiScope = "https://open.feishu.cn/open-apis/contact/v1/scope/get"
	//获取部门列表
	ApiDepartmentSimpleList = "https://open.feishu.cn/open-apis/contact/v1/department/simple/list"
	//获取部门详情
	ApiDepartmentInfoGet = "https://open.feishu.cn/open-apis/contact/v1/department/info/get"
	//获取部门用户列表
	ApiDepartmentUserList = "https://open.feishu.cn/open-apis/contact/v1/department/user/list"
	//获取用户详情
	ApiDepartmentUserDetailList = "https://open.feishu.cn/open-apis/contact/v1/department/user/detail/list"
	//批量获取用户信息
	ApiUserBatchGet = "https://open.feishu.cn/open-apis/contact/v1/user/batch_get"
	//批量获取用户信息v2
	ApiUserBatchGetV2 = "https://open.feishu.cn/open-apis/contact/v2/user/batch_get"

	//////////////////机器人发送消息
	//机器人发送消息
	ApiRobotSendMessage = "https://open.feishu.cn/open-apis/message/v4/send/"
	//机器人批量发送消息
	ApiRobotSendBatchMessage = "https://open.feishu.cn/open-apis/message/v4/batch_send/"

	//////////////////角色
	//获取角色列表
	ApiRoleList = "https://open.feishu.cn/open-apis/contact/v2/role/list"
	//获取角色成员列表
	ApiRoleMemberList = "https://open.feishu.cn/open-apis/contact/v2/role/members"

	/////////////////////////////////////////////////////////
	//创建日历
	ApiCalendarCreate = "https://open.feishu.cn/open-apis/calendar/v3/calendars"
	//获取日历
	ApiCalendarGet = "https://open.feishu.cn/open-apis/calendar/v3/calendar_list/%s"
	//获取日历列表
	ApiCalendarListGet = "https://open.feishu.cn/open-apis/calendar/v3/calendar_list"
	//更新日历
	ApiCalendarUpdate = "https://open.feishu.cn/open-apis/calendar/v3/calendars/%s"
	//创建日程
	ApiCalendarEventCreate = "https://open.feishu.cn/open-apis/calendar/v3/calendars/%s/events"
	//删除日程
	ApiCalendarEventDelete = "https://open.feishu.cn/open-apis/calendar/v3/calendars/%s/events/%s"
	//邀请/移除日程参与者
	ApiCalendarEventAttendeesUpdate = "https://open.feishu.cn/open-apis/calendar/v3/calendars/%s/events/%s/attendees"
	//获取访问控制列表
	ApiCalendarAttendeesGet = "https://open.feishu.cn/open-apis/calendar/v3/calendars/%s/acl"
	//删除访问空值
	ApiCalendarAttendeesDelete = "https://open.feishu.cn/open-apis/calendar/v3/calendars/%s/acl/%s"

	//搜索用户
	ApiSearchUser = "https://open.feishu.cn/open-apis/search/v1/user"

	//检验应用管理员
	ApiIsUserAdmin = "https://open.feishu.cn/open-apis/application/v3/is_user_admin"
)

//Other Const
const (
	TestAppId     = "cli_9d5e49aae9ae9101"
	TestAppSecret = "HDzPYfWmf8rmhsF2hHSvmhTffojOYCdI"
	TestTicket    = "03050a75e8a7914bd43bf7d705025c4ff53bce6f"
)

const (
	AccessRoleReader         = "reader"
	AccessRoleFreeBusyReader = "free_busy_reader"
	AccessRoleWriter         = "writer"
	AccessRoleOwner          = "owner"
)

const (
	ActionInvite = "invite"
	ActionRemove = "remove"
)
