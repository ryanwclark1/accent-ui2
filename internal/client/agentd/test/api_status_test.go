/*
accent-agentd

Testing StatusAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package agentd

import (
	"context"
	openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/agentd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_agentd_StatusAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StatusAPIService GetStatus", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StatusAPI.GetStatus(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
