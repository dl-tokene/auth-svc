package data

import "gitlab.com/distributed_lab/kit/pgdb"

type AddressQ interface {
	Get() *Address
	Select() []Address
	Insert(value Address) Address
	Update(value Address) Address
	Delete()

	Page(pageParams pgdb.OffsetPageParams) AddressQ
	FilterByAddress(addresses ...string) AddressQ
	SearchByAddress(address string) AddressQ
	FilterByUserID(userIds ...int64) AddressQ
	FilterByBusinessID(businessIds ...int64) AddressQ
}

type Address struct {
	Address string `db:"address" structs:"address"`
	UserID  int64  `db:"user_id" structs:"user_id"`
}
