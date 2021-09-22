package sdk

import (
	"fmt"
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
	"net/url"
)

//搜索用户 https://bytedance.feishu.cn/docs/doccnizryz7NKuUmVfkRJWeZGVc
func (u User) SearchUser(query string, pageSize int, pageToken string) (*vo.SearchUserResp, error) {
	queryParams := map[string]interface{}{}
	if query != "" {
		queryParams["query"] = query
	}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}
	if pageToken != "" {
		queryParams["page_token"] = pageToken
	}
	respBody, err := http.Get(consts.ApiSearchUser, queryParams, http.BuildTokenHeaderOptions(u.UserAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.SearchUserResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//使用手机号或邮箱获取用户ID https://open.feishu.cn/document/ukTMukTMukTM/uUzMyUjL1MjM14SNzITN
func (t Tenant) BatchGetId(emails []string, mobiles []string) (*vo.BatchGetIdResp, error) {
	query := ""
	for _, email := range emails {
		query += fmt.Sprintf("emails=%s&", email)
	}
	for _, mobile := range mobiles {
		query += fmt.Sprintf("mobiles=%s&", url.QueryEscape(mobile))
	}
	fullUrl := consts.ApiBatchGetUserId
	if len(query) > 0 {
		fullUrl += fmt.Sprintf("?%s", query[0:len(query)-1])
	}
	respBody, err := http.Get(fullUrl, nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.BatchGetIdResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
