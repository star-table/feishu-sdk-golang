package sdk

import (
	"github.com/polaris-team/feishu-sdk-golang/core/consts"
	"github.com/polaris-team/feishu-sdk-golang/core/model/vo"
	"github.com/polaris-team/feishu-sdk-golang/core/util/http"
	"github.com/polaris-team/feishu-sdk-golang/core/util/json"
	"github.com/polaris-team/feishu-sdk-golang/core/util/log"
)

//获取子部门列表 https://open.feishu.cn/document/ukTMukTMukTM/ugzN3QjL4czN04CO3cDN
func (t Tenant) GetDepartmentSimpleList(departmentId string, offset, pageSize int, fetchChild bool) (*vo.GetDepartmentSimpleListRespVo, error){
	queryParams := map[string]interface{}{
		"department_id": departmentId,
		"offset": offset,
		"page_size": pageSize,
		"fetch_child": fetchChild,
	}
	respBody, err := http.Get(consts.ApiDepartmentSimpleList, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetDepartmentSimpleListRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取部门详情 https://open.feishu.cn/document/ukTMukTMukTM/uAzNz4CM3MjLwczM
func (t Tenant) GetDepartmentInfo(departmentId string) (*vo.GetDepartmentInfoRespVo, error){
	queryParams := map[string]interface{}{
		"department_id": departmentId,
	}
	respBody, err := http.Get(consts.ApiDepartmentInfoGet, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetDepartmentInfoRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取部门用户列表 https://open.feishu.cn/document/ukTMukTMukTM/uEzNz4SM3MjLxczM
func (t Tenant) GetDepartmentUserList(departmentId string, offset, pageSize int, fetchChild bool) (*vo.GetDepartmentUserListRespVo, error){
	queryParams := map[string]interface{}{
		"department_id": departmentId,
		"offset": offset,
		"page_size": pageSize,
		"fetch_child": fetchChild,
	}
	respBody, err := http.Get(consts.ApiDepartmentUserList, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetDepartmentUserListRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}