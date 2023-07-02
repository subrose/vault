package vault

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("conflict")
	// ErrForbidden  = errors.New("forbidden")
	ErrIndexError = errors.New("index")
)

type ErrForbidden struct{ action Action }

func (f ErrForbidden) Error() string {
	return fmt.Sprintf("forbidden: principal %s doing %s on %s", f.action.Principal.Name, f.action.Action, f.action.Resource)
}

type ValueError struct {
	Err error
	Msg string
}

func (ve *ValueError) Error() string {
	return fmt.Sprintf("value error: %s", ve.Err)
}

func (ve *ValueError) Unwrap() error {
	return ve.Err
}

func newValueError(err error) *ValueError {
	return &ValueError{
		Err: err,
	}
}
