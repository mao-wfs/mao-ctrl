package device

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/internal/app/domain"
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
func (h *wfsHandler) Start(ctx context.Context) error {
	if err := h.correlator.Initialize(ctx); err != nil {
		return xerrors.Errorf("failed to initialize the correlator: %w", err)
	}
	if err := h.fg.Initialize(ctx); err != nil {
		return xerrors.Errorf("failed to initialize the FG: %w", err)
	}

	corrCh := make(chan error)
	fgCh := make(chan error)

	go func() {
		defer close(corrCh)
		corrCh <- h.correlator.Start(ctx)
	}()
	go func() {
		defer close(fgCh)
		fgCh <- h.fg.Start(ctx)
	}()

	// TODO: Refactor the error handling
	corrErr, fgErr := <-corrCh, <-fgCh
	if corrErr != nil || fgErr != nil {
		return xerrors.Errorf("correlator: %+v, FG: %+v", corrErr, fgErr)
	}
	return nil
}

// Halt halts MAO-WFS.
func (h *wfsHandler) Halt(ctx context.Context) error {
	corrCh := make(chan error)
	fgCh := make(chan error)

	go func() {
		defer close(corrCh)
		corrCh <- h.correlator.Halt(ctx)
	}()
	go func() {
		defer close(fgCh)
		fgCh <- h.fg.Halt(ctx)
	}()

	// TODO: Refactor the error handling
	corrErr, fgErr := <-corrCh, <-fgCh
	if corrErr != nil || fgErr != nil {
		return xerrors.Errorf("correlator: %+v, FG: %+v", corrErr, fgErr)
	}
	return nil
}
