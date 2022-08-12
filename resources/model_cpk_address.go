/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CpkAddress struct {
	Key
	Attributes CpkAddressAttributes `json:"attributes"`
}
type CpkAddressResponse struct {
	Data     CpkAddress `json:"data"`
	Included Included   `json:"included"`
}

type CpkAddressListResponse struct {
	Data     []CpkAddress `json:"data"`
	Included Included     `json:"included"`
	Links    *Links       `json:"links"`
}

// MustCpkAddress - returns CpkAddress from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCpkAddress(key Key) *CpkAddress {
	var cPKAddress CpkAddress
	if c.tryFindEntry(key, &cPKAddress) {
		return &cPKAddress
	}
	return nil
}
