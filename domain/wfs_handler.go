package domain

import (
	"context"
	"time"
)

// WFSHandler is the interface that describe the handler of MAO-WFS.
type WFSHandler interface {
	// Start starts MAO-WFS.
	Start(ctx context.Context, st, et time.Time) error
	// Halt halts MAO-WFS.
	Halt(ctx context.Context, ht time.Time) error
}
