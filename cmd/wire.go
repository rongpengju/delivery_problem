//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rongpengju/delivery_problem/internal/biz"
	"github.com/rongpengju/delivery_problem/internal/data"

	"github.com/rongpengju/delivery_problem/internal/pkg/etc"
)

type Entrypoint struct {
	Config  *etc.Config
	Courier *biz.Courier
	DB      *data.Database
}

func NewEntrypoint(config *etc.Config, courier *biz.Courier, db *data.Database) *Entrypoint {
	return &Entrypoint{
		Config:  config,
		Courier: courier,
		DB:      db,
	}
}

func initEntrypoint() (*Entrypoint, error) {
	panic(wire.Build(
		NewEntrypoint,
		newConfig,
		biz.ProviderSet,
		data.ProviderSet,
	))
}
