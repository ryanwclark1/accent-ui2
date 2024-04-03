/*
accent-chatd

Testing ConfigAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package chatd

import (
	"context"
	openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/chatd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_chatd_ConfigAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConfigAPIService GetConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.ConfigAPI.GetConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
