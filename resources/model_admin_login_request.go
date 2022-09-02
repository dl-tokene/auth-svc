/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AdminLoginRequest struct {
	Key
	Attributes AdminLoginRequestAttributes `json:"attributes"`
}
type AdminLoginRequestResponse struct {
	Data     AdminLoginRequest `json:"data"`
	Included Included          `json:"included"`
}

type AdminLoginRequestListResponse struct {
	Data     []AdminLoginRequest `json:"data"`
	Included Included            `json:"included"`
	Links    *Links              `json:"links"`
}

// MustAdminLoginRequest - returns AdminLoginRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAdminLoginRequest(key Key) *AdminLoginRequest {
	var adminLoginRequest AdminLoginRequest
	if c.tryFindEntry(key, &adminLoginRequest) {
		return &adminLoginRequest
	}
	return nil
}
