// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api"
	"github.com/larksuite/oapi-sdk-go/api/core/request"
	"github.com/larksuite/oapi-sdk-go/api/core/response"
	"github.com/larksuite/oapi-sdk-go/core"
	"github.com/larksuite/oapi-sdk-go/core/config"
	"io"
)

type Service struct {
	conf                     *config.Config
	AccessRecords            *AccessRecordService
	Users                    *UserService
	Devices                  *DeviceService
	AccessRecordAccessPhotos *AccessRecordAccessPhotoService
	UserFaces                *UserFaceService
}

func NewService(conf *config.Config) *Service {
	s := &Service{
		conf: conf,
	}
	s.AccessRecords = newAccessRecordService(s)
	s.Users = newUserService(s)
	s.Devices = newDeviceService(s)
	s.AccessRecordAccessPhotos = newAccessRecordAccessPhotoService(s)
	s.UserFaces = newUserFaceService(s)
	return s
}

type AccessRecordService struct {
	service *Service
}

func newAccessRecordService(service *Service) *AccessRecordService {
	return &AccessRecordService{
		service: service,
	}
}

type UserService struct {
	service *Service
}

func newUserService(service *Service) *UserService {
	return &UserService{
		service: service,
	}
}

type DeviceService struct {
	service *Service
}

func newDeviceService(service *Service) *DeviceService {
	return &DeviceService{
		service: service,
	}
}

type AccessRecordAccessPhotoService struct {
	service *Service
}

func newAccessRecordAccessPhotoService(service *Service) *AccessRecordAccessPhotoService {
	return &AccessRecordAccessPhotoService{
		service: service,
	}
}

type UserFaceService struct {
	service *Service
}

func newUserFaceService(service *Service) *UserFaceService {
	return &UserFaceService{
		service: service,
	}
}

type AccessRecordListReqCall struct {
	ctx           *core.Context
	accessRecords *AccessRecordService
	queryParams   map[string]interface{}
	optFns        []request.OptFn
}

func (rc *AccessRecordListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}
func (rc *AccessRecordListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *AccessRecordListReqCall) SetFrom(from int) {
	rc.queryParams["from"] = from
}
func (rc *AccessRecordListReqCall) SetTo(to int) {
	rc.queryParams["to"] = to
}
func (rc *AccessRecordListReqCall) SetDeviceId(deviceId int64) {
	rc.queryParams["device_id"] = deviceId
}
func (rc *AccessRecordListReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *AccessRecordListReqCall) Do() (*AccessRecordListResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &AccessRecordListResult{}
	req := request.NewRequest("/open-apis/acs/v1/access_records", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.accessRecords.service.conf, req)
	return result, err
}

func (accessRecords *AccessRecordService) List(ctx *core.Context, optFns ...request.OptFn) *AccessRecordListReqCall {
	return &AccessRecordListReqCall{
		ctx:           ctx,
		accessRecords: accessRecords,
		queryParams:   map[string]interface{}{},
		optFns:        optFns,
	}
}

type AccessRecordAccessPhotoGetReqCall struct {
	ctx                      *core.Context
	accessRecordAccessPhotos *AccessRecordAccessPhotoService
	pathParams               map[string]interface{}
	optFns                   []request.OptFn
	result                   io.Writer
}

func (rc *AccessRecordAccessPhotoGetReqCall) SetAccessRecordId(accessRecordId int64) {
	rc.pathParams["access_record_id"] = accessRecordId
}
func (rc *AccessRecordAccessPhotoGetReqCall) SetResponseStream(result io.Writer) {
	rc.result = result
}

func (rc *AccessRecordAccessPhotoGetReqCall) Do() (io.Writer, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetResponseStream())
	req := request.NewRequest("/open-apis/acs/v1/access_records/:access_record_id/access_photo", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, rc.result, rc.optFns...)
	err := api.Send(rc.ctx, rc.accessRecordAccessPhotos.service.conf, req)
	return rc.result, err
}

func (accessRecordAccessPhotos *AccessRecordAccessPhotoService) Get(ctx *core.Context, optFns ...request.OptFn) *AccessRecordAccessPhotoGetReqCall {
	return &AccessRecordAccessPhotoGetReqCall{
		ctx:                      ctx,
		accessRecordAccessPhotos: accessRecordAccessPhotos,
		pathParams:               map[string]interface{}{},
		optFns:                   optFns,
	}
}

type DeviceListReqCall struct {
	ctx     *core.Context
	devices *DeviceService
	optFns  []request.OptFn
}

func (rc *DeviceListReqCall) Do() (*DeviceListResult, error) {
	var result = &DeviceListResult{}
	req := request.NewRequest("/open-apis/acs/v1/devices", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.devices.service.conf, req)
	return result, err
}

func (devices *DeviceService) List(ctx *core.Context, optFns ...request.OptFn) *DeviceListReqCall {
	return &DeviceListReqCall{
		ctx:     ctx,
		devices: devices,
		optFns:  optFns,
	}
}

type UserPatchReqCall struct {
	ctx         *core.Context
	users       *UserService
	body        *User
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *UserPatchReqCall) SetUserId(userId string) {
	rc.pathParams["user_id"] = userId
}
func (rc *UserPatchReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *UserPatchReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &response.NoData{}
	req := request.NewRequest("/open-apis/acs/v1/users/:user_id", "PATCH",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.users.service.conf, req)
	return result, err
}

func (users *UserService) Patch(ctx *core.Context, body *User, optFns ...request.OptFn) *UserPatchReqCall {
	return &UserPatchReqCall{
		ctx:         ctx,
		users:       users,
		body:        body,
		pathParams:  map[string]interface{}{},
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type UserGetReqCall struct {
	ctx         *core.Context
	users       *UserService
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *UserGetReqCall) SetUserId(userId string) {
	rc.pathParams["user_id"] = userId
}
func (rc *UserGetReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *UserGetReqCall) Do() (*UserGetResult, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &UserGetResult{}
	req := request.NewRequest("/open-apis/acs/v1/users/:user_id", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.users.service.conf, req)
	return result, err
}

func (users *UserService) Get(ctx *core.Context, optFns ...request.OptFn) *UserGetReqCall {
	return &UserGetReqCall{
		ctx:         ctx,
		users:       users,
		pathParams:  map[string]interface{}{},
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type UserListReqCall struct {
	ctx         *core.Context
	users       *UserService
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *UserListReqCall) SetPageSize(pageSize int) {
	rc.queryParams["page_size"] = pageSize
}
func (rc *UserListReqCall) SetPageToken(pageToken string) {
	rc.queryParams["page_token"] = pageToken
}
func (rc *UserListReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *UserListReqCall) Do() (*UserListResult, error) {
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &UserListResult{}
	req := request.NewRequest("/open-apis/acs/v1/users", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.users.service.conf, req)
	return result, err
}

func (users *UserService) List(ctx *core.Context, optFns ...request.OptFn) *UserListReqCall {
	return &UserListReqCall{
		ctx:         ctx,
		users:       users,
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type UserFaceGetReqCall struct {
	ctx         *core.Context
	userFaces   *UserFaceService
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns      []request.OptFn
	result      io.Writer
}

func (rc *UserFaceGetReqCall) SetUserId(userId string) {
	rc.pathParams["user_id"] = userId
}
func (rc *UserFaceGetReqCall) SetIsCropped(isCropped bool) {
	rc.queryParams["is_cropped"] = isCropped
}
func (rc *UserFaceGetReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}
func (rc *UserFaceGetReqCall) SetResponseStream(result io.Writer) {
	rc.result = result
}

func (rc *UserFaceGetReqCall) Do() (io.Writer, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	rc.optFns = append(rc.optFns, request.SetResponseStream())
	req := request.NewRequest("/open-apis/acs/v1/users/:user_id/face", "GET",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, nil, rc.result, rc.optFns...)
	err := api.Send(rc.ctx, rc.userFaces.service.conf, req)
	return rc.result, err
}

func (userFaces *UserFaceService) Get(ctx *core.Context, optFns ...request.OptFn) *UserFaceGetReqCall {
	return &UserFaceGetReqCall{
		ctx:         ctx,
		userFaces:   userFaces,
		pathParams:  map[string]interface{}{},
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}

type UserFaceUpdateReqCall struct {
	ctx         *core.Context
	userFaces   *UserFaceService
	body        *request.FormData
	pathParams  map[string]interface{}
	queryParams map[string]interface{}
	optFns      []request.OptFn
}

func (rc *UserFaceUpdateReqCall) SetFiles(files *request.File) {
	rc.body.AddFile("files", files)
}
func (rc *UserFaceUpdateReqCall) SetFileType(fileType string) {
	rc.body.AddParam("file_type", fileType)
}
func (rc *UserFaceUpdateReqCall) SetFileName(fileName string) {
	rc.body.AddParam("file_name", fileName)
}
func (rc *UserFaceUpdateReqCall) SetUserId(userId string) {
	rc.pathParams["user_id"] = userId
}
func (rc *UserFaceUpdateReqCall) SetUserIdType(userIdType string) {
	rc.queryParams["user_id_type"] = userIdType
}

func (rc *UserFaceUpdateReqCall) Do() (*response.NoData, error) {
	rc.optFns = append(rc.optFns, request.SetPathParams(rc.pathParams))
	rc.optFns = append(rc.optFns, request.SetQueryParams(rc.queryParams))
	var result = &response.NoData{}
	req := request.NewRequest("/open-apis/acs/v1/users/:user_id/face", "PUT",
		[]request.AccessTokenType{request.AccessTokenTypeTenant}, rc.body, result, rc.optFns...)
	err := api.Send(rc.ctx, rc.userFaces.service.conf, req)
	return result, err
}

func (userFaces *UserFaceService) Update(ctx *core.Context, optFns ...request.OptFn) *UserFaceUpdateReqCall {
	return &UserFaceUpdateReqCall{
		ctx:         ctx,
		userFaces:   userFaces,
		body:        request.NewFormData(),
		pathParams:  map[string]interface{}{},
		queryParams: map[string]interface{}{},
		optFns:      optFns,
	}
}
