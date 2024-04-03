/*
accent-confd

Testing SoundsAPIService

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

func Test_confd_SoundsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SoundsAPIService CreateSounds", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SoundsAPI.CreateSounds(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SoundsAPIService DeleteSounds", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var soundCategory string

		httpRes, err := apiClient.SoundsAPI.DeleteSounds(context.Background(), soundCategory).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SoundsAPIService DeleteSoundsFiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var soundCategory string
		var soundFilename string

		httpRes, err := apiClient.SoundsAPI.DeleteSoundsFiles(context.Background(), soundCategory, soundFilename).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SoundsAPIService GetSounds", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var soundCategory string

		resp, httpRes, err := apiClient.SoundsAPI.GetSounds(context.Background(), soundCategory).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SoundsAPIService GetSoundsFiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var soundCategory string
		var soundFilename string

		httpRes, err := apiClient.SoundsAPI.GetSoundsFiles(context.Background(), soundCategory, soundFilename).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SoundsAPIService ListSounds", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SoundsAPI.ListSounds(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SoundsAPIService ListSoundsLanguages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SoundsAPI.ListSoundsLanguages(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SoundsAPIService UpdateSoundsFiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var soundCategory string
		var soundFilename string

		httpRes, err := apiClient.SoundsAPI.UpdateSoundsFiles(context.Background(), soundCategory, soundFilename).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
