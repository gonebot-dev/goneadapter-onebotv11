package onebotv11

import (
	"fmt"

	"github.com/gonebot-dev/gonebot/message"
)

type FaceType struct {
	ID string `json:"id"`
}

func (serializer FaceType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer FaceType) TypeName() string {
	return "face"
}

func (face FaceType) ToRawText(msg message.MessageSegment) string {
	result := msg.Data.(FaceType)
	return fmt.Sprintf("[OnebotV11:face,id=%s]", result.ID)
}

type AtType struct {
	QQ string `json:"qq"`
}

func (serializer AtType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer AtType) TypeName() string {
	return "at"
}

func (at AtType) ToRawText(msg message.MessageSegment) string {
	result := msg.Data.(AtType)
	return fmt.Sprintf("[OnebotV11:at,qq=%s]", result.QQ)
}

type RPSType struct{}

func (serializer RPSType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer RPSType) TypeName() string {
	return "rps"
}

func (rsp RPSType) ToRawText(msg message.MessageSegment) string {
	return "[OnebotV11:rps]"
}

type DiceType struct{}

func (serializer DiceType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer DiceType) TypeName() string {
	return "dice"
}

func (dice DiceType) ToRawText(msg message.MessageSegment) string {
	return "[OnebotV11:dice]"
}

type ShakeType struct{}

func (serializer ShakeType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer ShakeType) TypeName() string {
	return "shake"
}

func (shake ShakeType) ToRawText(msg message.MessageSegment) string {
	return "[OnebotV11:shake]"
}

type PokeType struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func (serializer PokeType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer PokeType) TypeName() string {
	return "poke"
}

func (poke PokeType) ToRawText(msg message.MessageSegment) string {
	result := msg.Data.(PokeType)
	return fmt.Sprintf("[OnebotV11:poke,qq=%s,type=%s]", result.ID, result.Type)
}

type ShareType struct {
	Url string `json:"url"`
}

func (serializer ShareType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer ShareType) TypeName() string {
	return "share"
}

func (share ShareType) ToRawText(msg message.MessageSegment) string {
	result := msg.Data.(ShareType)
	return fmt.Sprintf("[OnebotV11:share,url=%s]", result.Url)
}

type ContactType struct {
	// "qq" or "group"
	Type string `json:"type"`
	ID   string `json:"id"`
}

func (serializer ContactType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer ContactType) TypeName() string {
	return "contact"
}

func (contact ContactType) ToRawText(msg message.MessageSegment) string {
	result := msg.Data.(ContactType)
	return fmt.Sprintf("[OnebotV11:contact,type=%s,id=%s]", result.Type, result.ID)
}

type LocationType struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

func (serializer LocationType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer LocationType) TypeName() string {
	return "location"
}

func (location LocationType) ToRawText(msg message.MessageSegment) string {
	result := msg.Data.(LocationType)
	return fmt.Sprintf("[OnebotV11:location,lat=%s,lon=%s]", result.Lat, result.Lon)
}

type MusicType struct {
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

func (serializer MusicType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer MusicType) TypeName() string {
	return "music"
}

func (music MusicType) ToRawText(msg message.MessageSegment) string {
	result := msg.Data.(MusicType)
	return fmt.Sprintf("[OnebotV11:music,type=%s,id=%s,url=%s,audio=%s,title=%s,content=%s,image=%s]", result.Type, result.ID, result.Url, result.Audio, result.Title, result.Content, result.Image)
}

type ReplyType struct {
	ID string `json:"id"`
}

func (serializer ReplyType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer ReplyType) TypeName() string {
	return "reply"
}

func (reply ReplyType) ToRawText(msg message.MessageSegment) string {
	result := msg.Data.(ReplyType)
	return fmt.Sprintf("[OnebotV11:reply,id=%s]", result.ID)
}

type ForwardType struct {
	ID string `json:"id"`
}

func (serializer ForwardType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer ForwardType) TypeName() string {
	return "forward"
}

func (forward ForwardType) ToRawText(msg message.MessageSegment) string {
	result := msg.Data.(ForwardType)
	return fmt.Sprintf("[OnebotV11:forward,id=%s]", result.ID)
}

// Node for forward
type NodeType struct {
	// By ID
	ID string `json:"id"`
	// Custom
	UserID   string                   `json:"user_id"`
	Nickname string                   `json:"nickname"`
	Content  []message.MessageSegment `json:"content"`
}

func (serializer NodeType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer NodeType) TypeName() string {
	return "node"
}

func (node NodeType) ToRawText(msg message.MessageSegment) string {
	result := msg.Data.(NodeType)
	return fmt.Sprintf("[OnebotV11:node,id=%s,user_id=%s,nickname=%s,content=%#v]", result.ID, result.UserID, result.Nickname, result.Content)
}

type XMLType struct {
	Data string `json:"data"`
}

func (serializer XMLType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer XMLType) TypeName() string {
	return "xml"
}

func (xml XMLType) ToRawText(msg message.MessageSegment) string {
	result := msg.Data.(XMLType)
	return fmt.Sprintf("[OnebotV11:xml,data=%s]", result.Data)
}

type JSONType struct {
	Data string `json:"data"`
}

func (serializer JSONType) AdapterName() string {
	return OneBotV11.Name
}

func (serializer JSONType) TypeName() string {
	return "json"
}

func (json JSONType) ToRawText(msg message.MessageSegment) string {
	result := msg.Data.(JSONType)
	return fmt.Sprintf("[OnebotV11:json,data=%s]", result.Data)
}

// Convert message.MessageSegment.Data to message.MessageType
func ToMessageType(typeName string, msg any) message.MessageType {
	switch typeName {
	case "face":
		return msg.(FaceType)
	case "at":
		return msg.(AtType)
	case "rps":
		return msg.(RPSType)
	case "dice":
		return msg.(DiceType)
	case "shake":
		return msg.(ShakeType)
	case "poke":
		return msg.(PokeType)
	case "share":
		return msg.(ShareType)
	case "contact":
		return msg.(ContactType)
	case "location":
		return msg.(LocationType)
	case "music":
		return msg.(MusicType)
	case "reply":
		return msg.(ReplyType)
	case "forward":
		return msg.(ForwardType)
	case "node":
		return msg.(NodeType)
	case "xml":
		return msg.(XMLType)
	case "json":
		return msg.(JSONType)
	default:
		return nil
	}
}
