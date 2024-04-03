/*
accent-confd

Testing GroupsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package confd

import (
	"context"
	openapiclient "github.com/ryanwclark/accent-voice/confd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_confd_GroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GroupsAPIService AssociateGroupCallpermission", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string
		var callpermissionId int32

		httpRes, err := apiClient.GroupsAPI.AssociateGroupCallpermission(context.Background(), groupUuid, callpermissionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService AssociateGroupExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string
		var extensionId int32

		httpRes, err := apiClient.GroupsAPI.AssociateGroupExtension(context.Background(), groupUuid, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService AssociateGroupSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string
		var scheduleId int32

		httpRes, err := apiClient.GroupsAPI.AssociateGroupSchedule(context.Background(), groupUuid, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService CreateGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GroupsAPI.CreateGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService DeleteGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string

		httpRes, err := apiClient.GroupsAPI.DeleteGroup(context.Background(), groupUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService DissociateGroupCallpermission", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string
		var callpermissionId int32

		httpRes, err := apiClient.GroupsAPI.DissociateGroupCallpermission(context.Background(), groupUuid, callpermissionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService DissociateGroupExtension", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string
		var extensionId int32

		httpRes, err := apiClient.GroupsAPI.DissociateGroupExtension(context.Background(), groupUuid, extensionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService DissociateGroupSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string
		var scheduleId int32

		httpRes, err := apiClient.GroupsAPI.DissociateGroupSchedule(context.Background(), groupUuid, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService GetGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string

		resp, httpRes, err := apiClient.GroupsAPI.GetGroup(context.Background(), groupUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService GetGroupFallback", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string

		resp, httpRes, err := apiClient.GroupsAPI.GetGroupFallback(context.Background(), groupUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService ListGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.GroupsAPI.ListGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService UpdateCallPickupInterceptorGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callpickupId int32

		httpRes, err := apiClient.GroupsAPI.UpdateCallPickupInterceptorGroups(context.Background(), callpickupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService UpdateCallPickupTargetGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callpickupId int32

		httpRes, err := apiClient.GroupsAPI.UpdateCallPickupTargetGroups(context.Background(), callpickupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService UpdateGroup", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string

		httpRes, err := apiClient.GroupsAPI.UpdateGroup(context.Background(), groupUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService UpdateGroupFallback", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string

		httpRes, err := apiClient.GroupsAPI.UpdateGroupFallback(context.Background(), groupUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService UpdateGroupMemberExtensions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string

		httpRes, err := apiClient.GroupsAPI.UpdateGroupMemberExtensions(context.Background(), groupUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService UpdateGroupMemberUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string

		httpRes, err := apiClient.GroupsAPI.UpdateGroupMemberUsers(context.Background(), groupUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GroupsAPIService UpdateUserGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.GroupsAPI.UpdateUserGroups(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
