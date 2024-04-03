/*
accent-confd

Testing SchedulesAPIService

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

func Test_confd_SchedulesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SchedulesAPIService AssociateGroupSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string
		var scheduleId int32

		httpRes, err := apiClient.SchedulesAPI.AssociateGroupSchedule(context.Background(), groupUuid, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService AssociateIncallSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var incallId int32
		var scheduleId int32

		httpRes, err := apiClient.SchedulesAPI.AssociateIncallSchedule(context.Background(), incallId, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService AssociateOutcallSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var outcallId int32
		var scheduleId int32

		httpRes, err := apiClient.SchedulesAPI.AssociateOutcallSchedule(context.Background(), outcallId, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService AssociateQueueSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var queueId int32
		var scheduleId int32

		httpRes, err := apiClient.SchedulesAPI.AssociateQueueSchedule(context.Background(), queueId, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService AssociateUserSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var scheduleId int32

		httpRes, err := apiClient.SchedulesAPI.AssociateUserSchedule(context.Background(), userId, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService CreateSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SchedulesAPI.CreateSchedule(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService DeleteSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var scheduleId int32

		httpRes, err := apiClient.SchedulesAPI.DeleteSchedule(context.Background(), scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService DissociateGroupSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string
		var scheduleId int32

		httpRes, err := apiClient.SchedulesAPI.DissociateGroupSchedule(context.Background(), groupUuid, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService DissociateIncallSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var incallId int32
		var scheduleId int32

		httpRes, err := apiClient.SchedulesAPI.DissociateIncallSchedule(context.Background(), incallId, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService DissociateOutcallSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var outcallId int32
		var scheduleId int32

		httpRes, err := apiClient.SchedulesAPI.DissociateOutcallSchedule(context.Background(), outcallId, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService DissociateQueueSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var queueId int32
		var scheduleId int32

		httpRes, err := apiClient.SchedulesAPI.DissociateQueueSchedule(context.Background(), queueId, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService DissociateUserSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var scheduleId int32

		httpRes, err := apiClient.SchedulesAPI.DissociateUserSchedule(context.Background(), userId, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService GetSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var scheduleId int32

		resp, httpRes, err := apiClient.SchedulesAPI.GetSchedule(context.Background(), scheduleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService ListSchedules", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SchedulesAPI.ListSchedules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SchedulesAPIService UpdateSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var scheduleId int32

		httpRes, err := apiClient.SchedulesAPI.UpdateSchedule(context.Background(), scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
