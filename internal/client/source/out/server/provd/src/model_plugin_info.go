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

type PluginInfo struct {
	PluginInfo PluginInfoObject `json:"plugin_info,omitempty"`
}

// AssertPluginInfoRequired checks if the required fields are not zero-ed
func AssertPluginInfoRequired(obj PluginInfo) error {
	if err := AssertPluginInfoObjectRequired(obj.PluginInfo); err != nil {
		return err
	}
	return nil
}

// AssertPluginInfoConstraints checks if the values respects the defined constraints
func AssertPluginInfoConstraints(obj PluginInfo) error {
	return nil
}
