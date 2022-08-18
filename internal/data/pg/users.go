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

const usersTableName = "users"

func newUsersQ(db *pgdb.DB) data.UsersQ {
	return &usersQ{
		db:         db,
		sql:        sq.StatementBuilder,
		pageParams: nil,
	}
}

type usersQ struct {
	db         *pgdb.DB
	sql        sq.StatementBuilderType
	pageParams *pgdb.OffsetPageParams
}

func (q *usersQ) Get() (*data.User, error) {
	var result data.User
	stmt := q.sql.Select("*").From(usersTableName)
	if q.pageParams != nil {
		stmt = q.pageParams.ApplyTo(stmt, "address")
	}
	err := q.db.Get(&result, stmt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get address from db")
	}
	return &result, nil
}

func (q *usersQ) Select() ([]data.User, error) {
	var result []data.User
	stmt := q.sql.Select("*").From(usersTableName)
	if q.pageParams != nil {
		stmt = q.pageParams.ApplyTo(stmt, "address")
	}
	err := q.db.Select(&result, stmt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to select addresses from db")
	}
	return result, nil
}

func (q *usersQ) Insert(value data.User) (*data.User, error) {
	clauses := structs.Map(value)

	var result data.User
	stmt := sq.Insert(usersTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert addresses to db")
	}
	return &result, nil
}

func (q *usersQ) Update(value data.User) (*data.User, error) {
	clauses := structs.Map(value)

	var result data.User
	stmt := q.sql.Update(usersTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update address in db")
	}
	return &result, nil
}

func (q *usersQ) Delete() error {
	err := q.db.Exec(q.sql.Delete(usersTableName))
	if err != nil {
		return errors.Wrap(err, "failed to delete address from db")
	}
	return nil
}

func (q *usersQ) Page(pageParams pgdb.OffsetPageParams) data.UsersQ {
	q.pageParams = &pageParams
	return q
}

func (q *usersQ) FilterByAddress(addresses ...string) data.UsersQ {
	q.sql = q.sql.Where(sq.Eq{"address": addresses})
	return q
}

func (q *usersQ) SearchByAddress(address string) data.UsersQ {
	q.sql = q.sql.Where(sq.Like{"address": fmt.Sprint(address, "%")})
	return q
}

func (q *usersQ) FilterByUserID(userIds ...int64) data.UsersQ {
	q.sql = q.sql.Where(sq.Eq{"id": userIds})
	return q
}
