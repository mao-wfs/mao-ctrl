package device

import (
	"context"
	"time"
)

// FGHandler is the interface that describe the FG handler.
type FGHandler interface {
	Start(ctx context.Context, st, et time.Time) error
	Halt(ctx context.Context, ht time.Time) error
}
