package device

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/internal/app/domain"
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/gateway/device/correlator"
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/gateway/device/optswitch"
)

// handler represents the handler of MAO-WFS.
type handler struct {
	correlator correlator.Handler
	optswitch  optswitch.Handler
}

// NewHandler returns a new handler of MAO-WFS.
func NewHandler(corr correlator.Handler, sw optswitch.Handler) domain.Handler {
	return &handler{
		correlator: corr,
		optswitch:  sw,
	}
}

// Start starts MAO-WFS.
func (h *handler) Start(ctx context.Context) error {
	if err := h.correlator.Initialize(ctx); err != nil {
		return xerrors.Errorf("failed to initialize the correlator: %w", err)
	}
	if err := h.optswitch.Initialize(ctx); err != nil {
		return xerrors.Errorf("failed to initialize the optical switch: %w", err)
	}

	corrCh := make(chan error)
	swCh := make(chan error)

	go func() {
		defer close(corrCh)
		corrCh <- h.correlator.Start(ctx)
	}()
	go func() {
		defer close(swCh)
		swCh <- h.optswitch.Start(ctx)
	}()

	// TODO: Refactor the error handling
	corrErr, swErr := <-corrCh, <-swCh
	if corrErr != nil || swErr != nil {
		return xerrors.Errorf("correlator: %+v, optical switch: %+v", corrErr, swErr)
	}
	return nil
}

// Halt halts MAO-WFS.
func (h *handler) Halt(ctx context.Context) error {
	corrCh := make(chan error)
	swCh := make(chan error)

	go func() {
		defer close(corrCh)
		corrCh <- h.correlator.Halt(ctx)
	}()
	go func() {
		defer close(swCh)
		swCh <- h.optswitch.Halt(ctx)
	}()

	// TODO: Refactor the error handling
	corrErr, swErr := <-corrCh, <-swCh
	if corrErr != nil || swErr != nil {
		return xerrors.Errorf("correlator: %+v, optical switch: %+v", corrErr, swErr)
	}
	return nil
}
