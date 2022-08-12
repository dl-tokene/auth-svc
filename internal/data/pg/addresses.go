package pg

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
)

const addressTableName = "address"

func newAddressQ(db *pgdb.DB) data.AddressQ {
	return &addressQ{
		db:         db,
		sql:        sq.StatementBuilder,
		pageParams: nil,
	}
}

type addressQ struct {
	db         *pgdb.DB
	sql        sq.StatementBuilderType
	pageParams *pgdb.OffsetPageParams
}

func (q *addressQ) Get() *data.Address {
	var result data.Address
	stmt := q.sql.Select("*").From(addressTableName)
	if q.pageParams != nil {
		stmt = q.pageParams.ApplyTo(stmt, "address")
	}
	err := q.db.Get(&result, stmt)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(errors.Wrap(err, "failed to get address from db"))
	}
	return &result
}

func (q *addressQ) Select() []data.Address {
	var result []data.Address
	stmt := q.sql.Select("*").From(addressTableName)
	if q.pageParams != nil {
		stmt = q.pageParams.ApplyTo(stmt, "address")
	}
	err := q.db.Select(&result, stmt)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(errors.Wrap(err, "failed to select addresses from db"))
	}
	return result
}

func (q *addressQ) Insert(value data.Address) data.Address {
	clauses := structs.Map(value)

	var result data.Address
	stmt := sq.Insert(addressTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		panic(errors.Wrap(err, "failed to insert addresses to db"))
	}
	return result
}

func (q *addressQ) Update(value data.Address) data.Address {
	clauses := structs.Map(value)

	var result data.Address
	stmt := q.sql.Update(addressTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		panic(errors.Wrap(err, "failed to update address in db"))
	}
	return result
}

func (q *addressQ) Delete() {
	err := q.db.Exec(q.sql.Delete(addressTableName))
	if err != nil {
		panic(errors.Wrap(err, "failed to delete address from db"))
	}
}

func (q *addressQ) Page(pageParams pgdb.OffsetPageParams) data.AddressQ {
	q.pageParams = &pageParams
	return q
}

func (q *addressQ) FilterByAddress(addresses ...string) data.AddressQ {
	q.sql = q.sql.Where(sq.Eq{"address": addresses})
	return q
}

func (q *addressQ) SearchByAddress(address string) data.AddressQ {
	q.sql = q.sql.Where(sq.Like{"address": fmt.Sprint(address, "%")})
	return q
}

func (q *addressQ) FilterByUserID(userIds ...int64) data.AddressQ {
	q.sql = q.sql.Where(sq.Eq{"user_id": userIds})
	return q
}

func (q *addressQ) FilterByBusinessID(businessIds ...int64) data.AddressQ {
	q.sql = q.sql.Where(sq.Eq{"business_id": businessIds})
	return q
}
