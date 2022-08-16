package pg

import (
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokene/nonce-auth-svc/internal/data"
)

func NewMasterQ(db *pgdb.DB) data.MasterQ {
	return &masterQ{
		db: db.Clone(),
	}
}

type masterQ struct {
	db *pgdb.DB
}

func (m *masterQ) New() data.MasterQ {
	return NewMasterQ(m.db)
}

func (m *masterQ) User() data.UserQ {
	return newUserQ(m.db)
}

func (m *masterQ) Address() data.AddressQ {
	return newAddressQ(m.db)
}
func (m *masterQ) Nonce() data.NonceQ {
	return newNonceQ(m.db)
}
func (m *masterQ) Transaction(fn func(q data.MasterQ) error) error {
	return m.db.Transaction(func() error {
		return fn(m)
	})
}
