package device

import (
	"context"
	"time"

	"golang.org/x/xerrors"
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
func (h *WFSHandler) Start(ctx context.Context, startTime, endTime time.Time) error {
	if err := h.Correlator.Start(ctx, startTime, endTime); err != nil {
		return xerrors.Errorf("error in correlator: %w", err)
	}
	if err := h.FG.Start(ctx, startTime, endTime); err != nil {
		return xerrors.Errorf("error in switch: %w", err)
	}
	return nil
}

// Halt halts MAO-WFS.
func (h *WFSHandler) Halt(ctx context.Context, haltTime time.Time) error {
	if err := h.Correlator.Halt(ctx, haltTime); err != nil {
		return xerrors.Errorf("error in correlator: %w", err)
	}
	if err := h.FG.Halt(ctx, haltTime); err != nil {
		return xerrors.Errorf("error in switch: %w", err)
	}
	return nil
}
