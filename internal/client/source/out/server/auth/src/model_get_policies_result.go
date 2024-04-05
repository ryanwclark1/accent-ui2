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

type GetPoliciesResult struct {

	// A paginated list of policies
	Items []PolicyResult `json:"items"`

	// The number of policies matching the searched term
	Total int32 `json:"total"`
}

// AssertGetPoliciesResultRequired checks if the required fields are not zero-ed
func AssertGetPoliciesResultRequired(obj GetPoliciesResult) error {
	elements := map[string]interface{}{
		"items": obj.Items,
		"total": obj.Total,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Items {
		if err := AssertPolicyResultRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertGetPoliciesResultConstraints checks if the values respects the defined constraints
func AssertGetPoliciesResultConstraints(obj GetPoliciesResult) error {
	return nil
}
