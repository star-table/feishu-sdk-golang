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
	//刷新access_token
	ApiRefreshAccessToken = "https://open.feishu.cn/open-apis/authen/v1/refresh_access_token"

	//////////////////部门和用户
	//获取通讯录授权范围
	ApiScope = "https://open.feishu.cn/open-apis/contact/v1/scope/get"
	//获取通讯录授权范围v2
	ApiScopeV2 = "https://open.feishu.cn/open-apis/contact/v2/scope/get"

	//获取部门列表
	ApiDepartmentSimpleList = "https://open.feishu.cn/open-apis/contact/v1/department/simple/list"
	//获取部门列表 v2
	ApiDepartmentSimpleListV2 = "https://open.feishu.cn/open-apis/contact/v2/department/simple/list"

	//获取部门详情
	ApiDepartmentInfoGet = "https://open.feishu.cn/open-apis/contact/v1/department/info/get"
	//批量获取部门详情
	ApiDepartmentInfoBatchGet = "https://open.feishu.cn/open-apis/contact/v2/department/detail/batch_get"

	//获取部门用户列表
	ApiDepartmentUserList = "https://open.feishu.cn/open-apis/contact/v1/department/user/list"
	//获取部门用户列表v2
	ApiDepartmentUserListV2 = "https://open.feishu.cn/open-apis/contact/v2/department/user/list"
	//获取用户详情
	ApiDepartmentUserDetailList = "https://open.feishu.cn/open-apis/contact/v1/department/user/detail/list"
	//获取用户详情v2
	ApiDepartmentUserDetailListV2 = "https://open.feishu.cn/open-apis/contact/v2/department/user/detail/list"

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

	//搜索用户F
	ApiSearchUser = "https://open.feishu.cn/open-apis/search/v1/user"

	//检验应用管理员
	ApiIsUserAdmin = "https://open.feishu.cn/open-apis/application/v3/is_user_admin"

	////////用户群组
	//获取用户所在的群列表
	ApiUserGroupLIst = "https://open.feishu.cn/open-apis/user/v4/group_list"
	//获取群成员列表
	ApiChatMembers = "https://open.feishu.cn/open-apis/chat/v4/members"
	//搜索用户所在的群列表
	ApiChatSearch = "https://open.feishu.cn/open-apis/chat/v4/search"

	////////群信息和群管理
	//创建群
	ApiCreateChat = "https://open.feishu.cn/open-apis/chat/v4/create/"
	//获取群列表
	ApiChatList = "https://open.feishu.cn/open-apis/chat/v4/list"
	//获取群信息
	ApiChatInfo = "https://open.feishu.cn/open-apis/chat/v4/info"
	//更新群信息
	ApiUpdateChat = "https://open.feishu.cn/open-apis/chat/v4/update/"
	//拉用户进群
	ApiAddChatUser = "https://open.feishu.cn/open-apis/chat/v4/chatter/add/"
	//移除用户出群
	ApiRemoveChatUser = "https://open.feishu.cn/open-apis/chat/v4/chatter/delete/"
	//解散群
	ApiDisbandChat = "https://open.feishu.cn/open-apis/chat/v4/disband"
	//拉机器人进群
	ApiAddBot = "https://open.feishu.cn/open-apis/bot/v4/add"

	/////////订单
	//查询用户是否在应用开通范围
	ApiCheckUser    = "https://open.feishu.cn/open-apis/pay/v1/paid_scope/check_user"
	ApiGetOrderList = "https://open.feishu.cn/open-apis/pay/v1/order/list"
	ApiGetOrderInfo = "https://open.feishu.cn/open-apis/pay/v1/order/get"
)

//Other Const
const (
	TestAppId     = "cli_9d5e49aae9ae9101"
	TestAppSecret = "HDzPYfWmf8rmhsF2hHSvmhTffojOYCdI"
	TestTicket    = "e03716ea22bd01d6be12a367551151cfa71ef993"
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
