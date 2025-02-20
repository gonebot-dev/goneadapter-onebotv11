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
var OneBotV11Adapter adapter.GoneAdapter

var ws *websocket.Conn

func init() {
	OneBotV11Adapter.Name = "OneBotV11"
	OneBotV11Adapter.Description = "The adapter for onebot v11 protocol"
	OneBotV11Adapter.Version = "v2.0.alpha"
	OneBotV11Adapter.SupportedPlatform = "qq"
	OneBotV11Adapter.Connector = Connector
	ws = nil
}

func GetAdapter() adapter.GoneAdapter {
	return OneBotV11Adapter
}
