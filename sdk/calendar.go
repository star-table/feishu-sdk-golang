package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//获取日历 https://open.feishu.cn/document/ukTMukTMukTM/uMDN04yM0QjLzQDN?lang=zh-CN
func (t Tenant) GetCalendar(calendarId string) (*vo.CommonCalendarResp, error) {
	respBody, err := http.Get(consts.ApiCalendarGet + calendarId, nil, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonCalendarResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//创建日历 https://open.feishu.cn/document/ukTMukTMukTM/uQTM14CNxUjL0ETN?lang=zh-CN
func (t Tenant) CreateCalendar(bodyParams vo.CreateCalendarReq) (*vo.CommonCalendarResp, error) {
	respBody, err := http.Post(consts.ApiCalendarCrate, nil,  json.ToJsonIgnoreError(bodyParams), http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil{
		log.Error(err)
		return nil, err
	}
	respVo := &vo.CommonCalendarResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

func (t Tenant) UpdateCalendar() {

}