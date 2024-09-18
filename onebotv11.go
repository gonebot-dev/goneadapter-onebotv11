package onebotv11

import (
	"github.com/gonebot-dev/gonebot/adapter"
	"github.com/gonebot-dev/gonebot/message"
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

	// Register serializers
	// Message serializers
	message.RegisterSerializer("face", OneBotV11.Name, FaceType{})
	message.RegisterSerializer("at", OneBotV11.Name, AtType{})
	message.RegisterSerializer("rps", OneBotV11.Name, RPSType{})
	message.RegisterSerializer("dice", OneBotV11.Name, DiceType{})
	message.RegisterSerializer("shake", OneBotV11.Name, ShakeType{})
	message.RegisterSerializer("poke", OneBotV11.Name, PokeType{})
	message.RegisterSerializer("share", OneBotV11.Name, ShareType{})
	message.RegisterSerializer("contact", OneBotV11.Name, ContactType{})
	message.RegisterSerializer("location", OneBotV11.Name, LocationType{})
	message.RegisterSerializer("music", OneBotV11.Name, MusicType{})
	message.RegisterSerializer("reply", OneBotV11.Name, ReplyType{})
	message.RegisterSerializer("forward", OneBotV11.Name, ForwardType{})
	message.RegisterSerializer("node", OneBotV11.Name, NodeType{})
	message.RegisterSerializer("xml", OneBotV11.Name, XMLType{})
	message.RegisterSerializer("json", OneBotV11.Name, JSONType{})
	// Notice serializers
	message.RegisterSerializer("group_upload", OneBotV11.Name, GroupFileUpload{})
	message.RegisterSerializer("group_admin", OneBotV11.Name, AdminChange{})
	message.RegisterSerializer("group_decreasse", OneBotV11.Name, GroupMemberDecrease{})
	message.RegisterSerializer("group_increase", OneBotV11.Name, GroupMemberIncrease{})
	message.RegisterSerializer("group_ban", OneBotV11.Name, GroupBan{})
	message.RegisterSerializer("friend_add", OneBotV11.Name, FriendAdd{})
	message.RegisterSerializer("group_recall", OneBotV11.Name, GroupRecall{})
	message.RegisterSerializer("friend_recall", OneBotV11.Name, FriendRecall{})
	message.RegisterSerializer("group_poke", OneBotV11.Name, GroupPoke{})
	message.RegisterSerializer("lucky_king", OneBotV11.Name, RedPacketLuckyKing{})
	message.RegisterSerializer("honor", OneBotV11.Name, GroupHonorChange{})
}
