package pg

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
)

const smtClaimsTableName = "smt_claims"

func newSMTClaimsQ(db *pgdb.DB) data.SMTClaimsQ {
	return &smtClaimsQ{
		db:         db,
		sql:        sq.StatementBuilder,
		pageParams: nil,
	}
}

type smtClaimsQ struct {
	db         *pgdb.DB
	sql        sq.StatementBuilderType
	pageParams *pgdb.OffsetPageParams
}

func (q *smtClaimsQ) Get() *data.SMTClaims {
	var result data.SMTClaims
	stmt := q.sql.Select("*").From(smtClaimsTableName)
	if q.pageParams != nil {
		stmt = q.pageParams.ApplyTo(stmt, "user_id")
	}
	err := q.db.Get(&result, stmt)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(errors.Wrap(err, "failed to get SMT claims from db"))
	}
	return &result
}

func (q *smtClaimsQ) Select() []data.SMTClaims {
	var result []data.SMTClaims
	stmt := q.sql.Select("*").From(smtClaimsTableName)
	if q.pageParams != nil {
		stmt = q.pageParams.ApplyTo(stmt, "user_id")
	}
	err := q.db.Select(&result, stmt)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		panic(errors.Wrap(err, "failed to select SMT claims from db"))
	}
	return result
}

func (q *smtClaimsQ) Insert(value data.SMTClaims) data.SMTClaims {
	clauses := structs.Map(value)

	var result data.SMTClaims
	stmt := sq.Insert(smtClaimsTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		panic(errors.Wrap(err, "failed to insert SMT claims to db"))
	}
	return result
}

func (q *smtClaimsQ) Update(value data.SMTClaims) data.SMTClaims {
	clauses := structs.Map(value)

	var result data.SMTClaims
	stmt := q.sql.Update(smtClaimsTableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		panic(errors.Wrap(err, "failed to update SMT claims in db"))
	}
	return result
}

func (q *smtClaimsQ) Delete() {
	err := q.db.Exec(q.sql.Delete(smtClaimsTableName))
	if err != nil {
		panic(errors.Wrap(err, "failed to delete SMT claims from db"))
	}
	return
}

func (q *smtClaimsQ) Page(pageParams pgdb.OffsetPageParams) data.SMTClaimsQ {
	q.pageParams = &pageParams
	return q
}

func (q *smtClaimsQ) FilterByUserID(userIds ...int64) data.SMTClaimsQ {
	q.sql = q.sql.Where(sq.Eq{"user_id": userIds})
	return q
}
