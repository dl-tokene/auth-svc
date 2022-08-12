package pg

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
)

const userTableName = "users"

func newUserQ(db *pgdb.DB) data.UserQ {
	return &userQ{
		db:         db,
		sql:        sq.StatementBuilder,
		pageParams: nil,
	}
}

type userQ struct {
	db         *pgdb.DB
	sql        sq.StatementBuilderType
	pageParams *pgdb.OffsetPageParams
}

func (q *userQ) Get() *data.User {
	var result data.User
	stmt := q.sql.Select("*").From(userTableName)
	if q.pageParams != nil {
		stmt = q.pageParams.ApplyTo(stmt, "id")
	}
	err := q.db.Get(&result, stmt)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(errors.Wrap(err, "failed to get user from db"))
	}
	return &result
}

func (q *userQ) Select() []data.User {
	var result []data.User
	stmt := q.sql.Select("*").From(userTableName)
	if q.pageParams != nil {
		stmt = q.pageParams.ApplyTo(stmt, "id")
	}
	err := q.db.Select(&result, stmt)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(errors.Wrap(err, "failed to select users from db"))
	}
	return result
}

func (q *userQ) Insert(value data.User) data.User {
	clauses := structs.Map(value)

	var result data.User
	stmt := sq.Insert(userTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		panic(errors.Wrap(err, "failed to insert user to db"))
	}
	return result
}

func (q *userQ) Update(value data.User) data.User {
	clauses := structs.Map(value)

	var result data.User
	stmt := q.sql.Update(userTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		panic(errors.Wrap(err, "failed to update user in db"))
	}
	return result
}

func (q *userQ) Delete() {
	err := q.db.Exec(q.sql.Delete(userTableName))
	if err != nil {
		panic(errors.Wrap(err, "failed to delete user from db"))
	}
	return
}

func (q *userQ) Page(pageParams pgdb.OffsetPageParams) data.UserQ {
	q.pageParams = &pageParams
	return q
}

func (q *userQ) FilterByID(ids ...int64) data.UserQ {
	q.sql = q.sql.Where(sq.Eq{"id": ids})
	return q
}

func (q *userQ) FilterByRole(roles ...string) data.UserQ {
	q.sql = q.sql.Where(sq.Eq{"role": roles})
	return q
}

func (q *userQ) FilterByNFTRole(nftRoles ...string) data.UserQ {
	q.sql = q.sql.Where(sq.Eq{"nft_role": nftRoles})
	return q
}

func (q *userQ) FilterUnmatchedNFTRoles() data.UserQ {
	q.sql = q.sql.Where("role<>nft_role")
	return q
}
