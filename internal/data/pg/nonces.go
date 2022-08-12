package pg

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
)

const nonceTableName = "nonce"

func newNonceQ(db *pgdb.DB) data.NonceQ {
	return &nonceQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

type nonceQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func (q *nonceQ) Get() *data.Nonce {
	var result data.Nonce
	err := q.db.Get(&result, q.sql.Select("*").From(nonceTableName))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(errors.Wrap(err, "failed to get nonce from db"))
	}
	return &result
}

func (q *nonceQ) Select() []data.Nonce {
	var result []data.Nonce
	err := q.db.Select(&result, q.sql.Select("*").From(nonceTableName))
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(errors.Wrap(err, "failed to select nonces from db"))
	}
	return result
}

func (q *nonceQ) Insert(value data.Nonce) data.Nonce {
	clauses := structs.Map(value)

	var result data.Nonce
	stmt := sq.Insert(nonceTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		panic(errors.Wrap(err, "failed to insert nonce to db"))
	}
	return result
}

func (q *nonceQ) Update(value data.Nonce) data.Nonce {
	clauses := structs.Map(value)

	var result data.Nonce
	stmt := q.sql.Update(nonceTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		panic(errors.Wrap(err, "failed to update nonce in db"))
	}
	return result
}

func (q *nonceQ) Delete() {
	err := q.db.Exec(q.sql.Delete(nonceTableName))
	if err != nil {
		panic(errors.Wrap(err, "failed to delete nonces from db"))
	}
	return
}

func (q *nonceQ) FilterByAddress(addresses ...string) data.NonceQ {
	pred := sq.Eq{"address": addresses}
	q.sql = q.sql.Where(pred)
	return q
}

func (q *nonceQ) FilterExpired() data.NonceQ {
	q.sql = q.sql.Where("expires < localtimestamp")
	return q
}
