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

type GeneralConfigsObject struct {
	Params []GeneralConfigObject `json:"params,omitempty"`
}

// AssertGeneralConfigsObjectRequired checks if the required fields are not zero-ed
func AssertGeneralConfigsObjectRequired(obj GeneralConfigsObject) error {
	for _, el := range obj.Params {
		if err := AssertGeneralConfigObjectRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertGeneralConfigsObjectConstraints checks if the values respects the defined constraints
func AssertGeneralConfigsObjectConstraints(obj GeneralConfigsObject) error {
	return nil
}
