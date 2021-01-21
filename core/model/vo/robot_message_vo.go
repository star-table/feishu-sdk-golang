package vo

import "encoding/binary"

//定义参照: https://open.feishu.cn/open-apis/message/v4/send/
type MsgVo struct {
	OpenId      string `json:"open_id,omitempty"`
	UserId      string `json:"user_id,omitempty"`
	Email       string `json:"email,omitempty"`
	ChatId      string `json:"chat_id,omitempty"`
	MsgType     string `json:"msg_type"`
	RootId      string `json:"root_id,omitempty"`
	UpdateMulti bool   `json:"update_multi"`

	Card    *Card       `json:"card,omitempty"`
	Content *MsgContent `json:"content,omitempty"`
}

type BatchMsgVo struct {
	DepartmentIds []string `json:"department_ids"`
	OpenIds       []string `json:"open_ids"`
	UserIds       []string `json:"user_ids"`
	MsgType       string   `json:"msg_type"`

	Card    *Card       `json:"card,omitempty"`
	Content *MsgContent `json:"content,omitempty"`
}

type MsgContent struct {
	Text     string   `json:"text"`
	ImageKey string   `json:"image_key"`
	Post     *MsgPost `json:"post,omitempty"`
}

type MsgPost struct {
	ZhCn *MsgPostValue `json:"zh_cn,omitempty"`
	EnUs *MsgPostValue `json:"en_us,omitempty"`
	JaJp *MsgPostValue `json:"ja_jp,omitempty"`
}

type MsgPostValue struct {
	Title   string      `json:"title"`
	Content interface{} `json:"content"`
}

type MsgPostContentText struct {
	Tag      string `json:"tag"`
	UnEscape bool   `json:"un_escape"`
	Text     string `json:"text"`
}

type MsgPostContentA struct {
	Tag  string `json:"tag"`
	Text string `json:"text"`
	Href string `json:"href"`
}

type MsgPostContentAt struct {
	Tag    string `json:"tag"`
	UserId string `json:"user_id"`
}

type MsgPostContentImage struct {
	Tag      string  `json:"tag"`
	ImageKey string  `json:"image_key"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
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
	Config       *CardConfig     `json:"config,omitempty"`
	CardLink     *CardElementUrl `json:"card_link,omitempty"`
	Header       *CardHeader     `json:"header,omitempty"`
	Elements     []interface{}   `json:"elements"`
	I18nElements *I18nElement    `json:"i18n_elements,omitempty"`
}

type CardConfig struct {
	WideScreenMode bool `json:"wide_screen_mode"`
}

type CardHeader struct {
	Title *CardHeaderTitle `json:"title,omitempty"`
}

type CardHeaderTitle struct {
	Tag     string    `json:"tag"`
	Content string    `json:"content"`
	Lines   int       `json:"lines,omitempty"`
	I18n    *CardI18n `json:"i18n,omitempty"`
}

type CardI18n struct {
	ZhCn string `json:"zh_cn"`
	EnUs string `json:"en_us"`
	JaJp string `json:"ja_jp"`
}

type CardElementContentModule struct {
	Tag string `json:"tag"`

	Text   *CardElementText   `json:"text"`
	Fields []CardElementField `json:"fields"`
	Extra  *CardElementExtra  `json:"extra,omitempty"`
}

type CardElementBrModule struct {
	Tag string `json:"tag"`
}

type CardElementImageModule struct {
	Tag string `json:"tag"`

	ImgKey string           `json:"img_key"`
	Alt    CardElementText  `json:"alt"`
	Title  *CardElementText `json:"title,omitempty"`
}

type CardElementTextAlt struct {
	Tag     string                    `json:"tag"`
	Content string                    `json:"content"`
	Lines   int                       `json:"lines,omitempty"`
	Href    map[string]CardElementUrl `json:"href,omitempty"`
}

type CardElementActionModule struct {
	Tag     string        `json:"tag"`
	Layout  string        `json:"layout"`
	Actions []interface{} `json:"actions"`
}

type ActionButton struct {
	Tag string `json:"tag"`

	Text     CardElementText        `json:"text"`
	Url      string                 `json:"url,omitempty"`
	MultiUrl *CardElementUrl        `json:"multi_url"`
	Type     string                 `json:"type,omitempty"`
	Value    map[string]interface{} `json:"value,omitempty"`
	Confirm  *CardElementConfirm    `json:"confirm,omitempty"`
}

type ActionSelectMenu struct {
	Tag string `json:"tag"`

	Placeholder   *CardElementText       `json:"placeholder,omitempty"`
	InitialOption string                 `json:"initial_option,omitempty"`
	Options       []CardElementOption    `json:"options"`
	Value         map[string]interface{} `json:"value,omitempty"`
	Confirm       *CardElementConfirm    `json:"confirm,omitempty"`
}

type ActionOverflow struct {
	Tag string `json:"tag"`

	Options []CardElementOption    `json:"options"`
	Value   map[string]interface{} `json:"value,omitempty"`
	Confirm *CardElementConfirm    `json:"confirm,omitempty"`
}

type ActionDatePicker struct {
	Tag string `json:"tag"`

	InitialDate     string                 `json:"initial_date,omitempty"`
	InitialTime     string                 `json:"initial_time,omitempty"`
	InitialDatetime string                 `json:"initial_datetime,omitempty"`
	Placeholder     *CardElementText       `json:"placeholder,omitempty"`
	Value           map[string]interface{} `json:"value,omitempty"`
	Confirm         *CardElementConfirm    `json:"confirm,omitempty"`
}

type CardElementExtra struct {
	Tag    string           `json:"tag"`
	ImgKey string           `json:"img_key"`
	Alt    *CardElementText `json:"alt,omitempty"`
}

type CardElementField struct {
	IsShort bool            `json:"is_short"`
	Text    CardElementText `json:"text,omitempty"`
}

type I18nElement struct {
	ZhCn []interface{} `json:"zh_cn"`
	EnUs []interface{} `json:"en_us"`
	JaJp []interface{} `json:"ja_jp"`
}

type CardElementText struct {
	Tag     string                    `json:"tag"`
	Content string                    `json:"content"`
	Lines   int                       `json:"lines,omitempty"`
	Href    map[string]CardElementUrl `json:"href,omitempty"`
}

type CardElementAction struct {
	Tag  string           `json:"tag"`
	Text *CardElementText `json:"text,omitempty"`
	Type string           `json:"type"`
}

type CardElementUrl struct {
	Url        string `json:"url"`
	AndroidUrl string `json:"android_url"`
	IosUrl     string `json:"ios_url"`
	PcUrl      string `json:"pc_url"`
}

type CardElementConfirm struct {
	Title *CardHeaderTitle `json:"title,omitempty"`
	Text  *CardElementText `json:"text,omitempty"`
}

type CardElementOption struct {
	Text     *CardElementText `json:"text,omitempty"`
	Value    string           `json:"value"`
	Url      string           `json:"url,omitempty"`
	MultiUrl *CardElementUrl  `json:"multi_url,omitempty"`
}

type UpdateImageVo struct {
	Image binary.ByteOrder
}

type CardElementNote struct {
	Tag      string        `json:"tag"`
	Elements []interface{} `json:"elements"`
}
