package octadm

import (
	"golang.org/x/xerrors"
)

// Result codes of the correlator operations.
const (
	resultOK int = iota
	resultReset
	resultNotExecutable
	resultInvalidArgs
	resultUnknownError
	_
	resultConflict
	resultInvalidKwd
)

// Errors of the correlator operations.
var (
	errNotExecutable = xerrors.New("cannot executed due to the difference of the operation mode")
	errInvalidArgs   = xerrors.New("invalid arguments")
	errUnknown       = xerrors.New("error while executing this command")
	errConflict      = xerrors.New("cannot executed due to a contradiction or conflicting commands")
	errInvaildKwd    = xerrors.New("invalid keyword")
)
