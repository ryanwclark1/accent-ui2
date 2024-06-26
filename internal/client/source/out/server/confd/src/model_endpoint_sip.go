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

type EndpointSip struct {

	// A list of PJSIP AOR section options for this endpoint
	AorSectionOptions [][]string `json:"aor_section_options,omitempty"`

	// The ID of the Asterisk onto which this configuration applies.
	AsteriskId string `json:"asterisk_id,omitempty"`

	// A list of PJSIP auth section options for this endpoint
	AuthSectionOptions [][]string `json:"auth_section_options,omitempty"`

	// A list of PJSIP endpoint section options for this endpoint
	EndpointSectionOptions [][]string `json:"endpoint_section_options,omitempty"`

	// A list of PJSIP identify section options for this endpoint
	IdentifySectionOptions [][]string `json:"identify_section_options,omitempty"`

	// The human readable name for this configuration
	Label string `json:"label,omitempty"`

	// The name of the PJSIP entity, auto-generated if missing
	Name string `json:"name,omitempty"`

	// A list of PJSIP auth section options for this endpoint
	OutboundAuthSectionOptions [][]string `json:"outbound_auth_section_options,omitempty"`

	// A list of PJSIP auth section options for this endpoint
	RegistrationOutboundAuthSectionOptions [][]string `json:"registration_outbound_auth_section_options,omitempty"`

	// A list of PJSIP registration section options for this endpoint
	RegistrationSectionOptions [][]string `json:"registration_section_options,omitempty"`

	// A list of templates this configuration will inherit from.  The inheritance only applies to option sections. Not to the name, label and other fields. For a given list of templates [A, B, C] A will be applied over B. C will be applied over (A,B) etc.
	Templates []EndpointSipTemplatesRelation `json:"templates,omitempty"`

	// The UUID of the tenant
	TenantUuid string `json:"tenant_uuid,omitempty"`

	Transport SipTransportRelation `json:"transport,omitempty"`

	// The UUID of this resource
	Uuid string `json:"uuid,omitempty"`
}

// AssertEndpointSipRequired checks if the required fields are not zero-ed
func AssertEndpointSipRequired(obj EndpointSip) error {
	for _, el := range obj.Templates {
		if err := AssertEndpointSipTemplatesRelationRequired(el); err != nil {
			return err
		}
	}
	if err := AssertSipTransportRelationRequired(obj.Transport); err != nil {
		return err
	}
	return nil
}

// AssertEndpointSipConstraints checks if the values respects the defined constraints
func AssertEndpointSipConstraints(obj EndpointSip) error {
	return nil
}
