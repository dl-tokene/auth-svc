package data

import (
	"time"

	"gitlab.com/distributed_lab/kit/pgdb"
)

type UsersQ interface {
	Get() (*User, error)
	Select() ([]User, error)
	Insert(value User) (*User, error)
	Update(value User) (*User, error)
	Delete() error

	Page(pageParams pgdb.OffsetPageParams) UsersQ
	FilterByAddress(addresses ...string) UsersQ
	SearchByAddress(address string) UsersQ
	FilterByUserID(userIds ...int64) UsersQ
}

type User struct {
	Address   string    `db:"address" structs:"address"`
	ID        int64     `db:"id" structs:"-"`
	CreatedAt time.Time `db:"createdat" structs:"createdat"`
}
