package vo

type CreateCalendarV4Req struct {
	Summary      string `json:"summary"`
	Description  string `json:"description,omitempty"`
	Permissions  string `json:"permissions"`
	Color        int32  `json:"color"`
	SummaryAlias string `json:"summary_alias"`
}

type CommonCalendarV4Resp struct {
	CommonVo
	Data CalendarV4 `json:"data"`
}

type CalendarV4 struct {
	Calendar CalendarV4Data `json:"calendar"`
}

type CalendarV4Data struct {
	Description  string `json:"description"`
	Type         string `json:"type"`
	Summary      string `json:"summary"`
	Permissions  string `json:"permissions"`
	Color        int32  `json:"color"`
	SummaryAlias string `json:"summary_alias"`
	IsDeleted    bool   `json:"is_deleted"`
	Role         string `json:"role"`
	IsThirdParty bool   `json:"is_third_party"`
	CalendarId   string `json:"calendar_id"`
}

type VoidV4Resp struct {
	CommonVo
	Data string `json:"data"`
}

type DeleteCalendarV4Req struct {
	CalendarId string `json:"calendar_id"`
}

type GetCalendarV4Req struct {
	CalendarId string `json:"calendar_id"`
}

type GetCalendarV4Resp struct {
	CommonVo
	Data CalendarV4Data `json:"data"`
}

type GetCalendarListV4Req struct {
	PageSize int `json:"page_size"`
	PageToken string `json:"page_token"`
	SyncToken string `json:"sync_token"`
}

type GetCalendarListV4Resp struct {
	CommonVo
	Data GetCalendarListV4RespData `json:"data"`
}

type GetCalendarListV4RespData struct {
	PageToken string `json:"page_token"`
	SyncToken string `json:"sync_token"`
	HasMore bool `json:"has_more"`
	CalendarList []GetCalendarListV4RespDataCalendarItem `json:"calendar_list"`
}

type GetCalendarListV4RespDataCalendarItem struct {
	CalendarID  string `json:"calendar_id"`
	Color       int64  `json:"color"`
	Permissions string `json:"permissions"`
	Role        string `json:"role"`
	Summary     string `json:"summary"`
	Type        string `json:"type"`
}

type UpdateCalendarV4Req struct {
	Summary      string `json:"summary"`
	Description  string `json:"description"`
	Permissions  string `json:"permissions"`
	Color        int    `json:"color"`
	SummaryAlias string `json:"summary_alias"`
	CalendarId   string `json:"calendar_id"`
}

type SearchCalendarV4Req struct {
	Query string `json:"query"`
}

type SearchCalendarV4Resp struct {
	CommonVo
	Data SearchCalendarV4RespData `data`
}
type SearchCalendarV4RespData struct {
	PageToken string `json:"page_token"`
	Items []SearchCalendarV4RespDataItem `json:"items"`
}
type SearchCalendarV4RespDataItem struct {
	CalendarID string `json:"calendar_id"`
	Summary    string `json:"summary"`
	Type       string `json:"type"`
}

type UnsubscribeCalendarV4Req struct {
	CalendarId string `json:"calendar_id"`
}

type SubscribeCalendarV4Req struct {
	CalendarId string `json:"calendar_id"`
}

type SubscribeCalendarV4Resp struct {
	CommonVo
	Data SubscribeCalendarV4RespData `json:"data"`
}
type SubscribeCalendarV4RespData struct {
	Calendar SubscribeCalendarV4RespDataCalendar `json:"calendar"`
}
type SubscribeCalendarV4RespDataCalendar struct {
	CalendarID string `json:"calendar_id"`
	Color      int64  `json:"color"`
	Role       string `json:"role"`
	Summary    string `json:"summary"`
	Type       string `json:"type"`
}
