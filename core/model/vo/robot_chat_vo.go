package vo

type CreateChatReqVo struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	OpenIds     []string          `json:"open_ids"`
	UserIds     []string          `json:"user_ids"`
	I18nNames   map[string]string `json:"i18n_names"`
}

type CreateChatRespVo struct {
	CommonVo
	Data CreateChatData `json:"data"`
}

type CreateChatData struct {
	ChatId         string   `json:"chat_id"`
	InvalidOpenIds []string `json:"invalid_open_ids"`
	InvalidUserIds []string `json:"invalid_user_ids"`
}

type ChatInfoRespVo struct {
	CommonVo
	Data ChatInfoData `json:"data"`
}

type ChatInfoData struct {
	Avatar      string            `json:"avatar"`
	ChatId      string            `json:"chat_id"`
	Description string            `json:"description"`
	I18nNames   map[string]string `json:"i18n_names"`
	Members     []ChatMemberData  `json:"members"`
	Name        string            `json:"name"`
	Type        string            `json:"type"`
	OwnerOpenId string            `json:"owner_open_id"`
	OwnerUserId string            `json:"owner_user_id"`
}

type ChatMemberData struct {
	OpenId string `json:"open_id"`
	UserId string `json:"user_id"`
}

type UpdateChatReqVo struct {
	ChatId      string            `json:"chat_id"`
	OwnerUserId string            `json:"owner_user_id,omitempty"`
	OwnerOpenId string            `json:"owner_open_id,omitempty"`
	Name        string            `json:"name,omitempty"`
	I18nNames   map[string]string `json:"i18n_names,omitempty"`
}

type UpdateChatRespVo struct {
	CommonVo
	Data UpdateChatData `json:"data"`
}

type UpdateChatData struct {
	ChatId string `json:"chat_id"`
}

type UpdateChatMemberReqVo struct {
	ChatId  string   `json:"chat_id"`
	UserIds []string `json:"user_ids"`
	OpenIds []string `json:"open_ids"`
}

type UpdateChatMemberRespVo struct {
	CommonVo
	Data UpdateChatMemberRespData `json:"data"`
}

type UpdateChatMemberRespData struct {
	InvalidOpenIds []string `json:"invalid_open_ids"`
	InvalidUserIds []string `json:"invalid_user_ids"`
}
