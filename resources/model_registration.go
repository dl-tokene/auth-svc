/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Registration struct {
	Key
	Attributes RegistrationAttributes `json:"attributes"`
}
type RegistrationResponse struct {
	Data     Registration `json:"data"`
	Included Included     `json:"included"`
}

type RegistrationListResponse struct {
	Data     []Registration `json:"data"`
	Included Included       `json:"included"`
	Links    *Links         `json:"links"`
}

// MustRegistration - returns Registration from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustRegistration(key Key) *Registration {
	var registration Registration
	if c.tryFindEntry(key, &registration) {
		return &registration
	}
	return nil
}
