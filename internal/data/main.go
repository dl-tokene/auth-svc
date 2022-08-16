package data

type MasterQ interface {
	New() MasterQ

	User() UserQ
	Address() AddressQ
	Nonce() NonceQ

	Transaction(fn func(db MasterQ) error) error
}
