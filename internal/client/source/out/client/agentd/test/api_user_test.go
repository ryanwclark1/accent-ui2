/*
accent-agentd

Testing UserAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package agentd

import (
	"context"
	openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_agentd_UserAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserAPIService GetUserAgent", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UserAPI.GetUserAgent(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService LoginUserAgent", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.UserAPI.LoginUserAgent(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService LogoffUserAgent", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.UserAPI.LogoffUserAgent(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService PauseUserAgent", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.UserAPI.PauseUserAgent(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService UnpauseUserAgent", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.UserAPI.UnpauseUserAgent(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
