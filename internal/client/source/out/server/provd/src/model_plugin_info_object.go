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

type PluginInfoObject struct {
	Capabilities map[string]map[string]string `json:"capabilities,omitempty"`

	Description string `json:"description,omitempty"`

	Version string `json:"version,omitempty"`
}

// AssertPluginInfoObjectRequired checks if the required fields are not zero-ed
func AssertPluginInfoObjectRequired(obj PluginInfoObject) error {
	return nil
}

// AssertPluginInfoObjectConstraints checks if the values respects the defined constraints
func AssertPluginInfoObjectConstraints(obj PluginInfoObject) error {
	return nil
}
