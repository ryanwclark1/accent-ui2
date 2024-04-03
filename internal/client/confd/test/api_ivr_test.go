/*
accent-confd

Testing IvrAPIService

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

func Test_confd_IvrAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IvrAPIService CreateIvr", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IvrAPI.CreateIvr(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IvrAPIService DeleteIvr", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ivrId int32

		httpRes, err := apiClient.IvrAPI.DeleteIvr(context.Background(), ivrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IvrAPIService GetIvr", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ivrId int32

		resp, httpRes, err := apiClient.IvrAPI.GetIvr(context.Background(), ivrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IvrAPIService ListIvr", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IvrAPI.ListIvr(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IvrAPIService UpdateIvr", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var ivrId int32

		httpRes, err := apiClient.IvrAPI.UpdateIvr(context.Background(), ivrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
