package cloudio

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// combiErr combines the *semantically* correct error from the os package
// with the decorated error from the cloudstore/blob package;
// outside calls to Unwrap(), errors.Is() and errors.As() should work
// similar as if we worked directly with the io package.
//
// Notice that there is no way to
type combiErr struct {
	base error // os.ErrNotExist
	deco error // i.e. gocloud.dev/blob error
}

func (ce combiErr) Error() string {
	return fmt.Sprintf("base error: %v; decorated (secondary) error: %v", ce.base.Error(), ce.deco.Error())
}

func (ce combiErr) Unwrap() error {
	return fmt.Errorf("decorated (secondary) error: %v; base error: %w", ce.deco.Error(), ce.base)
}

func (ce combiErr) Base() error {
	return ce.base
}

func (ce combiErr) Deco() error {
	return ce.deco
}

// CreateNotExist returns a "not exists" - similar to io package behaviour.
func CreateNotExist(err error) error {
	return combiErr{base: os.ErrNotExist, deco: err}
}

// IsNotExist function for cloudio
//
// we could check err := errors.As(..., &combiErr{})
// and then peform
// 		errors.Is(err, cloudStoreErr1)
// 		errors.Is(err, cloudStoreErr2)
func IsNotExist(err error) bool {
	if errors.Is(err, os.ErrNotExist) || strings.Contains(err.Error(), "code=NotFound") {
		return true
	}
	return false
}
