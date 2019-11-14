package consts

//API Host const， v3
const(
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


	//////////////////部门和用户
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
)

//Other Const
const(
	TestAppId = "cli_9d40f5bf08f95108"
	TestAppSecret = "Apx5vdWeIxVzDBQ6ARte6grZgOCgbhgP"
	TestTicket = "c72293913862fdd6d423f61283adbe0b91980c5d"
)