package message

type MessageBox interface {
	MessageBody() []byte
	Targets() []string
	MsgType() string
}

type MessageBody struct {
	SendTargets []string // 消息发送的目标
	MessageType string   // 消息类型，text 普通消息，auth 登录验证
	Body        []byte   // 消息
}

func (mb *MessageBody) MessageBody() []byte { return nil }
func (mb *MessageBody) Targets() []string   { return mb.SendTargets }
func (mb *MessageBody) MsgType() string     { return mb.MessageType }
