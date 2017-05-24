// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apigateway

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

/* RPC stubs for Authorizer resource provider */

// AuthorizerToken is the type token corresponding to the Authorizer package type.
const AuthorizerToken = tokens.Type("aws:apigateway/authorizer:Authorizer")

// AuthorizerProviderOps is a pluggable interface for Authorizer-related management functionality.
type AuthorizerProviderOps interface {
    Check(ctx context.Context, obj *Authorizer) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *Authorizer) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Authorizer, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Authorizer, new *Authorizer, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Authorizer, new *Authorizer, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// AuthorizerProvider is a dynamic gRPC-based plugin for managing Authorizer resources.
type AuthorizerProvider struct {
    ops AuthorizerProviderOps
}

// NewAuthorizerProvider allocates a resource provider that delegates to a ops instance.
func NewAuthorizerProvider(ops AuthorizerProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &AuthorizerProvider{ops: ops}
}

func (p *AuthorizerProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(AuthorizerToken))
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

func (p *AuthorizerProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(AuthorizerToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        if req.Unknowns[Authorizer_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: obj.Name}, nil
}

func (p *AuthorizerProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(AuthorizerToken))
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

func (p *AuthorizerProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(AuthorizerToken))
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

func (p *AuthorizerProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(AuthorizerToken))
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
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *AuthorizerProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(AuthorizerToken))
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

func (p *AuthorizerProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(AuthorizerToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *AuthorizerProvider) Unmarshal(
    v *pbstruct.Struct) (*Authorizer, resource.PropertyMap, mapper.DecodeError) {
    var obj Authorizer
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable Authorizer structure(s) */

// Authorizer is a marshalable representation of its corresponding IDL type.
type Authorizer struct {
    Name string `json:"name"`
    Type AuthorizerType `json:"type"`
    AuthorizerCredentials *resource.ID `json:"authorizerCredentials,omitempty"`
    AuthorizerResultTTLInSeconds *float64 `json:"authorizerResultTTLInSeconds,omitempty"`
    AuthorizerURI *string `json:"authorizerURI,omitempty"`
    IdentitySource *string `json:"identitySource,omitempty"`
    IdentityValidationExpression *string `json:"identityValidationExpression,omitempty"`
    Providers *[]resource.ID `json:"providers,omitempty"`
    RestAPI *resource.ID `json:"restAPI,omitempty"`
}

// Authorizer's properties have constants to make dealing with diffs and property bags easier.
const (
    Authorizer_Name = "name"
    Authorizer_Type = "type"
    Authorizer_AuthorizerCredentials = "authorizerCredentials"
    Authorizer_AuthorizerResultTTLInSeconds = "authorizerResultTTLInSeconds"
    Authorizer_AuthorizerURI = "authorizerURI"
    Authorizer_IdentitySource = "identitySource"
    Authorizer_IdentityValidationExpression = "identityValidationExpression"
    Authorizer_Providers = "providers"
    Authorizer_RestAPI = "restAPI"
)

/* Typedefs */

type (
    AuthorizerType string
)

/* Constants */

const (
    CognitoAuthorizer AuthorizerType = "COGNITO_USER_POOLS"
    TokenAuthorizer AuthorizerType = "TOKEN"
)


