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

type DhcpInfoObject struct {
	DhcpInfo DhcpInfo `json:"dhcp_info,omitempty"`
}

// AssertDhcpInfoObjectRequired checks if the required fields are not zero-ed
func AssertDhcpInfoObjectRequired(obj DhcpInfoObject) error {
	if err := AssertDhcpInfoRequired(obj.DhcpInfo); err != nil {
		return err
	}
	return nil
}

// AssertDhcpInfoObjectConstraints checks if the values respects the defined constraints
func AssertDhcpInfoObjectConstraints(obj DhcpInfoObject) error {
	return nil
}
