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

type RawConfigurationObject struct {
	RawConfig map[string]interface{} `json:"raw_config,omitempty"`
}

// AssertRawConfigurationObjectRequired checks if the required fields are not zero-ed
func AssertRawConfigurationObjectRequired(obj RawConfigurationObject) error {
	return nil
}

// AssertRawConfigurationObjectConstraints checks if the values respects the defined constraints
func AssertRawConfigurationObjectConstraints(obj RawConfigurationObject) error {
	return nil
}