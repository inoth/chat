//go:build wireinject

package main

import (
	"chat/internal/controller"
	"chat/internal/provider"
	"chat/internal/service"

	"github.com/google/wire"
	"github.com/inoth/toybox"
	"github.com/inoth/toybox/config"
)

func initApp(conf config.ConfigMate) *toybox.ToyBox {
	panic(wire.Build(service.ProviderSet, controller.ProviderSet, provider.ProviderSet, newApp))
}
