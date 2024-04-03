/*
accent-confd

Testing RegistrarsAPIService

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

func Test_confd_RegistrarsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RegistrarsAPIService CreateRegistrar", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RegistrarsAPI.CreateRegistrar(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistrarsAPIService DeleteRegistrar", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var registrarId string

		httpRes, err := apiClient.RegistrarsAPI.DeleteRegistrar(context.Background(), registrarId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistrarsAPIService GetRegistrar", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var registrarId string

		resp, httpRes, err := apiClient.RegistrarsAPI.GetRegistrar(context.Background(), registrarId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistrarsAPIService GetRegistrars", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.RegistrarsAPI.GetRegistrars(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistrarsAPIService UpdateRegistrar", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var registrarId string

		httpRes, err := apiClient.RegistrarsAPI.UpdateRegistrar(context.Background(), registrarId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
