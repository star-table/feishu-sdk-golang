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
)

//Other Const
const(
	TestAppId = "xxx"
	TestAppSecret = "xxx"
)