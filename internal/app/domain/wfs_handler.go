package domain

import (
	"context"
)

// WFSHandler is the interface that describe the handler of MAO-WFS.
type WFSHandler interface {
	// Start starts MAO-WFS.
	Start(ctx context.Context) error

	// Halt halts MAO-WFS immediately.
	Halt(ctx context.Context) error
}
