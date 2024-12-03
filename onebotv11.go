package onebotv11

import (
	"github.com/gonebot-dev/gonebot/adapter"
	"github.com/gorilla/websocket"
)

// The adapter for "OneBot v11"
//
// # This requires a NTQQ protocol application running aside, like Lagrange.OneBot
//
// And the reverse socket server should be started on 127.0.0.1:21390(ONEBOTV11_HOST in .env file) by default.
//
// You can override the host address to your liking by setting ONEBOTV11_HOST in .env file.
//
// And you should be aware that the host is your gonebot server, not any NTQQ protocol
var OneBotV11 adapter.Adapter

var ws *websocket.Conn
var actionResult chan any

func init() {
	OneBotV11.Name = "OneBot v11"
	OneBotV11.Description = "The adapter for onebot v11 protocol"
	OneBotV11.Version = "0.1.0"
	OneBotV11.Start = start
	OneBotV11.Finalize = finalize
	ws = nil
	actionResult = make(chan any, 1)
}
