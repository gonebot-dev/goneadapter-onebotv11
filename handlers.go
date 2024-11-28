package onebotv11

import (
	"bytes"
	"encoding/json"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/gonebot-dev/gonebot/message"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

func messageDecoder(rawMsg string) {
	if !gjson.Valid(rawMsg) {
		log.Printf("[ONEBOTV11] | receiveHandler: Invalid message JSON %s\n", rawMsg)
		return
	}
	messageType := gjson.Get(rawMsg, "message_type")

	var msg message.Message
	var msgData []PayloadMessage

	msg.Self = gjson.Get(rawMsg, "self_id").String()
	msg.Sender = gjson.Get(rawMsg, "user_id").String()
	msg.Receiver = msg.Self

	switch messageType.String() {
	case "private":
		var msgInfo PrivateMessage
		_ = json.Unmarshal([]byte(rawMsg), &msgInfo)

		msg.IsToMe = true
		msgData = msgInfo.Message

	case "group":
		var msgInfo GroupMessage
		_ = json.Unmarshal([]byte(rawMsg), &msgInfo)

		msg.IsToMe = false
		msg.Group = strconv.FormatInt(msgInfo.GroupID, 10)
		msgData = msgInfo.Message

	default:
		log.Printf("[ONEBOTV11] | receiveHandler: Unsupported message type %s\n", messageType.String())
		return
	}

	for _, msgInterface := range msgData {
		useAdapter := OneBotV11.Name
		if msgInterface.Type == "text" ||
			msgInterface.Type == "image" ||
			msgInterface.Type == "voice" ||
			msgInterface.Type == "video" ||
			msgInterface.Type == "file" {
			useAdapter = ""
		}
		serializer := message.GetSerializer(msgInterface.Type, useAdapter)
		if serializer == nil {
			continue
		}
		if msgInterface.Type == "at" {
			atID := msgInterface.Data.(map[string]any)["qq"].(string)
			if atID == msg.Self {
				msg.IsToMe = true
			}
		}
		msg.AttachSegment(message.MessageSegment{
			Type:    msgInterface.Type,
			Adapter: useAdapter,
			Data:    serializer.Serialize(msgInterface.Data.(map[string]any), reflect.TypeOf(serializer)),
		}, serializer)
	}
	OneBotV11.ReceiveChannel.Push(msg, true)

	end := ""
	if len(rawMsg) > message.LOG_MESSAGE_LEN_THRESHOLD {
		rawMsg = rawMsg[:message.LOG_MESSAGE_LEN_THRESHOLD]
		end = "..."
	}
	log.Printf("[ONEBOTV11] | receiveHandler: Receive message %s%s\n", rawMsg, end)
}

func noticeDecoder(rawMsg string) {
	if !gjson.Valid(rawMsg) {
		log.Printf("[ONEBOTV11] | receiveHandler: Invalid notice JSON %s\n", rawMsg)
		return
	}
	noticeType := gjson.Get(rawMsg, "sub_type")

	var msg message.Message

	msg.Self = gjson.Get(rawMsg, "self_id").String()
	msg.Sender = gjson.Get(rawMsg, "user_id").String()

	switch noticeType.String() {
	case "group_upload":
		var noticeInfo GroupFileUpload
		_ = json.Unmarshal([]byte(rawMsg), &noticeInfo)

		msg.IsToMe = false
		msg.Group = strconv.FormatInt(noticeInfo.GroupID, 10)

	case "group_admin":
		var noticeInfo AdminChange
		_ = json.Unmarshal([]byte(rawMsg), &noticeInfo)

		msg.IsToMe = noticeInfo.UserID == noticeInfo.SelfID
		msg.Group = strconv.FormatInt(noticeInfo.GroupID, 10)

	case "group_decrease":
		var noticeInfo GroupMemberDecrease
		_ = json.Unmarshal([]byte(rawMsg), &noticeInfo)

		msg.IsToMe = false
		msg.Group = strconv.FormatInt(noticeInfo.GroupID, 10)

	case "group_increase":
		var noticeInfo GroupMemberIncrease
		_ = json.Unmarshal([]byte(rawMsg), &noticeInfo)

		msg.IsToMe = false
		msg.Group = strconv.FormatInt(noticeInfo.GroupID, 10)

	case "group_ban":
		var noticeInfo GroupBan
		_ = json.Unmarshal([]byte(rawMsg), &noticeInfo)

		msg.IsToMe = noticeInfo.UserID == noticeInfo.SelfID
		msg.Group = strconv.FormatInt(noticeInfo.GroupID, 10)

	case "friend_add":
		var noticeInfo FriendAdd
		_ = json.Unmarshal([]byte(rawMsg), &noticeInfo)

		msg.IsToMe = true

	case "group_recall":
		var noticeInfo GroupRecall
		_ = json.Unmarshal([]byte(rawMsg), &noticeInfo)

		msg.IsToMe = false
		msg.Group = strconv.FormatInt(noticeInfo.GroupID, 10)

	case "friend_recall":
		var noticeInfo FriendRecall
		_ = json.Unmarshal([]byte(rawMsg), &noticeInfo)

		msg.IsToMe = true

	case "poke":
		var noticeInfo GroupPoke
		_ = json.Unmarshal([]byte(rawMsg), &noticeInfo)

		msg.IsToMe = noticeInfo.TargetID == noticeInfo.SelfID
		msg.Group = strconv.FormatInt(noticeInfo.GroupID, 10)
		if msg.Group == "0" {
			msg.Group = ""
		}

	case "lucky_king":
		var noticeInfo RedPacketLuckyKing
		_ = json.Unmarshal([]byte(rawMsg), &noticeInfo)

		msg.IsToMe = false
		msg.Group = strconv.FormatInt(noticeInfo.GroupID, 10)

	case "honor":
		var noticeInfo GroupHonorChange
		_ = json.Unmarshal([]byte(rawMsg), &noticeInfo)

		msg.IsToMe = noticeInfo.UserID == noticeInfo.SelfID
		msg.Group = strconv.FormatInt(noticeInfo.GroupID, 10)

	default:
		log.Printf("[ONEBOTV11] | receiveHandler: Unsupported notice type %s\n", noticeType.String())
		return
	}
	notice := noticeType.String()
	// Special handle for poke 🤬
	if notice == "poke" {
		if gjson.Get(rawMsg, "group_id").String() == "" {
			notice = "friend_poke"
		} else {
			notice = "group_poke"
		}
	}
	serializer := message.GetSerializer(notice, OneBotV11.Name)
	msg.AttachSegment(message.MessageSegment{
		Type:    notice,
		Adapter: OneBotV11.Name,
		Data:    rawMsg,
	}, serializer)
	OneBotV11.ReceiveChannel.Push(msg, true)

	end := ""
	if len(rawMsg) > message.LOG_MESSAGE_LEN_THRESHOLD {
		rawMsg = rawMsg[:message.LOG_MESSAGE_LEN_THRESHOLD]
		end = "..."
	}
	log.Printf("[ONEBOTV11] | receiveHandler: Receive notice %s%s\n", rawMsg, end)
}

func handleReceive(rawMsg string) {
	if !gjson.Valid(rawMsg) {
		log.Printf("[ONEBOTV11] | Invalid message: %s\n", rawMsg)
		return
	}
	dataField := gjson.Get(rawMsg, "data")
	if dataField.Exists() {
		log.Printf("[ONEBOTV11] | receiveHandler: Receive action result %#v\n", dataField.Value())

		// Ignore send message result
		if gjson.Get(rawMsg, "data.message_id").Exists() && !gjson.Get(rawMsg, "data.time").Exists() {
			log.Println("[ONEBOTV11] | receiveHandler: Ignore send message result")
			return
		}
		actionResult <- dataField.Value()
		return
	}
	postType := gjson.Get(rawMsg, "post_type")
	if postType.Exists() {
		switch postType.String() {
		case "meta_event":
			metaEventType := gjson.Get(rawMsg, "meta_event_type")
			if metaEventType.Exists() {
				if metaEventType.String() == "heartbeat" {
					log.Println("[ONEBOTV11] | Receive heartbeat.")
				}
			}
		case "message":
			messageDecoder(rawMsg)
		case "notice":
			noticeDecoder(rawMsg)
		case "request":
			log.Println("[ONEBOTV11] | Request message currently unsupported.")
		default:
			log.Printf("[ONEBOTV11] | Unsupported post type %s\n", postType.String())
		}
	}
}

func receiveHandler() {
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("[ONEBOTV11] | Read message Error: \n%s\n", err)
		}
		re := regexp.MustCompile(`\\u([0-9a-fA-F]{4})`)
		escapedMsg := re.ReplaceAllStringFunc(string(msg), func(match string) string {
			codePoint, err := strconv.ParseUint(strings.Trim(match, "\\u"), 16, 32)
			if err != nil {
				return match
			}
			return string(rune(codePoint))
		})
		handleReceive(escapedMsg)
	}
}

func parseMessage(msg message.Message) (result []PayloadMessage) {
	result = make([]PayloadMessage, 0)
	for _, segment := range msg.GetSegments() {
		serializer := message.GetSerializer(segment.Type, segment.Adapter)
		if serializer == nil {
			log.Printf("[ONEBOTV11] | parseMessage: Unable to find serializer for %s\n", segment.Type)
			continue
		}
		result = append(result, PayloadMessage{
			Type: segment.Type,
			Data: serializer.Deserialize(segment.Data, reflect.TypeOf(serializer)),
		})
	}
	return
}

func sendHandler() {
	for {
		msg := OneBotV11.SendChannel.Pull()
		var result ActionPayload
		if msg.Group == "" {
			// Private message
			user, err := strconv.ParseInt(msg.Receiver, 10, 64)
			if err != nil {
				log.Printf("[ONEBOTV11] | sendHandler: Unable to parse user id %s\n", msg.Receiver)
				continue
			}
			result.Action = "send_private_msg"
			result.Params = SendPrivateMessage{
				UserID:     user,
				Message:    parseMessage(msg),
				AutoEscape: false,
			}
		} else {
			// Group message
			group, err := strconv.ParseInt(msg.Group, 10, 64)
			if err != nil {
				log.Printf("[ONEBOTV11] | sendHandler: Unable to parse group id %s\n", msg.Group)
				continue
			}
			result.Action = "send_group_msg"
			result.Params = SendGroupMessage{
				GroupID:    group,
				Message:    parseMessage(msg),
				AutoEscape: false,
			}
		}

		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		err := jsonEncoder.Encode(result)
		if err != nil {
			log.Printf("[ONEBOTV11] | sendHandler: Unable to marshal message %#v\n", result)
			continue
		}
		jsonResult := bf.String()

		err = ws.WriteMessage(websocket.TextMessage, []byte(jsonResult))

		end := ""
		if len(jsonResult) > message.LOG_MESSAGE_LEN_THRESHOLD {
			jsonResult = jsonResult[:message.LOG_MESSAGE_LEN_THRESHOLD]
			end = "..."
		}
		if err != nil {
			log.Printf("[ONEBOTV11] | sendHandler: Unable to send message %s%s\n", jsonResult, end)
			continue
		}

		log.Printf("[ONEBOTV11] | sendHandler: Send message %s%s\n", jsonResult, end)
	}
}

func actionHandler() {
	for {
		msg := OneBotV11.ActionChannel.Pull()
		var result ActionPayload

		result.Params = msg.Action
		switch msg.Action.(type) {
		case SendPrivateMessage:
			result.Action = "send_private_msg"
		case SendGroupMessage:
			result.Action = "send_group_msg"
		case DeleteMessage:
			result.Action = "delete_msg"
		case GetMessage:
			result.Action = "get_msg"
		case GetForwardMessage:
			result.Action = "get_forward_msg"
		case SendLike:
			result.Action = "send_like"
		case SetGroupKick:
			result.Action = "set_group_kick"
		case SetGroupBan:
			result.Action = "set_group_ban"
		case SetGroupWholeBan:
			result.Action = "set_group_whole_ban"
		case SetGroupAdmin:
			result.Action = "set_group_admin"
		case SetGroupCard:
			result.Action = "set_group_card"
		case SetGroupName:
			result.Action = "set_group_name"
		case SetGroupLeave:
			result.Action = "set_group_leave"
		case SetGroupSpecialTitle:
			result.Action = "set_group_special_title"
		case SetFriendAddRequest:
			result.Action = "set_friend_add_request"
		case SetGroupAddRequest:
			result.Action = "set_group_add_request"
		case GetLoginInfo:
			result.Action = "get_login_info"
		case GetStrangerInfo:
			result.Action = "get_stranger_info"
		case GetFriendList:
			result.Action = "get_friend_list"
		case GetGroupInfo:
			result.Action = "get_group_info"
		case GetGroupList:
			result.Action = "get_group_list"
		case GetGroupMemberInfo:
			result.Action = "get_group_member_info"
		case GetGroupMemberList:
			result.Action = "get_group_member_list"
		case GetGroupHonorInfo:
			result.Action = "get_group_honor_info"
		case GetCookies:
			result.Action = "get_cookies"
		case GetCsrfToken:
			result.Action = "get_csrf_token"
		case GetCredentials:
			result.Action = "get_credentials"
		case GetRecord:
			result.Action = "get_record"
		case GetImage:
			result.Action = "get_image"
		case CanSendImage:
			result.Action = "can_send_image"
		case CanSendRecord:
			result.Action = "can_send_record"
		case GetVersionInfo:
			result.Action = "get_version_info"
		case SetRestart:
			result.Action = "set_restart"
		case CleanCache:
			result.Action = "clean_cache"
		default:
			log.Printf("[ONEBOTV11] | actionHandler: Unknown action %#v\n", msg.Action)
			msg.ResultChannel <- nil
			continue
		}

		if msg.AdapterName != OneBotV11.Name {
			msg.ResultChannel <- nil
			log.Printf("[ONEBOTV11] | actionHandler: Ignore action for %s\n", msg.AdapterName)
			continue
		}

		jsonResult, err := json.Marshal(result)
		if err != nil {
			msg.ResultChannel <- nil
			log.Printf("[ONEBOTV11] | actionHandler: Unable to marshal action %#v\n", result)
			continue
		}

		if cap(actionResult) > 0 {
			<-actionResult
		}
		err = ws.WriteMessage(websocket.TextMessage, jsonResult)

		end := ""
		if len(jsonResult) > message.LOG_MESSAGE_LEN_THRESHOLD {
			jsonResult = jsonResult[:message.LOG_MESSAGE_LEN_THRESHOLD]
			end = "..."
		}
		if err != nil {
			msg.ResultChannel <- nil
			log.Printf("[ONEBOTV11] | actionHandler: Unable to send action %s%s...\n", jsonResult, end)
			continue
		}
		log.Printf("[ONEBOTV11] | actionHandler: Send action %s%s\n", jsonResult, end)
		reply := <-actionResult
		log.Printf("[ONEBOTV11] | actionHandler: Action %s%s received a result!%#v\n", jsonResult, end, reply)
		msg.ResultChannel <- reply
	}
}
