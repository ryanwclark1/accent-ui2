/*
accent-calld

Testing TrunksAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package calld

import (
	"context"
	openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_calld_TrunksAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TrunksAPIService ListTrunks", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TrunksAPI.ListTrunks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
