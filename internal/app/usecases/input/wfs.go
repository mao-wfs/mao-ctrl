package input

import (
	"context"

	"github.com/mao-wfs/mao-ctrl/internal/app/usecases/output"
)

// WFSInputPort is the interface that describe the input port of MAO-WFS.
type WFSInputPort interface {
	// Start starts MAO-WFS.
	Start(ctx context.Context) output.Error

	// Halt halts MAO-WFS.
	Halt(ctx context.Context) output.Error
}
