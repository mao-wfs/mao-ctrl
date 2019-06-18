package device

import (
	"context"
)

// FGHandler is the interface that describe the FG handler.
type FGHandler interface {
	Start(ctx context.Context) error
	Halt(ctx context.Context) error
	Initialize(ctx context.Context) error
}
