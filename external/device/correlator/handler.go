package correlator

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/adapters/gateway/device"
	"github.com/mao-wfs/mao-ctrl/config"
	"github.com/mao-wfs/mao-ctrl/external/device/client"
)

const defaultBufSize = 1024

// Handler represents the correlator handler of MAO-WFS.
type Handler struct {
	config *config.CorrelatorConfig
	conn   *client.TCPClient
}

// NewHandler returns a new correlator handler.
func NewHandler() (device.CorrelatorHandler, error) {
	conf, err := config.LoadCorrelatorConfig()
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

// Start starts the correlator of MAO-WFS.
func (h *Handler) Start(ctx context.Context) error {
	if err := h.start(ctx); err != nil {
		return xerrors.Errorf("error in start method: %w", err)
	}
	return nil
}

// Halt halts the correlator of MAO-WFS.
func (h *Handler) Halt(ctx context.Context) error {
	if err := h.halt(ctx); err != nil {
		return xerrors.Errorf("error in halt method: %w", err)
	}
	return nil
}

// Initialize initializes the correlator of MAO-WFS.
func (h *Handler) Initialize(ctx context.Context) error {
	return nil
}

func (h *Handler) start(ctx context.Context) error {
	return h.startCorrelation()
}

func (h *Handler) halt(ctx context.Context) error {
	return h.haltCorrelation()
}

func (h *Handler) startCorrelation() error {
	var st time.Time
	msg := fmt.Sprintf(
		"ctl_corstart=%04dy%03dd%02dh%02dm%02ds:0x10;",
		st.Year(),
		st.YearDay(),
		st.Hour(),
		st.Minute(),
		st.Second(),
	)
	return h.execCmd(msg)
}

func (h *Handler) haltCorrelation() error {
	var ht time.Time
	msg := fmt.Sprintf(
		"ctl_corstop=%04dy%03dd%02dh%02dm%02ds;",
		ht.Year(),
		ht.YearDay(),
		ht.Hour(),
		ht.Minute(),
		ht.Second(),
	)
	return h.execCmd(msg)
}

func (h *Handler) execCmd(msg string) error {
	buf, err := h.conn.Query(msg, defaultBufSize)
	if err != nil {
		return err
	}
	res := string(buf)
	return h.classifyResult(res)
}

func (h *Handler) classifyResult(res string) error {
	resCode := h.extractResultCode(res)
	switch resCode {
	case resultNotExecutable:
		return errNotExecutable
	case resultInvalidArgs:
		return errInvalidArgs
	case resultUnknownError:
		return errUnknown
	case resultConflict:
		return errConflict
	case resultInvalidKwd:
		return errInvaildKwd
	default:
		return nil
	}
}

func (h *Handler) extractResultCode(res string) int {
	re := regexp.MustCompile(`[0-9]`)
	resCode, _ := strconv.Atoi(re.FindString(res))
	return resCode
}
