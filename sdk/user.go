package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//获取部门用户列表 https://open.feishu.cn/document/ukTMukTMukTM/uEzNz4SM3MjLxczM
func (u User) SearchUser(query string, pageSize int, pageToken string) (*vo.SearchUserResp, error){
	queryParams := map[string]interface{}{
		"query": query,
	}
	if pageSize > 0{
		queryParams["page_size"] = pageSize
	}
	if pageToken != ""{
		queryParams["page_token"] = pageToken
	}
	respBody, err := http.Get(consts.ApiSearchUser, queryParams, http.BuildTokenHeaderOptions(u.UserAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.SearchUserResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}