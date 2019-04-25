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
	Conn *client.TCPClient
}

// NewHandler returns a new correlator handler.
func NewHandler(c config.CorrelatorConfig) (*Handler, error) {
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
	if err := h.startCorrelation(st, et); err != nil {
		return err
	}
	return nil
}

func (h *Handler) halt(ctx context.Context, ht time.Time) error {
	if err := h.haltCorrelation(ht); err != nil {
		return err
	}
	return nil
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
	if err := h.execCmd(msg); err != nil {
		return err
	}
	return nil
}

func (h *Handler) execCmd(msg string) error {
	buf, err := h.Conn.Query(msg, defaultBufSize)
	if err != nil {
		return err
	}
	res := string(buf)
	if err := h.classifyResult(res); err != nil {
		return err
	}
	return nil
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
