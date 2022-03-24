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
	//获取登录用户身份(通过user_access_token)
	ApiOAuth2GetUserInfoByAccessToken = "https://open.feishu.cn/connect/qrconnect/oauth2/user_info/"
	//code2session
	ApiTokenLoginValidate = "https://open.feishu.cn/open-apis/mina/v2/tokenLoginValidate"
	//刷新access_token
	ApiRefreshAccessToken = "https://open.feishu.cn/open-apis/authen/v1/refresh_access_token"
	//获取登录用户身份（新版）
	ApiAuthenAccessToken = "https://open.feishu.cn/open-apis/authen/v1/access_token"
	//JSAPI 临时授权凭证
	ApiJSApiTicket = "https://open.feishu.cn/open-apis/jssdk/ticket/get"

	//////////////////部门和用户
	//获取通讯录授权范围
	ApiScope = "https://open.feishu.cn/open-apis/contact/v1/scope/get"
	//获取通讯录授权范围v2
	ApiScopeV2 = "https://open.feishu.cn/open-apis/contact/v2/scope/get"
	//使用手机号或邮箱获取用户 ID
	ApiBatchGetUserId = "https://open.feishu.cn/open-apis/user/v1/batch_get_id"

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
	//获取用户列表v3
	ApiUserListV3 = "https://open.feishu.cn/open-apis/contact/v3/users"

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
	//创建日历v4版本
	ApiCalendarCreateV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars"
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

	//创建日程v4
	ApiCalendarEventCreateV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/events"
	//删除日程v4
	ApiCalendarEventDeleteV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/events/%s"
	//获取日程v4
	ApiCalendarEventGetV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/events/%s"
	//获取日程列表v4
	ApiCalendarEventBatchGetV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/events"
	//更新日程v4
	ApiCalendarEventUpdateV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/events/%s"
	//搜索日程v4
	ApiCalendarEventSearchV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/events/search"
	//创建日程参与人v4
	ApiCalendarEventAttendeesAddV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/events/%s/attendees"
	//删除日程参与人v4
	ApiCalendarEventAttendeesDeleteV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/events/%s/attendees/batch_delete"
	//获取日程参与人列表v4
	ApiCalendarEventAttendeesGetV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/events/%s/attendees"
	//获取参与人群成员列表v4
	ApiCalendarEventAttendeesChatMembersGetV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/events/%s/attendees/%s/chat_members"

	ApiCalendarGetV4          = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s"
	ApiCalendarDeleteV4       = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s"
	ApiCalendarGetListV4      = "https://open.feishu.cn/open-apis/calendar/v4/calendars"
	ApiCalendarUpdateV4       = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s"
	ApiCalendarSearchV4       = "https://open.feishu.cn/open-apis/calendar/v4/calendars/search"
	ApiCalendarUnsubscribeV4  = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/unsubscribe"
	ApiCalendarSubscribeV4    = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/subscribe"
	ApiCalendarSubscriptionV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/subscription"
	//创建访问控制
	ApiCalendarAddCalendarAclV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/acls"
	//删除访问控制
	ApiCalendarDeleteCalendarAclV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/acls/%s"
	//获取访问控制列表
	ApiCalendarCalendarAclGetV4 = "https://open.feishu.cn/open-apis/calendar/v4/calendars/%s/acls"

	//搜索用户
	ApiSearchUser = "https://open.feishu.cn/open-apis/search/v1/user"

	//检验应用管理员
	ApiIsUserAdmin = "https://open.feishu.cn/open-apis/application/v3/is_user_admin"
	//查询应用管理员列表
	ApiAdminUserList = "https://open.feishu.cn/open-apis/user/v4/app_admin_user/list"
	//获取企业信息
	ApiOrgInfo = "https://open.feishu.cn/open-apis/tenant/v2/tenant/query"

	////////用户群组
	//获取用户所在的群列表
	ApiUserGroupLIst = "https://open.feishu.cn/open-apis/user/v4/group_list"
	//获取群成员列表
	ApiChatMembers = "https://open.feishu.cn/open-apis/chat/v4/members"
	//搜索用户所在的群列表
	ApiChatSearch = "https://open.feishu.cn/open-apis/chat/v4/search"
	//获取用户或机器人所在的群列表
	ApiImChatList = "https://open.feishu.cn/open-apis/im/v1/chats"
	//获取群信息
	ApiImChatInfo = "https://open.feishu.cn/open-apis/im/v1/chats/"

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

	////////云文档
	//查询文档
	ApiSearchDocs = "https://open.feishu.cn/open-apis/suite/docs-api/search/object"
	//获取云文档信息
	ApiGetDocMeta = "https://open.feishu.cn/open-apis/doc/v2/meta"
	//获取图片
	ApiGetImage = "https://open.feishu.cn/open-apis/image/v4/get"

	/// 授权状态
	// https://bytedance.feishu.cn/docs/doccnHJx2UbLZh5kiWjNawICyNd#
	ApiGetScopes   = "https://open.feishu.cn/open-apis/application/v6/scopes"
	ApiApplyScopes = "https://open.feishu.cn/open-apis/application/v6/scopes/apply"
)

//Other Const
const (
	TestAppId     = "**********************"
	TestAppSecret = "******************************"
	TestTicket    = "**********************************"
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
