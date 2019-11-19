package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
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

//获取部门用户详情列表 https://open.feishu.cn/document/ukTMukTMukTM/uYzN3QjL2czN04iN3cDN?lang=zh-CN
func (t Tenant) GetDepartmentUserDetailList(departmentId string, offset, pageSize int, fetchChild bool) (*vo.GetDepartmentUserDetailListRespVo, error){
	queryParams := map[string]interface{}{
		"department_id": departmentId,
		"offset": offset,
		"page_size": pageSize,
		"fetch_child": fetchChild,
	}
	respBody, err := http.Get(consts.ApiDepartmentUserDetailList, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetDepartmentUserDetailListRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//批量获取用户信息 https://open.feishu.cn/document/ukTMukTMukTM/uIzNz4iM3MjLyczM
func (t Tenant) GetUserBatchGet(employeeIds []string, openIds []string) (*vo.GetUserBatchGetRespVo, error){
	queryParams := make([]http.QueryParameter, 0)
	if employeeIds != nil && len(employeeIds) > 0{
		for _, id := range employeeIds{
			queryParams = append(queryParams, http.QueryParameter{
				Key: "employee_ids",
				Value: id,
			})
		}
	}
	if openIds != nil && len(openIds) > 0{
		for _, id := range openIds{
			queryParams = append(queryParams, http.QueryParameter{
				Key: "open_ids",
				Value: id,
			})
		}
	}
	respBody, err := http.GetRepetition(consts.ApiDepartmentUserList, queryParams, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetUserBatchGetRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}