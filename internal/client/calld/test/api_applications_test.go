/*
accent-calld

Testing ApplicationsAPIService

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

func Test_calld_ApplicationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationsAPIService AnswerApplicationCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string

		httpRes, err := apiClient.ApplicationsAPI.AnswerApplicationCall(context.Background(), applicationUuid, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService CreateApplicationCallToNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var nodeUuid string

		resp, httpRes, err := apiClient.ApplicationsAPI.CreateApplicationCallToNode(context.Background(), applicationUuid, nodeUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService CreateApplicationCallToUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var nodeUuid string

		resp, httpRes, err := apiClient.ApplicationsAPI.CreateApplicationCallToUser(context.Background(), applicationUuid, nodeUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService CreateApplicationCalls", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string

		resp, httpRes, err := apiClient.ApplicationsAPI.CreateApplicationCalls(context.Background(), applicationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService CreateApplicationNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string

		resp, httpRes, err := apiClient.ApplicationsAPI.CreateApplicationNode(context.Background(), applicationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService DeleteApplicationCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string

		httpRes, err := apiClient.ApplicationsAPI.DeleteApplicationCall(context.Background(), applicationUuid, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService DeleteApplicationCallFromNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var nodeUuid string
		var callId string

		httpRes, err := apiClient.ApplicationsAPI.DeleteApplicationCallFromNode(context.Background(), applicationUuid, nodeUuid, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService DeleteApplicationNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var nodeUuid string

		httpRes, err := apiClient.ApplicationsAPI.DeleteApplicationNode(context.Background(), applicationUuid, nodeUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService DeletePlayback", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var playbackUuid string

		httpRes, err := apiClient.ApplicationsAPI.DeletePlayback(context.Background(), applicationUuid, playbackUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetApplication", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string

		resp, httpRes, err := apiClient.ApplicationsAPI.GetApplication(context.Background(), applicationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetApplicationCalls", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string

		resp, httpRes, err := apiClient.ApplicationsAPI.GetApplicationCalls(context.Background(), applicationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetApplicationNodes", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string

		resp, httpRes, err := apiClient.ApplicationsAPI.GetApplicationNodes(context.Background(), applicationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var nodeUuid string

		resp, httpRes, err := apiClient.ApplicationsAPI.GetNode(context.Background(), applicationUuid, nodeUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService GetSnoop", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var snoopUuid string

		resp, httpRes, err := apiClient.ApplicationsAPI.GetSnoop(context.Background(), applicationUuid, snoopUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService HoldApplicationCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string

		httpRes, err := apiClient.ApplicationsAPI.HoldApplicationCall(context.Background(), applicationUuid, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService InsertApplicationCallToNode", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var nodeUuid string
		var callId string

		httpRes, err := apiClient.ApplicationsAPI.InsertApplicationCallToNode(context.Background(), applicationUuid, nodeUuid, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService ListApplicationSnoops", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string

		resp, httpRes, err := apiClient.ApplicationsAPI.ListApplicationSnoops(context.Background(), applicationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService MuteApplicationCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string

		httpRes, err := apiClient.ApplicationsAPI.MuteApplicationCall(context.Background(), applicationUuid, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService PlayApplicationCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string

		resp, httpRes, err := apiClient.ApplicationsAPI.PlayApplicationCall(context.Background(), applicationUuid, callId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService ResumeApplicationCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string

		httpRes, err := apiClient.ApplicationsAPI.ResumeApplicationCall(context.Background(), applicationUuid, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService SendApplicationCallDTMF", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string

		httpRes, err := apiClient.ApplicationsAPI.SendApplicationCallDTMF(context.Background(), applicationUuid, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService SnoopApplicationCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string

		resp, httpRes, err := apiClient.ApplicationsAPI.SnoopApplicationCall(context.Background(), applicationUuid, callId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService StartApplicationCallMOH", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string
		var mohUuid string

		httpRes, err := apiClient.ApplicationsAPI.StartApplicationCallMOH(context.Background(), applicationUuid, callId, mohUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService StartApplicationCallProgress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string

		httpRes, err := apiClient.ApplicationsAPI.StartApplicationCallProgress(context.Background(), applicationUuid, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService StopApplicationCallMOH", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string

		httpRes, err := apiClient.ApplicationsAPI.StopApplicationCallMOH(context.Background(), applicationUuid, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService StopApplicationCallProgress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string

		httpRes, err := apiClient.ApplicationsAPI.StopApplicationCallProgress(context.Background(), applicationUuid, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService StopSnoop", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var snoopUuid string

		httpRes, err := apiClient.ApplicationsAPI.StopSnoop(context.Background(), applicationUuid, snoopUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService UnmuteApplicationCall", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var callId string

		httpRes, err := apiClient.ApplicationsAPI.UnmuteApplicationCall(context.Background(), applicationUuid, callId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ApplicationsAPIService UpdateSnoop", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationUuid string
		var snoopUuid string

		httpRes, err := apiClient.ApplicationsAPI.UpdateSnoop(context.Background(), applicationUuid, snoopUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
