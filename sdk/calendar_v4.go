package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//创建日历v4版本 https://open.feishu.cn/document/ukTMukTMukTM/uUDM3YjL1AzN24SNwcjN
func (t Tenant) CreateCalendarV4(bodyParams vo.CreateCalendarV4Req) (*vo.CommonCalendarV4Resp, error) {
	respBody, err := http.Post(consts.ApiCalendarCreateV4, nil, json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonCalendarV4Resp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
