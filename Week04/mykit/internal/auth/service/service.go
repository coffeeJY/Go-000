package service

import (
	"mykit/internal/auth/repository"
)

type Service struct {
	repo repository.Server
}

// NewService new a greeter service.
func NewService(r repository.Server) *Service {
	return &Service{
		repo: r,
	}
}
