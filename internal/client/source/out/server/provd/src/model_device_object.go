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

// DeviceObject - A device object
type DeviceObject struct {
	Device Device `json:"device,omitempty"`
}

// AssertDeviceObjectRequired checks if the required fields are not zero-ed
func AssertDeviceObjectRequired(obj DeviceObject) error {
	if err := AssertDeviceRequired(obj.Device); err != nil {
		return err
	}
	return nil
}

// AssertDeviceObjectConstraints checks if the values respects the defined constraints
func AssertDeviceObjectConstraints(obj DeviceObject) error {
	return nil
}
