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

type GeneralConfigObject struct {
	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`

	Links LinksObject `json:"links,omitempty"`

	Value string `json:"value,omitempty"`
}

// AssertGeneralConfigObjectRequired checks if the required fields are not zero-ed
func AssertGeneralConfigObjectRequired(obj GeneralConfigObject) error {
	if err := AssertLinksObjectRequired(obj.Links); err != nil {
		return err
	}
	return nil
}

// AssertGeneralConfigObjectConstraints checks if the values respects the defined constraints
func AssertGeneralConfigObjectConstraints(obj GeneralConfigObject) error {
	return nil
}
