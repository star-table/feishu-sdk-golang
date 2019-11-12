package vo

//定义参照: https://open.feishu.cn/open-apis/message/v4/send/
type MsgVo struct {
	OpenId string `json:"omitempty,open_id"`
	UserId string `json:"omitempty,user_id"`
	Email string `json:"omitempty,email"`
	ChatId string `json:"omitempty,chat_id"`
	MsgType string `json:"msg_type"`
	RootId string `json:"omitempty,root_id"`
	UpdateMulti bool `json:"update_multi"`

	Card Card `json:"card"`
}

type MsgResp struct {
	CommonVo
	
	Data MsgRespData `json:"data"`
}

type MsgRespData struct {
	MessageId string `json:"message_id"`
}

//机器人消息Card字段数据格式定义
type Card struct {
	Config *CardConfig `json:"config"`
	CardLink *CardElementUrl `json:"card_link"`
	Header *CardHeader `json:"header"`
	Elements []interface{} `json:"elements"`
	I18nElements *I18nElement `json:"i18n_elements"`
}

type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode"`
}

type CardHeader struct {
	Title *CardHeaderTitle `json:"title"`
}

type CardHeaderTitle struct {
	Tag string `json:"tag"`
	Content string `json:"content"`
	Lines int `json:"omitempty,lines"`
	I18n *CardI18n `json:"i18n"`
}

type CardI18n struct {
	ZhCn string `json:"zh_cn"`
	EnUs string `json:"en_us"`
	JaJp string `json:"ja_jp"`
}

type CardElement struct {
	Tag string `json:"tag"`

}

type CardElementContentModule struct {
	CardElement

	Text CardElementText `json:"text"`
	Fields []CardElementField `json:"fields"`
	Extra *CardElementExtra `json:"extra"`
}

type CardElementBrModule struct {
	CardElement
}

type CardElementImageModule struct {
	CardElement

	ImgKey string `json:"img_key"`
	Alt CardElementText `json:"alt"`
	Title *CardElementText `json:"alt"`
}

type CardElementActionModule struct {
	CardElement
}

type Action struct {
	Tag string `json:"tag"`
}

type ActionButton struct {
	Action

	Text CardElementText `json:"text"`
	Url string `json:"omitempty,url"`
	MultiUrl *CardElementUrl `json:"multi_url"`
	Type string `json:"omitempty,type"`
	Value string `json:"omitempty,value"`
	Confirm *CardElementConfirm `json:"confirm"`
}

type ActionSelectMenu struct {
	Action

	Placeholder *CardElementText `json:"placeholder"`
	InitialOption string `json:"omitempty,initial_option"`
	Options []CardElementOption `json:"options"`
	Value string `json:"omitempty,value"`
	Confirm *CardElementConfirm `json:"confirm"`
}

type ActionOverflow struct {
	Action

	Options []CardElementOption `json:"options"`
	Value string `json:"omitempty,value"`
	Confirm *CardElementConfirm `json:"confirm"`
}

type ActionDatePicker struct {
	Action

	InitialDate string `json:"omitempty,initial_date"`
	InitialTime string `json:"omitempty,initial_time"`
	InitialDatetime string `json:"omitempty,initial_datetime"`
	Placeholder string `json:"omitempty,placeholder"`
	Value string `json:"omitempty,value"`
	Confirm *CardElementConfirm `json:"confirm"`
}

type CardElementExtra struct {
	Tag string `json:"tag"`
	ImgKey string `json:"img_key"`
	Alt *CardElementText `json:"alt"`
}

type CardElementField struct {
	IsShort bool `json:"is_short"`
	Text *CardElementText `json:"text"`
}

type I18nElement struct {
	ZhCn []interface{} `json:"zh_cn"`
	EnUs []interface{} `json:"en_us"`
	JaJp []interface{} `json:"ja_jp"`
}

type CardElementText struct {
	Tag string `json:"tag"`
	Content string `json:"content"`
	Lines int `json:"omitempty,lines"`
}

type CardElementAction struct {
	Tag string `json:"tag"`
	Text *CardElementText `json:"text"`
	Type string `json:"type"`
}

type CardElementUrl struct {
	Url string `json:"url"`
	AndroidUrl string `json:"android_url"`
	IosUrl string `json:"ios_url"`
	PcUrl string `json:"pc_url"`
}

type CardElementConfirm struct {
	Title *CardHeaderTitle `json:"title"`
	Text *CardElementText `json:"text"`
}

type CardElementOption struct {
	Text *CardElementText `json:"text"`
	Value string `json:"value"`
	Url string `json:""omitempty,url"`
	MultiUrl *CardElementUrl `json:"multi_url"`
}