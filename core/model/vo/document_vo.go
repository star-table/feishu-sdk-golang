package vo

type SearchDocsReqVo struct {
	SearchKey *string   `json:"search_key"`
	Count     *int      `json:"count"`
	Offset    *int      `json:"offset"`
	OwnerIds  *[]string `json:"owner_ids"`
	ChatIds   *[]string `json:"chat_ids"`
	DocsTypes *[]string `json:"docs_types"`
}

type SearchDocsRespVo struct {
	CommonVo
	Data SearchDocsData `json:"data"`
}

type SearchDocsData struct {
	HasMore      bool          `json:"has_more"`
	Total        int64         `json:"total"`
	DocsEntities []DocEntities `json:"docs_entities"`
}

type DocEntities struct {
	DocsToken string `json:"docs_token"`
	DocsType  string `json:"docs_type"`
	Title     string `json:"title"`
	OwnerId   string `json:"owner_id"`
}

type GetDocMetaRespVo struct {
	CommonVo
	Data DocMetaData `json:"data"`
}

type DocMetaData struct {
	CreateDate     string `json:"create_date"`
	CreateTime     int64  `json:"create_time"`
	CreateUid      string `json:"create_uid"`
	CreateUserName string `json:"create_user_name"`
	DeleteFlag     int    `json:"delete_flag"`
	EditTime       int64  `json:"edit_time"`
	EditUserName   string `json:"edit_user_name"`
	IsExternal     bool   `json:"is_external"`
	IsPined        bool   `json:"is_pined"`
	IsStared       bool   `json:"is_stared"`
	ObjType        string `json:"obj_type"`
	OwnerId        string `json:"owner_id"`
	OwnerUserName  string `json:"owner_user_name"`
	ServerTime     int64  `json:"server_time"`
	TenantId       string `json:"tenant_id"`
	Title          string `json:"title"`
	Type           int    `json:"type"`
	Url            string `json:"url"`
}
