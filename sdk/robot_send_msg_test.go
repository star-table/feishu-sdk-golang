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
		OpenId:  "ou_433f2b1abf3e0ec316fd9e60d4cda654",
		MsgType: "text",
		Content: &vo.MsgContent{
			Text: "test msg",
		},
	})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)

	resp, err = tenant.SendMessage(vo.MsgVo{
		OpenId:  "ou_433f2b1abf3e0ec316fd9e60d4cda654",
		MsgType: "post",
		Content: &vo.MsgContent{
			Post: &vo.MsgPost{
				ZhCn: &vo.MsgPostValue{
					Title: "test title",
					Content: []interface{}{
						[]interface{}{
							vo.MsgPostContentText{
								Tag:  "text",
								Text: "测试一哈",
							},
							vo.MsgPostContentA{
								Tag:  "a",
								Text: "超链接",
								Href: "http://feishu.cn",
							},
						},
						[]interface{}{
							vo.MsgPostContentText{
								Tag:  "text",
								Text: "测试二行",
							},
							vo.MsgPostContentA{
								Tag:  "a",
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

//func TestNewFileUploadRequest(t *testing.T) {
//	data := make(map[string]string)
//	t.Log(NewFileUploadRequest("https://open.feishu.cn/open-apis/image/v4/put/",
//		data, "image", "C:\\Users\\admin\\Desktop\\11111.png"))
//}

func TestTenant_SendMessage_Card(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	//tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	tenant, e := BuildTenant(app.AppAccessToken, "2eb972ddc1cf1652") //测试企业二
	t.Log(e)

	resp, err := tenant.SendMessage(vo.MsgVo{
		//OpenId: "ou_0e40d4035ceb951467beb62bb1f3e5ba", //陈
		//OpenId:  "ou_dcff2de6a28060eff9f0d9665d2d28be",//我
		OpenId:  "ou_44fbc016f49fa848a6182e4ff3c48780", //我 测试企业二
		MsgType: "interactive",
		Card: &vo.Card{
			Config: &vo.CardConfig{
				WideScreenMode: true,
			},
			Header: &vo.CardHeader{
				Title: &vo.CardHeaderTitle{
					Tag:     "plain_text",
					Content: "测试任务",
				},
			},
			Elements: []interface{}{
				vo.CardElementContentModule{
					Tag: "div",
					Text: &vo.CardElementText{
						Tag:     "plain_text",
						Content: "任务标题",
					},
				},
				vo.CardElementContentModule{
					Tag: "div",
					Text: &vo.CardElementText{
						Tag:     "plain_text",
						Content: "负责人：aabc",
					},
				},
				vo.CardElementNote{
					Tag: "note",
					Elements: []interface{}{
						vo.CardElementTextAlt{
							Tag:     "plain_text",
							Content: "开始时间:无        ",
						},
						vo.CardElementTextAlt{
							Tag:     "plain_text",
							Content: "截止时间:1990年12月12日     ",
						},
						vo.CardElementTextAlt{
							Tag:     "plain_text",
							Content: "原截止时间:1990年12月12日",
						},
					},
				},
				vo.CardElementActionModule{
					Tag: "action",
					Actions: []interface{}{
						vo.ActionButton{
							Tag: "button",
							Text: vo.CardElementText{
								Tag:     "plain_text",
								Content: "任务详情",
							},
							Url:  "https://polaris.feishu.cn/docs/doccnldXu2b7ei0pmHbViQXXSGg",
							Type: "default",
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

func TestTenant_SendMessage_Card_Daily_Report(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2e99b3ab0b0f1654")
	t.Log(e)

	resp, err := tenant.SendMessage(vo.MsgVo{
		OpenId:  "ou_ce397f53085cb712373f71874e8eae78",
		MsgType: "interactive",
		Card: &vo.Card{
			Config: &vo.CardConfig{
				WideScreenMode: true,
			},
			Header: &vo.CardHeader{
				Title: &vo.CardHeaderTitle{
					Tag:     "plain_text",
					Content: "任务日报",
				},
			},
			Elements: []interface{}{
				vo.CardElementContentModule{
					Tag: "div",
					Fields: []vo.CardElementField{
						{
							Text: vo.CardElementText{
								Tag:     "lark_md",
								Content: "**今日完成:** 100",
							},
						},
						{
							Text: vo.CardElementText{
								Tag:     "lark_md",
								Content: "**剩余:** 15",
							},
						},
					},
				},
				vo.CardElementActionModule{
					Tag: "action",
					Actions: []interface{}{
						vo.ActionButton{
							Tag: "button",
							Text: vo.CardElementText{
								Tag:     "plain_text",
								Content: "去处理",
							},
							Type: "link",
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

func TestTenant_SendMessage_PC_Applet(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.SendMessageBatch(vo.BatchMsgVo{
		OpenIds: []string{"ou_e1b43c426e884c586d52751853896688", "ou_64c74b3ffeb5ad1693ab830384373b5a"},
		MsgType: "interactive",
		Card: &vo.Card{
			Config: &vo.CardConfig{
				WideScreenMode: true,
			},
			Header: &vo.CardHeader{
				Title: &vo.CardHeaderTitle{
					Tag:     "plain_text",
					Content: "Pc小程序预览页",
				},
			},
			Elements: []interface{}{
				vo.CardElementActionModule{
					Tag: "action",
					Actions: []interface{}{
						vo.ActionButton{
							Tag: "button",
							Text: vo.CardElementText{
								Tag:     "plain_text",
								Content: "点击预览",
							},
							MultiUrl: &vo.CardElementUrl{
								PcUrl: "https://applink.feishu.cn/client/mini_program/open?appId=cli_9d5e49aae9ae9101&mode=sidebar-semi&path=pages/PC/Home/index",
								//PcUrl: "https://applink.feishu.cn/client/mini_program/open?appId=cli_9d5e49aae9ae9101&mode=sidebar-semi",
								IosUrl:     "https://baidu.com",
								AndroidUrl: "http://app.bjx.cloud/project",
							},
							Type: "default",
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

func TestTenant_SendNewIssueCard(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.SendMessage(vo.MsgVo{
		OpenId:  "ou_e1b43c426e884c586d52751853896688",
		MsgType: "interactive",
		Card: &vo.Card{
			//Config: &vo.CardConfig{
			//	WideScreenMode: true,
			//},
			Header: &vo.CardHeader{
				Title: &vo.CardHeaderTitle{
					Tag:     "plain_text",
					Content: "⏰ 新任务",
				},
			},
			Elements: []interface{}{
				vo.CardElementContentModule{
					Tag: "div",
					Fields: []vo.CardElementField{
						{
							Text: vo.CardElementText{
								Tag:     "lark_md",
								Content: "**所属项目**\n个人安排",
							},
						},
						{
							Text: vo.CardElementText{
								Tag:     "lark_md",
								Content: "**\n任务标题**增加一个按钮",
								Lines:   3,
							},
						},
						{
							Text: vo.CardElementText{
								Tag:     "lark_md",
								Content: "**\n负责人: ** " + "符莎321312" + "\t\t **优先级: **" + "普通",
							},
						},
					},
				},
				vo.CardElementContentModule{
					Tag: "div",
					Fields: []vo.CardElementField{
						{
							Text: vo.CardElementText{
								Tag:     "lark_md",
								Content: "**计划开始时间:**",
							},
						},
					},
				},
				vo.CardElementActionModule{
					Tag: "action",
					Actions: []interface{}{
						vo.ActionDatePicker{
							Tag: "picker_datetime",
							Placeholder: &vo.CardElementText{
								Tag:     "plain_text",
								Content: "修改计划开始时间",
							},
							Confirm: &vo.CardElementConfirm{
								Title: &vo.CardHeaderTitle{
									Tag:     "plain_text",
									Content: "确认要修改这个任务的计划开始时间吗?",
								},
								Text: &vo.CardElementText{
									Tag:     "plain_text",
									Content: "⏰ 新任务",
								},
							},
						},
					},
				},
				vo.CardElementContentModule{
					Tag: "div",
					Fields: []vo.CardElementField{
						{
							Text: vo.CardElementText{
								Tag:     "lark_md",
								Content: "**计划截止时间:**",
							},
						},
					},
				},
				vo.CardElementActionModule{
					Tag: "action",
					Actions: []interface{}{
						vo.ActionDatePicker{
							Tag: "picker_datetime",
							Placeholder: &vo.CardElementText{
								Tag:     "plain_text",
								Content: "修改截止时间",
							},
							Confirm: &vo.CardElementConfirm{
								Title: &vo.CardHeaderTitle{
									Tag:     "plain_text",
									Content: "确认要修改这个任务的截止时间吗?",
								},
								Text: &vo.CardElementText{
									Tag:     "plain_text",
									Content: "确定修改",
								},
							},
						},
					},
				},
				vo.CardElementActionModule{
					Tag: "action",
					Actions: []interface{}{
						vo.ActionButton{
							Tag: "button",
							Text: vo.CardElementText{
								Tag:     "lark_md",
								Content: "查看任务详情",
							},
							Type: "default",
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

func TestTenant_SendMessageBatch(t *testing.T) {
	param := vo.MsgVo{
		OpenId:  "ou_dcff2de6a28060eff9f0d9665d2d28be",
		MsgType: "interactive",
		Card: &vo.Card{
			Elements: []interface{}{
				vo.CardElementImageModule{
					Tag: "img",
					Alt: vo.CardElementText{
						Tag:     "lark_md",
						Content: "11",
						//Lines:   0,
						//Href:    nil,
					},
					Title: &vo.CardElementText{
						Tag:     "lark_md",
						Content: "11",
					},
					ImgKey: "img_5bb79fbc-12f6-4a67-a4f9-dee33e5bd3fg",
				},
			},
		},
	}

	t.Log(json.ToJson(param))
}
