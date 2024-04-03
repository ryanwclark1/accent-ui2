/*
accent-call-logd

Testing StatusAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package calllogd

import (
	"context"
	openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calllogd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_calllogd_StatusAPIService(t *testing.T) {

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
