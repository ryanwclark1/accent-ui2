/*
accent-calld

Testing VoicemailsAPIService

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

func Test_calld_VoicemailsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test VoicemailsAPIService CheckUserVoicemailGreeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var greeting string

		httpRes, err := apiClient.VoicemailsAPI.CheckUserVoicemailGreeting(context.Background(), greeting).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService CheckVoicemailGreeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voicemailId int32
		var greeting string

		httpRes, err := apiClient.VoicemailsAPI.CheckVoicemailGreeting(context.Background(), voicemailId, greeting).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService CopyUserVoicemailGreeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var greeting string

		httpRes, err := apiClient.VoicemailsAPI.CopyUserVoicemailGreeting(context.Background(), greeting).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService CopyVoicemailGreeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voicemailId int32
		var greeting string

		httpRes, err := apiClient.VoicemailsAPI.CopyVoicemailGreeting(context.Background(), voicemailId, greeting).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService CreateUserVoicemailGreeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var greeting string

		httpRes, err := apiClient.VoicemailsAPI.CreateUserVoicemailGreeting(context.Background(), greeting).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService CreateVoicemailGreeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voicemailId int32
		var greeting string

		httpRes, err := apiClient.VoicemailsAPI.CreateVoicemailGreeting(context.Background(), voicemailId, greeting).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService DeleteUserVoicemailGreeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var greeting string

		httpRes, err := apiClient.VoicemailsAPI.DeleteUserVoicemailGreeting(context.Background(), greeting).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService DeleteUserVoicemailMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var messageId string

		httpRes, err := apiClient.VoicemailsAPI.DeleteUserVoicemailMessage(context.Background(), messageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService DeleteVoicemailGreeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voicemailId int32
		var greeting string

		httpRes, err := apiClient.VoicemailsAPI.DeleteVoicemailGreeting(context.Background(), voicemailId, greeting).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService DeleteVoicemailMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voicemailId int32
		var messageId string

		httpRes, err := apiClient.VoicemailsAPI.DeleteVoicemailMessage(context.Background(), voicemailId, messageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService GetUserVoicemailFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var folderId int32

		resp, httpRes, err := apiClient.VoicemailsAPI.GetUserVoicemailFolder(context.Background(), folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService GetUserVoicemailGreeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var greeting string

		httpRes, err := apiClient.VoicemailsAPI.GetUserVoicemailGreeting(context.Background(), greeting).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService GetUserVoicemailMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var messageId string

		resp, httpRes, err := apiClient.VoicemailsAPI.GetUserVoicemailMessage(context.Background(), messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService GetUserVoicemailMessageRecording", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var messageId string

		httpRes, err := apiClient.VoicemailsAPI.GetUserVoicemailMessageRecording(context.Background(), messageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService GetVoicemail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voicemailId int32

		resp, httpRes, err := apiClient.VoicemailsAPI.GetVoicemail(context.Background(), voicemailId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService GetVoicemailFolder", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voicemailId int32
		var folderId int32

		resp, httpRes, err := apiClient.VoicemailsAPI.GetVoicemailFolder(context.Background(), voicemailId, folderId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService GetVoicemailGreeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voicemailId int32
		var greeting string

		httpRes, err := apiClient.VoicemailsAPI.GetVoicemailGreeting(context.Background(), voicemailId, greeting).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService GetVoicemailMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voicemailId int32
		var messageId string

		resp, httpRes, err := apiClient.VoicemailsAPI.GetVoicemailMessage(context.Background(), voicemailId, messageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService GetVoicemailMessageRecording", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voicemailId int32
		var messageId string

		httpRes, err := apiClient.VoicemailsAPI.GetVoicemailMessageRecording(context.Background(), voicemailId, messageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService ListUserVoicemails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.VoicemailsAPI.ListUserVoicemails(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService UpdateUserVoicemailGreeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var greeting string

		httpRes, err := apiClient.VoicemailsAPI.UpdateUserVoicemailGreeting(context.Background(), greeting).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService UpdateUserVoicemailMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var messageId string

		httpRes, err := apiClient.VoicemailsAPI.UpdateUserVoicemailMessage(context.Background(), messageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService UpdateVoicemailGreeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voicemailId int32
		var greeting string

		httpRes, err := apiClient.VoicemailsAPI.UpdateVoicemailGreeting(context.Background(), voicemailId, greeting).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test VoicemailsAPIService UpdateVoicemailMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var voicemailId int32
		var messageId string

		httpRes, err := apiClient.VoicemailsAPI.UpdateVoicemailMessage(context.Background(), voicemailId, messageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
