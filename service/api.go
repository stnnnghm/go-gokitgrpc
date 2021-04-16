package service

import (
	"context"
	"log"
)

type service struct {
	logger log.Logger
}

// Service interface describes a service that adds numbers
type Service interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
}

// NewService returns a Service with all expected dependencies
func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

// Add func implements Service interface
func (s service) Add(ctx context.Context, numA, numB float32) (float32, error) {
	return numA + numB, nil
}
