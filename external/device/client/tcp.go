package client

import (
	"net"
)

// TCPClient represents the TCP client.
type TCPClient struct {
	net.Conn
}

// NewTCPClient returns a new TCP client.
func NewTCPClient(addr string) (*TCPClient, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	clt := &TCPClient{
		Conn: conn,
	}
	return clt, nil
}

// Write send the message to a device.
func (c *TCPClient) Write(msg string) error {
	byteMsg := []byte(msg)
	if _, err := c.Conn.Write(byteMsg); err != nil {
		return err
	}
	return nil
}

// Read receive the message from a device.
func (c *TCPClient) Read(bufSize int) ([]byte, error) {
	buf := make([]byte, bufSize)
	len, err := c.Conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:len], nil
}

// Query queries to a device.
func (c *TCPClient) Query(msg string, bufSize int) ([]byte, error) {
	if err := c.Write(msg); err != nil {
		return nil, err
	}
	buf, err := c.Read(bufSize)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
