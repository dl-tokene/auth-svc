/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AdminLogin struct {
	Key
	Attributes AdminLoginAttributes `json:"attributes"`
}
type AdminLoginResponse struct {
	Data     AdminLogin `json:"data"`
	Included Included   `json:"included"`
}

type AdminLoginListResponse struct {
	Data     []AdminLogin `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustAdminLogin - returns AdminLogin from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAdminLogin(key Key) *AdminLogin {
	var adminLogin AdminLogin
	if c.tryFindEntry(key, &adminLogin) {
		return &adminLogin
	}
	return nil
}
