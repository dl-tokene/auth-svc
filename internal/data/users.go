package data

import "gitlab.com/distributed_lab/kit/pgdb"

const (
	UserStatusNotVerified     = "not_verified"
	UserStatusAddressVerified = "address_verified"
	UserStatusKYCVerified     = "kyc_verified"
	UserStatusTier1Verified   = "tier1_verified"
	UserStatusPaymentVerified = "payment_verified"
	UserStatusDocSignVerified = "doc_sign_verified"
	UserStatusTier2Verified   = "tier2_verified"
	UserStatusSuspended       = "user_suspended"
	UserStatusRejected        = "user_rejected"

	KYCStatusNotStarted = "not_started"
	KYCStatusPending    = "pending"
	KYCStatusRejected   = "rejected"
	KYCStatusApproved   = "approved"
)

type UserQ interface {
	Get() *User
	Select() []User
	Insert(value User) User
	Update(value User) User
	Delete()

	Page(pageParams pgdb.OffsetPageParams) UserQ
	FilterByID(ids ...int64) UserQ
	FilterByRole(roles ...string) UserQ
	FilterByNFTRole(roles ...string) UserQ
	FilterUnmatchedNFTRoles() UserQ
}

type User struct {
	ID      int64  `db:"id" structs:"-"`
	Options string `db:"options" structs:"options"`
}
