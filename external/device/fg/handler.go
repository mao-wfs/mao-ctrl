package fg

import (
	"context"
	"strings"
	"time"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/config"
	"github.com/mao-wfs/mao-ctrl/external/device/client"
)

const defaultBufSize = 1024

// Handler represents the FG handler of MAO-WFS.
type Handler struct {
	*client.TCPClient
}

// NewHandler returns a new FG handler.
func NewHandler(c config.FGConfig) (*Handler, error) {
	addr := c.GetAddr()
	clt, err := client.NewTCPClient(addr)
	if err != nil {
		return nil, err
	}
	h := &Handler{
		TCPClient: clt,
	}
	return h, nil
}

// Start starts the FG for the correlator.
func (h *Handler) Start(ctx context.Context, st, et time.Time) error {
	if err := h.start(ctx, st, et); err != nil {
		return xerrors.Errorf("error in start method: %w", err)
	}
	return nil
}

// Halt halts the FG for the correlator.
func (h *Handler) Halt(ctx context.Context, ht time.Time) error {
	if err := h.halt(ctx, ht); err != nil {
		return xerrors.Errorf("error in halt method: %w", err)
	}
	return nil
}

func (h *Handler) start(ctx context.Context, st, et time.Time) error {
	return nil
}

func (h *Handler) halt(ctx context.Context, ht time.Time) error {
	return nil
}

func (h *Handler) startOutput() error {
	msg := "OUTP ON\n"
	return h.execCmd(msg)
}

func (h *Handler) haltOutput() error {
	msg := "OUTP OFF\n"
	return h.execCmd(msg)
}

func (h *Handler) execCmd(msg string) error {
	if err := h.Write(msg); err != nil {
		return err
	}
	return h.classifyResult()
}

func (h *Handler) classifyResult() error {
	msg := "SYST:ERR?\n"
	buf, err := h.Query(msg, defaultBufSize)
	if err != nil {
		return err
	}
	res := string(buf)
	if res != "+0,\"No error\"\n" {
		errMsg := strings.TrimRight(res, "\n")
		return xerrors.New(errMsg)
	}
	return nil
}
