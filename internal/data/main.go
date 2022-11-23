package data

type MasterQ interface {
	New() MasterQ

	Nonce() NonceQ

	Transaction(fn func(db MasterQ) error) error
}
