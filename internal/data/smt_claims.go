package data

import (
	"gitlab.com/distributed_lab/kit/pgdb"
)

type SMTClaimsQ interface {
	Get() *SMTClaims
	Select() []SMTClaims
	Insert(value SMTClaims) SMTClaims
	Update(value SMTClaims) SMTClaims
	Delete()

	Page(pageParams pgdb.OffsetPageParams) SMTClaimsQ
	FilterByUserID(userIds ...int64) SMTClaimsQ
}

type SMTClaims struct {
	UserID int64   `db:"user_id" structs:"user_id"`
	Amount float64 `db:"amount" structs:"amount"`
}
