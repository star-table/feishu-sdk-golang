package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//校验应用管理员 https://open.feishu.cn/document/ukTMukTMukTM/uITN1EjLyUTNx4iM1UTM
func (t Tenant) IsUserAdmin(openId string, employeeId string) (*vo.IsUserAdminResp, error) {
	queryParams := map[string]interface{}{}
	if openId != "" {
		queryParams["open_id"] = openId
	}
	if employeeId != "" {
		queryParams["employee_id"] = employeeId
	}
	respBody, err := http.Get(consts.ApiIsUserAdmin, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.IsUserAdminResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//查询应用管理员列表 https://open.feishu.cn/document/ukTMukTMukTM/ucDOwYjL3gDM24yN4AjN
func (t Tenant) AdminUserList() (*vo.AdminUserListResp, error) {
	respBody, err := http.Get(consts.ApiAdminUserList, nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.AdminUserListResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取企业信息 https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/tenant-v2/tenant/query
func (t Tenant) OrgInfo() (*vo.OrgInfoResp, error) {
	respBody, err := http.Get(consts.ApiOrgInfo, nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.OrgInfoResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

func (t Tenant) GetJsTicket() (*vo.GetJsTicketResp, error) {
	respBody, err := http.Post(consts.ApiJSApiTicket, nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetJsTicketResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
