/*
accent-confd

Testing ExtensionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package confd

import (
	"context"
	openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_confd_ExtensionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ExtensionsAPIService AssociateConferenceExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var conferenceId int32
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.AssociateConferenceExtension(context.Background(), conferenceId, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService AssociateGroupExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.AssociateGroupExtension(context.Background(), groupUuid, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService AssociateIncallExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var incallId int32
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.AssociateIncallExtension(context.Background(), incallId, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService AssociateLineExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineId int32
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.AssociateLineExtension(context.Background(), lineId, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService AssociateOutcallExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var outcallId int32
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.AssociateOutcallExtension(context.Background(), outcallId, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService AssociateParkingLotExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var parkingLotId int32
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.AssociateParkingLotExtension(context.Background(), parkingLotId, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService AssociateQueueExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var queueId int32
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.AssociateQueueExtension(context.Background(), queueId, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService CreateExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExtensionsAPI.CreateExtension(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService CreateLineExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineId int32

		resp, httpRes, err := apiClient.ExtensionsAPI.CreateLineExtension(context.Background(), lineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService DeleteExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.DeleteExtension(context.Background(), extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService DissociateConferenceExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var conferenceId int32
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.DissociateConferenceExtension(context.Background(), conferenceId, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService DissociateGroupExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.DissociateGroupExtension(context.Background(), groupUuid, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService DissociateIncallExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var incallId int32
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.DissociateIncallExtension(context.Background(), incallId, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService DissociateLineExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineId int32
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.DissociateLineExtension(context.Background(), lineId, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService DissociateOutcallExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var outcallId int32
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.DissociateOutcallExtension(context.Background(), outcallId, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService DissociateParkingLotExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var parkingLotId int32
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.DissociateParkingLotExtension(context.Background(), parkingLotId, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService DissociateQueueExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var queueId int32
		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.DissociateQueueExtension(context.Background(), queueId, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService GetExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var extensionId int32

		resp, httpRes, err := apiClient.ExtensionsAPI.GetExtension(context.Background(), extensionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService GetExtensionFeature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var extensionUuid string

		resp, httpRes, err := apiClient.ExtensionsAPI.GetExtensionFeature(context.Background(), extensionUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService ListExtensions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExtensionsAPI.ListExtensions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService ListExtensionsFeatures", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ExtensionsAPI.ListExtensionsFeatures(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService UpdateExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var extensionId int32

		httpRes, err := apiClient.ExtensionsAPI.UpdateExtension(context.Background(), extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ExtensionsAPIService UpdateExtensionFeature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var extensionUuid string

		httpRes, err := apiClient.ExtensionsAPI.UpdateExtensionFeature(context.Background(), extensionUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
