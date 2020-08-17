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

func TestNewFileUploadRequest(t *testing.T) {
	data := make(map[string]string)
	t.Log(NewFileUploadRequest("https://open.feishu.cn/open-apis/image/v4/put/",
		data, "image", "C:\\Users\\admin\\Desktop\\物料宣传.jpg"))
}

func TestTenant_SendMessage_Card(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2e499c21bf8f1658")
	t.Log(e)

	resp, err := tenant.SendMessage(vo.MsgVo{
		OpenId:  "ou_34afd0445a227c382e54e756842f834e",
		MsgType: "interactive",
		Card: &vo.Card{
			Config: &vo.CardConfig{
				WideScreenMode: false,
			},
			Header: &vo.CardHeader{
				Title: &vo.CardHeaderTitle{
					Tag:     "plain_text",
					Content: "「北极星协作」近期更新汇总 ",
				},
			},
			Elements: []interface{}{
				vo.CardElementImageModule{
					Tag: "img",
					Alt: vo.CardElementText{
						Tag:     "lark_md",
						Content: "北极星协作",
					},
					ImgKey: "img_5bb79fbc-12f6-4a67-a4f9-dee33e5bd3fg",
				},
				vo.CardElementContentModule{
					Tag: "div",
					Text: &vo.CardElementText{
						Tag:     "plain_text",
						Content: "1. 小程序：大幅优化了性能和体验\n2. 工作台改版：方便快速定位到近期截止且与我相关的任务\n3. 新增任务筛选器：组织内与我相关的所有任务一目了然\n4. 新增任务表格视图：更直观的查看任务信息，支持批量操作任务\n5. 新增敏捷研发项目：需求的分解 –> 任务的推进 –> QA缺陷的修复，迭代规划，高效推动产品的Scrum敏捷研发",
					},
				},

				vo.CardElementActionModule{
					Tag: "action",
					Actions: []interface{}{
						vo.ActionButton{
							Tag: "button",
							Text: vo.CardElementText{
								Tag:     "plain_text",
								Content: "查看更多",
							},
							Url:  "https://polaris.feishu.cn/docs/doccn5PXJnUDDhLee2Uksn337mb#0v0O1m",
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
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.SendMessage(vo.MsgVo{
		OpenId:  "ou_433f2b1abf3e0ec316fd9e60d4cda654",
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
							MultiUrl: &vo.CardElementUrl{
								PcUrl:      "http://feishu.cn",
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
