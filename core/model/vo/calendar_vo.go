package vo

type Calendar struct {
	Id                string `json:"id"`
	Summary           string `json:"summary"`
	Description       string `json:"description"`
	IsPrivate         bool `json:"is_private"`
	DefaultAccessRole string `json:"default_access_role"`
}

type CommonCalendarResp struct {
	CommonVo
	Data Calendar `json:"data"`
}

type CreateCalendarReq struct {
	Summary           string `json:"summary"`
	Description       string `json:"description,omitempty"`
	IsPrivate         bool `json:"is_private,omitempty"`
	DefaultAccessRole string `json:"default_access_role,omitempty"`
}

type UpdateCalendarReq struct {
	CalendarId string `json:"calendarId,omitempty"`
	Summary           string `json:"summary,omitempty"`
	Description       string `json:"description,omitempty"`
	IsPrivate         bool `json:"is_private,omitempty"`
	DefaultAccessRole string `json:"default_access_role,omitempty"`
}
