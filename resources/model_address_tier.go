/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type AddressTier struct {
	Key
	Attributes AddressTierAttributes `json:"attributes"`
}
type AddressTierResponse struct {
	Data     AddressTier `json:"data"`
	Included Included    `json:"included"`
}

type AddressTierListResponse struct {
	Data     []AddressTier `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustAddressTier - returns AddressTier from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustAddressTier(key Key) *AddressTier {
	var addressTier AddressTier
	if c.tryFindEntry(key, &addressTier) {
		return &addressTier
	}
	return nil
}
