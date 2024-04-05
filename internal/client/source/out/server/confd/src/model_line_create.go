/*
 * accent-confd
 *
 * Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.
 *
 * API version: 1.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package confdserver

type LineCreate struct {

	// Name to display when calling
	CallerIdName string `json:"caller_id_name,omitempty"`

	// Number to display when calling
	CallerIdNum string `json:"caller_id_num,omitempty"`

	// The name of an internal context
	Context string `json:"context"`

	// Line ID
	Id int32 `json:"id,omitempty"`

	// Line's position on the device
	Position int32 `json:"position,omitempty"`

	// Code used to provision a device
	ProvisioningCode string `json:"provisioning_code,omitempty"`

	// Name of the template line used by the device
	Registrar string `json:"registrar,omitempty"`

	Extensions []ExtensionRelationBase `json:"extensions,omitempty"`

	EndpointSip EndpointSip `json:"endpoint_sip,omitempty"`

	EndpointSccp EndpointSccp `json:"endpoint_sccp,omitempty"`

	EndpointCustom EndpointCustom `json:"endpoint_custom,omitempty"`
}

// AssertLineCreateRequired checks if the required fields are not zero-ed
func AssertLineCreateRequired(obj LineCreate) error {
	elements := map[string]interface{}{
		"context": obj.Context,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Extensions {
		if err := AssertExtensionRelationBaseRequired(el); err != nil {
			return err
		}
	}
	if err := AssertEndpointSipRequired(obj.EndpointSip); err != nil {
		return err
	}
	if err := AssertEndpointSccpRequired(obj.EndpointSccp); err != nil {
		return err
	}
	if err := AssertEndpointCustomRequired(obj.EndpointCustom); err != nil {
		return err
	}
	return nil
}

// AssertLineCreateConstraints checks if the values respects the defined constraints
func AssertLineCreateConstraints(obj LineCreate) error {
	return nil
}
