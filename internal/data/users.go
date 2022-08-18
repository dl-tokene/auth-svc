package data

import "gitlab.com/distributed_lab/kit/pgdb"

type UsersQ interface {
	Get() *User
	Select() []User
	Insert(value User) User
	Update(value User) User
	Delete()

	Page(pageParams pgdb.OffsetPageParams) UsersQ
	FilterByAddress(addresses ...string) UsersQ
	SearchByAddress(address string) UsersQ
	FilterByUserID(userIds ...int64) UsersQ
}

type User struct {
	Address string `db:"address" structs:"address"`
	ID      int64  `db:"id" structs:"-"`
}
