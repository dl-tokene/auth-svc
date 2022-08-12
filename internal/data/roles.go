package data

import "gitlab.com/distributed_lab/kit/pgdb"

const (
	RoleUndefined = "undefined" // NOTE(rufus): Reserved as a fallback.
	RoleUser      = "user"      // NOTE(rufus): Deprecated, do not use. To be removed after the next DB wipe.
	RoleAdmin     = "admin"
	RoleTier0     = "tier0"
	RoleTier1     = "tier1"
	RoleTier2     = "tier2"
	RoleSuspended = "tier90"
	RoleRejected  = "tier99"
)

type RoleQ interface {
	Get() *Role
	Select() []Role
	Insert(value Role) Role
	Update(value Role) Role
	Delete()

	Page(pageParams pgdb.OffsetPageParams) RoleQ
	FilterByName(name ...string) RoleQ
}

type Role struct {
	Name        string `db:"name" structs:"name"`
	Description string `db:"description" structs:"description"`
}
