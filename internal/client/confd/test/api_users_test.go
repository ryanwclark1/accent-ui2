/*
accent-confd

Testing UsersAPIService

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

func Test_confd_UsersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UsersAPIService AssociateUserAgent", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var agentId int32

		httpRes, err := apiClient.UsersAPI.AssociateUserAgent(context.Background(), userId, agentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService AssociateUserCallpermission", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var callpermissionId int32

		httpRes, err := apiClient.UsersAPI.AssociateUserCallpermission(context.Background(), userId, callpermissionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService AssociateUserFuncKeyTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var templateId int32

		httpRes, err := apiClient.UsersAPI.AssociateUserFuncKeyTemplate(context.Background(), userId, templateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService AssociateUserLine", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var lineId int32

		httpRes, err := apiClient.UsersAPI.AssociateUserLine(context.Background(), userId, lineId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService AssociateUserLines", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.UsersAPI.AssociateUserLines(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService AssociateUserSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var scheduleId int32

		httpRes, err := apiClient.UsersAPI.AssociateUserSchedule(context.Background(), userId, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService AssociateUserVoicemail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var voicemailId int32

		httpRes, err := apiClient.UsersAPI.AssociateUserVoicemail(context.Background(), userId, voicemailId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService CreateUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.CreateUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService CreateUserExternalApp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string
		var appName string

		resp, httpRes, err := apiClient.UsersAPI.CreateUserExternalApp(context.Background(), userUuid, appName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService CreateUserMeeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.CreateUserMeeting(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService CreateUserVoicemail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UsersAPI.CreateUserVoicemail(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DeleteUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.UsersAPI.DeleteUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DeleteUserExternalApp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string
		var appName string

		httpRes, err := apiClient.UsersAPI.DeleteUserExternalApp(context.Background(), userUuid, appName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DeleteUserFuncKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var position int32

		httpRes, err := apiClient.UsersAPI.DeleteUserFuncKey(context.Background(), userId, position).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DeleteUserMeeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var meetingUuid string

		httpRes, err := apiClient.UsersAPI.DeleteUserMeeting(context.Background(), meetingUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DissociateUserAgent", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.UsersAPI.DissociateUserAgent(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DissociateUserCallpermission", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var callpermissionId int32

		httpRes, err := apiClient.UsersAPI.DissociateUserCallpermission(context.Background(), userId, callpermissionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DissociateUserFuncKeyTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var templateId int32

		httpRes, err := apiClient.UsersAPI.DissociateUserFuncKeyTemplate(context.Background(), userId, templateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DissociateUserLine", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var lineId int32

		httpRes, err := apiClient.UsersAPI.DissociateUserLine(context.Background(), userId, lineId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DissociateUserQueue", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var queueId int32
		var userId string

		httpRes, err := apiClient.UsersAPI.DissociateUserQueue(context.Background(), queueId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DissociateUserSchedule", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var scheduleId int32

		httpRes, err := apiClient.UsersAPI.DissociateUserSchedule(context.Background(), userId, scheduleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService DissociateUserVoicemail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.UsersAPI.DissociateUserVoicemail(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ExportUsersCsv", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.ExportUsersCsv(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UsersAPI.GetUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserExternalApp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string
		var appName string

		resp, httpRes, err := apiClient.UsersAPI.GetUserExternalApp(context.Background(), userUuid, appName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserFallback", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UsersAPI.GetUserFallback(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserForward", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var forwardName string

		resp, httpRes, err := apiClient.UsersAPI.GetUserForward(context.Background(), userId, forwardName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserFuncKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var position int32

		resp, httpRes, err := apiClient.UsersAPI.GetUserFuncKey(context.Background(), userId, position).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserLineAssociatedEndpointsSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string
		var lineId int32

		resp, httpRes, err := apiClient.UsersAPI.GetUserLineAssociatedEndpointsSip(context.Background(), userUuid, lineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserLineMainAssociatedEndpointsSip", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string

		resp, httpRes, err := apiClient.UsersAPI.GetUserLineMainAssociatedEndpointsSip(context.Background(), userUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserMeeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var meetingUuid string

		resp, httpRes, err := apiClient.UsersAPI.GetUserMeeting(context.Background(), meetingUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserService", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var serviceName string

		resp, httpRes, err := apiClient.UsersAPI.GetUserService(context.Background(), userId, serviceName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserServices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UsersAPI.GetUserServices(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUserVoicemail", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UsersAPI.GetUserVoicemail(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService GetUsersSubscriptions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.GetUsersSubscriptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService HeadUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.UsersAPI.HeadUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ImportUsersCsv", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.ImportUsersCsv(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListFuncKeyTemplateUserAssociations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var templateId int32

		resp, httpRes, err := apiClient.UsersAPI.ListFuncKeyTemplateUserAssociations(context.Background(), templateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.ListUser(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListUserExternalApps", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string

		resp, httpRes, err := apiClient.UsersAPI.ListUserExternalApps(context.Background(), userUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListUserForwards", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UsersAPI.ListUserForwards(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListUserFuncKeyTemplateAssociations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UsersAPI.ListUserFuncKeyTemplateAssociations(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListUserFuncKeys", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.UsersAPI.ListUserFuncKeys(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListUserMeetingAuthorizations", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var meetingUuid string

		resp, httpRes, err := apiClient.UsersAPI.ListUserMeetingAuthorizations(context.Background(), meetingUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService ListUserMeetings", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.ListUserMeetings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService PutUserMeetingAuthorizationAccept", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var meetingUuid string
		var authorizationUuid string

		resp, httpRes, err := apiClient.UsersAPI.PutUserMeetingAuthorizationAccept(context.Background(), meetingUuid, authorizationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService PutUserMeetingAuthorizationReject", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var meetingUuid string
		var authorizationUuid string

		resp, httpRes, err := apiClient.UsersAPI.PutUserMeetingAuthorizationReject(context.Background(), meetingUuid, authorizationUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateCallFilterCallerUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callfilterId int32

		httpRes, err := apiClient.UsersAPI.UpdateCallFilterCallerUsers(context.Background(), callfilterId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateCallFilterMemberUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callfilterId int32

		httpRes, err := apiClient.UsersAPI.UpdateCallFilterMemberUsers(context.Background(), callfilterId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateCallPickupInterceptorUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callpickupId int32

		httpRes, err := apiClient.UsersAPI.UpdateCallPickupInterceptorUsers(context.Background(), callpickupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateCallPickupTargetUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var callpickupId int32

		httpRes, err := apiClient.UsersAPI.UpdateCallPickupTargetUsers(context.Background(), callpickupId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateGroupMemberUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var groupUuid string

		httpRes, err := apiClient.UsersAPI.UpdateGroupMemberUsers(context.Background(), groupUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdatePagingCallerUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var pagingId int32

		httpRes, err := apiClient.UsersAPI.UpdatePagingCallerUsers(context.Background(), pagingId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdatePagingMemberUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var pagingId int32

		httpRes, err := apiClient.UsersAPI.UpdatePagingMemberUsers(context.Background(), pagingId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateSwitchboardMemberUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var switchboardUuid string

		httpRes, err := apiClient.UsersAPI.UpdateSwitchboardMemberUsers(context.Background(), switchboardUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.UsersAPI.UpdateUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserExternalApp", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userUuid string
		var appName string

		httpRes, err := apiClient.UsersAPI.UpdateUserExternalApp(context.Background(), userUuid, appName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserFallback", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.UsersAPI.UpdateUserFallback(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserForward", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var forwardName string

		httpRes, err := apiClient.UsersAPI.UpdateUserForward(context.Background(), userId, forwardName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserForwards", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.UsersAPI.UpdateUserForwards(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserFuncKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var position int32

		httpRes, err := apiClient.UsersAPI.UpdateUserFuncKey(context.Background(), userId, position).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserFuncKeys", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.UsersAPI.UpdateUserFuncKeys(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserGroups", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.UsersAPI.UpdateUserGroups(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserMeeting", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var meetingUuid string

		httpRes, err := apiClient.UsersAPI.UpdateUserMeeting(context.Background(), meetingUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserQueueAssociation", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var queueId int32
		var userId string

		httpRes, err := apiClient.UsersAPI.UpdateUserQueueAssociation(context.Background(), queueId, userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserService", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var serviceName string

		httpRes, err := apiClient.UsersAPI.UpdateUserService(context.Background(), userId, serviceName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUserServices", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		httpRes, err := apiClient.UsersAPI.UpdateUserServices(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersAPIService UpdateUsersCsv", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.UsersAPI.UpdateUsersCsv(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
