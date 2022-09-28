/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type JwtPair struct {
	Key
	Attributes JwtPairAttributes `json:"attributes"`
}
type JwtPairResponse struct {
	Data     JwtPair  `json:"data"`
	Included Included `json:"included"`
}

type JwtPairListResponse struct {
	Data     []JwtPair `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustJwtPair - returns JwtPair from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustJwtPair(key Key) *JwtPair {
	var jwtPair JwtPair
	if c.tryFindEntry(key, &jwtPair) {
		return &jwtPair
	}
	return nil
}
