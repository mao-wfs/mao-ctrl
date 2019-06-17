package fg

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/internal/app/interfaces/gateway/device"
	"github.com/mao-wfs/mao-ctrl/internal/pkg/config"
	"github.com/mao-wfs/mao-ctrl/internal/app/infrastructure/device/client"
)

const defaultBufSize = 1024

// Handler represents the FG handler of MAO-WFS.
type Handler struct {
	config *config.FGConfig
	conn   *client.TCPClient
}

// NewHandler returns a new FG handler.
func NewHandler() (device.FGHandler, error) {
	conf, err := config.LoadFGConfig()
	if err != nil {
		return nil, err
	}
	clt, err := client.NewTCPClient(conf.GetAddr())
	if err != nil {
		return nil, err
	}
	h := &Handler{
		config: conf,
		conn:   clt,
	}
	return h, nil
}

// Start starts the FG of MAO-WFS.
func (h *Handler) Start(ctx context.Context) error {
	return h.startOutput()
}

// Halt halts the FG of MAO-WFS.
func (h *Handler) Halt(ctx context.Context) error {
	return h.haltOutput()
}

// Initialize initializes the FG of the MAO-WFS.
func (h *Handler) Initialize(ctx context.Context) error {
	if err := h.enableDigitalPattern(); err != nil {
		return err
	}

	if err := h.setFuncPatternVolatile(); err != nil {
		return err
	}
	if err := h.setDigitalPattern(); err != nil {
		return err
	}

	if err := h.setDigitalPatternTrigerSlopePositive(); err != nil {
		return err
	}
	return h.setDigitalPatternTrigerSourceExternal()
}

func (h *Handler) startOutput() error {
	msg := "OUTP ON"
	return h.execCmd(msg)
}

func (h *Handler) haltOutput() error {
	msg := "OUTP OFF"
	return h.execCmd(msg)
}

func (h *Handler) enableDigitalPattern() error {
	msg := "DIG:PATT ON"
	return h.execCmd(msg)
}

func (h *Handler) setDigitalPatternTrigerSourceExternal() error {
	msg := "DIG:PATT:TRIG:SOUR EXT"
	return h.execCmd(msg)
}

func (h *Handler) setDigitalPatternTrigerSlopePositive() error {
	msg := "DIG:PATT:TRIG:SLOP POS"
	return h.execCmd(msg)
}

func (h *Handler) setFuncPatternVolatile() error {
	msg := "FUNC:PATT VOLATILE"
	return h.execCmd(msg)
}

func (h *Handler) setDigitalPattern() error {
	o := h.config.GetOrder()
	oStr := make([]string, len(o))
	for i, v := range o {
		oStr[i] = strconv.Itoa(int(v))
	}
	msgPatt := strings.Join(oStr, ", ")
	msg := fmt.Sprintf("DATA:PATTERN VOLATILE, %s", msgPatt)
	return h.execCmd(msg)
}

func (h *Handler) execCmd(msg string) error {
	if err := h.conn.Write(msg + "\n"); err != nil {
		return err
	}
	return h.classifyResult()
}

func (h *Handler) classifyResult() error {
	msg := "SYST:ERR?"
	buf, err := h.conn.Query(msg + "\n", defaultBufSize)
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