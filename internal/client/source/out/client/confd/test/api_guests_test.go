/*
accent-confd

Testing GuestsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package confd

import (
	"context"
	openapiclient "github.com/ryanwclark/accent-voice/confd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_confd_GuestsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GuestsAPIService CreateGuestMeetingAuthorization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var guestUuid string
		var meetingUuid string

		resp, httpRes, err := apiClient.GuestsAPI.CreateGuestMeetingAuthorization(context.Background(), guestUuid, meetingUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GuestsAPIService DeleteUserMeetingAuthorization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var guestUuid string
		var meetingUuid string
		var authorizationUuid string

		httpRes, err := apiClient.GuestsAPI.DeleteUserMeetingAuthorization(context.Background(), guestUuid, meetingUuid, authorizationUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GuestsAPIService GetGuestMeeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var meetingUuid string

		resp, httpRes, err := apiClient.GuestsAPI.GetGuestMeeting(context.Background(), meetingUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GuestsAPIService GetGuestMeetingAuthorization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var guestUuid string
		var meetingUuid string
		var authorizationUuid string

		resp, httpRes, err := apiClient.GuestsAPI.GetGuestMeetingAuthorization(context.Background(), guestUuid, meetingUuid, authorizationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GuestsAPIService GetUserMeetingAuthorization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var guestUuid string
		var meetingUuid string
		var authorizationUuid string

		resp, httpRes, err := apiClient.GuestsAPI.GetUserMeetingAuthorization(context.Background(), guestUuid, meetingUuid, authorizationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
