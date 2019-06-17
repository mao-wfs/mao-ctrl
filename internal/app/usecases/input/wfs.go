package input

import (
	"context"
)

// WFSInputPort is the interface that describe the input port of MAO-WFS.
type WFSInputPort interface {
	// Start starts MAO-WFS.
	Start(ctx context.Context) error

	// Halt halts MAO-WFS.
	Halt(ctx context.Context) error
}
