# Go API client for calld

Control your calls from a REST API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 0.0.1
- Generator version: 7.5.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://accentvoice.io/](https://accentvoice.io/)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import calld "github.com/ryanwclark1/accent-ui2/internal/client/calld"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `calld.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), calld.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `calld.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), calld.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `calld.ContextOperationServerIndices` and `calld.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), calld.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), calld.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://api.accentvoice.io/1.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdhocConferencesAPI* | [**AddParticipantToAdhocConference**](docs/AdhocConferencesAPI.md#addparticipanttoadhocconference) | **Put** /users/me/conferences/adhoc/{conference_id}/participants/{call_id} | Add a participant into an adhoc conference
*AdhocConferencesAPI* | [**CreateAdhocConference**](docs/AdhocConferencesAPI.md#createadhocconference) | **Post** /users/me/conferences/adhoc | Create an adhoc conference
*AdhocConferencesAPI* | [**DeleteAdhocConference**](docs/AdhocConferencesAPI.md#deleteadhocconference) | **Delete** /users/me/conferences/adhoc/{conference_id} | Delete an adhoc conference
*AdhocConferencesAPI* | [**RemoveParticipantFromAdhocConference**](docs/AdhocConferencesAPI.md#removeparticipantfromadhocconference) | **Delete** /users/me/conferences/adhoc/{conference_id}/participants/{call_id} | Remove a participant from an adhoc conference
*ApplicationsAPI* | [**AnswerApplicationCall**](docs/ApplicationsAPI.md#answerapplicationcall) | **Put** /applications/{application_uuid}/calls/{call_id}/answer | Answer a call
*ApplicationsAPI* | [**CreateApplicationCallToNode**](docs/ApplicationsAPI.md#createapplicationcalltonode) | **Post** /applications/{application_uuid}/nodes/{node_uuid}/calls | Make a new call to the node
*ApplicationsAPI* | [**CreateApplicationCallToUser**](docs/ApplicationsAPI.md#createapplicationcalltouser) | **Post** /applications/{application_uuid}/nodes/{node_uuid}/calls/user | Initiate a call to a user and insert it in the node
*ApplicationsAPI* | [**CreateApplicationCalls**](docs/ApplicationsAPI.md#createapplicationcalls) | **Post** /applications/{application_uuid}/calls | Make a new call to the application
*ApplicationsAPI* | [**CreateApplicationNode**](docs/ApplicationsAPI.md#createapplicationnode) | **Post** /applications/{application_uuid}/nodes | Make a new node and add calls
*ApplicationsAPI* | [**DeleteApplicationCall**](docs/ApplicationsAPI.md#deleteapplicationcall) | **Delete** /applications/{application_uuid}/calls/{call_id} | Hangup a call from the application
*ApplicationsAPI* | [**DeleteApplicationCallFromNode**](docs/ApplicationsAPI.md#deleteapplicationcallfromnode) | **Delete** /applications/{application_uuid}/nodes/{node_uuid}/calls/{call_id} | Remove call from the node
*ApplicationsAPI* | [**DeleteApplicationNode**](docs/ApplicationsAPI.md#deleteapplicationnode) | **Delete** /applications/{application_uuid}/nodes/{node_uuid} | Delete node and hangup all calls
*ApplicationsAPI* | [**DeletePlayback**](docs/ApplicationsAPI.md#deleteplayback) | **Delete** /applications/{application_uuid}/playbacks/{playback_uuid} | Stop and remove playback
*ApplicationsAPI* | [**GetApplication**](docs/ApplicationsAPI.md#getapplication) | **Get** /applications/{application_uuid} | Show an application
*ApplicationsAPI* | [**GetApplicationCalls**](docs/ApplicationsAPI.md#getapplicationcalls) | **Get** /applications/{application_uuid}/calls | List calls from the application
*ApplicationsAPI* | [**GetApplicationNodes**](docs/ApplicationsAPI.md#getapplicationnodes) | **Get** /applications/{application_uuid}/nodes | List nodes from the application
*ApplicationsAPI* | [**GetNode**](docs/ApplicationsAPI.md#getnode) | **Get** /applications/{application_uuid}/nodes/{node_uuid} | Show a node
*ApplicationsAPI* | [**GetSnoop**](docs/ApplicationsAPI.md#getsnoop) | **Get** /applications/{application_uuid}/snoops/{snoop_uuid} | View snooping parameters
*ApplicationsAPI* | [**HoldApplicationCall**](docs/ApplicationsAPI.md#holdapplicationcall) | **Put** /applications/{application_uuid}/calls/{call_id}/hold/start | Place a call on hold
*ApplicationsAPI* | [**InsertApplicationCallToNode**](docs/ApplicationsAPI.md#insertapplicationcalltonode) | **Put** /applications/{application_uuid}/nodes/{node_uuid}/calls/{call_id} | Insert call to the node
*ApplicationsAPI* | [**ListApplicationSnoops**](docs/ApplicationsAPI.md#listapplicationsnoops) | **Get** /applications/{application_uuid}/snoops | List active snoops
*ApplicationsAPI* | [**MuteApplicationCall**](docs/ApplicationsAPI.md#muteapplicationcall) | **Put** /applications/{application_uuid}/calls/{call_id}/mute/start | Mute a call
*ApplicationsAPI* | [**PlayApplicationCall**](docs/ApplicationsAPI.md#playapplicationcall) | **Post** /applications/{application_uuid}/calls/{call_id}/playbacks | Play file to the call
*ApplicationsAPI* | [**ResumeApplicationCall**](docs/ApplicationsAPI.md#resumeapplicationcall) | **Put** /applications/{application_uuid}/calls/{call_id}/hold/stop | Resume a call that has been placed on hold
*ApplicationsAPI* | [**SendApplicationCallDTMF**](docs/ApplicationsAPI.md#sendapplicationcalldtmf) | **Put** /applications/{application_uuid}/calls/{call_id}/dtmf | Simulate a user pressing DTMF keys
*ApplicationsAPI* | [**SnoopApplicationCall**](docs/ApplicationsAPI.md#snoopapplicationcall) | **Post** /applications/{application_uuid}/calls/{call_id}/snoops | Start snooping on a call
*ApplicationsAPI* | [**StartApplicationCallMOH**](docs/ApplicationsAPI.md#startapplicationcallmoh) | **Put** /applications/{application_uuid}/calls/{call_id}/moh/{moh_uuid}/start | Starts playing a music on hold
*ApplicationsAPI* | [**StartApplicationCallProgress**](docs/ApplicationsAPI.md#startapplicationcallprogress) | **Put** /applications/{application_uuid}/calls/{call_id}/progress/start | Play the progress ringing tone
*ApplicationsAPI* | [**StopApplicationCallMOH**](docs/ApplicationsAPI.md#stopapplicationcallmoh) | **Put** /applications/{application_uuid}/calls/{call_id}/moh/stop | Stops playing a music on hold
*ApplicationsAPI* | [**StopApplicationCallProgress**](docs/ApplicationsAPI.md#stopapplicationcallprogress) | **Put** /applications/{application_uuid}/calls/{call_id}/progress/stop | Stop playing the progress ringing tone.
*ApplicationsAPI* | [**StopSnoop**](docs/ApplicationsAPI.md#stopsnoop) | **Delete** /applications/{application_uuid}/snoops/{snoop_uuid} | Stop snooping
*ApplicationsAPI* | [**UnmuteApplicationCall**](docs/ApplicationsAPI.md#unmuteapplicationcall) | **Put** /applications/{application_uuid}/calls/{call_id}/mute/stop | Unmute a call
*ApplicationsAPI* | [**UpdateSnoop**](docs/ApplicationsAPI.md#updatesnoop) | **Put** /applications/{application_uuid}/snoops/{snoop_uuid} | Change snooping parameters
*CallsAPI* | [**AnswerCall**](docs/CallsAPI.md#answercall) | **Put** /calls/{call_id}/answer | Answer a call
*CallsAPI* | [**AnswerUserCall**](docs/CallsAPI.md#answerusercall) | **Put** /users/me/calls/{call_id}/answer | Answer a call from user
*CallsAPI* | [**ConnectCallToUser**](docs/CallsAPI.md#connectcalltouser) | **Put** /calls/{call_id}/user/{user_uuid} | Connect a call to a user
*CallsAPI* | [**CreateCall**](docs/CallsAPI.md#createcall) | **Post** /calls | Make a new call
*CallsAPI* | [**CreateUserCall**](docs/CallsAPI.md#createusercall) | **Post** /users/me/calls | Make a new call from a user
*CallsAPI* | [**DeleteCall**](docs/CallsAPI.md#deletecall) | **Delete** /calls/{call_id} | Hangup a call
*CallsAPI* | [**GetCall**](docs/CallsAPI.md#getcall) | **Get** /calls/{call_id} | Show a call
*CallsAPI* | [**HangupUserCall**](docs/CallsAPI.md#hangupusercall) | **Delete** /users/me/calls/{call_id} | Hangup a call from a user
*CallsAPI* | [**HoldCall**](docs/CallsAPI.md#holdcall) | **Put** /calls/{call_id}/hold/start | Hold a call
*CallsAPI* | [**HoldUserCall**](docs/CallsAPI.md#holdusercall) | **Put** /users/me/calls/{call_id}/hold/start | Hold a call from user
*CallsAPI* | [**ListCalls**](docs/CallsAPI.md#listcalls) | **Get** /calls | List calls
*CallsAPI* | [**ListUserCalls**](docs/CallsAPI.md#listusercalls) | **Get** /users/me/calls | List calls of a user
*CallsAPI* | [**MuteCall**](docs/CallsAPI.md#mutecall) | **Put** /calls/{call_id}/mute/start | Mute a call
*CallsAPI* | [**MuteUserCall**](docs/CallsAPI.md#muteusercall) | **Put** /users/me/calls/{call_id}/mute/start | Mute a call from user
*CallsAPI* | [**SendCallDTMF**](docs/CallsAPI.md#sendcalldtmf) | **Put** /calls/{call_id}/dtmf | Simulate a user pressing DTMF keys
*CallsAPI* | [**SendUserDTMF**](docs/CallsAPI.md#senduserdtmf) | **Put** /users/me/calls/{call_id}/dtmf | Simulate a user pressing DTMF keys
*CallsAPI* | [**StartCurrentUserRecording**](docs/CallsAPI.md#startcurrentuserrecording) | **Put** /users/me/calls/{call_id}/record/start | Start recording a call
*CallsAPI* | [**StartRecording**](docs/CallsAPI.md#startrecording) | **Put** /calls/{call_id}/record/start | Start recording a call
*CallsAPI* | [**StopCurrentUserRecording**](docs/CallsAPI.md#stopcurrentuserrecording) | **Put** /users/me/calls/{call_id}/record/stop | Stop recording a call
*CallsAPI* | [**StopRecording**](docs/CallsAPI.md#stoprecording) | **Put** /calls/{call_id}/record/stop | Stop recording a call
*CallsAPI* | [**UnholdCall**](docs/CallsAPI.md#unholdcall) | **Put** /calls/{call_id}/hold/stop | Unhold a call
*CallsAPI* | [**UnholdUserCall**](docs/CallsAPI.md#unholdusercall) | **Put** /users/me/calls/{call_id}/hold/stop | Unhold a call from user
*CallsAPI* | [**UnmuteCall**](docs/CallsAPI.md#unmutecall) | **Put** /calls/{call_id}/mute/stop | Unmute a call
*CallsAPI* | [**UnmuteUserCall**](docs/CallsAPI.md#unmuteusercall) | **Put** /users/me/calls/{call_id}/mute/stop | Unmute a call from user
*ConferencesAPI* | [**KickParticipant**](docs/ConferencesAPI.md#kickparticipant) | **Delete** /conferences/{conference_id}/participants/{participant_id} | Kick participant from a conference
*ConferencesAPI* | [**ListConferenceParticipants**](docs/ConferencesAPI.md#listconferenceparticipants) | **Get** /conferences/{conference_id}/participants | List participants of a conference
*ConferencesAPI* | [**ListUserConferenceParticipants**](docs/ConferencesAPI.md#listuserconferenceparticipants) | **Get** /users/me/conferences/{conference_id}/participants | List participants of a conference as a user
*ConferencesAPI* | [**MuteParticipant**](docs/ConferencesAPI.md#muteparticipant) | **Put** /conferences/{conference_id}/participants/{participant_id}/mute | Mute a participant in a conference
*ConferencesAPI* | [**StartConferenceRecording**](docs/ConferencesAPI.md#startconferencerecording) | **Post** /conferences/{conference_id}/record | Record a conference
*ConferencesAPI* | [**StopConferenceRecording**](docs/ConferencesAPI.md#stopconferencerecording) | **Delete** /conferences/{conference_id}/record | Stop recording a conference
*ConferencesAPI* | [**UnmuteParticipant**](docs/ConferencesAPI.md#unmuteparticipant) | **Put** /conferences/{conference_id}/participants/{participant_id}/unmute | Unmute a participant in a conference
*ConfigAPI* | [**GetConfig**](docs/ConfigAPI.md#getconfig) | **Get** /config | Show the current configuration
*ConfigAPI* | [**PatchConfig**](docs/ConfigAPI.md#patchconfig) | **Patch** /config | Update the current configuration.
*FaxesAPI* | [**SendFax**](docs/FaxesAPI.md#sendfax) | **Post** /faxes | Send a fax
*FaxesAPI* | [**SendUserFax**](docs/FaxesAPI.md#senduserfax) | **Post** /users/me/faxes | Send a fax as the user detected from the token
*LinesAPI* | [**ListLines**](docs/LinesAPI.md#listlines) | **Get** /lines | List line endpoint statuses
*MeetingsAPI* | [**GetGuestMeetingStatus**](docs/MeetingsAPI.md#getguestmeetingstatus) | **Get** /guests/me/meetings/{meeting_uuid}/status | Get the status of a meeting
*MeetingsAPI* | [**KickMeetingParticipant**](docs/MeetingsAPI.md#kickmeetingparticipant) | **Delete** /meetings/{meeting_uuid}/participants/{participant_id} | Kick a participant from a meeting
*MeetingsAPI* | [**KickUserMeetingParticipant**](docs/MeetingsAPI.md#kickusermeetingparticipant) | **Delete** /users/me/meetings/{meeting_uuid}/participants/{participant_id} | Kick a participant from a meeting as a user
*MeetingsAPI* | [**ListMeetingParticipants**](docs/MeetingsAPI.md#listmeetingparticipants) | **Get** /meetings/{meeting_uuid}/participants | List participants of a meeting
*MeetingsAPI* | [**ListUserMeetingParticipants**](docs/MeetingsAPI.md#listusermeetingparticipants) | **Get** /users/me/meetings/{meeting_uuid}/participants | List participants of a meeting as a user
*RelocatesAPI* | [**CancelRelocate**](docs/RelocatesAPI.md#cancelrelocate) | **Put** /users/me/relocates/{relocate_uuid}/cancel | Cancel a relocate
*RelocatesAPI* | [**CompleteRelocate**](docs/RelocatesAPI.md#completerelocate) | **Put** /users/me/relocates/{relocate_uuid}/complete | Complete a relocate
*RelocatesAPI* | [**GetUserRelocate**](docs/RelocatesAPI.md#getuserrelocate) | **Get** /users/me/relocates/{relocate_uuid} | Get details of a relocate
*RelocatesAPI* | [**InitiateRelocate**](docs/RelocatesAPI.md#initiaterelocate) | **Post** /users/me/relocates | Initiate a relocate from the authenticated user
*RelocatesAPI* | [**ListUserRelocates**](docs/RelocatesAPI.md#listuserrelocates) | **Get** /users/me/relocates | Get the relocates of the authenticated user
*StatusAPI* | [**GetStatus**](docs/StatusAPI.md#getstatus) | **Get** /status | Print infos about internal status of accent-calld
*SwitchboardsAPI* | [**AnswerHeldCall**](docs/SwitchboardsAPI.md#answerheldcall) | **Put** /switchboards/{switchboard_uuid}/calls/held/{call_id}/answer | Answer the specified held call
*SwitchboardsAPI* | [**AnswerQueuedCall**](docs/SwitchboardsAPI.md#answerqueuedcall) | **Put** /switchboards/{switchboard_uuid}/calls/queued/{call_id}/answer | Answer the specified queued call
*SwitchboardsAPI* | [**HoldSwitchboardCall**](docs/SwitchboardsAPI.md#holdswitchboardcall) | **Put** /switchboards/{switchboard_uuid}/calls/held/{call_id} | Put the specified call on hold in the switchboard
*SwitchboardsAPI* | [**ListSwitchboardHeldCalls**](docs/SwitchboardsAPI.md#listswitchboardheldcalls) | **Get** /switchboards/{switchboard_uuid}/calls/held | List calls held in the switchboard
*SwitchboardsAPI* | [**ListSwitchboardQueuedCalls**](docs/SwitchboardsAPI.md#listswitchboardqueuedcalls) | **Get** /switchboards/{switchboard_uuid}/calls/queued | List calls queued in the switchboard
*TransfersAPI* | [**CancelTransfer**](docs/TransfersAPI.md#canceltransfer) | **Delete** /transfers/{transfer_id} | Cancel a transfer
*TransfersAPI* | [**CancelUserTransfer**](docs/TransfersAPI.md#cancelusertransfer) | **Delete** /users/me/transfers/{transfer_id} | Cancel a transfer
*TransfersAPI* | [**CompleteTransfer**](docs/TransfersAPI.md#completetransfer) | **Put** /transfers/{transfer_id}/complete | Complete a transfer
*TransfersAPI* | [**CompleteUserTransfer**](docs/TransfersAPI.md#completeusertransfer) | **Put** /users/me/transfers/{transfer_id}/complete | Complete a transfer
*TransfersAPI* | [**GetTransfer**](docs/TransfersAPI.md#gettransfer) | **Get** /transfers/{transfer_id} | Get details of a transfer
*TransfersAPI* | [**InitiateTransfer**](docs/TransfersAPI.md#initiatetransfer) | **Post** /transfers | Initiate a transfer
*TransfersAPI* | [**InitiateUserTransfer**](docs/TransfersAPI.md#initiateusertransfer) | **Post** /users/me/transfers | Initiate a transfer from the authenticated user
*TransfersAPI* | [**ListUserTransfers**](docs/TransfersAPI.md#listusertransfers) | **Get** /users/me/transfers | Get the transfers of the authenticated user
*TrunksAPI* | [**ListTrunks**](docs/TrunksAPI.md#listtrunks) | **Get** /trunks | List trunk endpoint statuses
*UsersAPI* | [**AnswerUserCall**](docs/UsersAPI.md#answerusercall) | **Put** /users/me/calls/{call_id}/answer | Answer a call from user
*UsersAPI* | [**CancelRelocate**](docs/UsersAPI.md#cancelrelocate) | **Put** /users/me/relocates/{relocate_uuid}/cancel | Cancel a relocate
*UsersAPI* | [**CancelUserTransfer**](docs/UsersAPI.md#cancelusertransfer) | **Delete** /users/me/transfers/{transfer_id} | Cancel a transfer
*UsersAPI* | [**CheckUserVoicemailGreeting**](docs/UsersAPI.md#checkuservoicemailgreeting) | **Head** /users/me/voicemails/greetings/{greeting} | Check if greeting exists
*UsersAPI* | [**CompleteRelocate**](docs/UsersAPI.md#completerelocate) | **Put** /users/me/relocates/{relocate_uuid}/complete | Complete a relocate
*UsersAPI* | [**CompleteUserTransfer**](docs/UsersAPI.md#completeusertransfer) | **Put** /users/me/transfers/{transfer_id}/complete | Complete a transfer
*UsersAPI* | [**CopyUserVoicemailGreeting**](docs/UsersAPI.md#copyuservoicemailgreeting) | **Post** /users/me/voicemails/greetings/{greeting}/copy | Copy a custom greeting
*UsersAPI* | [**CreateUserCall**](docs/UsersAPI.md#createusercall) | **Post** /users/me/calls | Make a new call from a user
*UsersAPI* | [**CreateUserVoicemailGreeting**](docs/UsersAPI.md#createuservoicemailgreeting) | **Post** /users/me/voicemails/greetings/{greeting} | Create a custom greeting
*UsersAPI* | [**DeleteUserVoicemailGreeting**](docs/UsersAPI.md#deleteuservoicemailgreeting) | **Delete** /users/me/voicemails/greetings/{greeting} | Delete a custom greeting
*UsersAPI* | [**DeleteUserVoicemailMessage**](docs/UsersAPI.md#deleteuservoicemailmessage) | **Delete** /users/me/voicemails/messages/{message_id} | Delete a mesage
*UsersAPI* | [**GetUserVoicemailFolder**](docs/UsersAPI.md#getuservoicemailfolder) | **Get** /users/me/voicemails/folders/{folder_id} | Get details of a folder
*UsersAPI* | [**GetUserVoicemailGreeting**](docs/UsersAPI.md#getuservoicemailgreeting) | **Get** /users/me/voicemails/greetings/{greeting} | Get a custom greeting
*UsersAPI* | [**GetUserVoicemailMessage**](docs/UsersAPI.md#getuservoicemailmessage) | **Get** /users/me/voicemails/messages/{message_id} | Get a message
*UsersAPI* | [**GetUserVoicemailMessageRecording**](docs/UsersAPI.md#getuservoicemailmessagerecording) | **Get** /users/me/voicemails/messages/{message_id}/recording | Get a message&#39;s recording
*UsersAPI* | [**HangupUserCall**](docs/UsersAPI.md#hangupusercall) | **Delete** /users/me/calls/{call_id} | Hangup a call from a user
*UsersAPI* | [**HoldUserCall**](docs/UsersAPI.md#holdusercall) | **Put** /users/me/calls/{call_id}/hold/start | Hold a call from user
*UsersAPI* | [**InitiateRelocate**](docs/UsersAPI.md#initiaterelocate) | **Post** /users/me/relocates | Initiate a relocate from the authenticated user
*UsersAPI* | [**InitiateUserTransfer**](docs/UsersAPI.md#initiateusertransfer) | **Post** /users/me/transfers | Initiate a transfer from the authenticated user
*UsersAPI* | [**KickUserMeetingParticipant**](docs/UsersAPI.md#kickusermeetingparticipant) | **Delete** /users/me/meetings/{meeting_uuid}/participants/{participant_id} | Kick a participant from a meeting as a user
*UsersAPI* | [**ListUserCalls**](docs/UsersAPI.md#listusercalls) | **Get** /users/me/calls | List calls of a user
*UsersAPI* | [**ListUserConferenceParticipants**](docs/UsersAPI.md#listuserconferenceparticipants) | **Get** /users/me/conferences/{conference_id}/participants | List participants of a conference as a user
*UsersAPI* | [**ListUserMeetingParticipants**](docs/UsersAPI.md#listusermeetingparticipants) | **Get** /users/me/meetings/{meeting_uuid}/participants | List participants of a meeting as a user
*UsersAPI* | [**ListUserRelocates**](docs/UsersAPI.md#listuserrelocates) | **Get** /users/me/relocates | Get the relocates of the authenticated user
*UsersAPI* | [**ListUserTransfers**](docs/UsersAPI.md#listusertransfers) | **Get** /users/me/transfers | Get the transfers of the authenticated user
*UsersAPI* | [**ListUserVoicemails**](docs/UsersAPI.md#listuservoicemails) | **Get** /users/me/voicemails | Get details of the voicemail of the authenticated user
*UsersAPI* | [**MuteUserCall**](docs/UsersAPI.md#muteusercall) | **Put** /users/me/calls/{call_id}/mute/start | Mute a call from user
*UsersAPI* | [**SendUserDTMF**](docs/UsersAPI.md#senduserdtmf) | **Put** /users/me/calls/{call_id}/dtmf | Simulate a user pressing DTMF keys
*UsersAPI* | [**StartCurrentUserRecording**](docs/UsersAPI.md#startcurrentuserrecording) | **Put** /users/me/calls/{call_id}/record/start | Start recording a call
*UsersAPI* | [**StopCurrentUserRecording**](docs/UsersAPI.md#stopcurrentuserrecording) | **Put** /users/me/calls/{call_id}/record/stop | Stop recording a call
*UsersAPI* | [**UnholdUserCall**](docs/UsersAPI.md#unholdusercall) | **Put** /users/me/calls/{call_id}/hold/stop | Unhold a call from user
*UsersAPI* | [**UnmuteUserCall**](docs/UsersAPI.md#unmuteusercall) | **Put** /users/me/calls/{call_id}/mute/stop | Unmute a call from user
*UsersAPI* | [**UpdateUserVoicemailGreeting**](docs/UsersAPI.md#updateuservoicemailgreeting) | **Put** /users/me/voicemails/greetings/{greeting} | Update a custom greeting
*UsersAPI* | [**UpdateUserVoicemailMessage**](docs/UsersAPI.md#updateuservoicemailmessage) | **Put** /users/me/voicemails/messages/{message_id} | Update a message
*VoicemailsAPI* | [**CheckUserVoicemailGreeting**](docs/VoicemailsAPI.md#checkuservoicemailgreeting) | **Head** /users/me/voicemails/greetings/{greeting} | Check if greeting exists
*VoicemailsAPI* | [**CheckVoicemailGreeting**](docs/VoicemailsAPI.md#checkvoicemailgreeting) | **Head** /voicemails/{voicemail_id}/greetings/{greeting} | Check if greeting exists
*VoicemailsAPI* | [**CopyUserVoicemailGreeting**](docs/VoicemailsAPI.md#copyuservoicemailgreeting) | **Post** /users/me/voicemails/greetings/{greeting}/copy | Copy a custom greeting
*VoicemailsAPI* | [**CopyVoicemailGreeting**](docs/VoicemailsAPI.md#copyvoicemailgreeting) | **Post** /voicemails/{voicemail_id}/greetings/{greeting}/copy | Copy a custom greeting
*VoicemailsAPI* | [**CreateUserVoicemailGreeting**](docs/VoicemailsAPI.md#createuservoicemailgreeting) | **Post** /users/me/voicemails/greetings/{greeting} | Create a custom greeting
*VoicemailsAPI* | [**CreateVoicemailGreeting**](docs/VoicemailsAPI.md#createvoicemailgreeting) | **Post** /voicemails/{voicemail_id}/greetings/{greeting} | Create a custom greeting
*VoicemailsAPI* | [**DeleteUserVoicemailGreeting**](docs/VoicemailsAPI.md#deleteuservoicemailgreeting) | **Delete** /users/me/voicemails/greetings/{greeting} | Delete a custom greeting
*VoicemailsAPI* | [**DeleteUserVoicemailMessage**](docs/VoicemailsAPI.md#deleteuservoicemailmessage) | **Delete** /users/me/voicemails/messages/{message_id} | Delete a mesage
*VoicemailsAPI* | [**DeleteVoicemailGreeting**](docs/VoicemailsAPI.md#deletevoicemailgreeting) | **Delete** /voicemails/{voicemail_id}/greetings/{greeting} | Delete a custom greeting
*VoicemailsAPI* | [**DeleteVoicemailMessage**](docs/VoicemailsAPI.md#deletevoicemailmessage) | **Delete** /voicemails/{voicemail_id}/messages/{message_id} | Delete a mesage
*VoicemailsAPI* | [**GetUserVoicemailFolder**](docs/VoicemailsAPI.md#getuservoicemailfolder) | **Get** /users/me/voicemails/folders/{folder_id} | Get details of a folder
*VoicemailsAPI* | [**GetUserVoicemailGreeting**](docs/VoicemailsAPI.md#getuservoicemailgreeting) | **Get** /users/me/voicemails/greetings/{greeting} | Get a custom greeting
*VoicemailsAPI* | [**GetUserVoicemailMessage**](docs/VoicemailsAPI.md#getuservoicemailmessage) | **Get** /users/me/voicemails/messages/{message_id} | Get a message
*VoicemailsAPI* | [**GetUserVoicemailMessageRecording**](docs/VoicemailsAPI.md#getuservoicemailmessagerecording) | **Get** /users/me/voicemails/messages/{message_id}/recording | Get a message&#39;s recording
*VoicemailsAPI* | [**GetVoicemail**](docs/VoicemailsAPI.md#getvoicemail) | **Get** /voicemails/{voicemail_id} | Get details of a voicemail
*VoicemailsAPI* | [**GetVoicemailFolder**](docs/VoicemailsAPI.md#getvoicemailfolder) | **Get** /voicemails/{voicemail_id}/folders/{folder_id} | Get details of a folder
*VoicemailsAPI* | [**GetVoicemailGreeting**](docs/VoicemailsAPI.md#getvoicemailgreeting) | **Get** /voicemails/{voicemail_id}/greetings/{greeting} | Get a custom greeting
*VoicemailsAPI* | [**GetVoicemailMessage**](docs/VoicemailsAPI.md#getvoicemailmessage) | **Get** /voicemails/{voicemail_id}/messages/{message_id} | Get a message
*VoicemailsAPI* | [**GetVoicemailMessageRecording**](docs/VoicemailsAPI.md#getvoicemailmessagerecording) | **Get** /voicemails/{voicemail_id}/messages/{message_id}/recording | Get a message&#39;s recording
*VoicemailsAPI* | [**ListUserVoicemails**](docs/VoicemailsAPI.md#listuservoicemails) | **Get** /users/me/voicemails | Get details of the voicemail of the authenticated user
*VoicemailsAPI* | [**UpdateUserVoicemailGreeting**](docs/VoicemailsAPI.md#updateuservoicemailgreeting) | **Put** /users/me/voicemails/greetings/{greeting} | Update a custom greeting
*VoicemailsAPI* | [**UpdateUserVoicemailMessage**](docs/VoicemailsAPI.md#updateuservoicemailmessage) | **Put** /users/me/voicemails/messages/{message_id} | Update a message
*VoicemailsAPI* | [**UpdateVoicemailGreeting**](docs/VoicemailsAPI.md#updatevoicemailgreeting) | **Put** /voicemails/{voicemail_id}/greetings/{greeting} | Update a custom greeting
*VoicemailsAPI* | [**UpdateVoicemailMessage**](docs/VoicemailsAPI.md#updatevoicemailmessage) | **Put** /voicemails/{voicemail_id}/messages/{message_id} | Update a message


## Documentation For Models

 - [AdhocConference](docs/AdhocConference.md)
 - [AdhocConferenceCreation](docs/AdhocConferenceCreation.md)
 - [Application](docs/Application.md)
 - [ApplicationCall](docs/ApplicationCall.md)
 - [ApplicationCallRequestToExten](docs/ApplicationCallRequestToExten.md)
 - [ApplicationCallRequestToUser](docs/ApplicationCallRequestToUser.md)
 - [ApplicationCalls](docs/ApplicationCalls.md)
 - [ApplicationNode](docs/ApplicationNode.md)
 - [ApplicationNodeCallRequest](docs/ApplicationNodeCallRequest.md)
 - [ApplicationNodeRequest](docs/ApplicationNodeRequest.md)
 - [ApplicationNodes](docs/ApplicationNodes.md)
 - [ApplicationPlayback](docs/ApplicationPlayback.md)
 - [ApplicationSnoop](docs/ApplicationSnoop.md)
 - [ApplicationSnoopPut](docs/ApplicationSnoopPut.md)
 - [ApplicationSnoops](docs/ApplicationSnoops.md)
 - [Call](docs/Call.md)
 - [CallID](docs/CallID.md)
 - [CallRequest](docs/CallRequest.md)
 - [CallRequestDestination](docs/CallRequestDestination.md)
 - [CallRequestSource](docs/CallRequestSource.md)
 - [ComponentWithStatus](docs/ComponentWithStatus.md)
 - [ConfigPatchItem](docs/ConfigPatchItem.md)
 - [ConnectCallToUserRequest](docs/ConnectCallToUserRequest.md)
 - [EndpointLine](docs/EndpointLine.md)
 - [EndpointLines](docs/EndpointLines.md)
 - [EndpointTrunk](docs/EndpointTrunk.md)
 - [EndpointTrunks](docs/EndpointTrunks.md)
 - [Error](docs/Error.md)
 - [Fax](docs/Fax.md)
 - [GreetingCopy](docs/GreetingCopy.md)
 - [ListCalls200Response](docs/ListCalls200Response.md)
 - [LocationLine](docs/LocationLine.md)
 - [MeetingStatus](docs/MeetingStatus.md)
 - [Participant](docs/Participant.md)
 - [ParticipantList](docs/ParticipantList.md)
 - [PluginsStatus](docs/PluginsStatus.md)
 - [Relocate](docs/Relocate.md)
 - [RelocateCompletion](docs/RelocateCompletion.md)
 - [RelocateList](docs/RelocateList.md)
 - [StatusSummary](docs/StatusSummary.md)
 - [StatusValue](docs/StatusValue.md)
 - [SwitchboardHeldCall](docs/SwitchboardHeldCall.md)
 - [SwitchboardHeldCalls](docs/SwitchboardHeldCalls.md)
 - [SwitchboardQueuedCall](docs/SwitchboardQueuedCall.md)
 - [SwitchboardQueuedCalls](docs/SwitchboardQueuedCalls.md)
 - [TalkingTo](docs/TalkingTo.md)
 - [Transfer](docs/Transfer.md)
 - [TransferFlow](docs/TransferFlow.md)
 - [TransferList](docs/TransferList.md)
 - [TransferRequest](docs/TransferRequest.md)
 - [UserCallRequest](docs/UserCallRequest.md)
 - [UserRelocateLocation](docs/UserRelocateLocation.md)
 - [UserRelocateRequest](docs/UserRelocateRequest.md)
 - [UserTransferRequest](docs/UserTransferRequest.md)
 - [Voicemail](docs/Voicemail.md)
 - [VoicemailFolder](docs/VoicemailFolder.md)
 - [VoicemailFolderBase](docs/VoicemailFolderBase.md)
 - [VoicemailMessage](docs/VoicemailMessage.md)
 - [VoicemailMessageBase](docs/VoicemailMessageBase.md)
 - [VoicemailMessageUpdate](docs/VoicemailMessageUpdate.md)
 - [VoicemailsStatus](docs/VoicemailsStatus.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### accent_auth_token

- **Type**: API key
- **API key parameter name**: X-Auth-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Auth-Token and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		calld.ContextAPIKeys,
		map[string]calld.APIKey{
			"X-Auth-Token": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

help@accentvoice.io

