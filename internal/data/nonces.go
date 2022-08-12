package data

import (
	"time"
)

type NonceQ interface {
	Get() *Nonce
	Select() []Nonce
	Insert(value Nonce) Nonce
	Update(value Nonce) Nonce
	Delete()

	FilterByAddress(addresses ...string) NonceQ
	FilterExpired() NonceQ
}

type Nonce struct {
	ID      int64      `db:"id" structs:"-"`
	Message string     `db:"message" structs:"message"`
	Expires *time.Time `db:"expires" structs:"expires"`
	Address string     `db:"address" structs:"address"`
}
