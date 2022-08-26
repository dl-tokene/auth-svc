/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AuthNonce struct {
	Key
	Attributes AuthNonceAttributes `json:"attributes"`
}
type AuthNonceResponse struct {
	Data     AuthNonce `json:"data"`
	Included Included  `json:"included"`
}

type AuthNonceListResponse struct {
	Data     []AuthNonce `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustAuthNonce - returns AuthNonce from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAuthNonce(key Key) *AuthNonce {
	var authNonce AuthNonce
	if c.tryFindEntry(key, &authNonce) {
		return &authNonce
	}
	return nil
}
