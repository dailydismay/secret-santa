package entity

import "github.com/rs/xid"

type ID string

func NewID() ID {
	return ID(xid.New().String())
}
