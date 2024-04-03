/*
accent-auth

Testing SessionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package auth

import (
	"context"
	openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_auth_SessionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SessionsAPIService DeleteSession", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sessionUuid string

		httpRes, err := apiClient.SessionsAPI.DeleteSession(context.Background(), sessionUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SessionsAPIService GetSessions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SessionsAPI.GetSessions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SessionsAPIService GetUserSessions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string

		resp, httpRes, err := apiClient.SessionsAPI.GetUserSessions(context.Background(), userUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SessionsAPIService UserDeleteSession", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string
		var sessionUuid string

		httpRes, err := apiClient.SessionsAPI.UserDeleteSession(context.Background(), userUuid, sessionUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
