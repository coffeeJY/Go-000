//+build wireinject

package main

import (
	"github.com/google/wire"
	"mykit/internal/auth/repository"
	"mykit/internal/auth/service"
)

func CreateConcatService() *service.Service {
	wire.Build(
		repository.NewRepository,
		service.NewService,
	)
	return &service.Service{}
}
