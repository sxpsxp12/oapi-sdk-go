// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/api/core/response"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
)

type Service struct {
	conf            *config.Config
	Apps            *AppService
	AppTables       *AppTableService
	AppTableFields  *AppTableFieldService
	AppTableRecords *AppTableRecordService
	AppTableViews   *AppTableViewService
}

func NewService(conf *config.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.Apps = newAppService(s)
	s.AppTables = newAppTableService(s)
	s.AppTableFields = newAppTableFieldService(s)
	s.AppTableRecords = newAppTableRecordService(s)
	s.AppTableViews = newAppTableViewService(s)
	return s
}

type AppService struct {
	service *Service
}

func newAppService(service *Service) *AppService {
	return &AppService{
		service: service,
	}
}

type AppTableService struct {
	service *Service
}

func newAppTableService(service *Service) *AppTableService {
	return &AppTableService{
		service: service,
	}
}

type AppTableFieldService struct {
	service *Service
}

func newAppTableFieldService(service *Service) *AppTableFieldService {
	return &AppTableFieldService{
		service: service,
	}
}

type AppTableRecordService struct {
	service *Service
}

func newAppTableRecordService(service *Service) *AppTableRecordService {
	return &AppTableRecordService{
		service: service,
	}
}

type AppTableViewService struct {
	service *Service
}

func newAppTableViewService(service *Service) *AppTableViewService {
	return &AppTableViewService{
		service: service,
	}
}

type AppGetReqCall struct {
	ctx        *core.Context
	apps       *AppService
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *AppGetReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}

func (rc *AppGetReqCall) Do() (*AppGetResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &AppGetResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.apps.service.conf, req)
	return result, err
}

func (apps *AppService) Get(ctx *core.Context, optFns ...request.OptFn) *AppGetReqCall {
	return &AppGetReqCall{
		ctx:        ctx,
		apps:       apps,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type AppTableListReqCall struct {
	ctx         *core.Context
	appTables   *AppTableService
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *AppTableListReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *AppTableListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *AppTableListReqCall) Do() (*AppTableListResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AppTableListResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTables.service.conf, req)
	return result, err
}

func (appTables *AppTableService) List(ctx *core.Context, optFns ...request.OptFn) *AppTableListReqCall {
	return &AppTableListReqCall{
		ctx:         ctx,
		appTables:   appTables,
		pathParams:  map[string]interface{}{},
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type AppTableBatchCreateReqCall struct {
	ctx         *core.Context
	appTables   *AppTableService
	body        *AppTableBatchCreateReqBody
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *AppTableBatchCreateReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableBatchCreateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *AppTableBatchCreateReqCall) Do() (*AppTableBatchCreateResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AppTableBatchCreateResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/batch_create", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTables.service.conf, req)
	return result, err
}

func (appTables *AppTableService) BatchCreate(ctx *core.Context, body *AppTableBatchCreateReqBody, optFns ...request.OptFn) *AppTableBatchCreateReqCall {
	return &AppTableBatchCreateReqCall{
		ctx:         ctx,
		appTables:   appTables,
		body:        body,
		pathParams:  map[string]interface{}{},
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type AppTableCreateReqCall struct {
	ctx         *core.Context
	appTables   *AppTableService
	body        *AppTableCreateReqBody
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *AppTableCreateReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableCreateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *AppTableCreateReqCall) Do() (*AppTableCreateResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AppTableCreateResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTables.service.conf, req)
	return result, err
}

func (appTables *AppTableService) Create(ctx *core.Context, body *AppTableCreateReqBody, optFns ...request.OptFn) *AppTableCreateReqCall {
	return &AppTableCreateReqCall{
		ctx:         ctx,
		appTables:   appTables,
		body:        body,
		pathParams:  map[string]interface{}{},
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type AppTableDeleteReqCall struct {
	ctx        *core.Context
	appTables  *AppTableService
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *AppTableDeleteReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableDeleteReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}

func (rc *AppTableDeleteReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &response.NoData{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id", "DELETE",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTables.service.conf, req)
	return result, err
}

func (appTables *AppTableService) Delete(ctx *core.Context, optFns ...request.OptFn) *AppTableDeleteReqCall {
	return &AppTableDeleteReqCall{
		ctx:        ctx,
		appTables:  appTables,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type AppTableBatchDeleteReqCall struct {
	ctx        *core.Context
	appTables  *AppTableService
	body       *AppTableBatchDeleteReqBody
	pathParams map[string]interface{}
	optFns     []request.OptFn
}

func (rc *AppTableBatchDeleteReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}

func (rc *AppTableBatchDeleteReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &response.NoData{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/batch_delete", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTables.service.conf, req)
	return result, err
}

func (appTables *AppTableService) BatchDelete(ctx *core.Context, body *AppTableBatchDeleteReqBody, optFns ...request.OptFn) *AppTableBatchDeleteReqCall {
	return &AppTableBatchDeleteReqCall{
		ctx:        ctx,
		appTables:  appTables,
		body:       body,
		pathParams: map[string]interface{}{},
		optFns:     optFns,
	}
}

type AppTableFieldListReqCall struct {
	ctx            *core.Context
	appTableFields *AppTableFieldService
	pathParams     map[string]interface{}
	queryParams    map[string]interface{}
	optFns         []request.OptFn
}

func (rc *AppTableFieldListReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableFieldListReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableFieldListReqCall) SetViewId(viewId string) {
	rc.queryParams["view_id"] = viewId
}
func (rc *AppTableFieldListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *AppTableFieldListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}

func (rc *AppTableFieldListReqCall) Do() (*AppTableFieldListResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AppTableFieldListResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableFields.service.conf, req)
	return result, err
}

func (appTableFields *AppTableFieldService) List(ctx *core.Context, optFns ...request.OptFn) *AppTableFieldListReqCall {
	return &AppTableFieldListReqCall{
		ctx:            ctx,
		appTableFields: appTableFields,
		pathParams:     map[string]interface{}{},
		queryParams:    map[string]interface{}{},
		optFns:         optFns,
	}
}

type AppTableFieldCreateReqCall struct {
	ctx            *core.Context
	appTableFields *AppTableFieldService
	body           *AppTableField
	pathParams     map[string]interface{}
	queryParams    map[string]interface{}
	optFns         []request.OptFn
}

func (rc *AppTableFieldCreateReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableFieldCreateReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableFieldCreateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *AppTableFieldCreateReqCall) Do() (*AppTableFieldCreateResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AppTableFieldCreateResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableFields.service.conf, req)
	return result, err
}

func (appTableFields *AppTableFieldService) Create(ctx *core.Context, body *AppTableField, optFns ...request.OptFn) *AppTableFieldCreateReqCall {
	return &AppTableFieldCreateReqCall{
		ctx:            ctx,
		appTableFields: appTableFields,
		body:           body,
		pathParams:     map[string]interface{}{},
		queryParams:    map[string]interface{}{},
		optFns:         optFns,
	}
}

type AppTableFieldDeleteReqCall struct {
	ctx            *core.Context
	appTableFields *AppTableFieldService
	pathParams     map[string]interface{}
	optFns         []request.OptFn
}

func (rc *AppTableFieldDeleteReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableFieldDeleteReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableFieldDeleteReqCall) SetFieldId(fieldId string) {
	rc.pathParams["field_id"] = fieldId
}

func (rc *AppTableFieldDeleteReqCall) Do() (*AppTableFieldDeleteResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &AppTableFieldDeleteResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id", "DELETE",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableFields.service.conf, req)
	return result, err
}

func (appTableFields *AppTableFieldService) Delete(ctx *core.Context, optFns ...request.OptFn) *AppTableFieldDeleteReqCall {
	return &AppTableFieldDeleteReqCall{
		ctx:            ctx,
		appTableFields: appTableFields,
		pathParams:     map[string]interface{}{},
		optFns:         optFns,
	}
}

type AppTableFieldUpdateReqCall struct {
	ctx            *core.Context
	appTableFields *AppTableFieldService
	body           *AppTableField
	pathParams     map[string]interface{}
	optFns         []request.OptFn
}

func (rc *AppTableFieldUpdateReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableFieldUpdateReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableFieldUpdateReqCall) SetFieldId(fieldId string) {
	rc.pathParams["field_id"] = fieldId
}

func (rc *AppTableFieldUpdateReqCall) Do() (*AppTableFieldUpdateResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &AppTableFieldUpdateResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/fields/:field_id", "PUT",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableFields.service.conf, req)
	return result, err
}

func (appTableFields *AppTableFieldService) Update(ctx *core.Context, body *AppTableField, optFns ...request.OptFn) *AppTableFieldUpdateReqCall {
	return &AppTableFieldUpdateReqCall{
		ctx:            ctx,
		appTableFields: appTableFields,
		body:           body,
		pathParams:     map[string]interface{}{},
		optFns:         optFns,
	}
}

type AppTableRecordBatchDeleteReqCall struct {
	ctx             *core.Context
	appTableRecords *AppTableRecordService
	body            *AppTableRecordBatchDeleteReqBody
	pathParams      map[string]interface{}
	optFns          []request.OptFn
}

func (rc *AppTableRecordBatchDeleteReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableRecordBatchDeleteReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}

func (rc *AppTableRecordBatchDeleteReqCall) Do() (*AppTableRecordBatchDeleteResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &AppTableRecordBatchDeleteResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_delete", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableRecords.service.conf, req)
	return result, err
}

func (appTableRecords *AppTableRecordService) BatchDelete(ctx *core.Context, body *AppTableRecordBatchDeleteReqBody, optFns ...request.OptFn) *AppTableRecordBatchDeleteReqCall {
	return &AppTableRecordBatchDeleteReqCall{
		ctx:             ctx,
		appTableRecords: appTableRecords,
		body:            body,
		pathParams:      map[string]interface{}{},
		optFns:          optFns,
	}
}

type AppTableRecordBatchCreateReqCall struct {
	ctx             *core.Context
	appTableRecords *AppTableRecordService
	body            *AppTableRecordBatchCreateReqBody
	pathParams      map[string]interface{}
	queryParams     map[string]interface{}
	optFns          []request.OptFn
}

func (rc *AppTableRecordBatchCreateReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableRecordBatchCreateReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableRecordBatchCreateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *AppTableRecordBatchCreateReqCall) Do() (*AppTableRecordBatchCreateResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AppTableRecordBatchCreateResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_create", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableRecords.service.conf, req)
	return result, err
}

func (appTableRecords *AppTableRecordService) BatchCreate(ctx *core.Context, body *AppTableRecordBatchCreateReqBody, optFns ...request.OptFn) *AppTableRecordBatchCreateReqCall {
	return &AppTableRecordBatchCreateReqCall{
		ctx:             ctx,
		appTableRecords: appTableRecords,
		body:            body,
		pathParams:      map[string]interface{}{},
		queryParams:     map[string]interface{}{},
		optFns:          optFns,
	}
}

type AppTableRecordGetReqCall struct {
	ctx             *core.Context
	appTableRecords *AppTableRecordService
	pathParams      map[string]interface{}
	queryParams     map[string]interface{}
	optFns          []request.OptFn
}

func (rc *AppTableRecordGetReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableRecordGetReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableRecordGetReqCall) SetRecordId(recordId string) {
	rc.pathParams["record_id"] = recordId
}
func (rc *AppTableRecordGetReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *AppTableRecordGetReqCall) Do() (*AppTableRecordGetResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AppTableRecordGetResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableRecords.service.conf, req)
	return result, err
}

func (appTableRecords *AppTableRecordService) Get(ctx *core.Context, optFns ...request.OptFn) *AppTableRecordGetReqCall {
	return &AppTableRecordGetReqCall{
		ctx:             ctx,
		appTableRecords: appTableRecords,
		pathParams:      map[string]interface{}{},
		queryParams:     map[string]interface{}{},
		optFns:          optFns,
	}
}

type AppTableRecordUpdateReqCall struct {
	ctx             *core.Context
	appTableRecords *AppTableRecordService
	body            *AppTableRecord
	pathParams      map[string]interface{}
	queryParams     map[string]interface{}
	optFns          []request.OptFn
}

func (rc *AppTableRecordUpdateReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableRecordUpdateReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableRecordUpdateReqCall) SetRecordId(recordId string) {
	rc.pathParams["record_id"] = recordId
}
func (rc *AppTableRecordUpdateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *AppTableRecordUpdateReqCall) Do() (*AppTableRecordUpdateResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AppTableRecordUpdateResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id", "PUT",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableRecords.service.conf, req)
	return result, err
}

func (appTableRecords *AppTableRecordService) Update(ctx *core.Context, body *AppTableRecord, optFns ...request.OptFn) *AppTableRecordUpdateReqCall {
	return &AppTableRecordUpdateReqCall{
		ctx:             ctx,
		appTableRecords: appTableRecords,
		body:            body,
		pathParams:      map[string]interface{}{},
		queryParams:     map[string]interface{}{},
		optFns:          optFns,
	}
}

type AppTableRecordDeleteReqCall struct {
	ctx             *core.Context
	appTableRecords *AppTableRecordService
	pathParams      map[string]interface{}
	optFns          []request.OptFn
}

func (rc *AppTableRecordDeleteReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableRecordDeleteReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableRecordDeleteReqCall) SetRecordId(recordId string) {
	rc.pathParams["record_id"] = recordId
}

func (rc *AppTableRecordDeleteReqCall) Do() (*DeleteRecord, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &DeleteRecord{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/:record_id", "DELETE",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableRecords.service.conf, req)
	return result, err
}

func (appTableRecords *AppTableRecordService) Delete(ctx *core.Context, optFns ...request.OptFn) *AppTableRecordDeleteReqCall {
	return &AppTableRecordDeleteReqCall{
		ctx:             ctx,
		appTableRecords: appTableRecords,
		pathParams:      map[string]interface{}{},
		optFns:          optFns,
	}
}

type AppTableRecordListReqCall struct {
	ctx             *core.Context
	appTableRecords *AppTableRecordService
	pathParams      map[string]interface{}
	queryParams     map[string]interface{}
	optFns          []request.OptFn
}

func (rc *AppTableRecordListReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableRecordListReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableRecordListReqCall) SetViewId(viewId string) {
	rc.queryParams["view_id"] = viewId
}
func (rc *AppTableRecordListReqCall) SetFilter(filter string) {
	rc.queryParams["filter"] = filter
}
func (rc *AppTableRecordListReqCall) SetSort(sort string) {
	rc.queryParams["sort"] = sort
}
func (rc *AppTableRecordListReqCall) SetFieldNames(fieldNames string) {
	rc.queryParams["field_names"] = fieldNames
}
func (rc *AppTableRecordListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *AppTableRecordListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}
func (rc *AppTableRecordListReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *AppTableRecordListReqCall) Do() (*AppTableRecordListResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AppTableRecordListResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableRecords.service.conf, req)
	return result, err
}

func (appTableRecords *AppTableRecordService) List(ctx *core.Context, optFns ...request.OptFn) *AppTableRecordListReqCall {
	return &AppTableRecordListReqCall{
		ctx:             ctx,
		appTableRecords: appTableRecords,
		pathParams:      map[string]interface{}{},
		queryParams:     map[string]interface{}{},
		optFns:          optFns,
	}
}

type AppTableRecordBatchUpdateReqCall struct {
	ctx             *core.Context
	appTableRecords *AppTableRecordService
	body            *AppTableRecordBatchUpdateReqBody
	pathParams      map[string]interface{}
	queryParams     map[string]interface{}
	optFns          []request.OptFn
}

func (rc *AppTableRecordBatchUpdateReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableRecordBatchUpdateReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableRecordBatchUpdateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *AppTableRecordBatchUpdateReqCall) Do() (*AppTableRecordBatchUpdateResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AppTableRecordBatchUpdateResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records/batch_update", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableRecords.service.conf, req)
	return result, err
}

func (appTableRecords *AppTableRecordService) BatchUpdate(ctx *core.Context, body *AppTableRecordBatchUpdateReqBody, optFns ...request.OptFn) *AppTableRecordBatchUpdateReqCall {
	return &AppTableRecordBatchUpdateReqCall{
		ctx:             ctx,
		appTableRecords: appTableRecords,
		body:            body,
		pathParams:      map[string]interface{}{},
		queryParams:     map[string]interface{}{},
		optFns:          optFns,
	}
}

type AppTableRecordCreateReqCall struct {
	ctx             *core.Context
	appTableRecords *AppTableRecordService
	body            *AppTableRecord
	pathParams      map[string]interface{}
	queryParams     map[string]interface{}
	optFns          []request.OptFn
}

func (rc *AppTableRecordCreateReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableRecordCreateReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableRecordCreateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *AppTableRecordCreateReqCall) Do() (*AppTableRecordCreateResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AppTableRecordCreateResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/records", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableRecords.service.conf, req)
	return result, err
}

func (appTableRecords *AppTableRecordService) Create(ctx *core.Context, body *AppTableRecord, optFns ...request.OptFn) *AppTableRecordCreateReqCall {
	return &AppTableRecordCreateReqCall{
		ctx:             ctx,
		appTableRecords: appTableRecords,
		body:            body,
		pathParams:      map[string]interface{}{},
		queryParams:     map[string]interface{}{},
		optFns:          optFns,
	}
}

type AppTableViewCreateReqCall struct {
	ctx           *core.Context
	appTableViews *AppTableViewService
	body          *AppTableView
	pathParams    map[string]interface{}
	optFns        []request.OptFn
}

func (rc *AppTableViewCreateReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableViewCreateReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}

func (rc *AppTableViewCreateReqCall) Do() (*AppTableViewCreateResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &AppTableViewCreateResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views", "POST",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableViews.service.conf, req)
	return result, err
}

func (appTableViews *AppTableViewService) Create(ctx *core.Context, body *AppTableView, optFns ...request.OptFn) *AppTableViewCreateReqCall {
	return &AppTableViewCreateReqCall{
		ctx:           ctx,
		appTableViews: appTableViews,
		body:          body,
		pathParams:    map[string]interface{}{},
		optFns:        optFns,
	}
}

type AppTableViewDeleteReqCall struct {
	ctx           *core.Context
	appTableViews *AppTableViewService
	pathParams    map[string]interface{}
	optFns        []request.OptFn
}

func (rc *AppTableViewDeleteReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableViewDeleteReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableViewDeleteReqCall) SetViewId(viewId string) {
	rc.pathParams["view_id"] = viewId
}

func (rc *AppTableViewDeleteReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	var result = &response.NoData{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views/:view_id", "DELETE",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableViews.service.conf, req)
	return result, err
}

func (appTableViews *AppTableViewService) Delete(ctx *core.Context, optFns ...request.OptFn) *AppTableViewDeleteReqCall {
	return &AppTableViewDeleteReqCall{
		ctx:           ctx,
		appTableViews: appTableViews,
		pathParams:    map[string]interface{}{},
		optFns:        optFns,
	}
}

type AppTableViewListReqCall struct {
	ctx           *core.Context
	appTableViews *AppTableViewService
	pathParams    map[string]interface{}
	queryParams   map[string]interface{}
	optFns        []request.OptFn
}

func (rc *AppTableViewListReqCall) SetAppToken(appToken string) {
	rc.pathParams["app_token"] = appToken
}
func (rc *AppTableViewListReqCall) SetTableId(tableId string) {
	rc.pathParams["table_id"] = tableId
}
func (rc *AppTableViewListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}
func (rc *AppTableViewListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}

func (rc *AppTableViewListReqCall) Do() (*AppTableViewListResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AppTableViewListResult{}
	req := request.NewRequest("/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeUser, request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.appTableViews.service.conf, req)
	return result, err
}

func (appTableViews *AppTableViewService) List(ctx *core.Context, optFns ...request.OptFn) *AppTableViewListReqCall {
	return &AppTableViewListReqCall{
		ctx:           ctx,
		appTableViews: appTableViews,
		pathParams:    map[string]interface{}{},
		queryParams:   map[string]interface{}{},
		optFns:        optFns,
	}
}
