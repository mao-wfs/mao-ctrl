package octadm

type CorrelationMode uint

const (
	Auto1   CorrelationMode = 0x01
	Auto2   CorrelationMode = 0x02
	Cross12 CorrelationMode = 0x10
)

type IntegTime int

const (
	FiveMillisecond IntegTime = 5
	TenMillisecond  IntegTime = 10
)

type MaskTime int

// WindowFunc is a window function.
type WindowFunc string

// Window functions that can be used in OCTAD-M.
const (
	None     WindowFunc = "none"
	Hamming  WindowFunc = "hamming"
	Hanning  WindowFunc = "hanning"
	Blackman WindowFunc = "blackman"
)
