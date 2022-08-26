/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Address struct {
	Key
	Attributes AddressAttributes `json:"attributes"`
}
type AddressResponse struct {
	Data     Address  `json:"data"`
	Included Included `json:"included"`
}

type AddressListResponse struct {
	Data     []Address `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustAddress - returns Address from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAddress(key Key) *Address {
	var address Address
	if c.tryFindEntry(key, &address) {
		return &address
	}
	return nil
}
