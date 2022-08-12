/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AuthNonceRequest struct {
	Key
	Attributes AuthNonceRequestAttributes `json:"attributes"`
}
type AuthNonceRequestResponse struct {
	Data     AuthNonceRequest `json:"data"`
	Included Included         `json:"included"`
}

type AuthNonceRequestListResponse struct {
	Data     []AuthNonceRequest `json:"data"`
	Included Included           `json:"included"`
	Links    *Links             `json:"links"`
}

// MustAuthNonceRequest - returns AuthNonceRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAuthNonceRequest(key Key) *AuthNonceRequest {
	var authNonceRequest AuthNonceRequest
	if c.tryFindEntry(key, &authNonceRequest) {
		return &authNonceRequest
	}
	return nil
}
