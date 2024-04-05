/*
 * accent-provd
 *
 * Provisioning application REST API
 *
 * API version: 0.2
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package provdserver

// Param - A configuration parameter
type Param struct {
	Param ParamObject `json:"param,omitempty"`
}

// AssertParamRequired checks if the required fields are not zero-ed
func AssertParamRequired(obj Param) error {
	if err := AssertParamObjectRequired(obj.Param); err != nil {
		return err
	}
	return nil
}

// AssertParamConstraints checks if the values respects the defined constraints
func AssertParamConstraints(obj Param) error {
	return nil
}
