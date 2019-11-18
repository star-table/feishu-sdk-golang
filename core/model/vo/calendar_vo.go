package vo

type Calendar struct {
	Id                string `json:"id"`
	Summary           string `json:"summary"`
	Description       string `json:"description"`
	IsPrivate         bool   `json:"is_private"`
	DefaultAccessRole string `json:"default_access_role"`
}

type CommonCalendarResp struct {
	CommonVo
	Data Calendar `json:"data"`
}

type CalendarListResp struct {
	CommonVo
	PageToken string       `json:"page_token"`
	SyncToken string       `json:"sync_token"`
	Data      CalendarList `json:"data"`
}

type CalendarList struct {
	Items []Calendar `json:"items"`
}

type CreateCalendarReq struct {
	Summary           string `json:"summary"`
	Description       string `json:"description,omitempty"`
	IsPrivate         bool   `json:"is_private,omitempty"`
	DefaultAccessRole string `json:"default_access_role,omitempty"`
}

type UpdateCalendarReq struct {
	CalendarId        string `json:"calendarId,omitempty"`
	Summary           string `json:"summary,omitempty"`
	Description       string `json:"description,omitempty"`
	IsPrivate         bool   `json:"is_private,omitempty"`
	DefaultAccessRole string `json:"default_access_role,omitempty"`
}

type CreateCalendarEventReq struct {
	Summary     string       `json:"summary"`
	Description string       `json:"description,omitempty"`
	Start       TimeFormat   `json:"start"`
	End         TimeFormat   `json:"end"`
	Attendees   *[]Attendees `json:"attendees,omitempty"`
	Visibility  string       `json:"visibility,omitempty"`
}

type TimeFormat struct {
	Date      string `json:"date,omitempty"`
	TimeStamp int    `json:"time_stamp,omitempty"`
	TimeZone  string `json:"time_zone,omitempty"`
}

type Attendees struct {
	OpenId      string `json:"open_id,omitempty"`
	EmployeeId  string `json:"employee_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Optional    bool   `json:"optional,omitempty"`
}

type CommonCalendarEventResp struct {
	CommonVo
	Data CalendarEvent `json:"data"`
}

type CalendarEvent struct {
	Id          string      `json:"id"`
	Summary     string      `json:"summary"`
	Description string      `json:"description"`
	Start       TimeFormat  `json:"start"`
	End         TimeFormat  `json:"end"`
	Attendees   []Attendees `json:"attendees"`
	Visibility  string      `json:"visibility"`
}

type CalendarEventListResp struct {
	CommonVo
	PageToken string          `json:"page_token"`
	SyncToken string          `json:"sync_token"`
	Data      []CalendarEvent `json:"data"`
}

type AttendeesResp struct {
	Attendees
	Status string `json:"status"`
}

type UpdateCalendarEventAtendeesReq struct {
	Attendees []AttendeesResp `json:"attendees"`
}

type UpdateCalendarEventAtendeesResp struct {
	CommonVo
	Data []Attendees `json:"data"`
}
