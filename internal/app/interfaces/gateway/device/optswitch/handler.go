package optswitch

import (
	"context"
)

// Handler is the interface that describe the FG handler.
type Handler interface {
	Start(ctx context.Context) error
	Halt(ctx context.Context) error
	Initialize(ctx context.Context) error
}
