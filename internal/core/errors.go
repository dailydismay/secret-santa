package core

import (
	"errors"
	"fmt"
)

var (
	ErrEntityNotFound      = errors.New("entity not found")
	ErrInvalidParam        = errors.New("invalid param")
	ErrNotEnoughRights     = errors.New("not enough rights for resource")
	ErrResourceDuplication = errors.New("resource duplication")
)

func NewErrInvalidParam(resource, field string) error {
	return fmt.Errorf("%w: %s %s", ErrInvalidParam, resource, field)
}

func NewEntityNotFoundError(resource, searchPredicate string) error {
	return fmt.Errorf("%w: %s `%s`", ErrEntityNotFound, resource, searchPredicate)
}

func NewNotEnoughRightsError(subject, id string) error {
	return fmt.Errorf("%w for %s with id `%s`", ErrNotEnoughRights, subject, id)
}
