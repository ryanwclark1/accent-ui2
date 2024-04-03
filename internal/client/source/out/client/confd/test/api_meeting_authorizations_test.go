/*
accent-confd

Testing MeetingAuthorizationsAPIService

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

func Test_confd_MeetingAuthorizationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MeetingAuthorizationsAPIService CreateGuestMeetingAuthorization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var guestUuid string
		var meetingUuid string

		resp, httpRes, err := apiClient.MeetingAuthorizationsAPI.CreateGuestMeetingAuthorization(context.Background(), guestUuid, meetingUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MeetingAuthorizationsAPIService DeleteUserMeetingAuthorization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var guestUuid string
		var meetingUuid string
		var authorizationUuid string

		httpRes, err := apiClient.MeetingAuthorizationsAPI.DeleteUserMeetingAuthorization(context.Background(), guestUuid, meetingUuid, authorizationUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MeetingAuthorizationsAPIService GetGuestMeetingAuthorization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var guestUuid string
		var meetingUuid string
		var authorizationUuid string

		resp, httpRes, err := apiClient.MeetingAuthorizationsAPI.GetGuestMeetingAuthorization(context.Background(), guestUuid, meetingUuid, authorizationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MeetingAuthorizationsAPIService GetUserMeetingAuthorization", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var guestUuid string
		var meetingUuid string
		var authorizationUuid string

		resp, httpRes, err := apiClient.MeetingAuthorizationsAPI.GetUserMeetingAuthorization(context.Background(), guestUuid, meetingUuid, authorizationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MeetingAuthorizationsAPIService ListUserMeetingAuthorizations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var meetingUuid string

		resp, httpRes, err := apiClient.MeetingAuthorizationsAPI.ListUserMeetingAuthorizations(context.Background(), meetingUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MeetingAuthorizationsAPIService PutUserMeetingAuthorizationAccept", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var meetingUuid string
		var authorizationUuid string

		resp, httpRes, err := apiClient.MeetingAuthorizationsAPI.PutUserMeetingAuthorizationAccept(context.Background(), meetingUuid, authorizationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MeetingAuthorizationsAPIService PutUserMeetingAuthorizationReject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var meetingUuid string
		var authorizationUuid string

		resp, httpRes, err := apiClient.MeetingAuthorizationsAPI.PutUserMeetingAuthorizationReject(context.Background(), meetingUuid, authorizationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
