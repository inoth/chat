package core

import (
	"bytes"
	"chat/apps/imchat/message"
	"chat/toybox/components/logger"
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Id   string
	hub  *ChatHubServer
	conn *websocket.Conn
	send chan []byte
}

func NewClient(uid string, hub *ChatHubServer, w http.ResponseWriter, r *http.Request) (*Client, error) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	c := &Client{
		Id:   uid,
		hub:  hub,
		conn: wsConn,
		send: make(chan []byte),
	}
	hub.register <- c
	return c, nil
}

func (c *Client) Run() {
	go c.readPump(c.hub.ctx)
	go c.writePump(c.hub.ctx)
}

func (c *Client) readPump(ctx context.Context) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, msg, err := c.conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					logger.Log.Errorf("err: %v", err)
					return
				}
				return
			}
			msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
			m, err := message.NewMsgBoxWithByte(msg)
			if err != nil {
				logger.Log.Error(err)
				continue
			}
			m.SetSource(c.Id)
			c.hub.broadcastInput <- m
		}
	}
}

func (c *Client) writePump(ctx context.Context) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
