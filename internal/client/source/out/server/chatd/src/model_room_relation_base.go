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

type RoomRelationBase struct {

	// The UUID of the room
	Uuid string `json:"uuid,omitempty"`
}

// AssertRoomRelationBaseRequired checks if the required fields are not zero-ed
func AssertRoomRelationBaseRequired(obj RoomRelationBase) error {
	return nil
}

// AssertRoomRelationBaseConstraints checks if the values respects the defined constraints
func AssertRoomRelationBaseConstraints(obj RoomRelationBase) error {
	return nil
}
