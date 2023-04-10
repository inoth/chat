package main

import (
	"chat/src/router"

	"github.com/inoth/toybox"
	"github.com/inoth/toybox/components/cache"
	"github.com/inoth/toybox/components/config"
	logzap "github.com/inoth/toybox/components/logger/log_zap"
	"github.com/inoth/toybox/middleware"
	"github.com/inoth/toybox/services/ginsvc"
)

func main() {
	toybox.New(
		&cache.CacheComponent{},
		&config.ViperComponent{},
		&logzap.ZapComponent{},
	).Run(
		ginsvc.New(":8000").
			SetMiddleware(
				middleware.Recovery(),
				middleware.RequestLog(),
				// middleware.Cors(),
			).SetRouter(&router.BaseRouter{}))
}
