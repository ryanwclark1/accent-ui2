/*
accent-calld

Testing CallsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package calld

import (
	"context"
	openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_calld_CallsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CallsAPIService AnswerCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.AnswerCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService AnswerUserCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.AnswerUserCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService ConnectCallToUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string
		var userUuid string

		resp, httpRes, err := apiClient.CallsAPI.ConnectCallToUser(context.Background(), callId, userUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService CreateCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CallsAPI.CreateCall(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService CreateUserCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CallsAPI.CreateUserCall(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService DeleteCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.DeleteCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService GetCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		resp, httpRes, err := apiClient.CallsAPI.GetCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService HangupUserCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.HangupUserCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService HoldCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.HoldCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService HoldUserCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.HoldUserCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService ListCalls", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CallsAPI.ListCalls(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService ListUserCalls", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CallsAPI.ListUserCalls(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService MuteCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.MuteCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService MuteUserCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.MuteUserCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService SendCallDTMF", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.SendCallDTMF(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService SendUserDTMF", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.SendUserDTMF(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService StartCurrentUserRecording", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.StartCurrentUserRecording(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService StartRecording", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.StartRecording(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService StopCurrentUserRecording", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.StopCurrentUserRecording(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService StopRecording", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.StopRecording(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService UnholdCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.UnholdCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService UnholdUserCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.UnholdUserCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService UnmuteCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.UnmuteCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CallsAPIService UnmuteUserCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callId string

		httpRes, err := apiClient.CallsAPI.UnmuteUserCall(context.Background(), callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
