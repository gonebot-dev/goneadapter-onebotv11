# 消息类型

[消息段类型](#facetype)
- [QQ 表情](#facetype)

- [@某人](#attype)

- [猜拳魔法表情](#rpstype)

- [掷骰子魔法表情](#dicetype)

- [窗口抖动（戳一戳）](#shaketype)

- [戳一戳大表情](#poketype)

- [链接分享](#sharetyoe)

- [推荐好友/群](#contacttype)

- [位置](#locationtype)

- [音乐分享/自定义分享](#musictype)

- [回复](#replytype)

- [合并转发](#forwardtype)

- [合并转发节点/自定义节点](#nodetype)

- [XML 消息](#xmltype)

- [JSON 消息](#jsontype)

**OneBotV11 适配器实现了所有 OneBotV11 协议支持的消息段类型。所有的结构体定义如下：**

### FaceType
QQ 表情
```go
type FaceType struct {
	ID string `json:"id"`
}
```

### AtType
@某人
```go
type AtType struct {
	QQ string `json:"qq"`
}
```

### RPSType
猜拳魔法表情
```go
type RPSType struct{}
```

### DiceType
掷骰子魔法表情
```go
type DiceType struct{}
```

### ShakeType
窗口抖动
```go
type ShakeType struct{}
```

### PokeType
戳一戳大表情
```go
type PokeType struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
```

### ShareTyoe
链接分享
```go
type ShareType struct {
	Url string `json:"url"`
}
```

### ContactType
好友/群聊推荐
```go
type ContactType struct {
	// "qq" or "group"
	Type string `json:"type"`
	ID   string `json:"id"`
}
```

### LocationType
位置信息
```go
type LocationType struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}
```

### MusicType
音乐分享/自定义分享
```go
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
```

### ReplyType
消息回复
```go
type ReplyType struct {
	ID string `json:"id"`
}
```

### ForwardType
合并转发
```go
type ForwardType struct {
	ID string `json:"id"`
}
```

### NodeType
合并转发节点
接收时，此消息段不会直接出现在消息事件的 message 中，需通过 GetForwardMsg 获取。
```go
// Node for forward
type NodeType struct {
	// By ID
	ID string `json:"id"`
	// Custom
	UserID   string                   `json:"user_id"`
	Nickname string                   `json:"nickname"`
	Content  []message.MessageSegment `json:"content"`
}
```

### XMLType
XML 消息
```go
type XMLType struct {
	Data string `json:"data"`
}
```

### JSONType
JSON 消息
```go
type JSONType struct {
	Data string `json:"data"`
}
```
