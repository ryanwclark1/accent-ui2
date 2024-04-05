/*
 * accent-auth
 *
 * Accent's authentication service
 *
 * API version: 0.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package authserver

type UserList struct {

	// The number of users matching the searched term
	Filtered int32 `json:"filtered,omitempty"`

	// A paginated list of users
	Items []UserResult `json:"items,omitempty"`

	// The number of users
	Total int32 `json:"total,omitempty"`
}

// AssertUserListRequired checks if the required fields are not zero-ed
func AssertUserListRequired(obj UserList) error {
	for _, el := range obj.Items {
		if err := AssertUserResultRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertUserListConstraints checks if the values respects the defined constraints
func AssertUserListConstraints(obj UserList) error {
	return nil
}
