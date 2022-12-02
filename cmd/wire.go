//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rongpengju/delivery_problem/internal/biz"

	"github.com/rongpengju/delivery_problem/internal/pkg/etc"
)

type Entrypoint struct {
	Config  *etc.Config
	Courier *biz.Courier
}

func NewEntrypoint(config *etc.Config, courier *biz.Courier) *Entrypoint {
	return &Entrypoint{
		Config:  config,
		Courier: courier,
	}
}

func initEntrypoint() (*Entrypoint, error) {
	panic(wire.Build(
		NewEntrypoint,
		newConfig,
		biz.ProviderSet,
	))
}
