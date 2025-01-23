package provider

import (
	"chat/internal/controller"

	"github.com/google/wire"
	"github.com/inoth/toybox/ginserver"
	"github.com/inoth/toybox/ginserver/middleware"
	"github.com/inoth/toybox/wsserver"
)

var ProviderSet = wire.NewSet(
	NewHttpServer,
	NewWebsocketServer,
)

func NewHttpServer(
	cc *controller.ChatController,
	idx *controller.IndexController,
) *ginserver.GinHttpServer {
	return ginserver.NewHttp(
		ginserver.WithMiddleware(
			middleware.SetTraceId(),
		),
		ginserver.WithHandlers(cc, idx),
	)
}

func NewWebsocketServer(ws *controller.MessageController) *wsserver.WebsocketServer {
	return wsserver.New(
		wsserver.WithHandler(
			ws.Handler(),
		),
	)
}
