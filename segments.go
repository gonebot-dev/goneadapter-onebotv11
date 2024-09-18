package onebotv11

import (
	"fmt"
	"reflect"

	"github.com/gonebot-dev/gonebot/message"
)

type FaceType struct {
	message.MessageType
	ID string `json:"id"`
}

func (face FaceType) Matcher(typeName, adapterName string) bool {
	return typeName == "face" && adapterName == OneBotV11.Name
}

func (face FaceType) ToRawText(msg message.MessageSegment) string {
	result := face.Deserialize(msg.Data, reflect.TypeOf(face)).(FaceType)
	return fmt.Sprintf("[OnebotV11:face,id=%s]", result.ID)
}

type AtType struct {
	message.MessageType
	QQ string `json:"qq"`
}

func (at AtType) Matcher(typeName, adapterName string) bool {
	return typeName == "at" && adapterName == OneBotV11.Name
}

func (at AtType) ToRawText(msg message.MessageSegment) string {
	result := at.Deserialize(msg.Data, reflect.TypeOf(at)).(AtType)
	return fmt.Sprintf("[OnebotV11:at,qq=%s]", result.QQ)
}

type RPSType struct {
	message.MessageType
}

func (rsp RPSType) Matcher(typeName, adapterName string) bool {
	return typeName == "rps" && adapterName == OneBotV11.Name
}

func (rsp RPSType) ToRawText(msg message.MessageSegment) string {
	return "[OnebotV11:rps]"
}

type DiceType struct {
	message.MessageType
}

func (dice DiceType) Matcher(typeName, adapterName string) bool {
	return typeName == "dice" && adapterName == OneBotV11.Name
}

func (dice DiceType) ToRawText(msg message.MessageSegment) string {
	return "[OnebotV11:dice]"
}

type ShakeType struct {
	message.MessageType
}

func (shake ShakeType) Matcher(typeName, adapterName string) bool {
	return typeName == "shake" && adapterName == OneBotV11.Name
}

func (shake ShakeType) ToRawText(msg message.MessageSegment) string {
	return "[OnebotV11:shake]"
}

type PokeType struct {
	message.MessageType
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (poke PokeType) Matcher(typeName, adapterName string) bool {
	return typeName == "poke" && adapterName == OneBotV11.Name
}

func (poke PokeType) ToRawText(msg message.MessageSegment) string {
	result := poke.Deserialize(msg.Data, reflect.TypeOf(poke)).(PokeType)
	return fmt.Sprintf("[OnebotV11:poke,qq=%s,type=%s]", result.ID, result.Type)
}

type ShareType struct {
	message.MessageType
	Url string `json:"url"`
}

func (share ShareType) Matcher(typeName, adapterName string) bool {
	return typeName == "share" && adapterName == OneBotV11.Name
}

func (share ShareType) ToRawText(msg message.MessageSegment) string {
	result := share.Deserialize(msg.Data, reflect.TypeOf(share)).(ShareType)
	return fmt.Sprintf("[OnebotV11:share,url=%s]", result.Url)
}

type ContactType struct {
	message.MessageType
	// "qq" or "group"
	Type string `json:"type"`
	ID   string `json:"id"`
}

func (contact ContactType) Matcher(typeName, adapterName string) bool {
	return typeName == "contact" && adapterName == OneBotV11.Name
}

func (contact ContactType) ToRawText(msg message.MessageSegment) string {
	result := contact.Deserialize(msg.Data, reflect.TypeOf(contact)).(ContactType)
	return fmt.Sprintf("[OnebotV11:contact,type=%s,id=%s]", result.Type, result.ID)
}

type LocationType struct {
	message.MessageType
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func (location LocationType) Matcher(typeName, adapterName string) bool {
	return typeName == "location" && adapterName == OneBotV11.Name
}

func (location LocationType) ToRawText(msg message.MessageSegment) string {
	result := location.Deserialize(msg.Data, reflect.TypeOf(location)).(LocationType)
	return fmt.Sprintf("[OnebotV11:location,lat=%s,lon=%s]", result.Lat, result.Lon)
}

type MusicType struct {
	message.MessageType
	// "qq", "163", "xm" or "custom"
	Type string `json:"type"`
	// Official
	ID string `json:"id"`
	// Custom
	Url     string `json:"url"`
	Audio   string `json:"audio"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}

func (music MusicType) Matcher(typeName, adapterName string) bool {
	return typeName == "music" && adapterName == OneBotV11.Name
}

func (music MusicType) ToRawText(msg message.MessageSegment) string {
	result := music.Deserialize(msg.Data, reflect.TypeOf(music)).(MusicType)
	return fmt.Sprintf("[OnebotV11:music,type=%s,id=%s,url=%s,audio=%s,title=%s,content=%s,image=%s]", result.Type, result.ID, result.Url, result.Audio, result.Title, result.Content, result.Image)
}

type ReplyType struct {
	message.MessageType
	ID string `json:"id"`
}

func (reply ReplyType) Matcher(typeName, adapterName string) bool {
	return typeName == "reply" && adapterName == OneBotV11.Name
}

func (reply ReplyType) ToRawText(msg message.MessageSegment) string {
	result := reply.Deserialize(msg.Data, reflect.TypeOf(reply)).(ReplyType)
	return fmt.Sprintf("[OnebotV11:reply,id=%s]", result.ID)
}

type ForwardType struct {
	message.MessageType
	ID string `json:"id"`
}

func (forward ForwardType) Matcher(typeName, adapterName string) bool {
	return typeName == "forward" && adapterName == OneBotV11.Name
}

func (forward ForwardType) ToRawText(msg message.MessageSegment) string {
	result := forward.Deserialize(msg.Data, reflect.TypeOf(forward)).(ForwardType)
	return fmt.Sprintf("[OnebotV11:forward,id=%s]", result.ID)
}

// Node for forward
type NodeType struct {
	message.MessageType
	// By ID
	ID string `json:"id"`
	// Custom
	UserID   string                   `json:"user_id"`
	Nickname string                   `json:"nickname"`
	Content  []message.MessageSegment `json:"content"`
}

func (NodeType) Matcher(typeName, adapterName string) bool {
	return typeName == "node" && adapterName == OneBotV11.Name
}

func (node NodeType) ToRawText(msg message.MessageSegment) string {
	result := node.Deserialize(msg.Data, reflect.TypeOf(node)).(NodeType)
	return fmt.Sprintf("[OnebotV11:node,id=%s,user_id=%s,nickname=%s,content=%#v]", result.ID, result.UserID, result.Nickname, result.Content)
}

type XMLType struct {
	message.MessageType
	Data string `json:"data"`
}

func (xml XMLType) Matcher(typeName, adapterName string) bool {
	return typeName == "xml" && adapterName == OneBotV11.Name
}

func (xml XMLType) ToRawText(msg message.MessageSegment) string {
	result := xml.Deserialize(msg.Data, reflect.TypeOf(xml)).(XMLType)
	return fmt.Sprintf("[OnebotV11:xml,data=%s]", result.Data)
}

type JSONType struct {
	message.MessageType
	Data string `json:"data"`
}

func (json JSONType) Matcher(typeName, adapterName string) bool {
	return typeName == "json" && adapterName == OneBotV11.Name
}

func (json JSONType) ToRawText(msg message.MessageSegment) string {
	result := json.Deserialize(msg.Data, reflect.TypeOf(json)).(JSONType)
	return fmt.Sprintf("[OnebotV11:json,data=%s]", result.Data)
}
