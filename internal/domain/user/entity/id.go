package entity

import (
	"github.com/rs/xid"
)

type UserID string

func NewUserID() UserID {
	return UserID(xid.New().String())
}
