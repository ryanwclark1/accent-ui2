/*
accent-confd

Testing EndpointsAPIService

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

func Test_confd_EndpointsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EndpointsAPIService AssociateLineEndpointCustom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineId int32
		var customId int32

		httpRes, err := apiClient.EndpointsAPI.AssociateLineEndpointCustom(context.Background(), lineId, customId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService AssociateLineEndpointSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineId int32
		var sccpId int32

		httpRes, err := apiClient.EndpointsAPI.AssociateLineEndpointSccp(context.Background(), lineId, sccpId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService AssociateLineEndpointSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineId int32
		var sipUuid string

		httpRes, err := apiClient.EndpointsAPI.AssociateLineEndpointSip(context.Background(), lineId, sipUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService AssociateTrunkEndpointCustom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var trunkId int32
		var customId int32

		httpRes, err := apiClient.EndpointsAPI.AssociateTrunkEndpointCustom(context.Background(), trunkId, customId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService AssociateTrunkEndpointIax", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var trunkId int32
		var iaxId int32

		httpRes, err := apiClient.EndpointsAPI.AssociateTrunkEndpointIax(context.Background(), trunkId, iaxId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService AssociateTrunkEndpointSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var trunkId int32
		var sipUuid string

		httpRes, err := apiClient.EndpointsAPI.AssociateTrunkEndpointSip(context.Background(), trunkId, sipUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService CreateEndpointCustom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EndpointsAPI.CreateEndpointCustom(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService CreateEndpointIax", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EndpointsAPI.CreateEndpointIax(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService CreateEndpointSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EndpointsAPI.CreateEndpointSccp(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService CreateEndpointSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EndpointsAPI.CreateEndpointSip(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService CreateEndpointSipTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EndpointsAPI.CreateEndpointSipTemplate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService DeleteEndpointCustom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customId int32

		httpRes, err := apiClient.EndpointsAPI.DeleteEndpointCustom(context.Background(), customId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService DeleteEndpointIax", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var iaxId int32

		httpRes, err := apiClient.EndpointsAPI.DeleteEndpointIax(context.Background(), iaxId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService DeleteEndpointSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sccpId int32

		httpRes, err := apiClient.EndpointsAPI.DeleteEndpointSccp(context.Background(), sccpId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService DeleteEndpointSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sipUuid string

		httpRes, err := apiClient.EndpointsAPI.DeleteEndpointSip(context.Background(), sipUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService DeleteEndpointSipTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var templateUuid string

		httpRes, err := apiClient.EndpointsAPI.DeleteEndpointSipTemplate(context.Background(), templateUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService DissociateLineEndpointCustom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineId int32
		var customId int32

		httpRes, err := apiClient.EndpointsAPI.DissociateLineEndpointCustom(context.Background(), lineId, customId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService DissociateLineEndpointSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineId int32
		var sccpId int32

		httpRes, err := apiClient.EndpointsAPI.DissociateLineEndpointSccp(context.Background(), lineId, sccpId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService DissociateLineEndpointSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var lineId int32
		var sipUuid string

		httpRes, err := apiClient.EndpointsAPI.DissociateLineEndpointSip(context.Background(), lineId, sipUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService DissociateTrunkEndpointCustom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var trunkId int32
		var customId int32

		httpRes, err := apiClient.EndpointsAPI.DissociateTrunkEndpointCustom(context.Background(), trunkId, customId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService DissociateTrunkEndpointIax", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var trunkId int32
		var iaxId int32

		httpRes, err := apiClient.EndpointsAPI.DissociateTrunkEndpointIax(context.Background(), trunkId, iaxId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService DissociateTrunkEndpointSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var trunkId int32
		var sipUuid string

		httpRes, err := apiClient.EndpointsAPI.DissociateTrunkEndpointSip(context.Background(), trunkId, sipUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService GetEndpointCustom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customId int32

		resp, httpRes, err := apiClient.EndpointsAPI.GetEndpointCustom(context.Background(), customId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService GetEndpointIax", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var iaxId int32

		resp, httpRes, err := apiClient.EndpointsAPI.GetEndpointIax(context.Background(), iaxId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService GetEndpointSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sccpId int32

		resp, httpRes, err := apiClient.EndpointsAPI.GetEndpointSccp(context.Background(), sccpId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService GetEndpointSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sipUuid string

		resp, httpRes, err := apiClient.EndpointsAPI.GetEndpointSip(context.Background(), sipUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService GetEndpointSipTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var templateUuid string

		resp, httpRes, err := apiClient.EndpointsAPI.GetEndpointSipTemplate(context.Background(), templateUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService GetUserLineAssociatedEndpointsSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string
		var lineId int32

		resp, httpRes, err := apiClient.EndpointsAPI.GetUserLineAssociatedEndpointsSip(context.Background(), userUuid, lineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService GetUserLineMainAssociatedEndpointsSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string

		resp, httpRes, err := apiClient.EndpointsAPI.GetUserLineMainAssociatedEndpointsSip(context.Background(), userUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService ListEndpointsCustom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EndpointsAPI.ListEndpointsCustom(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService ListEndpointsIax", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EndpointsAPI.ListEndpointsIax(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService ListEndpointsSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EndpointsAPI.ListEndpointsSccp(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService ListEndpointsSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EndpointsAPI.ListEndpointsSip(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService ListEndpointsSipTemplates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.EndpointsAPI.ListEndpointsSipTemplates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService UpdateEndpointCustom", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var customId int32

		httpRes, err := apiClient.EndpointsAPI.UpdateEndpointCustom(context.Background(), customId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService UpdateEndpointIax", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var iaxId int32

		httpRes, err := apiClient.EndpointsAPI.UpdateEndpointIax(context.Background(), iaxId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService UpdateEndpointSccp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sccpId int32

		httpRes, err := apiClient.EndpointsAPI.UpdateEndpointSccp(context.Background(), sccpId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService UpdateEndpointSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var sipUuid string

		httpRes, err := apiClient.EndpointsAPI.UpdateEndpointSip(context.Background(), sipUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EndpointsAPIService UpdateEndpointSipTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var templateUuid string

		httpRes, err := apiClient.EndpointsAPI.UpdateEndpointSipTemplate(context.Background(), templateUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
