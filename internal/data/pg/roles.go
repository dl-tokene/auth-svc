package pg

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
)

const roleTableName = "role"

func newRoleQ(db *pgdb.DB) data.RoleQ {
	return &roleQ{
		db:         db,
		sql:        sq.StatementBuilder,
		pageParams: nil,
	}
}

type roleQ struct {
	db         *pgdb.DB
	sql        sq.StatementBuilderType
	pageParams *pgdb.OffsetPageParams
}

func (q *roleQ) Get() *data.Role {
	var result data.Role

	stmt := q.sql.Select("*").From(roleTableName)
	if q.pageParams != nil {
		stmt = q.pageParams.ApplyTo(stmt, "name")
	}
	err := q.db.Get(&result, stmt)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(errors.Wrap(err, "failed to get role from db"))
	}
	return &result
}

func (q *roleQ) Select() []data.Role {
	var result []data.Role

	stmt := q.sql.Select("*").From(roleTableName)
	if q.pageParams != nil {
		stmt = q.pageParams.ApplyTo(stmt, "name")
	}
	err := q.db.Select(&result, stmt)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(errors.Wrap(err, "failed to select roles from db"))
	}
	return result
}

func (q *roleQ) Insert(value data.Role) data.Role {
	clauses := structs.Map(value)

	var result data.Role
	stmt := sq.Insert(roleTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		panic(errors.Wrap(err, "failed to insert role to db"))
	}
	return result
}

func (q *roleQ) Update(value data.Role) data.Role {
	clauses := structs.Map(value)

	var result data.Role
	stmt := q.sql.Update(roleTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		panic(errors.Wrap(err, "failed to update roles in db"))
	}
	return result
}

func (q *roleQ) Delete() {
	err := q.db.Exec(q.sql.Delete(roleTableName))
	if err != nil {
		panic(errors.Wrap(err, "failed to delete roles from db"))
	}
	return
}

func (q *roleQ) Page(pageParams pgdb.OffsetPageParams) data.RoleQ {
	q.pageParams = &pageParams
	return q
}

func (q *roleQ) FilterByName(names ...string) data.RoleQ {
	q.sql = q.sql.Where(sq.Eq{"name": names})
	return q
}
