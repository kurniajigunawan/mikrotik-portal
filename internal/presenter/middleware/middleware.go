package middleware

import (
	serviceRepo "github.com/kurniajigunawan/mikrotik-portal/internal/repository/service"
)

type Middleware struct {
	serviceRepo serviceRepo.Interface
}

func NewMiddleware(serviceRepo serviceRepo.Interface) *Middleware {
	return &Middleware{
		serviceRepo: serviceRepo,
	}
}
