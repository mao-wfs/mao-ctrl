package correlator

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/xerrors"

	"github.com/mao-wfs/mao-ctrl/config"
	"github.com/mao-wfs/mao-ctrl/external/device/client"
)

const defaultBufSize = 1024

// Handler represents the correlator handler of MAO-WFS.
type Handler struct {
	Config *config.CorrelatorConfig
	Conn   *client.TCPClient
}

// NewHandler returns a new correlator handler.
func NewHandler() (*Handler, error) {
	conf, err := config.LoadCorrelatorConfig()
	if err != nil {
		return nil, err
	}
	clt, err := client.NewTCPClient(conf.GetAddr())
	if err != nil {
		return nil, err
	}
	h := &Handler{
		Config: conf,
		Conn:   clt,
	}
	return h, nil
}

// Start starts the correlator of MAO-WFS.
func (h *Handler) Start(ctx context.Context, st, et time.Time) error {
	if err := h.start(ctx, st, et); err != nil {
		return xerrors.Errorf("error in start method: %w", err)
	}
	return nil
}

// Halt halts the correlator of MAO-WFS.
func (h *Handler) Halt(ctx context.Context, ht time.Time) error {
	if err := h.halt(ctx, ht); err != nil {
		return xerrors.Errorf("error in halt method: %w", err)
	}
	return nil
}

func (h *Handler) start(ctx context.Context, st, et time.Time) error {
	return h.startCorrelation(st, et)
}

func (h *Handler) halt(ctx context.Context, ht time.Time) error {
	return h.haltCorrelation(ht)
}

func (h *Handler) startCorrelation(st, et time.Time) error {
	stMsg := fmt.Sprintf(
		"ctl_corstart=%04d%03d%02d%02d%02d:0x10;",
		st.Year(),
		st.YearDay(),
		st.Hour(),
		st.Minute(),
		st.Second(),
	)
	if err := h.execCmd(stMsg); err != nil {
		return err
	}

	if !et.IsZero() || et.Before(st) {
		edMsg := fmt.Sprintf(
			"ctl_corstop=%04d%03d%02d%02d%02d;",
			et.Year(),
			et.YearDay(),
			et.Hour(),
			et.Minute(),
			et.Second(),
		)
		if err := h.execCmd(edMsg); err != nil {
			return err
		}
	}

	return nil
}

func (h *Handler) haltCorrelation(ht time.Time) error {
	msg := fmt.Sprintf(
		"ctl_corstop=%04d%03d%02d%02d%02d;",
		ht.Year(),
		ht.YearDay(),
		ht.Hour(),
		ht.Minute(),
		ht.Second(),
	)
	return h.execCmd(msg)
}

func (h *Handler) execCmd(msg string) error {
	buf, err := h.Conn.Query(msg, defaultBufSize)
	if err != nil {
		return err
	}
	res := string(buf)
	return h.classifyResult(res)
}

func (h *Handler) classifyResult(res string) error {
	resCode := h.extractResultCode(res)
	switch resCode {
	case resNotExecutable:
		return errNotExecutable
	case resInvalidArgs:
		return errInvalidArgs
	case resUnknownError:
		return errUnknown
	case resConflict:
		return errConflict
	case resInvalidKwd:
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
