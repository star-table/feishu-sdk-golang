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

type CreateCalendarEventV4Req struct {
	Summary         string                  `json:"summary"`
	Description     string                  `json:"description"`
	StartTime       CalendarEventTime       `json:"start_time"`
	EndTime         CalendarEventTime       `json:"end_time"`
	Vchat           Vchat                   `json:"vchat"`
	Visibility      string                  `json:"visibility"`       //default(忙碌),public(公开),private（私密）
	AttendeeAbility string                  `json:"attendee_ability"` //none,can_see_other,can_invite_others,can_modify_event
	FreeBusyStatus  string                  `json:"free_busy_status"`
	Location        Location                `json:"location"`
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
