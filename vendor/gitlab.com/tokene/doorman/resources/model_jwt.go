/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Jwt struct {
	Key
	Attributes JwtAttributes `json:"attributes"`
}
type JwtResponse struct {
	Data     Jwt      `json:"data"`
	Included Included `json:"included"`
}

type JwtListResponse struct {
	Data     []Jwt    `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustJwt - returns Jwt from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustJwt(key Key) *Jwt {
	var jwt Jwt
	if c.tryFindEntry(key, &jwt) {
		return &jwt
	}
	return nil
}
