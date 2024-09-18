package onebotv11

import (
	"fmt"
	"reflect"

	"github.com/gonebot-dev/gonebot/message"
)

// Message event

// You shouldn't assume that every field in this struct does exists in actual json string
type SenderObject struct {
	UserID   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
	// If is a group message, SenderObject may also contain the following fields
	Card  string `json:"card"`
	Area  string `json:"area"`
	Level string `json:"level"`
	// "owner", "admin" or "member"
	Role  string `json:"Role"`
	Title string `json:"Title"`
}

type PrivateMessage struct {
	Time        int64  `json:"time"`
	SelfID      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	// "friend", "group" or "other"
	SubType    string           `json:"sub_type"`
	MessageID  int32            `json:"message_id"`
	UserID     int64            `json:"user_id"`
	Message    []PayloadMessage `json:"message"`
	RawMessage string           `json:"raw_message"`
	Font       int32            `json:"font"`
	Sender     SenderObject     `json:"sender"`
}

type GroupMessage struct {
	Time        int64  `json:"time"`
	SelfID      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	// "normal", "notice" or "active"
	SubType    string           `json:"sub_type"`
	MessageID  int32            `json:"message_id"`
	GroupID    int64            `json:"group_id"`
	UserID     int64            `json:"user_id"`
	Message    []PayloadMessage `json:"message"`
	RawMessage string           `json:"raw_message"`
	Font       int32            `json:"font"`
	Sender     SenderObject     `json:"sender"`
}

// Notice event, post type will always be "notice"

type FileObject struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Busid int64  `json:"busid"`
}

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

func (fileUpload GroupFileUpload) Matcher(typeName, adapterName string) bool {
	return typeName == "group_upload" && adapterName == OneBotV11.Name
}

func (fileUpload GroupFileUpload) ToRawText(msg message.MessageSegment) string {
	result := fileUpload.Deserialize(msg.Data, reflect.TypeOf(fileUpload)).(GroupFileUpload)
	return fmt.Sprintf("[OnebotV11:group_upload,time=%d,self_id=%d,group_id=%d,user_id=%d]", result.Time, result.SelfID, result.GroupID, result.UserID)
}

type AdminChange struct {
	message.MessageType
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

func (adminChanage AdminChange) Matcher(typeName, adapterName string) bool {
	return typeName == "group_admin" && adapterName == OneBotV11.Name
}

func (adminChanage AdminChange) ToRawText(msg message.MessageSegment) string {
	result := adminChanage.Deserialize(msg.Data, reflect.TypeOf(adminChanage)).(AdminChange)
	return fmt.Sprintf("[OnebotV11:group_admin,time=%d,self_id=%d,group_id=%d,user_id=%d]", result.Time, result.SelfID, result.GroupID, result.UserID)
}

type GroupMemberDecrease struct {
	message.MessageType
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

func (groupMemberDecrease GroupMemberDecrease) Matcher(typeName, adapterName string) bool {
	return typeName == "group_decrease" && adapterName == OneBotV11.Name
}

func (groupMemberDecrease GroupMemberDecrease) ToRawText(msg message.MessageSegment) string {
	result := groupMemberDecrease.Deserialize(msg.Data, reflect.TypeOf(groupMemberDecrease)).(GroupMemberDecrease)
	return fmt.Sprintf("[OnebotV11:group_decrease,time=%d,self_id=%d,group_id=%d,user_id=%d]", result.Time, result.SelfID, result.GroupID, result.UserID)
}

type GroupMemberIncrease struct {
	message.MessageType
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

func (groupMemberIncrease GroupMemberIncrease) Matcher(typeName, adapterName string) bool {
	return typeName == "group_increase" && adapterName == OneBotV11.Name
}

func (groupMemberIncrease GroupMemberIncrease) ToRawText(msg message.MessageSegment) string {
	result := groupMemberIncrease.Deserialize(msg.Data, reflect.TypeOf(groupMemberIncrease)).(GroupMemberIncrease)
	return fmt.Sprintf("[OnebotV11:group_increase,time=%d,self_id=%d,group_id=%d,user_id=%d]", result.Time, result.SelfID, result.GroupID, result.UserID)
}

type GroupBan struct {
	message.MessageType
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

func (groupBan GroupBan) Matcher(typeName, adapterName string) bool {
	return typeName == "group_ban" && adapterName == OneBotV11.Name
}

func (groupBan GroupBan) ToRawText(msg message.MessageSegment) string {
	result := groupBan.Deserialize(msg.Data, reflect.TypeOf(groupBan)).(GroupBan)
	return fmt.Sprintf("[OnebotV11:group_ban,time=%d,self_id=%d,group_id=%d,user_id=%d]", result.Time, result.SelfID, result.GroupID, result.UserID)
}

type FriendAdd struct {
	message.MessageType
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "friend_add"
	NoticeType string `json:"notice_type"`
	UserID     int64  `json:"user_id"`
}

func (friendAdd FriendAdd) Matcher(typeName, adapterName string) bool {
	return typeName == "friend_add" && adapterName == OneBotV11.Name
}

func (friendAdd FriendAdd) ToRawText(msg message.MessageSegment) string {
	result := friendAdd.Deserialize(msg.Data, reflect.TypeOf(friendAdd)).(FriendAdd)
	return fmt.Sprintf("[OnebotV11:friend_add,time=%d,self_id=%d,user_id=%d]", result.Time, result.SelfID, result.UserID)
}

type GroupRecall struct {
	message.MessageType
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

func (groupRecall GroupRecall) Matcher(typeName, adapterName string) bool {
	return typeName == "group_recall" && adapterName == OneBotV11.Name
}

func (groupRecall GroupRecall) ToRawText(msg message.MessageSegment) string {
	result := groupRecall.Deserialize(msg.Data, reflect.TypeOf(groupRecall)).(GroupRecall)
	return fmt.Sprintf("[OnebotV11:group_recall,time=%d,self_id=%d,group_id=%d,user_id=%d,message_id=%d]", result.Time, result.SelfID, result.GroupID, result.UserID, result.MessageID)
}

type FriendRecall struct {
	message.MessageType
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "friend_recall"
	NoticeType string `json:"notice_type"`
	// Recaller ID
	UserID    int64 `json:"user_id"`
	MessageID int64 `json:"message_id"`
}

func (friendRecall FriendRecall) Matcher(typeName, adapterName string) bool {
	return typeName == "friend_recall" && adapterName == OneBotV11.Name
}

func (friendRecall FriendRecall) ToRawText(msg message.MessageSegment) string {
	result := friendRecall.Deserialize(msg.Data, reflect.TypeOf(friendRecall)).(FriendRecall)
	return fmt.Sprintf("[OnebotV11:friend_recall,time=%d,self_id=%d,user_id=%d,message_id=%d]", result.Time, result.SelfID, result.UserID, result.MessageID)
}

type GroupPoke struct {
	message.MessageType
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

func (groupPoke GroupPoke) Matcher(typeName, adapterName string) bool {
	return typeName == "poke" && adapterName == OneBotV11.Name
}

func (groupPoke GroupPoke) ToRawText(msg message.MessageSegment) string {
	result := groupPoke.Deserialize(msg.Data, reflect.TypeOf(groupPoke)).(GroupPoke)
	return fmt.Sprintf("[OnebotV11:poke,time=%d,self_id=%d,group_id=%d,user_id=%d,target_id=%d]", result.Time, result.SelfID, result.GroupID, result.UserID, result.TargetID)
}

type RedPacketLuckyKing struct {
	message.MessageType
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

func (redPacketLuckyKing RedPacketLuckyKing) Matcher(typeName, adapterName string) bool {
	return typeName == "lucky_king" && adapterName == OneBotV11.Name
}

func (redPacketLuckyKing RedPacketLuckyKing) ToRawText(msg message.MessageSegment) string {
	result := redPacketLuckyKing.Deserialize(msg.Data, reflect.TypeOf(redPacketLuckyKing)).(RedPacketLuckyKing)
	return fmt.Sprintf("[OnebotV11:lucky_king,time=%d,self_id=%d,group_id=%d,user_id=%d,target_id=%d]", result.Time, result.SelfID, result.GroupID, result.UserID, result.TargetID)
}

type GroupHonorChange struct {
	message.MessageType
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

func (groupHonorChange GroupHonorChange) Matcher(typeName, adapterName string) bool {
	return typeName == "honor" && adapterName == OneBotV11.Name
}

func (groupHonorChange GroupHonorChange) ToRawText(msg message.MessageSegment) string {
	result := groupHonorChange.Deserialize(msg.Data, reflect.TypeOf(groupHonorChange)).(GroupHonorChange)
	return fmt.Sprintf("[OnebotV11:honor,time=%d,self_id=%d,group_id=%d,user_id=%d,honor_type=%s]", result.Time, result.SelfID, result.GroupID, result.UserID, result.HonorType)
}

// Request event, post type will always be "request"

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

type GroupRequest struct {
	Time     int64  `json:"time"`
	SelfID   int64  `json:"self_id"`
	PostType string `json:"post_type"`
	// "group"
	RequestType string `json:"request_type"`
	// "add" or "invite"
	SubType string `json:"sub_type"`
	GroupId int64  `json:"group_id"`
	// Inviter ID
	UserId int64 `json:"user_id"`
	// Request comment
	Comment string `json:"comment"`
	// Flag for handling this request
	Flag string `json:"flag"`
}
