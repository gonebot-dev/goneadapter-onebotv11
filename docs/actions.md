# 行为

[全部行为](#sendprivatemessage)：
	- [发送私聊消息](#sendprivatemessage)
	- [发送群聊消息](#sendgroupmessage)
	- [撤回消息](#deletemessage)
	- [获取消息](#getmessage)
	- [获取合并转发消息](#getforwardmessage)
	- [发送好友赞](#sendlike)
	- [踢出群成员](#setgroupkick)
	- [禁言群成员](#setgroupban)
	- [全体禁言](#setgroupwholeban)
	- [设置管理员](#setgroupadmin)
	- [设置群名片](#setgroupcard)
	- [设置群名](#setgroupname)
	- [退出（或解散）群聊](#setgroupleave)
	- [设置专属头衔](#setgroupspecialtitle)
	- [处理加好友请求](#setfriendaddrequest)
	- [处理加群请求／邀请](#setgroupaddrequest)
	- [获取登录信息](#getlogininfo)
	- [获取陌生人信息](#getstrangerinfo)
	- [获取好友列表](#getfriendlist)
	- [获取群聊信息](#getgroupinfo)
	- [获取群成员列表](#getgrouplist)
	- [获取群成员信息](#getgroupmemberinfo)
	- [获取群荣誉信息](#getgrouphonorinfo)
	- [获取 Cookies](#getcookies)
	- [获取 CSRF Token](#getcsrftoken)
	- [获取 Credentials](#getcredentials)
	- [获取语音](#getrecord)
	- [获取图片](#getimage)
	- [检查是否可以发送图片](#canuploadimage)
	- [检查是否可以发送语音](#canuploadrecord)
	- [获取机器人信息](#getstatus)
	- [获取版本信息](#getversioninfo)
	- [设置重启定时任务](#setrestart)
	- [清理缓存](#cleancache)

**OneBotV11 适配器实现了所有 OneBotV11 协议支持的行为。所有的结构体定义如下：**

### **SendPrivateMessage:**
根据给出的信息，机器人会尝试发送一条私聊消息给指定的账号。
```go
type SendPrivateMessage struct {
	UserID  int64                    `json:"user_id"`
	Message []message.MessageSegment `json:"message"`
	// True if message is string, false otherwise
	AutoEscape bool `json:"auto_escape"`
}
```
### **SendGroupMessage**
根据给出的信息，机器人会尝试发送一条群聊消息给指定的群。
```go
type SendGroupMessage struct {
	GroupID int64                    `json:"group_id"`
	Message []message.MessageSegment `json:"message"`
	// True if message is string, false otherwise
	AutoEscape bool `json:"auto_escape"`
}
```
### **DeleteMessage**
机器人会尝试撤回给定信息中的消息 ID
```go
type DeleteMessage struct {
	MessageID int64 `json:"message_id"`
}
```
### **GetMessage**
机器人会尝试获取给定信息中的消息 ID
```go
type GetMessage struct {
	MessageID int64 `json:"message_id"`
}
```
&emsp;&emsp;同时，我们还定义了 GetMessage 的返回结果类型：
```go
type GetMessageResult struct {
	Time        int64                    `json:"time"`
	MessageType string                   `json:"message_type"`
	MessageID   int64                    `json:"message_id"`
	RealID      int64                    `json:"real_id"`
	Sender      SenderObject             `json:"sender"`
	Message     []message.MessageSegment `json:"message"`
}
```
### **GetForwardMessage**
机器人会尝试获取指定 ID 的合并转发消息内容
```go
type GetForwardMessage struct {
	ID string `json:"id"`
}
```
&emsp;&emsp;同时，我们定义了其返回类型：
```go
type GetForwardMessageResult struct {
	Message []message.MessageSegment `json:"Message"`
}
```
### **SendLike**
机器人会尝试发送 Times 次好友赞，为 UserID 指定的好友点赞，每天最多 10 次
```go
type SendLike struct {
	UserID int64 `json:"user_id"`
	Times  int64 `json:"times"`
}
```
### **SetGroupKick**
机器人会尝试从指定的群中踢出指定的成员
```go
type SetGroupKick struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	// Whether to always reject the add request of the user
	RejectAddRequest bool `json:"reject_add_request"`
}
```
### **SetGroupBan**
机器人会尝试将指定的成员禁言，时间单位为秒
```go
type SetGroupBan struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	// Use second as unit, value 0 means cancel the ban
	Duration int64 `json:"duration"`
}
```
### **SetGroupWholeBan**
机器人会尝试开启指定群聊的全体禁言
```go
type SetGroupWholeBan struct {
	GroupID int64 `json:"group_id"`
	// True to ban, false to cancel
	Enable bool `json:"enable"`
}
```
### **SetGroupAdmin**
机器人会尝试将指定的成员设置为管理员
```go
type SetGroupAdmin struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	// True to set as admin, false to cancel
	Enable bool `json:"enable"`
}
```
### **SetGroupCard**
机器人会尝试设置指定群聊的群名片
```go
type SetGroupCard struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	// Sets the group card, leave it empty to remove the group card
	Card string `json:"card"`
}
```
### **SetGroupName**
机器人会尝试设置指定群聊的群名
```go
type SetGroupName struct {
	GroupID int64  `json:"group_id"`
	Name    string `json:"name"`
}
```
### **SetGroupLeave**
机器人会尝试退出指定的群聊，如果机器人是群主，还可以设置是否直接解散群聊
```go
type SetGroupLeave struct {
	GroupID int64 `json:"group_id"`
	// Dissolve the group if true(and the bot is the owner)
	IsDismiss bool `json:"is_dismiss"`
}
```
### **SetGroupSpecialTitle**
机器人会尝试设置指定群聊的群成员专属头衔
```go
type SetGroupSpecialTitle struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	// Leave it empty to remove the special title
	SpecialTitle string `json:"special_title"`
	// Expiration time, use second as unit, value -1 means no expiration, seems not functioning
	Duration int64 `json:"duration"`
}
```
### **SetFriendAddRequest**
机器人会尝试处理加好友请求
```go
type SetFriendAddRequest struct {
	Flag string `json:"flag"`
	// True to approve, false to refuse
	Approve bool `json:"approve"`
	// Optional, sets the friend remark, only works when approve is true
	Remark string `json:"remark"`
}
```
### **SetGroupAddRequest**
机器人会尝试处理加群邀请/请求
```go
type SetGroupAddRequest struct {
	Flag string `json:"flag"`
	// "add" or "invite", which means the type of the request is add or invite
	SubType string `json:"sub_type"`
	// True to approve, false to refuse
	Approve bool `json:"approve"`
	// Optional, reason for rejecting the request, only works when approve is false
	Reason string `json:"reason"`
}
```
### **GetLoginInfo**
尝试获取机器人的登录信息，为了和其他类型区分开，添加了一个 int8 类型的字段，使用时具体值随意指定即可。
```go
type GetLoginInfo struct {
	GetLoginInfo int8 `json:"-"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetLoginInfoResult struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
}
```
### **GetStrangerInfo**
尝试获取指定陌生用户的信息
```go
type GetStrangerInfo struct {
	UserID int64 `json:"user_id"`
	// Whether to use cache(using cache leads to faster response, but may be outdated)
	NoCache bool `json:"no_cache"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetStrangerInfoResult struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int64  `json:"age"`
}
```
### **GetFriendList**
尝试获取机器人的好友列表，为了和其他类型区分开，添加了一个 int8 类型的字段，使用时具体值随意指定即可。
```go
type GetFriendList struct {
	GetFriendList int8 `json:"-"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetFriendListResult struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
}
```
### **GetGroupInfo**
尝试获取指定群聊的信息
```go
type GetGroupInfo struct {
	GroupID int64 `json:"group_id"`
	// Whether to use cache(using cache leads to faster response, but may be outdated)
	NoCache bool `json:"no_cache"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetGroupInfoResult struct {
	GroupID        int64  `json:"group_id"`
	GroupName      string `json:"group_name"`
	MemberCount    int64  `json:"member_count"`
	MaxMemberCount int64  `json:"max_member_count"`
}
```
### **GetGroupList**
尝试获取机器人的群聊列表，为了和其他类型区分开，添加了一个 int8 类型的字段，使用时具体值随意指定即可。
```go
type GetGroupList struct {
	GetGroupList int8 `json:"-"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetGroupListResult []GetGroupInfoResult
```
### **GetGroupMemberInfo**
尝试获取指定群聊中的指定群成员的信息
```go
type GetGroupMemberInfo struct {
	GroupID int64 `json:"group_id"`
	UserID  int64 `json:"user_id"`
	// Whether to use cache(using cache leads to faster response, but may be outdated)
	NoCache bool `json:"no_cache"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetGroupMemberInfoResult struct {
	// User info
	GroupID  int64  `json:"group_id"`
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Card     string `json:"card"`
	Sex      string `json:"sex"`
	Age      int64  `json:"age"`
	Area     string `json:"area"`
	// Group message info
	JoinTime     int64  `json:"join_time"`
	LastSentTime int64  `json:"last_sent_time"`
	Level        string `json:"level"`
	// Group user info
	// "owner", "admin" or "member"
	Role            string `json:"role"`
	Unfriendly      bool   `json:"unfriendly"`
	Title           string `json:"title"`
	TitleExpireTime int64  `json:"title_expire_time"`
	CardChangeable  bool   `json:"card_changeable"`
}
```
### **GetGroupMemberList**
尝试获取指定群聊中的群成员列表
```go
type GetGroupMemberList struct {
	GroupID int64 `json:"group_id"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetGroupMemberListResult []GetGroupMemberInfoResult
```
### **GetGroupHonorInfo**
尝试获取指定群聊的群荣誉信息
```go
type GetGroupHonorInfo struct {
	GroupID int64 `json:"group_id"`
	Type    string
}
```
&emsp;&emsp;同时，定义了其返回值和其中的子类型：
```go
type GroupTalkativeInfo struct {
	UserID   int64 `json:"user_id"`
	Nickname int64 `json:"nickname"`
	// Url of the avatar
	Avatar string `json:"avatar"`
	// How long it have lasted
	DayCount int64 `json:"day_count"`
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
```
### **GetCookies**
获取指定域名的 Cookies
```go
type GetCookies struct {
	Domain string `json:"domain"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetCookiesResult struct {
	Cookies string `json:"cookies"`
}
```
### **GetCsrfToken**
获取 CSRF Token
```go
type GetCsrfToken struct {
	GetCsrfToken int8 `json:"-"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetCsrfTokenResult struct {
	CsrfToken int64 `json:"csrf_token"`
}
```
### **GetCredentials**
上面两个接口的合并
```go
type GetCredentials struct {
	Domain string `json:"domain"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetCredentialsResult struct {
	Cookies   string `json:"cookies"`
	CsrfToken int64  `json:"csrf_token"`
}
```
### **GetRecord**
尝试获取指定消息的语音文件（协议服务端需安装 ffmpeg）
```go
// To use this, you may need to get ffmpeg installed(the app you are using must support this)
type GetRecord struct {
	// File name received
	File string `json:"file"`
	// Target format
	OutFormat string `json:"out_format"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetRecordResult struct {
	// Local file path
	File string `json:"file"`
}
```
### **GetImage**
尝试获取图片信息
```go
type GetImage struct {
	// File name received
	File string `json:"file"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetImageResult struct {
	// Local file path
	File string `json:"file"`
}
```
### **CanSendImage**
获取是否可以发送图片
```go
type CanSendImage struct {
	CanSendImage int8 `json:"-"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type CanSendImageResult struct {
	// Whether the bot can send image
	Yes bool `json:"yes"`
}
```
### **CanSendRecord**
获取是否可以发送语音
```go
type CanSendRecord struct {
	CanSendRecord int8 `json:"-"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type CanSendRecordResult struct {
	// Whether the bot can send record
	Yes bool `json:"yes"`
}
```
### **GetStatus**
获取机器人状态
```go
type GetStatus struct {
	GetStatus int8 `json:"-"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetStatusResult struct {
	// Will be null if can't query(I wonder how to express this)
	Online bool `json:"online"`
	// Whether everything is functioning well
	Good bool `json:"good"`
}
```
### **GetVersionInfo**
获取版本信息
```go
type GetVersionInfo struct {
	GetVersionInfo int8 `json:"-"`
}
```
&emsp;&emsp;同时，定义了其返回值：
```go
type GetVersionInfoResult struct {
	// Name of the app
	AppName string `json:"app_name"`
	// Version of the app
	AppVersion string `json:"app_version"`
	// OneBot version
	ProtocolVersion string `json:"protocol_version"`
}
```
### **SetRestart**
设置重启定时任务
```go
type SetRestart struct {
	// Milliseconds before restart, if cannot restart normally, try set it to about 2000
	Delay int64 `json:"delay"`
}
```
### **CleanCache**
清理缓存
```go
type CleanCache struct {
	CleanCache int8 `json:"-"`
}
```

另外，如果一些行为没有返回值，它们应当返回这个类型：
```go
// No result should return this
type EmptyResult struct{}
```
