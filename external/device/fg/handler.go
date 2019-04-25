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
	Conn *client.TCPClient
}

// NewHandler returns a new FG handler.
func NewHandler(c config.FGConfig) (*Handler, error) {
	addr := c.GetAddr()
	clt, err := client.NewTCPClient(addr)
	if err != nil {
		return nil, err
	}
	h := &Handler{
		Conn: clt,
	}
	return h, nil
}

// Start starts the FG for the FG.
func (h *Handler) Start(ctx context.Context, st, et time.Time) error {
	if err := h.start(ctx, st, et); err != nil {
		return xerrors.Errorf("error in start method: %w", err)
	}
	return nil
}

// Halt halts the FG for the FG.
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

func (h *Handler) classifyResult() error {
	msg := "SYST:ERR?\n"
	buf, err := h.Conn.Query(msg, defaultBufSize)
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
