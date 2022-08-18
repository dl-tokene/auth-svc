package data

type MasterQ interface {
	New() MasterQ

	Users() UsersQ
	Nonce() NonceQ

	Transaction(fn func(db MasterQ) error) error
}
