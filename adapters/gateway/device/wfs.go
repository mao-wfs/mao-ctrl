package device

import (
	"context"
	"time"
)

// WFSHandler represents the handler of MAO-WFS.
type WFSHandler struct {
	Correlator CorrelatorHandler
	FG         FGHandler
}

// NewWFSHandler returns a new handler of MAO-WFS.
func NewWFSHandler(corrHan CorrelatorHandler, fgHan FGHandler) *WFSHandler {
	return &WFSHandler{
		Correlator: corrHan,
		FG:         fgHan,
	}
}

// Start starts MAO-WFS.
// TODO: Implement a function to start the correlator and the FG synchronization
func (h *WFSHandler) Start(ctx context.Context, sensT time.Duration) (time.Time, error) {
	var t time.Time
	return t, nil
}

// Halt halts MAO-WFS.
// TODO: Implement a function to halt the correlator and the FG
func (h *WFSHandler) Halt(ctx context.Context) (time.Time, error) {
	now := time.Now()
	return now, nil
}
