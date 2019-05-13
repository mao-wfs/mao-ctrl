package domain

import (
	"context"
	"time"
)

// WFSHandler is the interface that describe the handler of MAO-WFS.
type WFSHandler interface {
	// Start starts MAO-WFS.
	// It senses for the specified time.
	Start(ctx context.Context, sensT time.Duration) (time.Time, error)

	// Halt halts MAO-WFS immediately.
	Halt(ctx context.Context) (time.Time, error)
}
