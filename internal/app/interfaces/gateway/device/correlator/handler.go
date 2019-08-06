package correlator

import (
	"context"
)

// Handler is the interface that describe the correlator handler.
type Handler interface {
	// Close closes the connection to the correlator.
	Close() error

	// Start starts the correlation.
	Start(ctx context.Context) error

	// Halt halts the correlation.
	Halt(ctx context.Context) error

	// Initialize initiate the correlator.
	Initialize(ctx context.Context) error
}
