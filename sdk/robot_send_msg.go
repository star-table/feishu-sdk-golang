package sdk

import (
	"fmt"
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
)

//发送消息卡片 https://open.feishu.cn/document/ukTMukTMukTM/uYTNwUjL2UDM14iN1ATN
func (t Tenant) SendMessage(msg vo.MsgVo) (*vo.MsgResp, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	fmt.Println(1111, reqBody)
	respBody, err := http.Post(consts.ApiRobotSendMessage, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.MsgResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}

//发送消息卡片 https://open.feishu.cn/document/ukTMukTMukTM/uYTNwUjL2UDM14iN1ATN
func (t Tenant) SendMessageBatch(msg vo.BatchMsgVo) (*vo.MsgResp, error) {
	reqBody := json.ToJsonIgnoreError(msg)
	respBody, err := http.Post(consts.ApiRobotSendBatchMessage, nil, reqBody, http.BuildTokenHeaderOptions(t.TenantAccessToken))
	if err != nil {
		log.Error(err)
		return nil, err
	}
	respVo := &vo.MsgResp{}
	json.FromJsonIgnoreError(respBody, respVo)
	return respVo, nil
}
