/*
accent-confd

Testing IngressesAPIService

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

func Test_confd_IngressesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IngressesAPIService CreateHttpIngress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IngressesAPI.CreateHttpIngress(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IngressesAPIService DeleteHttpIngress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var httpIngressUuid string

		httpRes, err := apiClient.IngressesAPI.DeleteHttpIngress(context.Background(), httpIngressUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IngressesAPIService GetHttpIngress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var httpIngressUuid string

		resp, httpRes, err := apiClient.IngressesAPI.GetHttpIngress(context.Background(), httpIngressUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IngressesAPIService ListHttpIngresses", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.IngressesAPI.ListHttpIngresses(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IngressesAPIService UpdateHttpIngress", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var httpIngressUuid string

		httpRes, err := apiClient.IngressesAPI.UpdateHttpIngress(context.Background(), httpIngressUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
