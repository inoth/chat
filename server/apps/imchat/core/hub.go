package core

import (
	"chat/apps/imchat"
	"chat/apps/imchat/message"
	"chat/apps/imchat/pipline/output"
	"chat/apps/imchat/pipline/process"
	"chat/toybox/components/config"
	"chat/toybox/components/logger"
	"chat/toybox/server"
	"context"
	"fmt"
	"sync"
)

var (
	serverName = "chathub"
	hub        *ChatHubServer
)

type ChatHubServer struct {
	m sync.RWMutex

	ctx    context.Context
	cancel context.CancelFunc

	clients map[string]*Client

	process   []imchat.MessageProcess
	msgoutput []imchat.MessageOutPut

	register   chan *Client
	unregister chan *Client

	broadcastInput  chan message.MessageBox
	broadcastOutput chan message.MessageBox

	RequiredComponents []string
}

func init() {
	server.Add(serverName, func() server.Server {
		cs := &ChatHubServer{
			m:                  sync.RWMutex{},
			clients:            make(map[string]*Client),
			register:           make(chan *Client),
			unregister:         make(chan *Client),
			broadcastInput:     make(chan message.MessageBox, 100),
			broadcastOutput:    make(chan message.MessageBox, 100),
			RequiredComponents: []string{"config", "logger"},
		}
		cs.ctx, cs.cancel = context.WithCancel(context.Background())
		return cs
	})
}

func SendMessage(msg message.MessageBox) {
	hub.broadcastOutput <- msg
}

func GetChatHub() *ChatHubServer {
	return hub
}

func (cs *ChatHubServer) init() {
	hub = cs
}

func (cs *ChatHubServer) Name() string { return serverName }
func (cs *ChatHubServer) RequiredComponent() []string {
	return cs.RequiredComponents
}

func (cs *ChatHubServer) Close() error {
	cs.cancel()
	return nil
}

func (cs *ChatHubServer) joinClients(client *Client) error {
	cs.m.RLock()
	defer cs.m.RUnlock()
	if _, ok := cs.clients[client.Id]; !ok {
		cs.clients[client.Id] = client

		sysmsg := fmt.Sprintf("client %v join to connect", client.Id)
		logger.Log.Infof(sysmsg)

		msg, err := message.NewMsgBoxWithString(sysmsg, message.SysMsg, message.SysMsg, []string{""})
		if err != nil {
			logger.Log.Warnf("create sys msg err: %v", err)
			return nil
		}

		cs.broadcastInput <- msg
		return nil
	}
	return fmt.Errorf("user %v already join", client.Id)
}

func (cs *ChatHubServer) removeClient(client *Client) error {
	cs.m.RLock()
	defer cs.m.RUnlock()
	if _, ok := cs.clients[client.Id]; ok {
		delete(cs.clients, client.Id)

		sysmsg := fmt.Sprintf("client %v exit to connect", client.Id)
		logger.Log.Infof(sysmsg)

		msg, err := message.NewMsgBoxWithString(sysmsg, message.SysMsg, message.SysMsg, []string{""})
		if err != nil {
			logger.Log.Warnf("create sys msg err: %v", err)
			return nil
		}

		cs.broadcastInput <- msg
		return nil
	}
	logger.Log.Warnf("user %v already not exists", client.Id)
	// return fmt.Errorf("user %v already not exists", client.Id)
	return nil
}

func (cs *ChatHubServer) loadProcess() error {
	// 加载配置，获取处理器顺序
	processList := config.Cfg.GetStringSlice("chat_process")
	cs.process = make([]imchat.MessageProcess, 0, len(processList))

	// 管道装载从output端开始加载，进行倒叙排列
	for i := len(processList) - 1; i >= 0; i-- {
		processCreator, ok := process.Process[processList[i]]
		if !ok {
			fmt.Printf("processor %v not loaded correctly\n", processList[i])
			continue
		}
		cs.process = append(cs.process, processCreator())
		fmt.Printf("loaded processor %v\n", processList[i])
	}
	return nil
}

func (cs *ChatHubServer) loadOutput() error {
	cs.msgoutput = make([]imchat.MessageOutPut, 0, len(output.Outputs))
	for k, output := range output.Outputs {
		cs.msgoutput = append(cs.msgoutput, output())
		fmt.Printf("loaded output %v\n", k)
	}
	return nil
}

func (cs *ChatHubServer) Start() error {
	// 保存一份到全局中，后续client注册使用
	cs.init()

	err := cs.loadProcess()
	if err != nil {
		return err
	}
	err = cs.loadOutput()
	if err != nil {
		return err
	}

	// 消息处理管道
	go cs.runPipline()

	for {
		select {
		case <-cs.ctx.Done():
			return fmt.Errorf("chathub services end")
		case reg := <-cs.register:
			err = cs.joinClients(reg)
			if err != nil {
				logger.Log.Errorf("client join err: %v", err)
				continue
			}
		case reg := <-cs.unregister:
			err = cs.removeClient(reg)
			if err != nil {
				logger.Log.Errorf("client exit err: %v", err)
				continue
			}
		case body := <-cs.broadcastOutput:
			if body.MsgType() == message.SysMsg {
				for _, client := range cs.clients {
					client.send <- body.String()
				}
			} else {
				for _, target := range body.Targets() {
					if _, ok := cs.clients[target]; ok {
						cs.clients[target].send <- body.String()
					}
				}
			}
		}
	}
}
