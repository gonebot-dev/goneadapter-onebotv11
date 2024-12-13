# 事件

[前置类型](#payloadmessage)：
- [消息段](#payloadmessage)
- [消息发送者](#senderobject)
- [文件信息](#fileobject)
[事件列表](#privatemessage)：
- [私聊消息](#privatemessage)
- [群聊消息](#groupmessage)
- [群文件上传](#groupfileupload)
- [群管理员变动](#adminchange)
- [群成员减少](#groupmemberdecrease)
- [群成员增加](#groupmemberincrease)
- [群禁言](#groupban)
- [好友添加](#friendadd)
- [群消息撤回](#grouprecall)
- [好友消息撤回](#friendrecall)
- [群内戳一戳](#grouppoke)
- [好友戳一戳](#friendpoke)
- [红包运气王](#redpacketluckyking)
- [群荣誉变更](#grouphonorchange)
- [好友添加请求](#friendrequest)
- [群添加请求](#grouprequest)

**OneBotV11 适配器实现了所有 OneBotV11 协议支持的事件。所有的结构体定义如下：**

首先，我们定义了一些子类型和辅助类型：
### **PayloadMessage:**
消息段中转类型，用于表示协议服务器返回的消息段内容。
```go
type PayloadMessage struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
```
### **SenderObject:**
消息发送者的具体信息，里面定义的字段根据协议服务器的不同，不一定存在：
```go
// You shouldn't assume that every field in this struct does exists in actual json string
type SenderObject struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int64  `json:"age"`
	// If is a group message, SenderObject may also contain the following fields
	Card  string `json:"card"`
	Area  string `json:"area"`
	Level string `json:"level"`
	// "owner", "admin" or "member"
	Role  string `json:"Role"`
	Title string `json:"Title"`
}
```
### **FileObject**
表示文件信息
```go
type FileObject struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Busid int64  `json:"busid"`
}
```
**之后是所有的消息类型：**
### **PrivateMessage**
私聊消息事件
```go
type PrivateMessage struct {
	Time        int64  `json:"time"`
	SelfID      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	// "friend", "group" or "other"
	SubType    string           `json:"sub_type"`
	MessageID  int64            `json:"message_id"`
	UserID     int64            `json:"user_id"`
	Message    []PayloadMessage `json:"message"`
	RawMessage string           `json:"raw_message"`
	Font       int64            `json:"font"`
	Sender     SenderObject     `json:"sender"`
}
```
### **GroupMessage**
群聊消息事件
```go
type GroupMessage struct {
	Time        int64  `json:"time"`
	SelfID      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	// "normal", "notice" or "active"
	SubType    string           `json:"sub_type"`
	MessageID  int64            `json:"message_id"`
	GroupID    int64            `json:"group_id"`
	UserID     int64            `json:"user_id"`
	Message    []PayloadMessage `json:"message"`
	RawMessage string           `json:"raw_message"`
	Font       int64            `json:"font"`
	Sender     SenderObject     `json:"sender"`
}
```
### **GroupFileUpload**
群聊文件上传事件
```go
type GroupFileUpload struct {
	message.MessageType
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "group_upload"
	NoticeType string `json:"notice_type"`
	GroupID    int64  `json:"group_id"`
	// Sender ID
	UserID int64      `json:"user_id"`
	File   FileObject `json:"file"`
}
```
### **AdminChange**
群管理员变动事件
```go
type AdminChange struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "group_admin"
	NoticeType string `json:"notice_type"`
	// "set" or "unset"
	SubType string `json:"sub_type"`
	GroupID int64  `json:"group_id"`
	// Victim ID
	UserID int64 `json:"user_id"`
}
```
### **GroupMemberDecrease**
群成员减少事件
```go
type GroupMemberDecrease struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "group_decrease"
	NoticeType string `json:"notice_type"`
	// "leave" or "kick" or "kick_me"
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	// Victim ID
	UserID int64 `json:"user_id"`
}
```
### **GroupMemberIncrease**
群成员增加事件
```go
type GroupMemberIncrease struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "group_increase"
	NoticeType string `json:"notice_type"`
	// "approve" or "invite"
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	// Joiner ID
	UserID int64 `json:"user_id"`
}
```
### **GroupBan**
群禁言事件
```go
type GroupBan struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "group_ban"
	NoticeType string `json:"notice_type"`
	// "ban" or "lift_ban"
	SubType    string `json:"sub_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	// Victim ID
	UserID   int64 `json:"user_id"`
	Duration int64 `json:"duration"`
}
```
### **FriendAdd**
好友添加事件
```go
type FriendAdd struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "friend_add"
	NoticeType string `json:"notice_type"`
	UserID     int64  `json:"user_id"`
}
```
### **GroupRecall**
群消息撤回事件
```go
type GroupRecall struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "group_recall"
	NoticeType string `json:"notice_type"`
	GroupID    int64  `json:"group_id"`
	OperatorID int64  `json:"operator_id"`
	// Recaller ID
	UserID int64 `json:"user_id"`
	// Message ID
	MessageID int64 `json:"message_id"`
}
```
### **FriendRecall**
好友消息撤回事件
```go
type FriendRecall struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "friend_recall"
	NoticeType string `json:"notice_type"`
	// Recaller ID
	UserID    int64 `json:"user_id"`
	MessageID int64 `json:"message_id"`
}
```
### **GroupPoke**
群内戳一戳事件
```go
type GroupPoke struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "notify"
	NoticeType string `json:"notice_type"`
	GroupID    int64  `json:"group_id"`
	// "poke"
	SubType string `json:"sub_type"`
	// Poker ID
	UserID int64 `json:"user_id"`
	// Pokee ID
	TargetID int64 `json:"target_id"`
}
```
### **FriendPoke**
好友戳一戳事件
```go
type FriendPoke struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "notify"
	NoticeType string `json:"notice_type"`
	// "poke"
	SubType string `json:"sub_type"`
	// Poker ID
	UserID int64 `json:"user_id"`
	// Pokee ID
	TargetID int64 `json:"target_id"`
}
```
### **RedPacketLuckyKing**
红包运气王事件
```go
type RedPacketLuckyKing struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "notify"
	NoticeType string `json:"notice_type"`
	// "lucky_king"
	SubType string `json:"sub_type"`
	// Group ID
	GroupID int64 `json:"group_id"`
	// Red packet sender ID
	UserID int64 `json:"user_id"`
	// Lucky king ID
	TargetID int64 `json:"target_id"`
}
```
### **GroupHonorChange**
群荣誉变更事件
```go
type GroupHonorChange struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "notify"
	NoticeType string `json:"notice_type"`
	// "honor"
	SubType string `json:"sub_type"`
	// Group ID
	GroupID int64 `json:"group_id"`
	// Honoree ID
	UserID int64 `json:"user_id"`
	// "talkative", "performer" or "emotion"
	HonorType string `json:"honor_type"`
}
```
### **FriendRequest**
好友添加请求事件
```go
type FriendRequest struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "friend"
	RequestType string `json:"request_type"`
	// New friend ID
	UserId int64 `json:"user_id"`
	// Request comment
	Comment string `json:"comment"`
	// Flag for handling this request
	Flag string `json:"flag"`
}
```
### **GroupRequest**
群添加请求事件
```go
type FriendRequest struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "friend"
	RequestType string `json:"request_type"`
	// New friend ID
	UserId int64 `json:"user_id"`
	// Request comment
	Comment string `json:"comment"`
	// Flag for handling this request
	Flag string `json:"flag"`
}
```
