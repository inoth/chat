package message

import "encoding/json"

const (
	SysMsg  = "sys"
	TextMsg = "text"
	AuthMsg = "auth"
)

type MessageBox interface {
	Body() []byte

	Targets() []string

	MsgType() string
	MsgSouce() string
	String() []byte

	SetTargets(string, ...string)
	SetMsg(string, string)
	SetSource(string)
}

type MessageBody struct {
	// 目标类型，uid || rid
	TargetType string `json:"targetType"`
	// 消息发送的目标
	SendTargets []string `json:"sendTargets"`
	// 消息类型，text 普通消息，sys 系统消息，auth 登录验证
	MessageType string `json:"messageType"`
	// 发送消息内容
	SendBody string `json:"sendBody"`
	// 消息来源 uid
	MessageSource string `json:"messageSource"`

	// 消息原数据
	body []byte `json:"-"`
}

func (mb *MessageBody) Body() []byte {
	buf, _ := json.Marshal(mb)
	return buf
}
func (mb *MessageBody) Targets() []string { return mb.SendTargets }
func (mb *MessageBody) MsgType() string   { return mb.MessageType }
func (mb *MessageBody) MsgSouce() string  { return mb.MessageSource }

func (mb *MessageBody) String() []byte {
	buf, _ := json.Marshal(mb)
	return buf
}

func (mb *MessageBody) SetTargets(targetType string, targets ...string) {
	mb.TargetType = targetType
	mb.SendTargets = targets
}

func (mb *MessageBody) SetSource(uid string) {
	mb.MessageSource = uid
}

func (mb *MessageBody) SetMsg(msgType, msg string) {
	mb.MessageType = msgType
	mb.SendBody = msg
}

// 用于从 websocket 字节中读取内容进行序列化
func NewMsgBoxWithByte(buf []byte) (MessageBox, error) {
	var box MessageBody
	err := json.Unmarshal(buf, &box)
	if err != nil {
		return nil, err
	}
	box.body = buf
	return &box, nil
}

// 用于系统内部构建消息内容
func NewMsgBoxWithString(msg, source, msgType string, targets []string) (MessageBox, error) {
	return &MessageBody{
		TargetType:    "uid",
		SendTargets:   targets,
		MessageType:   msgType,
		SendBody:      msg,
		MessageSource: source,
	}, nil
}
