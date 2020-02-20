package vo

type GroupListRespVo struct {
	CommonVo
	Data *UserGroupListDat `json:"data"`
}

type UserGroupListDat struct {
	HasMore   bool        `json:"has_more"`
	PageToken string      `json:"page_token"`
	Groups    []GroupData `json:"groups"`
}

type GroupData struct {
	Avatar      string `json:"avatar"`
	ChatId      string `json:"chat_id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	OwnerOpenId string `json:"owner_open_id"`
	OwnerUserId string `json:"owner_user_id"`
}

type ChatMembersRespVo struct {
	CommonVo
	Data *ChatGroupData `json:"data"`
}

type ChatGroupData struct {
	ChatId  string       `json:"chat_id"`
	HasMore bool         `json:"has_more"`
	Members []MemberData `json:"members"`
}

type MemberData struct {
	OpenId string `json:"open_id"`
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}
