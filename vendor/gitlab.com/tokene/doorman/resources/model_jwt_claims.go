/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type JwtClaims struct {
	Key
	Attributes JwtClaimsAttributes `json:"attributes"`
}
type JwtClaimsResponse struct {
	Data     JwtClaims `json:"data"`
	Included Included  `json:"included"`
}

type JwtClaimsListResponse struct {
	Data     []JwtClaims `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustJwtClaims - returns JwtClaims from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustJwtClaims(key Key) *JwtClaims {
	var jwtClaims JwtClaims
	if c.tryFindEntry(key, &jwtClaims) {
		return &jwtClaims
	}
	return nil
}
