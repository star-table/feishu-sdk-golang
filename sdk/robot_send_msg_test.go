package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"
	"testing"
)

func TestTenant_SendMessage(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.SendMessage(vo.MsgVo{
		OpenId: "ou_433f2b1abf3e0ec316fd9e60d4cda654",
		MsgType: "text",
		Content: &vo.MsgContent{
			Text: "test msg",
		},
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)

	resp, err = tenant.SendMessage(vo.MsgVo{
		OpenId: "ou_433f2b1abf3e0ec316fd9e60d4cda654",
		MsgType: "post",
		Content: &vo.MsgContent{
			Post: &vo.MsgPost{
				ZhCn: &vo.MsgPostValue{
					Title: "test title",
					Content: []interface{}{
						[]interface{}{
							vo.MsgPostContentText{
								Tag: "text",
								Text: "测试一哈",
							},
							vo.MsgPostContentA{
								Tag: "a",
								Text: "超链接",
								Href: "http://feishu.cn",
							},
						},
						[]interface{}{
							vo.MsgPostContentText{
								Tag: "text",
								Text: "测试二行",
							},
							vo.MsgPostContentA{
								Tag: "a",
								Text: "超链接",
								Href: "http://feishu.cn",
							},
						},
					},
				},
			},
		},
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}