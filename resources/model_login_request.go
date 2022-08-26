/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type LoginRequest struct {
	Key
	Attributes LoginRequestAttributes `json:"attributes"`
}
type LoginRequestResponse struct {
	Data     LoginRequest `json:"data"`
	Included Included     `json:"included"`
}

type LoginRequestListResponse struct {
	Data     []LoginRequest `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustLoginRequest - returns LoginRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustLoginRequest(key Key) *LoginRequest {
	var loginRequest LoginRequest
	if c.tryFindEntry(key, &loginRequest) {
		return &loginRequest
	}
	return nil
}
