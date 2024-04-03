/*
accent-auth

Testing AdminAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package auth

import (
	"context"
	openapiclient "github.com/ryanwclark/accent-voice/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_auth_AdminAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AdminAPIService UpdateAllUserEmails", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string

		httpRes, err := apiClient.AdminAPI.UpdateAllUserEmails(context.Background(), userUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
