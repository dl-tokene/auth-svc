/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AuthNonceRequestAttributes struct {
	// Account Address from MetaMask.
	Address string `json:"address"`
	// Optional hash of Terms of Service to be included in the nonce message.
	TermsHash *string `json:"terms_hash,omitempty"`
}
