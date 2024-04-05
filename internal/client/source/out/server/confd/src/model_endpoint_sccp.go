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

type EndpointSccp struct {

	// SCCP Endpoint ID
	Id int32 `json:"id,omitempty"`

	Trunk TrunkRelationBase `json:"trunk,omitempty"`

	Line []LineRelationBase `json:"line,omitempty"`

	// Advanced configuration options. Options are appended at the end of a  SCCP account in the file 'sccp.conf' used by asterisk. Please consult the asterisk documentation for further details on available parameters. Because of database limitations, only the following options are allowed:   * cid_name  * cid_num  * allow  * disallow   Options must have the following the form:  ``` {   \"options\": [     [\"name1\", \"value1\"],     [\"name2\", \"value2\"]   ] } ```  The resulting configuration in sip.conf will have the following form:  ``` [1000] name1=value1 name2=value2 ```
	Options [][]string `json:"options,omitempty"`

	// The UUID of the tenant
	TenantUuid string `json:"tenant_uuid,omitempty"`
}

// AssertEndpointSccpRequired checks if the required fields are not zero-ed
func AssertEndpointSccpRequired(obj EndpointSccp) error {
	if err := AssertTrunkRelationBaseRequired(obj.Trunk); err != nil {
		return err
	}
	for _, el := range obj.Line {
		if err := AssertLineRelationBaseRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertEndpointSccpConstraints checks if the values respects the defined constraints
func AssertEndpointSccpConstraints(obj EndpointSccp) error {
	return nil
}