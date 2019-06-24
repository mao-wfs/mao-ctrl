package correlator

import (
	"context"
)

// Handler is the interface that describe the correlator handler.
type Handler interface {
	Start(ctx context.Context) error
	Halt(ctx context.Context) error
	Initialize(ctx context.Context) error
}
