package device

// defaultBufSize is the default buffer size.
const defaultBufSize = 1024

// Client is the interface that describe the client of MAO-WFS devices.
type Client interface {
	Write(msg string) error
	Read(bufSize int) ([]byte, error)
	Query(msg string, bufSize int) ([]byte, error)
}
