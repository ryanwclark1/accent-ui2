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

type QueueFallbacks struct {
	BusyDestination DestinationType `json:"busy_destination,omitempty"`

	CongestionDestination DestinationType `json:"congestion_destination,omitempty"`

	FailDestination DestinationType `json:"fail_destination,omitempty"`

	NoanswerDestination DestinationType `json:"noanswer_destination,omitempty"`
}

// AssertQueueFallbacksRequired checks if the required fields are not zero-ed
func AssertQueueFallbacksRequired(obj QueueFallbacks) error {
	if err := AssertDestinationTypeRequired(obj.BusyDestination); err != nil {
		return err
	}
	if err := AssertDestinationTypeRequired(obj.CongestionDestination); err != nil {
		return err
	}
	if err := AssertDestinationTypeRequired(obj.FailDestination); err != nil {
		return err
	}
	if err := AssertDestinationTypeRequired(obj.NoanswerDestination); err != nil {
		return err
	}
	return nil
}

// AssertQueueFallbacksConstraints checks if the values respects the defined constraints
func AssertQueueFallbacksConstraints(obj QueueFallbacks) error {
	return nil
}
