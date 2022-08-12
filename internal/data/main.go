package data

type MasterQ interface {
	New() MasterQ

	User() UserQ
	Address() AddressQ
	Nonce() NonceQ
	SMTClaims() SMTClaimsQ

	Transaction(fn func(db MasterQ) error) error
}
