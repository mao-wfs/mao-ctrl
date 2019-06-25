package device

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/internal/app/domain"
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/gateway/device/correlator"
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/gateway/device/fg"
)

// handler represents the handler of MAO-WFS.
type handler struct {
	correlator correlator.Handler
	fg         fg.Handler
}

// NewHandler returns a new handler of MAO-WFS.
func NewHandler(corr correlator.Handler, fg fg.Handler) domain.Handler {
	return &handler{
		correlator: corr,
		fg:         fg,
	}
}

// Start starts MAO-WFS.
func (h *handler) Start(ctx context.Context) error {
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
func (h *handler) Halt(ctx context.Context) error {
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
