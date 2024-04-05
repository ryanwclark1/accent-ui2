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

type GetGroupsResult struct {

	// The number of groups matching the searched term.
	Filtered int32 `json:"filtered"`

	// A paginated list of groups
	Items []GroupResult `json:"items"`

	// The number of groups.
	Total int32 `json:"total"`
}

// AssertGetGroupsResultRequired checks if the required fields are not zero-ed
func AssertGetGroupsResultRequired(obj GetGroupsResult) error {
	elements := map[string]interface{}{
		"filtered": obj.Filtered,
		"items":    obj.Items,
		"total":    obj.Total,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Items {
		if err := AssertGroupResultRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertGetGroupsResultConstraints checks if the values respects the defined constraints
func AssertGetGroupsResultConstraints(obj GetGroupsResult) error {
	return nil
}
