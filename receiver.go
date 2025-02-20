package onebotv11

import (
	"log"

	"github.com/gonebot-dev/gonebot/message"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
)

// This file converts onebot11 json payload into MessageStruct.

// Handle raw websocket payload
func messageHandler(msg string) {
	if !gjson.Valid(msg) {
		log.Printf("Receive invalid JSON.\n")
		return
	}

	postType := gjson.Get(msg, "post_type")
	if postType.Exists() {
		switch postType.String() {
		//Meta event
		case "meta_event":
			metaEventType := gjson.Get(msg, "meta_event_type")
			if metaEventType.Exists() {
				//heartbeat
				if metaEventType.String() == "heartbeat" {
					log.Printf("Receive ws Heartbeat.\n")
				}
			}
		//message
		case "message":
			messageDecoder(msg)
		}
	}
}

// Format onebot message json and push into fifo queue.
func messageDecoder(rawMessage string) {
	//log.Printf("Receive raw message: %s\n", rawMessage)
	var newMsg message.Message

	//Who am i?
	newMsg.ReceiverID = gjson.Get(rawMessage, "self_id").String()

	//Is private message?
	if gjson.Get(rawMessage, "message_type").String() == "private" {
		newMsg.IsToMe = true
		newMsg.IsGroup = false
	} else {
		newMsg.IsGroup = true
	}
	//Is to me?
	selfID := gjson.Get(rawMessage, "self_id").String()
	newMsg.ReceiverID = selfID
	atUsers := gjson.GetMany(rawMessage, "message.#(type==\"at\")#.data.qq")
	for _, value := range atUsers {
		if value.String() == selfID {
			newMsg.IsToMe = true
		}
	}

	// TODO

	//Extract all text from message.
	textMessages := gjson.Get(rawMessage, "message.#(type==\"text\")#.data.text")
	textMessages.ForEach(func(_, value gjson.Result) bool {
		newMsg.AddTextSegment(value.String())
		return true // keep iterating, gjson
	})

	//Who send the message?
	newMsg.SenderID = gjson.Get(rawMessage, "sender.user_id").String()

	//Advanced handler
	// newMsg.RawMessage = rawMessage

	//Push message into messages queue.
	message.PushIncomingMsg(newMsg)
}

func ReadingMessage(ws *websocket.Conn) {
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Read message Error:\n%s\n", err)
		}
		msg := string(message)
		messageHandler(msg)
	}
}
