// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ec2

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/lumi/pkg/resource"
    "github.com/pulumi/lumi/pkg/tokens"
    "github.com/pulumi/lumi/pkg/util/contract"
    "github.com/pulumi/lumi/pkg/util/mapper"
    "github.com/pulumi/lumi/sdk/go/pkg/lumirpc"
)

/* RPC stubs for VPCGatewayAttachment resource provider */

// VPCGatewayAttachmentToken is the type token corresponding to the VPCGatewayAttachment package type.
const VPCGatewayAttachmentToken = tokens.Type("aws:ec2/vpcGatewayAttachment:VPCGatewayAttachment")

// VPCGatewayAttachmentProviderOps is a pluggable interface for VPCGatewayAttachment-related management functionality.
type VPCGatewayAttachmentProviderOps interface {
    Check(ctx context.Context, obj *VPCGatewayAttachment) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *VPCGatewayAttachment) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*VPCGatewayAttachment, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *VPCGatewayAttachment, new *VPCGatewayAttachment, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *VPCGatewayAttachment, new *VPCGatewayAttachment, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// VPCGatewayAttachmentProvider is a dynamic gRPC-based plugin for managing VPCGatewayAttachment resources.
type VPCGatewayAttachmentProvider struct {
    ops VPCGatewayAttachmentProviderOps
}

// NewVPCGatewayAttachmentProvider allocates a resource provider that delegates to a ops instance.
func NewVPCGatewayAttachmentProvider(ops VPCGatewayAttachmentProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &VPCGatewayAttachmentProvider{ops: ops}
}

func (p *VPCGatewayAttachmentProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(VPCGatewayAttachmentToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *VPCGatewayAttachmentProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(VPCGatewayAttachmentToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        if req.Unknowns[VPCGatewayAttachment_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: obj.Name}, nil
}

func (p *VPCGatewayAttachmentProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(VPCGatewayAttachmentToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &lumirpc.CreateResponse{
        Id:   string(id),
    }, nil
}

func (p *VPCGatewayAttachmentProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(VPCGatewayAttachmentToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &lumirpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *VPCGatewayAttachmentProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(VPCGatewayAttachmentToken))
    id := resource.ID(req.GetId())
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
        if diff.Changed("vpc") {
            replaces = append(replaces, "vpc")
        }
        if diff.Changed("internetGateway") {
            replaces = append(replaces, "internetGateway")
        }
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *VPCGatewayAttachmentProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(VPCGatewayAttachmentToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *VPCGatewayAttachmentProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(VPCGatewayAttachmentToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *VPCGatewayAttachmentProvider) Unmarshal(
    v *pbstruct.Struct) (*VPCGatewayAttachment, resource.PropertyMap, mapper.DecodeError) {
    var obj VPCGatewayAttachment
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable VPCGatewayAttachment structure(s) */

// VPCGatewayAttachment is a marshalable representation of its corresponding IDL type.
type VPCGatewayAttachment struct {
    Name string `json:"name"`
    VPC resource.ID `json:"vpc"`
    InternetGateway resource.ID `json:"internetGateway"`
}

// VPCGatewayAttachment's properties have constants to make dealing with diffs and property bags easier.
const (
    VPCGatewayAttachment_Name = "name"
    VPCGatewayAttachment_VPC = "vpc"
    VPCGatewayAttachment_InternetGateway = "internetGateway"
)


