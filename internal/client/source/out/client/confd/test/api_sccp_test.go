/*
accent-confd

Testing SccpAPIService

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

func Test_confd_SccpAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SccpAPIService AssociateLineEndpointSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineId int32
		var sccpId int32

		httpRes, err := apiClient.SccpAPI.AssociateLineEndpointSccp(context.Background(), lineId, sccpId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SccpAPIService CreateEndpointSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SccpAPI.CreateEndpointSccp(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SccpAPIService DeleteEndpointSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sccpId int32

		httpRes, err := apiClient.SccpAPI.DeleteEndpointSccp(context.Background(), sccpId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SccpAPIService DissociateLineEndpointSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineId int32
		var sccpId int32

		httpRes, err := apiClient.SccpAPI.DissociateLineEndpointSccp(context.Background(), lineId, sccpId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SccpAPIService GetEndpointSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sccpId int32

		resp, httpRes, err := apiClient.SccpAPI.GetEndpointSccp(context.Background(), sccpId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SccpAPIService ListAsteriskSccpGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SccpAPI.ListAsteriskSccpGeneral(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SccpAPIService ListEndpointsSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SccpAPI.ListEndpointsSccp(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SccpAPIService UpdateAsteriskSccpGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.SccpAPI.UpdateAsteriskSccpGeneral(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SccpAPIService UpdateEndpointSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sccpId int32

		httpRes, err := apiClient.SccpAPI.UpdateEndpointSccp(context.Background(), sccpId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
