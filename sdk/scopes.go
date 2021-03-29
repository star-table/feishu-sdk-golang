package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//1.查询租户授权状态 https://bytedance.feishu.cn/docs/doccnHJx2UbLZh5kiWjNawICyNd#dCNL6V
func (t Tenant) GetScopes() (*vo.GetScopesResp, error){
	respBody, err := http.Get(consts.ApiGetScopes, nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetScopesResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

// 申请授权 https://bytedance.feishu.cn/docs/doccnHJx2UbLZh5kiWjNawICyNd#kHHiAa
func (t Tenant) ApplyScopes() (*vo.ApplyScopesResp, error){
	respBody, err := http.Post(consts.ApiApplyScopes, nil, "", http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.ApplyScopesResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
