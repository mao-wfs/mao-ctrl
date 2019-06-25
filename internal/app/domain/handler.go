package domain

import (
	"context"
)

// Handler is the handler of MAO-WFS.
type Handler interface {
	// Start starts MAO-WFS.
	Start(ctx context.Context) error

	// Halt halts MAO-WFS.
	Halt(ctx context.Context) error
}
