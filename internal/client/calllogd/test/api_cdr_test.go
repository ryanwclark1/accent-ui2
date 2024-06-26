/*
accent-call-logd

Testing CdrAPIService

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

func Test_calllogd_CdrAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CdrAPIService CreateCDRRecordingsMediaExport", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CdrAPI.CreateCDRRecordingsMediaExport(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CdrAPIService DeleteCDRRecordingMedia", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cdrId int32
		var recordingUuid int32

		httpRes, err := apiClient.CdrAPI.DeleteCDRRecordingMedia(context.Background(), cdrId, recordingUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CdrAPIService DeleteCDRRecordingsMedia", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.CdrAPI.DeleteCDRRecordingsMedia(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CdrAPIService GetCDR", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cdrId int32

		resp, httpRes, err := apiClient.CdrAPI.GetCDR(context.Background(), cdrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CdrAPIService GetCDRRecordingMedia", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var cdrId int32
		var recordingUuid int32

		httpRes, err := apiClient.CdrAPI.GetCDRRecordingMedia(context.Background(), cdrId, recordingUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CdrAPIService GetCDRs", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CdrAPI.GetCDRs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CdrAPIService GetCurrentUserCDR", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CdrAPI.GetCurrentUserCDR(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CdrAPIService GetUserCDR", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string

		resp, httpRes, err := apiClient.CdrAPI.GetUserCDR(context.Background(), userUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
