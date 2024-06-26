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

// ProvisioningKey - Provisioning Key configuration
type ProvisioningKey struct {

	// Provisioning key used to authenticate request to fetch provisioning configuration file
	Param string `json:"param,omitempty"`
}

// AssertProvisioningKeyRequired checks if the required fields are not zero-ed
func AssertProvisioningKeyRequired(obj ProvisioningKey) error {
	return nil
}

// AssertProvisioningKeyConstraints checks if the values respects the defined constraints
func AssertProvisioningKeyConstraints(obj ProvisioningKey) error {
	return nil
}
