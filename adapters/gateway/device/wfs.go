package device

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/domain"
)

// wfsHandler represents the handler of MAO-WFS.
type wfsHandler struct {
	correlator CorrelatorHandler
	fg         FGHandler
}

// NewWFSHandler returns a new handler of MAO-WFS.
func NewWFSHandler(corr CorrelatorHandler, fg FGHandler) domain.WFSHandler {
	return &wfsHandler{
		correlator: corr,
		fg:         fg,
	}
}

// Start starts MAO-WFS.
// TODO: Implement a function to start the correlator and the FG synchronization
func (h *wfsHandler) Start(ctx context.Context) error {
	return nil
}

// Halt halts MAO-WFS.
func (h *wfsHandler) Halt(ctx context.Context) error {
	if err := h.correlator.Halt(ctx); err != nil {
		return xerrors.Errorf("failed to halt the correlator: %w", err)
	}
	if err := h.fg.Halt(ctx); err != nil {
		return xerrors.Errorf("failed to halt the FG: %w", err)
	}
	return nil
}
