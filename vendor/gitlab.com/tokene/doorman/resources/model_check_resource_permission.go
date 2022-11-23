/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CheckResourcePermission struct {
	Key
	Attributes CheckResourcePermissionAttributes `json:"attributes"`
}
type CheckResourcePermissionResponse struct {
	Data     CheckResourcePermission `json:"data"`
	Included Included                `json:"included"`
}

type CheckResourcePermissionListResponse struct {
	Data     []CheckResourcePermission `json:"data"`
	Included Included                  `json:"included"`
	Links    *Links                    `json:"links"`
}

// MustCheckResourcePermission - returns CheckResourcePermission from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustCheckResourcePermission(key Key) *CheckResourcePermission {
	var checkResourcePermission CheckResourcePermission
	if c.tryFindEntry(key, &checkResourcePermission) {
		return &checkResourcePermission
	}
	return nil
}
