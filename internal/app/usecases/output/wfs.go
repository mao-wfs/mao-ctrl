package output

// WFSOutputPort is the output port of MAO-WFS.
type WFSOutputPort interface {
	ResponseError(code int, err error) Error
}
