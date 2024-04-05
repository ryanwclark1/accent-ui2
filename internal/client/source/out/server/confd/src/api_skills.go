/*
 * accent-confd
 *
 * Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.
 *
 * API version: 1.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package confdserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// SkillsAPIController binds http requests to an api service and writes the service results to the http response
type SkillsAPIController struct {
	service      SkillsAPIServicer
	errorHandler ErrorHandler
}

// SkillsAPIOption for how the controller is set up.
type SkillsAPIOption func(*SkillsAPIController)

// WithSkillsAPIErrorHandler inject ErrorHandler into controller
func WithSkillsAPIErrorHandler(h ErrorHandler) SkillsAPIOption {
	return func(c *SkillsAPIController) {
		c.errorHandler = h
	}
}

// NewSkillsAPIController creates a default api controller
func NewSkillsAPIController(s SkillsAPIServicer, opts ...SkillsAPIOption) Router {
	controller := &SkillsAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the SkillsAPIController
func (c *SkillsAPIController) Routes() Routes {
	return Routes{
		"AssociateAgentSkill": Route{
			strings.ToUpper("Put"),
			"/1.1/agents/{agent_id}/skills/{skill_id}",
			c.AssociateAgentSkill,
		},
		"CreateSkill": Route{
			strings.ToUpper("Post"),
			"/1.1/agents/skills",
			c.CreateSkill,
		},
		"CreateSkillRule": Route{
			strings.ToUpper("Post"),
			"/1.1/queues/skillrules",
			c.CreateSkillRule,
		},
		"DeleteSkill": Route{
			strings.ToUpper("Delete"),
			"/1.1/agents/skills/{skill_id}",
			c.DeleteSkill,
		},
		"DeleteSkillRule": Route{
			strings.ToUpper("Delete"),
			"/1.1/queues/skillrules/{skillrule_id}",
			c.DeleteSkillRule,
		},
		"DissociateAgentSkill": Route{
			strings.ToUpper("Delete"),
			"/1.1/agents/{agent_id}/skills/{skill_id}",
			c.DissociateAgentSkill,
		},
		"GetSkill": Route{
			strings.ToUpper("Get"),
			"/1.1/agents/skills/{skill_id}",
			c.GetSkill,
		},
		"GetSkillRule": Route{
			strings.ToUpper("Get"),
			"/1.1/queues/skillrules/{skillrule_id}",
			c.GetSkillRule,
		},
		"ListSkillRules": Route{
			strings.ToUpper("Get"),
			"/1.1/queues/skillrules",
			c.ListSkillRules,
		},
		"ListSkills": Route{
			strings.ToUpper("Get"),
			"/1.1/agents/skills",
			c.ListSkills,
		},
		"UpdateSkill": Route{
			strings.ToUpper("Put"),
			"/1.1/agents/skills/{skill_id}",
			c.UpdateSkill,
		},
		"UpdateSkillRule": Route{
			strings.ToUpper("Put"),
			"/1.1/queues/skillrules/{skillrule_id}",
			c.UpdateSkillRule,
		},
	}
}

// AssociateAgentSkill - Associate agent and skill
func (c *SkillsAPIController) AssociateAgentSkill(w http.ResponseWriter, r *http.Request) {
	agentIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "agent_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skillIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "skill_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	bodyParam := AgentSkill{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil && !errors.Is(err, io.EOF) {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAgentSkillRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertAgentSkillConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AssociateAgentSkill(r.Context(), agentIdParam, skillIdParam, accentTenantParam, bodyParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateSkill - Create skill
func (c *SkillsAPIController) CreateSkill(w http.ResponseWriter, r *http.Request) {
	bodyParam := Skill{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSkillRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSkillConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.CreateSkill(r.Context(), bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// CreateSkillRule - Create skill rule
func (c *SkillsAPIController) CreateSkillRule(w http.ResponseWriter, r *http.Request) {
	bodyParam := SkillRule{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSkillRuleRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSkillRuleConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.CreateSkillRule(r.Context(), bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteSkill - Delete skill
func (c *SkillsAPIController) DeleteSkill(w http.ResponseWriter, r *http.Request) {
	skillIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "skill_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteSkill(r.Context(), skillIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteSkillRule - Delete skill rule
func (c *SkillsAPIController) DeleteSkillRule(w http.ResponseWriter, r *http.Request) {
	skillruleIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "skillrule_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DeleteSkillRule(r.Context(), skillruleIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// DissociateAgentSkill - Dissociate agent and skill
func (c *SkillsAPIController) DissociateAgentSkill(w http.ResponseWriter, r *http.Request) {
	agentIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "agent_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skillIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "skill_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.DissociateAgentSkill(r.Context(), agentIdParam, skillIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetSkill - Get skill
func (c *SkillsAPIController) GetSkill(w http.ResponseWriter, r *http.Request) {
	skillIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "skill_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetSkill(r.Context(), skillIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// GetSkillRule - Get skill rule
func (c *SkillsAPIController) GetSkillRule(w http.ResponseWriter, r *http.Request) {
	skillruleIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "skillrule_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.GetSkillRule(r.Context(), skillruleIdParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListSkillRules - List skill rule
func (c *SkillsAPIController) ListSkillRules(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	var recurseParam bool
	if query.Has("recurse") {
		param, err := parseBoolParameter(
			query.Get("recurse"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		recurseParam = param
	} else {
		var param bool = false
		recurseParam = param
	}
	var orderParam string
	if query.Has("order") {
		param := query.Get("order")

		orderParam = param
	} else {
	}
	var directionParam string
	if query.Has("direction") {
		param := query.Get("direction")

		directionParam = param
	} else {
	}
	var limitParam int32
	if query.Has("limit") {
		param, err := parseNumericParameter[int32](
			query.Get("limit"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		limitParam = param
	} else {
	}
	var offsetParam int32
	if query.Has("offset") {
		param, err := parseNumericParameter[int32](
			query.Get("offset"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		offsetParam = param
	} else {
	}
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	result, err := c.service.ListSkillRules(r.Context(), accentTenantParam, recurseParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// ListSkills - List skill
func (c *SkillsAPIController) ListSkills(w http.ResponseWriter, r *http.Request) {
	query, err := parseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	var recurseParam bool
	if query.Has("recurse") {
		param, err := parseBoolParameter(
			query.Get("recurse"),
			WithParse[bool](parseBool),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		recurseParam = param
	} else {
		var param bool = false
		recurseParam = param
	}
	var orderParam string
	if query.Has("order") {
		param := query.Get("order")

		orderParam = param
	} else {
	}
	var directionParam string
	if query.Has("direction") {
		param := query.Get("direction")

		directionParam = param
	} else {
	}
	var limitParam int32
	if query.Has("limit") {
		param, err := parseNumericParameter[int32](
			query.Get("limit"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		limitParam = param
	} else {
	}
	var offsetParam int32
	if query.Has("offset") {
		param, err := parseNumericParameter[int32](
			query.Get("offset"),
			WithParse[int32](parseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &ParsingError{Err: err}, nil)
			return
		}

		offsetParam = param
	} else {
	}
	var searchParam string
	if query.Has("search") {
		param := query.Get("search")

		searchParam = param
	} else {
	}
	result, err := c.service.ListSkills(r.Context(), accentTenantParam, recurseParam, orderParam, directionParam, limitParam, offsetParam, searchParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateSkill - Update skill
func (c *SkillsAPIController) UpdateSkill(w http.ResponseWriter, r *http.Request) {
	skillIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "skill_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := Skill{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSkillRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSkillConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateSkill(r.Context(), skillIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}

// UpdateSkillRule - Update skill rule
func (c *SkillsAPIController) UpdateSkillRule(w http.ResponseWriter, r *http.Request) {
	skillruleIdParam, err := parseNumericParameter[int32](
		chi.URLParam(r, "skillrule_id"),
		WithRequire[int32](parseInt32),
	)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	bodyParam := SkillRule{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bodyParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertSkillRuleRequired(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertSkillRuleConstraints(bodyParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	accentTenantParam := r.Header.Get("Accent-Tenant")
	result, err := c.service.UpdateSkillRule(r.Context(), skillruleIdParam, bodyParam, accentTenantParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
}