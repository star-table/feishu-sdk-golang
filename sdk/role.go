package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//获取角色列表 https://open.feishu.cn/document/ukTMukTMukTM/uYzMwUjL2MDM14iNzATN?lang=zh-CN
func (t Tenant) GetRoleList() (*vo.RoleListResp, error){
	respBody, err := http.Get(consts.ApiRoleList, nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.RoleListResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取角色成员列表 https://open.feishu.cn/document/ukTMukTMukTM/uczMwUjL3MDM14yNzATN?lang=zh-CN
func (t Tenant) GetRoleMemberList() (*vo.RoleMemberListResp, error){
	respBody, err := http.Get(consts.ApiRoleMemberList, nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.RoleMemberListResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}