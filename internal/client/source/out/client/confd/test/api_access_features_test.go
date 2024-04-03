/*
accent-confd

Testing AccessFeaturesAPIService

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

func Test_confd_AccessFeaturesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccessFeaturesAPIService CreateAccessFeature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AccessFeaturesAPI.CreateAccessFeature(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessFeaturesAPIService DeleteAccessFeature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accessFeatureId int32

		httpRes, err := apiClient.AccessFeaturesAPI.DeleteAccessFeature(context.Background(), accessFeatureId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessFeaturesAPIService GetAccessFeature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accessFeatureId int32

		resp, httpRes, err := apiClient.AccessFeaturesAPI.GetAccessFeature(context.Background(), accessFeatureId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessFeaturesAPIService ListAccessFeatures", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AccessFeaturesAPI.ListAccessFeatures(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessFeaturesAPIService UpdateAccessFeature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var accessFeatureId int32

		httpRes, err := apiClient.AccessFeaturesAPI.UpdateAccessFeature(context.Background(), accessFeatureId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
