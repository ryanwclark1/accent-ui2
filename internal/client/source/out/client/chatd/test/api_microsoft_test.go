/*
accent-chatd

Testing MicrosoftAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package chatd

import (
	"context"
	openapiclient "github.com/ryanwclark/accent-voice/chatd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_chatd_MicrosoftAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MicrosoftAPIService UpdateTeamsPresence", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string

		resp, httpRes, err := apiClient.MicrosoftAPI.UpdateTeamsPresence(context.Background(), userUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
