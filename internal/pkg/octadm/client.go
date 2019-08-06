package octadm

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"time"
)

var responseRegexp = regexp.MustCompile(`[0-9]`)

// Client is a client of OCTAD-M.
type Client interface {
	// Close closes the connection.
	Close() error

	// Exec executes a SCPI command.
	Exec(cmd string) error

	// ExecContext executes a SCPI command.
	ExecContext(ctx context.Context, cmd string) error

	// Ping verifies the connection to the device is still alive,
	// establishing a connection if necessary.
	Ping() error

	// PingContext verifies the connection to the device is still alive,
	// establishing a connection if necessary.
	PingContext(ctx context.Context) error

	// Query queries the device for the results of the specified command.
	Query(cmd string) (res []byte, err error)

	// QueryContext queries the device for the results of the specified command.
	QueryContext(ctx context.Context, cmd string) (res []byte, err error)
}

func newClient(addr string, timeout time.Duration) (Client, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}

	d := net.Dialer{
		Timeout: timeout,
	}
	conn, err := d.Dial("tcp", tcpAddr.String())
	if err != nil {
		return nil, err
	}
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		return nil, fmt.Errorf("failed to cast %T to *net.TCPConn", conn)
	}
	c := &client{
		conn: tcpConn,
	}
	return c, nil
}

type client struct {
	conn *net.TCPConn
}

func (c *client) Close() error {
	return c.conn.Close()
}

func (c *client) Exec(cmd string) error {
	return c.ExecContext(context.Background(), cmd)
}

func (c *client) ExecContext(ctx context.Context, cmd string) error {
	return c.exec(ctx, cmd)
}

func (c *client) exec(ctx context.Context, cmd string) error {
	b := []byte(cmd + ";")
	if _, err := c.conn.Write(b); err != nil {
		return err
	}

	buf := make([]byte, 1024)
	l, err := c.conn.Read(buf)
	if err != nil {
		return err
	}

	code, err := strconv.Atoi(responseRegexp.FindString(string(buf[:l])))
	if err != nil {
		return err
	}
	switch code {
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

func (c *client) Ping() error {
	return c.PingContext(context.Background())
}

func (c *client) PingContext(ctx context.Context) error {
	// BUG(scizorman): PingContext is not implemented yet.
	return nil
}

func (c *client) Query(cmd string) (res []byte, err error) {
	return c.QueryContext(context.Background(), cmd)
}

func (c *client) QueryContext(ctx context.Context, cmd string) (res []byte, err error) {
	if err := c.ExecContext(ctx, cmd); err != nil {
		return nil, err
	}

	res = make([]byte, 65536)
	l, err := c.conn.Read(res)
	if err != nil {
		return nil, err
	}
	return res[:l], nil
}
