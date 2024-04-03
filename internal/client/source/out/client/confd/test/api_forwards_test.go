/*
accent-confd

Testing ForwardsAPIService

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

func Test_confd_ForwardsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ForwardsAPIService GetUserForward", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var forwardName string

		resp, httpRes, err := apiClient.ForwardsAPI.GetUserForward(context.Background(), userId, forwardName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForwardsAPIService ListUserForwards", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.ForwardsAPI.ListUserForwards(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForwardsAPIService UpdateUserForward", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var forwardName string

		httpRes, err := apiClient.ForwardsAPI.UpdateUserForward(context.Background(), userId, forwardName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ForwardsAPIService UpdateUserForwards", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.ForwardsAPI.UpdateUserForwards(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
