/*
 * accent-chatd
 *
 * Control your message and presence from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package chatdserver

type Room struct {

	// The UUID of the room
	Uuid string `json:"uuid,omitempty"`

	// The name of the room
	Name string `json:"name,omitempty"`

	Users []RoomUser `json:"users"`
}

// AssertRoomRequired checks if the required fields are not zero-ed
func AssertRoomRequired(obj Room) error {
	elements := map[string]interface{}{
		"users": obj.Users,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Users {
		if err := AssertRoomUserRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRoomConstraints checks if the values respects the defined constraints
func AssertRoomConstraints(obj Room) error {
	return nil
}