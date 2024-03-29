// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/catalog/volume.proto

package catalog

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Volume service

func NewVolumeEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Volume service

type VolumeService interface {
	// 创建
	Create(ctx context.Context, in *VolumeCreateRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 更新
	Update(ctx context.Context, in *VolumeUpdateRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 删除
	Delete(ctx context.Context, in *VolumeDeleteRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 获取
	Get(ctx context.Context, in *VolumeGetRequest, opts ...client.CallOption) (*VolumeGetResponse, error)
	// 列举
	List(ctx context.Context, in *VolumeListRequest, opts ...client.CallOption) (*VolumeListResponse, error)
	// 搜索
	Search(ctx context.Context, in *VolumeSearchRequest, opts ...client.CallOption) (*VolumeListResponse, error)
	// 导出
	Export(ctx context.Context, in *VolumeExportRequest, opts ...client.CallOption) (*VolumeExportResponse, error)
}

type volumeService struct {
	c    client.Client
	name string
}

func NewVolumeService(name string, c client.Client) VolumeService {
	return &volumeService{
		c:    c,
		name: name,
	}
}

func (c *volumeService) Create(ctx context.Context, in *VolumeCreateRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Volume.Create", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeService) Update(ctx context.Context, in *VolumeUpdateRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Volume.Update", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeService) Delete(ctx context.Context, in *VolumeDeleteRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Volume.Delete", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeService) Get(ctx context.Context, in *VolumeGetRequest, opts ...client.CallOption) (*VolumeGetResponse, error) {
	req := c.c.NewRequest(c.name, "Volume.Get", in)
	out := new(VolumeGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeService) List(ctx context.Context, in *VolumeListRequest, opts ...client.CallOption) (*VolumeListResponse, error) {
	req := c.c.NewRequest(c.name, "Volume.List", in)
	out := new(VolumeListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeService) Search(ctx context.Context, in *VolumeSearchRequest, opts ...client.CallOption) (*VolumeListResponse, error) {
	req := c.c.NewRequest(c.name, "Volume.Search", in)
	out := new(VolumeListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumeService) Export(ctx context.Context, in *VolumeExportRequest, opts ...client.CallOption) (*VolumeExportResponse, error) {
	req := c.c.NewRequest(c.name, "Volume.Export", in)
	out := new(VolumeExportResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Volume service

type VolumeHandler interface {
	// 创建
	Create(context.Context, *VolumeCreateRequest, *UuidResponse) error
	// 更新
	Update(context.Context, *VolumeUpdateRequest, *UuidResponse) error
	// 删除
	Delete(context.Context, *VolumeDeleteRequest, *UuidResponse) error
	// 获取
	Get(context.Context, *VolumeGetRequest, *VolumeGetResponse) error
	// 列举
	List(context.Context, *VolumeListRequest, *VolumeListResponse) error
	// 搜索
	Search(context.Context, *VolumeSearchRequest, *VolumeListResponse) error
	// 导出
	Export(context.Context, *VolumeExportRequest, *VolumeExportResponse) error
}

func RegisterVolumeHandler(s server.Server, hdlr VolumeHandler, opts ...server.HandlerOption) error {
	type volume interface {
		Create(ctx context.Context, in *VolumeCreateRequest, out *UuidResponse) error
		Update(ctx context.Context, in *VolumeUpdateRequest, out *UuidResponse) error
		Delete(ctx context.Context, in *VolumeDeleteRequest, out *UuidResponse) error
		Get(ctx context.Context, in *VolumeGetRequest, out *VolumeGetResponse) error
		List(ctx context.Context, in *VolumeListRequest, out *VolumeListResponse) error
		Search(ctx context.Context, in *VolumeSearchRequest, out *VolumeListResponse) error
		Export(ctx context.Context, in *VolumeExportRequest, out *VolumeExportResponse) error
	}
	type Volume struct {
		volume
	}
	h := &volumeHandler{hdlr}
	return s.Handle(s.NewHandler(&Volume{h}, opts...))
}

type volumeHandler struct {
	VolumeHandler
}

func (h *volumeHandler) Create(ctx context.Context, in *VolumeCreateRequest, out *UuidResponse) error {
	return h.VolumeHandler.Create(ctx, in, out)
}

func (h *volumeHandler) Update(ctx context.Context, in *VolumeUpdateRequest, out *UuidResponse) error {
	return h.VolumeHandler.Update(ctx, in, out)
}

func (h *volumeHandler) Delete(ctx context.Context, in *VolumeDeleteRequest, out *UuidResponse) error {
	return h.VolumeHandler.Delete(ctx, in, out)
}

func (h *volumeHandler) Get(ctx context.Context, in *VolumeGetRequest, out *VolumeGetResponse) error {
	return h.VolumeHandler.Get(ctx, in, out)
}

func (h *volumeHandler) List(ctx context.Context, in *VolumeListRequest, out *VolumeListResponse) error {
	return h.VolumeHandler.List(ctx, in, out)
}

func (h *volumeHandler) Search(ctx context.Context, in *VolumeSearchRequest, out *VolumeListResponse) error {
	return h.VolumeHandler.Search(ctx, in, out)
}

func (h *volumeHandler) Export(ctx context.Context, in *VolumeExportRequest, out *VolumeExportResponse) error {
	return h.VolumeHandler.Export(ctx, in, out)
}
