package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//文档搜索 https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/ugDM4UjL4ADO14COwgTN
func (t Tenant) SearchDocs(userAccessToken string, req vo.SearchDocsReqVo) (*vo.SearchDocsRespVo, error) {
	respBody, err := http.Post(consts.ApiSearchDocs, nil, json.ToJsonIgnoreError(req), http.BuildTokenHeaderOptions(userAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.SearchDocsRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//获取文档元信息 https://open.feishu.cn/document/ugTM5UjL4ETO14COxkTN/uczN3UjL3czN14yN3cTN
func (t Tenant) GetDocMeta(userAccessToken string, docToken string) (*vo.GetDocMetaRespVo, error) {
	respBody, err := http.Get(consts.ApiGetDocMeta+"/"+docToken, nil, http.BuildTokenHeaderOptions(userAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.GetDocMetaRespVo{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
