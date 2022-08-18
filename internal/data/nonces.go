package data

import (
	"time"
)

type NonceQ interface {
	Get() (*Nonce, error)
	Select() ([]Nonce, error)
	Insert(value Nonce) (*Nonce, error)
	Update(value Nonce) (*Nonce, error)
	Delete() error

	FilterByAddress(addresses ...string) NonceQ
	FilterExpired() NonceQ
}

type Nonce struct {
	ID      int64      `db:"id" structs:"-"`
	Message string     `db:"message" structs:"message"`
	Expires *time.Time `db:"expires" structs:"expires"`
	Address string     `db:"address" structs:"address"`
}
