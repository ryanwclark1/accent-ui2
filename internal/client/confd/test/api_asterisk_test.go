/*
accent-confd

Testing AsteriskAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package confd

import (
	"context"
	openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_confd_AsteriskAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AsteriskAPIService ListAsteriskConfbridgeAccentDefaultBridge", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskConfbridgeAccentDefaultBridge(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskConfbridgeAccentDefaultUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskConfbridgeAccentDefaultUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskFeaturesApplicationmap", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskFeaturesApplicationmap(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskFeaturesFeaturemap", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskFeaturesFeaturemap(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskFeaturesGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskFeaturesGeneral(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskHepGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskHepGeneral(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskIaxCallnumberlimits", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskIaxCallnumberlimits(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskIaxGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskIaxGeneral(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskPjsipGlobal", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskPjsipGlobal(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskPjsipSystem", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskPjsipSystem(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskQueueGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskQueueGeneral(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskRtpGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskRtpGeneral(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskRtpIceHostCandidates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskRtpIceHostCandidates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskSccpGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskSccpGeneral(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskVoicemailGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskVoicemailGeneral(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ListAsteriskVoicemailZonemessages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ListAsteriskVoicemailZonemessages(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService ShowPjsipDoc", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.AsteriskAPI.ShowPjsipDoc(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskConfbridgeAccentDefaultBridge", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskConfbridgeAccentDefaultBridge(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskConfbridgeAccentDefaultUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskConfbridgeAccentDefaultUser(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskFeaturesApplicationmap", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskFeaturesApplicationmap(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskFeaturesFeaturemap", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskFeaturesFeaturemap(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskFeaturesGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskFeaturesGeneral(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskHepGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskHepGeneral(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskIaxCallnumberlimits", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskIaxCallnumberlimits(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskIaxGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskIaxGeneral(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskPjsipGlobal", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskPjsipGlobal(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskPjsipSystem", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskPjsipSystem(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskQueueGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskQueueGeneral(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskRtpGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskRtpGeneral(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskRtpIceHostCandidates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskRtpIceHostCandidates(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskSccpGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskSccpGeneral(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskVoicemailGeneral", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskVoicemailGeneral(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AsteriskAPIService UpdateAsteriskVoicemailZonemessages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.AsteriskAPI.UpdateAsteriskVoicemailZonemessages(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
