/*
accent-chatd

Testing MessagesAPIService

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

func Test_chatd_MessagesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MessagesAPIService CreateRoomMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roomUuid string

		resp, httpRes, err := apiClient.MessagesAPI.CreateRoomMessage(context.Background(), roomUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessagesAPIService ListRoomMessage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roomUuid string

		resp, httpRes, err := apiClient.MessagesAPI.ListRoomMessage(context.Background(), roomUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MessagesAPIService ListRoomsMessages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MessagesAPI.ListRoomsMessages(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
