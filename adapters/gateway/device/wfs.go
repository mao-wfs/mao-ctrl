package device

import (
	"context"
	"sync"

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
func (h *wfsHandler) Start(ctx context.Context) error {
	if err := h.correlator.Initialize(ctx); err != nil {
		return xerrors.Errorf("failed to initialize the correlator: %w", err)
	}
	if err := h.fg.Initialize(ctx); err != nil {
		return xerrors.Errorf("failed to initialize the FG: %w", err)
	}

	corrCh := make(chan error)
	fgCh := make(chan error)

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(corrCh)
		corrCh <- h.correlator.Start(ctx)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(fgCh)
		fgCh <- h.fg.Start(ctx)
	}()
	wg.Wait()

	// TODO: Fix error handlings
	errs := make([]error, 2)
	select {
	case err := <-corrCh:
		if err != nil {
			errs = append(errs, err)
		}
	case err := <-fgCh:
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return xerrors.Errorf("failed to start MAO-WFS: %+v", errs)
	}
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
