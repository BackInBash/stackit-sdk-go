/*
Resource Manager API

API v2 to manage resource containers - organizations, folders, projects incl. labels  ### Resource Management STACKIT resource management handles the terms _Organization_, _Folder_, _Project_, _Label_, and the hierarchical structure between them. Technically, organizations,  folders, and projects are _Resource Containers_ to which a _Label_ can be attached to. The STACKIT _Resource Manager_ provides CRUD endpoints to query and to modify the state.  ### Resource Hierarchy STACKIT resource hierarchy defines the relationship of resource containers as tree structure with organizations as the root node. Folders and projects are optional  child elements referencing the organization as parent. STACKIT resource hierarchy allows to structure cloud-resources clearly providing flexibility and individuality  for fine grained access control, access inheritance, and policies.  The STACKIT resource hierarchy model can be compared to the folder concept of a Unix file system. A folder can have exactly one parent. Folder nesting allows to organize and to structure content while defining  fine grained access permissions per folder. Within STACKIT resource hierarchy access is inherited, meaning if you have access to a resource container, you also have access to its child containers. - Users are assigned to resource containers as members by role - A user inherits permissions to all child containers  ### Organizations STACKIT organizations are the base element to create and to use cloud-resources. An organization is bound to one customer account. Organizations have a lifecycle. - Organizations are always the root node in resource hierarchy and do not have a parent  ### Folders STACKIT folders allow to organize cloud-resources and to define fine-grained access control. A folder might represent departments, teams, user groups, components etc.  Folders do not have a lifecycle as they act as structural element only. - Folders are optional - A folder can be created having either an organization, or a folder as parent - Folder names under the same parent must be unique (case insensitive) - Root organization cannot be changed  ### Projects STACKIT projects are needed to use cloud-resources. Projects serve as wrapper for underlying technical structures and processes. Projects have a lifecycle. Projects compared to folders may have different policies. - Projects are optional, but mandatory for cloud-resource usage - A project can be created having either an organization, or a folder as parent - A project must not have a project as parent - Project names under the same parent must not be unique - Root organization cannot be changed  ### Label STACKIT labels are key-value pairs including a resource container reference. Labels can be defined and attached freely to resource containers by which resources can be organized and queried. - Policy-based, immutable labels may exists  ### Current Limits - Vertically - Maximum folder nesting level: 5 - Horizontally - Maximum number of folders under one parent: 150 - Maximum number of projects under one organization: 2.500 - Maximum number of labels attached to one container: 100

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resourcemanager

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiCreateProjectRequest struct {
	ctx                  context.Context
	apiService           *DefaultApiService
	createProjectPayload *CreateProjectPayload
}

// func (r ApiCreateProjectRequest) CreateProjectPayload(createProjectPayload CreateProjectPayload) ApiCreateProjectRequest {
func (r ApiCreateProjectRequest) CreateProjectPayload(createProjectPayload CreateProjectPayload) ApiCreateProjectRequest {
	r.createProjectPayload = &createProjectPayload
	return r
}

func (r ApiCreateProjectRequest) Execute() (*ProjectResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProjectResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateProject")
	if err != nil {
		return localVarReturnValue, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/projects"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createProjectPayload
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

/*
CreateProject Create Project

Create a new project.
- The request is synchronous, but the workflow-based creation is asynchronous.
- Lifecycle state remains in CREATING, until workflow completes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateProjectRequest
*/
func (a *APIClient) CreateProject(ctx context.Context) ApiCreateProjectRequest {
	return ApiCreateProjectRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
	}
}

func (a *APIClient) CreateProjectExecute(ctx context.Context) (*ProjectResponse, error) {
	r := ApiCreateProjectRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
	}
	return r.Execute()
}

type ApiDeleteProjectRequest struct {
	ctx         context.Context
	apiService  *DefaultApiService
	containerId string
}

func (r ApiDeleteProjectRequest) Execute() error {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DeleteProject")
	if err != nil {
		return &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/projects/{containerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(ParameterValueToString(r.containerId, "containerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return newErr
	}

	return nil
}

/*
DeleteProject Delete Project

Triggers the deletion of a project.
- The request is synchronous, but the workflow-based deletion is asynchronous
- Lifecycle state remains in DELETING, until workflow completes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param containerId Project identifier - containerId as well as UUID identifier is supported.
	@return ApiDeleteProjectRequest
*/
func (a *APIClient) DeleteProject(ctx context.Context, containerId string) ApiDeleteProjectRequest {
	return ApiDeleteProjectRequest{
		apiService:  a.defaultApi,
		ctx:         ctx,
		containerId: containerId,
	}
}

func (a *APIClient) DeleteProjectExecute(ctx context.Context, containerId string) error {
	r := ApiDeleteProjectRequest{
		apiService:  a.defaultApi,
		ctx:         ctx,
		containerId: containerId,
	}
	return r.Execute()
}

type ApiGetProjectRequest struct {
	ctx            context.Context
	apiService     *DefaultApiService
	containerId    string
	includeParents *bool
}

// func (r ApiGetProjectRequest) IncludeParents(includeParents bool) ApiGetProjectRequest {
func (r ApiGetProjectRequest) IncludeParents(includeParents bool) ApiGetProjectRequest {
	r.includeParents = &includeParents
	return r
}

func (r ApiGetProjectRequest) Execute() (*ProjectResponseWithParents, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProjectResponseWithParents
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetProject")
	if err != nil {
		return localVarReturnValue, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/projects/{containerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(ParameterValueToString(r.containerId, "containerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeParents != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeParents", r.includeParents, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

/*
GetProject Get Project Details

Returns the project and its metadata.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param containerId Project identifier - containerId as well as UUID identifier is supported.
	@return ApiGetProjectRequest
*/
func (a *APIClient) GetProject(ctx context.Context, containerId string) ApiGetProjectRequest {
	return ApiGetProjectRequest{
		apiService:  a.defaultApi,
		ctx:         ctx,
		containerId: containerId,
	}
}

func (a *APIClient) GetProjectExecute(ctx context.Context, containerId string) (*ProjectResponseWithParents, error) {
	r := ApiGetProjectRequest{
		apiService:  a.defaultApi,
		ctx:         ctx,
		containerId: containerId,
	}
	return r.Execute()
}

type ApiGetProjectsRequest struct {
	ctx               context.Context
	apiService        *DefaultApiService
	containerParentId *string
	containerIds      *[]string
	member            *string
	offset            *float32
	limit             *float32
	creationTimeStart *time.Time
}

// Identifier of the parent resource container - containerId as well as UUID identifier is supported.

// func (r ApiGetProjectsRequest) ContainerParentId(containerParentId string) ApiGetProjectsRequest {
func (r ApiGetProjectsRequest) ContainerParentId(containerParentId string) ApiGetProjectsRequest {
	r.containerParentId = &containerParentId
	return r
}

// List of container identifiers - containerId as well as UUID identifier is supported.

// func (r ApiGetProjectsRequest) ContainerIds(containerIds []string) ApiGetProjectsRequest {
func (r ApiGetProjectsRequest) ContainerIds(containerIds []string) ApiGetProjectsRequest {
	r.containerIds = &containerIds
	return r
}

// E-Mail address of the user for whom the visible resource containers should be filtered.

// func (r ApiGetProjectsRequest) Member(member string) ApiGetProjectsRequest {
func (r ApiGetProjectsRequest) Member(member string) ApiGetProjectsRequest {
	r.member = &member
	return r
}

// The offset of the first item in the collection to return.

// func (r ApiGetProjectsRequest) Offset(offset float32) ApiGetProjectsRequest {
func (r ApiGetProjectsRequest) Offset(offset float32) ApiGetProjectsRequest {
	r.offset = &offset
	return r
}

// The maximum number of projects to return in the response. If not present, an appropriate default will be used. If maximum is exceeded, maximum is used.

// func (r ApiGetProjectsRequest) Limit(limit float32) ApiGetProjectsRequest {
func (r ApiGetProjectsRequest) Limit(limit float32) ApiGetProjectsRequest {
	r.limit = &limit
	return r
}

// A timestamp to specify the beginning of the creationTime from which entries should be returned. If not given, defaults to the beginning of time.

// func (r ApiGetProjectsRequest) CreationTimeStart(creationTimeStart time.Time) ApiGetProjectsRequest {
func (r ApiGetProjectsRequest) CreationTimeStart(creationTimeStart time.Time) ApiGetProjectsRequest {
	r.creationTimeStart = &creationTimeStart
	return r
}

func (r ApiGetProjectsRequest) Execute() (*AllProjectsResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *AllProjectsResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.GetProjects")
	if err != nil {
		return localVarReturnValue, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/projects"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.containerParentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "containerParentId", r.containerParentId, "")
	}
	if r.containerIds != nil {
		t := *r.containerIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "containerIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "containerIds", t, "multi")
		}
	}
	if r.member != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "member", r.member, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.creationTimeStart != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "creation-time-start", r.creationTimeStart, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

/*
GetProjects Get All Projects

Returns all projects and their metadata that:
- Are children of the specific containerParentId
- Match the given containerIds
- User is member of

Filter:
- Either containerParentId OR containerIds OR member must be passed
- If containerId and containerParentId are given, both are used for filtering - containers must point to the same parent
- If member and containerParentId are given, both are used for filtering
- If member is given, containers must not point to the same container parent

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetProjectsRequest
*/
func (a *APIClient) GetProjects(ctx context.Context) ApiGetProjectsRequest {
	return ApiGetProjectsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
	}
}

func (a *APIClient) GetProjectsExecute(ctx context.Context) (*AllProjectsResponse, error) {
	r := ApiGetProjectsRequest{
		apiService: a.defaultApi,
		ctx:        ctx,
	}
	return r.Execute()
}

type ApiUpdateProjectRequest struct {
	ctx                  context.Context
	apiService           *DefaultApiService
	containerId          string
	updateProjectPayload *UpdateProjectPayload
}

// func (r ApiUpdateProjectRequest) UpdateProjectPayload(updateProjectPayload UpdateProjectPayload) ApiUpdateProjectRequest {
func (r ApiUpdateProjectRequest) UpdateProjectPayload(updateProjectPayload UpdateProjectPayload) ApiUpdateProjectRequest {
	r.updateProjectPayload = &updateProjectPayload
	return r
}

func (r ApiUpdateProjectRequest) Execute() (*ProjectResponse, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProjectResponse
	)
	a := r.apiService
	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.UpdateProject")
	if err != nil {
		return localVarReturnValue, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/projects/{containerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(ParameterValueToString(r.containerId, "containerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.updateProjectPayload
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			statusCode: localVarHTTPResponse.StatusCode,
			body:       localVarBody,
			error:      err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

/*
UpdateProject Update Project

Update the project and its metadata.
- Update project name
- Update project labels
- Update project parent (folder or organization)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param containerId Project identifier - containerId as well as UUID identifier is supported.
	@return ApiUpdateProjectRequest
*/
func (a *APIClient) UpdateProject(ctx context.Context, containerId string) ApiUpdateProjectRequest {
	return ApiUpdateProjectRequest{
		apiService:  a.defaultApi,
		ctx:         ctx,
		containerId: containerId,
	}
}

func (a *APIClient) UpdateProjectExecute(ctx context.Context, containerId string) (*ProjectResponse, error) {
	r := ApiUpdateProjectRequest{
		apiService:  a.defaultApi,
		ctx:         ctx,
		containerId: containerId,
	}
	return r.Execute()
}