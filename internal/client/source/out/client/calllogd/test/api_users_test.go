/*
accent-call-logd

Testing UsersAPIService

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

func Test_calllogd_UsersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UsersAPIService GetCurrentUserCDR", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.GetCurrentUserCDR(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserCDR", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string

		resp, httpRes, err := apiClient.UsersAPI.GetUserCDR(context.Background(), userUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}