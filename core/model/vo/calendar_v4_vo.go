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
