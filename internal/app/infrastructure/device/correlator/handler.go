package correlator

import (
	"context"
	"time"

	"github.com/mao-wfs/mao-ctrl/internal/app/configs"
	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/gateway/device/correlator"
	"github.com/mao-wfs/mao-ctrl/internal/pkg/octadm"
)

const dialTimeout = 5 * time.Second

type handler struct {
	config     *configs.CorrelatorConfig
	correlator octadm.Handler
}

// NewHandler returns a new correlator handler.
func NewHandler() (correlator.Handler, error) {
	conf, err := configs.LoadCorrelatorConfig()
	if err != nil {
		return nil, err
	}
	corr, err := octadm.NewHandler(conf.Addr(), dialTimeout)
	if err != nil {
		return nil, err
	}
	h := &handler{
		config:     conf,
		correlator: corr,
	}
	return h, nil
}

// Start implements the Handler Start method.
func (h *handler) Start(ctx context.Context) error {
	var t time.Time
	return h.correlator.Start(t, octadm.Cross12)
}

// Halt implements the Handler Halt method.
func (h *handler) Halt(ctx context.Context) error {
	var t time.Time
	return h.correlator.Halt(t)
}

// Initialize implements the Handler Initialize method.
func (h *handler) Initialize(ctx context.Context) error {
	return h.correlator.SyncExternal()
}
