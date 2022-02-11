package vo

type GroupListRespVo struct {
	CommonVo
	Data *UserGroupListData `json:"data"`
}

type UserGroupListData struct {
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

type ImChatListRespVo struct {
	CommonVo
	Data ImChatListData `json:"data"`
}

type ImChatListData struct {
	Items     []SimpleChatInfo `json:"items"`
	PageToken string           `json:"page_token"`
	HasMore   bool             `json:"has_more"`
}

type SimpleChatInfo struct {
	ChatId      string `json:"chat_id"`
	Avatar      string `json:"avatar"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerId     string `json:"owner_id"`
	OwnerIdType string `json:"owner_id_type"`
}

type ImChatInfoRespVo struct {
	CommonVo
	Data ImChatInfoData `json:"data"`
}

type ImChatInfoData struct {
	Avatar                 string      `json:"avatar"`
	Name                   string      `json:"name"`
	Description            string      `json:"description"`
	I18nElements           I18nElement `json:"i18n_elements,omitempty"`
	AddMemberPermission    string      `json:"add_member_permission"`
	ShareCardPermission    string      `json:"share_card_permission"`
	AtAllPermission        string      `json:"at_all_permission"`
	EditPermission         string      `json:"edit_permission"`
	OwnerIdType            string      `json:"owner_id_type"`
	OwnerId                string      `json:"owner_id"`
	ChatMode               string      `json:"chat_mode"`
	ChatType               string      `json:"chat_type"`
	ChatTag                string      `json:"chat_tag"`
	JoinMessageVisibility  string      `json:"join_message_visibility"`
	LeaveMessageVisibility string      `json:"leave_message_visibility"`
	MembershipApproval     string      `json:"membership_approval"`
	ModerationPermission   string      `json:"moderation_permission"`
	External               bool        `json:"external"`
	TenantKey              string      `json:"tenant_key"`
}
