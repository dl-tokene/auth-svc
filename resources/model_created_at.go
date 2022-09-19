/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreatedAt struct {
	Key
	Attributes CreatedAtAttributes `json:"attributes"`
}
type CreatedAtResponse struct {
	Data     CreatedAt `json:"data"`
	Included Included  `json:"included"`
}

type CreatedAtListResponse struct {
	Data     []CreatedAt `json:"data"`
	Included Included    `json:"included"`
	Links    *Links      `json:"links"`
}

// MustCreatedAt - returns CreatedAt from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCreatedAt(key Key) *CreatedAt {
	var createdAt CreatedAt
	if c.tryFindEntry(key, &createdAt) {
		return &createdAt
	}
	return nil
}
