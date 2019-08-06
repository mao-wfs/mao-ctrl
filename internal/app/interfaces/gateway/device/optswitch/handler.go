package optswitch

import (
	"context"
)

// Handler is the interface that describe the optical switch handler.
type Handler interface {
	// Close closes the connection to the optical switch.
	Close() error

	// Start starts the swithcing.
	Start(ctx context.Context) error

	// Halt halts the switching.
	Halt(ctx context.Context) error

	// Initialize initiate the optical switch.
	Initialize(ctx context.Context) error
}
