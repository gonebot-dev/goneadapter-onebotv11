package onebotv11

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/gonebot-dev/gonebot/message"

	"github.com/gorilla/websocket"
)

// The third thread, Sender.
func SendingMessage(ws *websocket.Conn) {
	log.Printf("Sender started.\n")
	for {
		//get a result from cache channel
		resultMsg := message.GetResultMsg()

		var payload APIPayload
		var convMessage []interface{} // converted Message

		// TODO
		//text part
		for _, msg := range resultMsg.Segments {
			switch msg.Type {
			case "text":
				convMessage = append(convMessage,
					MessageSegmentText{
						Type: "text",
						Data: struct {
							Text string "json:\"text\""
						}{Text: msg.Content}})
			case "image":
				convMessage = append(convMessage,
					MessageSegmentImg{
						Type: "image",
						Data: struct {
							Uri string "json:\"file\""
						}{Uri: msg.Content}})
			}

			if msg.Type == "text" {

			}
		}

		//private message
		if !resultMsg.IsGroup {
			userid, _ := strconv.Atoi(resultMsg.ReceiverID)
			params := APISendPrivateMessage{}
			params.UserID = userid
			params.Message = convMessage
			payload.Action = "send_private_msg"
			payload.Params = params

		} else {
			//group message
			groupID, _ := strconv.Atoi(resultMsg.GroupID)
			params := APISendGroupMessage{}
			params.GroupID = groupID
			params.Message = convMessage
			payload.Action = "send_group_msg"
			payload.Params = params

		}

		// send
		jsonResult, _ := json.Marshal(payload)
		err := ws.WriteMessage(websocket.TextMessage, jsonResult)
		if err != nil {
			log.Println(err)
		}

		// log
		if len(jsonResult) < 100 {
			log.Printf("Sending message: %s\n", jsonResult)
		} else {
			log.Printf("Sending message: %s...\n", jsonResult[:100])
		}
	}
}
