package device

import (
	"context"
	"time"
)

// CorrelatorHandler is the interface that describe the correlator handler.
type CorrelatorHandler interface {
	Start(ctx context.Context, st, et time.Time) error
	Halt(ctx context.Context, ht time.Time) error
}
