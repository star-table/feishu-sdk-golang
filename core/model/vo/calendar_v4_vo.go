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
	PageSize  int    `json:"page_size"`
	PageToken string `json:"page_token"`
	SyncToken string `json:"sync_token"`
}

type GetCalendarListV4Resp struct {
	CommonVo
	Data GetCalendarListV4RespData `json:"data"`
}

type GetCalendarListV4RespData struct {
	PageToken    string                                  `json:"page_token"`
	SyncToken    string                                  `json:"sync_token"`
	HasMore      bool                                    `json:"has_more"`
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
	PageToken string                         `json:"page_token"`
	Items     []SearchCalendarV4RespDataItem `json:"items"`
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

type CreateCalendarEventV4Req struct {
	Summary         string                  `json:"summary"`
	Description     string                  `json:"description"`
	StartTime       CalendarEventTime       `json:"start_time"`
	EndTime         CalendarEventTime       `json:"end_time"`
	Vchat           *Vchat                  `json:"vchat"`
	Visibility      *string                 `json:"visibility"`       //default(忙碌),public(公开),private（私密）
	AttendeeAbility *string                 `json:"attendee_ability"` //none,can_see_other,can_invite_others,can_modify_event
	FreeBusyStatus  *string                 `json:"free_busy_status"`
	Location        *Location               `json:"location"`
	Color           int32                   `json:"color"`
	Reminders       []CalendarEventReminder `json:"reminders"`
	Recurrence      string                  `json:"recurrence"`
	Schemas         []CalendarEventSchema   `json:"schemas"`
}

type CalendarEventTime struct {
	Date      string `json:"date"`
	Timestamp string `json:"timestamp"`
	Timezone  string `json:"timezone"`
}

type Vchat struct {
	MeetingUrl string `json:"meeting_url"`
}

type Location struct {
	Name      string  `json:"name"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CalendarEventReminder struct {
	Minutes int64 `json:"minutes"`
}

type CalendarEventSchema struct {
	UiName   string `json:"ui_name"`
	UiStatus string `json:"ui_status"`
}

type CalendarEventDataV4 struct {
	EventId         string                  `json:"event_id"`
	Summary         string                  `json:"summary"`
	Description     string                  `json:"description"`
	StartTime       CalendarEventTime       `json:"start_time"`
	EndTime         CalendarEventTime       `json:"end_time"`
	Vchat           Vchat                   `json:"vchat"`
	Visibility      string                  `json:"visibility"`
	AttendeeAbility string                  `json:"attendee_ability"`
	FreeBusyStatus  string                  `json:"free_busy_status"`
	Location        Location                `json:"location"`
	Color           int32                   `json:"color"`
	Reminders       []CalendarEventReminder `json:"reminders"`
	Recurrence      string                  `json:"recurrence"`
	Schemas         []CalendarEventSchema   `json:"schemas"`
}

type CalendarEventInfoV4Resp struct {
	CommonVo
	Data CalendarEventV4 `json:"data"`
}

type CalendarEventV4 struct {
	Event CalendarEventDataV4 `json:"event"`
}

type GetCalendarEventListV4Resp struct {
	CommonVo
	Data GetCalendarEventListV4 `json:"data"`
}

type GetCalendarEventListV4 struct {
	PageToken string                `json:"page_token"`
	SyncToken string                `json:"sync_token"`
	HasMore   bool                  `json:"has_more"`
	Items     []CalendarEventDataV4 `json:"items"`
}

type AddCalendarEventAttendeesV4Req struct {
	Attendees []AttendeesV4 `json:"attendees"`
}

type AddCalendarEventAttendeesV4Resp struct {
	CommonVo
	Data AddCalendarEventAttendeesDataV4 `json:"data"`
}

type AddCalendarEventAttendeesDataV4 struct {
	Attendees []AttendeesV4 `json:"attendees"`
}

type AttendeesV4 struct {
	Type        string `json:"type"`
	IsOptional  bool   `json:"is_optional"`
	UserId      string `json:"user_id"`
	AttendeeId  string `json:"attendee_id"`
	RsvpStatus  string `json:"rsvp_status"`
	IsOrganizer bool   `json:"is_organizer"`
	IsExternal  bool   `json:"is_external"`
	DisplayName string `json:"display_name"`
}

type DeleteCalendarEventAttendeesV4Req struct {
	AttendeeIds []string `json:"attendee_ids"`
}

type GetCalendarEventAttendeesV4Resp struct {
	CommonVo
	Data GetCalendarEventAttendeesDataV4 `json:"data"`
}

type GetCalendarEventAttendeesDataV4 struct {
	HasMore   bool          `json:"has_more"`
	PageToken string        `json:"page_token"`
	Items     []AttendeesV4 `json:"items"`
}

type UpdateCalendarEventV4Req struct {
	Summary         *string                  `json:"summary"`
	Description     *string                  `json:"description"`
	StartTime       *CalendarEventTime       `json:"start_time"`
	EndTime         *CalendarEventTime       `json:"end_time"`
	Vchat           *Vchat                   `json:"vchat"`
	Visibility      *string                  `json:"visibility"`       //default(忙碌),public(公开),private（私密）
	AttendeeAbility *string                  `json:"attendee_ability"` //none,can_see_other,can_invite_others,can_modify_event
	FreeBusyStatus  *string                  `json:"free_busy_status"`
	Location        *Location                `json:"location"`
	Color           *int32                   `json:"color"`
	Reminders       *[]CalendarEventReminder `json:"reminders"`
	Recurrence      *string                  `json:"recurrence"`
	Schemas         *[]CalendarEventSchema   `json:"schemas"`
}

type AddCalendarAclV4Req struct {
	Role  string   `json:"role"` //unknown(未知)，free_busy_reader(游客)，reader(订阅者)，writer(编辑者)，owner(管理员)
	Scope AclScope `json:"scope"`
}

type AclScope struct {
	Type   string `json:"type"` //user(用户)
	UserId string `json:"user_id"`
}

type AddCalendarAclV4Resp struct {
	CommonVo
	Data AddCalendarAclRespData `json:"data"`
}

type AddCalendarAclRespData struct {
	Role  string   `json:"role"`
	AclId string   `json:"acl_id"`
	Scope AclScope `json:"scope"`
}

type GetCalendarAclListV4Resp struct {
	CommonVo
	Data GetCalendarAclListV4Data `json:"data"`
}

type GetCalendarAclListV4Data struct {
	Acls      []AddCalendarAclRespData `json:"acls"`
	HasMore   bool                     `json:"has_more"`
	PageToken string                   `json:"page_token"`
}
