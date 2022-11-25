/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type JwtValidation struct {
	Key
	Attributes JwtValidationAttributes `json:"attributes"`
}
type JwtValidationResponse struct {
	Data     JwtValidation `json:"data"`
	Included Included      `json:"included"`
}

type JwtValidationListResponse struct {
	Data     []JwtValidation `json:"data"`
	Included Included        `json:"included"`
	Links    *Links          `json:"links"`
}

// MustJwtValidation - returns JwtValidation from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustJwtValidation(key Key) *JwtValidation {
	var jwtValidation JwtValidation
	if c.tryFindEntry(key, &jwtValidation) {
		return &jwtValidation
	}
	return nil
}
