/*
accent-chatd

Testing PresencesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package chatd

import (
	"context"
	openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/chatd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_chatd_PresencesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PresencesAPIService GetUserPresence", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string

		resp, httpRes, err := apiClient.PresencesAPI.GetUserPresence(context.Background(), userUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PresencesAPIService ListPresences", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.PresencesAPI.ListPresences(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PresencesAPIService UpdateUserPresence", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string

		httpRes, err := apiClient.PresencesAPI.UpdateUserPresence(context.Background(), userUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
