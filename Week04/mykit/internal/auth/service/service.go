package service

import (
	"mykit/internal/auth/repository"
)

type Service struct {
	repo repository.Server
}

// NewGreeterService new a greeter service.
func NewGreeterService(r repository.Server) *Service {
	return &Service{
		repo: r,
	}
}
