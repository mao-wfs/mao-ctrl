package device

import (
	"context"
)

// CorrelatorHandler is the interface that describe the correlator handler.
type CorrelatorHandler interface {
	Start(ctx context.Context) error
	Halt(ctx context.Context) error
	Initialize(ctx context.Context) error
}
