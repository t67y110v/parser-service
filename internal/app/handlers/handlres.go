package handlers

import (
	"github.com/t67y110v/parser-service/internal/app/logging"
)

type Handlers struct {
	logger logging.Logger
}

func NewHandlers(logger logging.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
