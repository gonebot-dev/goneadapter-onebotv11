package onebotv11

type ActionPayload struct {
	Action string `json:"action"`
	Params any    `json:"params"`
}

type PayloadMessage struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}

type SendPrivateMessage struct {
	UserID  int64            `json:"user_id"`
	Message []PayloadMessage `json:"message"`
	// True if message is string, false otherwise
	AutoEscape bool `json:"auto_escape"`
}

type SendGroupMessage struct {
	GroupID int64            `json:"group_id"`
	Message []PayloadMessage `json:"message"`
	// True if message is string, false otherwise
	AutoEscape bool `json:"auto_escape"`
}

type DeleteMessage struct {
	MessageID int32 `json:"message_id"`
}

type GetMessage struct {
	MessageID int32 `json:"message_id"`
}

type GetMessageResult struct {
	Time        int32            `json:"time"`
	MessageType string           `json:"message_type"`
	MessageID   int32            `json:"message_id"`
	RealID      int32            `json:"real_id"`
	Sender      SenderObject     `json:"sender"`
	Message     []PayloadMessage `json:"message"`
}

type GetForwardMessage struct {
	ID string `json:"id"`
}

type GetForwardMessageResult struct {
	Message []PayloadMessage `json:"Message"`
}

type SendLike struct {
	UserID int64 `json:"user_id"`
	Times  int64 `json:"times"`
}

type SetGroupKick struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	// Whether to always reject the add request of the user
	RejectAddRequest bool `json:"reject_add_request"`
}

type SetGroupBan struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	// Use second as unit, value 0 means cancel the ban
	Duration int64 `json:"duration"`
}

type SetGroupWholeBan struct {
	GroupID int64 `json:"group_id"`
	// True to ban, false to cancel
	Enable bool `json:"enable"`
}

type SetGroupAdmin struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	// True to set as admin, false to cancel
	Enable bool `json:"enable"`
}

type SetGroupCard struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	// Sets the group card, leave it empty to remove the group card
	Card string `json:"card"`
}

type SetGroupName struct {
	GroupID int64  `json:"group_id"`
	Name    string `json:"name"`
}

type SetGroupLeave struct {
	GroupID int64 `json:"group_id"`
	// Dissolve the group if true(and the bot is the owner)
	IsDismiss bool `json:"is_dismiss"`
}

type SetGroupSpecialTitle struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	// Leave it empty to remove the special title
	SpecialTitle string `json:"special_title"`
	// Expiration time, use second as unit, value -1 means no expiration, seems not functioning
	Duration int64 `json:"duration"`
}

type SetFriendAddRequest struct {
	Flag string `json:"flag"`
	// True to approve, false to refuse
	Approve bool `json:"approve"`
	// Optional, sets the friend remark, only works when approve is true
	Remark string `json:"remark"`
}

type SetGroupAddRequest struct {
	Flag string `json:"flag"`
	// "add" or "invite", which means the type of the request is add or invite
	SubType string `json:"sub_type"`
	// True to approve, false to refuse
	Approve bool `json:"approve"`
	// Optional, reason for rejecting the request, only works when approve is false
	Reason string `json:"reason"`
}

type GetLoginInfo struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
}

type GetStrangerInfo struct {
	UserID int64 `json:"user_id"`
	// Whether to use cache(using cache leads to faster response, but may be outdated)
	NoCache bool `json:"no_cache"`
}

type GetFriendList struct {
	GetFriendList int8 `json:"-"`
}

type GetFriendListResult struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
}

type GetGroupInfo struct {
	GroupID int64 `json:"group_id"`
	// Whether to use cache(using cache leads to faster response, but may be outdated)
	NoCache bool `json:"no_cache"`
}

type GetGroupInfoResult struct {
	GroupID        int64  `json:"group_id"`
	GroupName      string `json:"group_name"`
	MemberCount    int32  `json:"member_count"`
	MaxMemberCount int32  `json:"max_member_count"`
}

type GetGroupList struct {
	GetGroupList int8 `json:"-"`
}

type GetGroupListResult []GetGroupInfoResult

type GetGroupMemberInfo struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	// Whether to use cache(using cache leads to faster response, but may be outdated)
	NoCache bool `json:"no_cache"`
}

type GetGroupMemberInfoResult struct {
	// User info
	GroupID  int64  `json:"group_id"`
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Card     string `json:"card"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
	Area     string `json:"area"`
	// Group message info
	JoinTime     int32  `json:"join_time"`
	LastSentTime int32  `json:"last_sent_time"`
	Level        string `json:"level"`
	// Group user info
	// "owner", "admin" or "member"
	Role            string `json:"role"`
	Unfriendly      bool   `json:"unfriendly"`
	Title           string `json:"title"`
	TitleExpireTime int32  `json:"title_expire_time"`
	CardChangeable  bool   `json:"card_changeable"`
}

type GetGroupMemberList struct {
	GroupID int64 `json:"group_id"`
}

type GetGroupMemberListResult []GetGroupMemberInfoResult

type GetGroupHonorInfo struct {
	GroupID int64 `json:"group_id"`
	// Type of the honor you want to know
	//
	// "talkative", "performer", "legend", "strong_newbie", "emotion", or "all" to get all the types
	Type string `json:"type"`
}

type GroupTalkativeInfo struct {
	UserID   int64 `json:"user_id"`
	Nickname int64 `json:"nickname"`
	// Url of the avatar
	Avatar string `json:"avatar"`
	// How long it have lasted
	DayCount int32 `json:"day_count"`
}

type GroupHonorInfo struct {
	UserID   int64 `json:"user_id"`
	Nickname int64 `json:"nickname"`
	// Url of the avatar
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
}

type GetGroupHonorInfoResult struct {
	GroupID          int64              `json:"group_id"`
	CurrentTalkative GroupTalkativeInfo `json:"current_talkative"`
	TalkativeList    []GroupHonorInfo   `json:"talkative_list"`
	PerformerList    []GroupHonorInfo   `json:"performer_list"`
	LegendList       []GroupHonorInfo   `json:"legend_list"`
	StrongNewbieList []GroupHonorInfo   `json:"strong_newbie_list"`
	EmotionList      []GroupHonorInfo   `json:"emotion_list"`
}

type GetCookies struct {
	Domain string `json:"domain"`
}

type GetCookiesResult struct {
	Cookies string `json:"cookies"`
}

type GetCsrfToken struct {
	GetCsrfToken int8 `json:"-"`
}

type GetCsrfTokenResult struct {
	CsrfToken int32 `json:"csrf_token"`
}

type GetCredentials struct {
	Domain string `json:"domain"`
}

type GetCredentialsResult struct {
	Cookies   string `json:"cookies"`
	CsrfToken int32  `json:"csrf_token"`
}

// To use this, you may need to get ffmpeg installed(the app you are using must support this)
type GetRecord struct {
	// File name received
	File string `json:"file"`
	// Target format
	OutFormat string `json:"out_format"`
}

type GetRecordResult struct {
	// Local file path
	File string `json:"file"`
}

type GetImage struct {
	// File name received
	File string `json:"file"`
}

type GetImageResult struct {
	// Local file path
	File string `json:"file"`
}

type CanSendImage struct {
	CanSendImage int8 `json:"-"`
}

type CanSendImageResult struct {
	// Whether the bot can send image
	Yes bool `json:"yes"`
}

type CanSendRecord struct {
	CanSendRecord int8 `json:"-"`
}

type CanSendRecordResult struct {
	// Whether the bot can send record
	Yes bool `json:"yes"`
}

type GetStatus struct {
	GetStatus int8 `json:"-"`
}

type GetStatusResult struct {
	// Will be null if can't query(I wonder how to express this)
	Online bool `json:"online"`
	// Whether everything is functioning well
	Good bool `json:"good"`
}

type GetVersionInfo struct {
	GetVersionInfo int8 `json:"-"`
}

type GetVersionInfoResult struct {
	// Name of the app
	AppName string `json:"app_name"`
	// Version of the app
	AppVersion string `json:"app_version"`
	// OneBot version
	ProtocolVersion string `json:"protocol_version"`
}

type SetRestart struct {
	// Milliseconds before restart, if cannot restart normally, try set it to about 2000
	Delay int32 `json:"delay"`
}

type CleanCache struct {
	CleanCache int8 `json:"-"`
}

// No result should return this
type EmptyResult struct{}
