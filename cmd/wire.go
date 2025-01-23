//go:build wireinject

package main

import (
	"chat/internal/controller"
	"chat/internal/provider"

	"github.com/google/wire"
	"github.com/inoth/toybox"
	"github.com/inoth/toybox/config"
)

func initApp(conf config.ConfigMate) *toybox.ToyBox {
	panic(wire.Build(controller.ProviderSet, provider.ProviderSet, newApp))
}
