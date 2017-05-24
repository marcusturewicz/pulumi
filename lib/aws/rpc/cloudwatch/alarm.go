// *** WARNING: this file was generated by the Lumi IDL Compiler (LUMIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudwatch

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

    __sns "github.com/pulumi/lumi/lib/aws/rpc/sns"
)

/* RPC stubs for ActionTarget resource provider */

// ActionTargetToken is the type token corresponding to the ActionTarget package type.
const ActionTargetToken = tokens.Type("aws:cloudwatch/alarm:ActionTarget")

// ActionTargetProviderOps is a pluggable interface for ActionTarget-related management functionality.
type ActionTargetProviderOps interface {
    Check(ctx context.Context, obj *ActionTarget) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *ActionTarget) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*ActionTarget, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *ActionTarget, new *ActionTarget, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *ActionTarget, new *ActionTarget, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// ActionTargetProvider is a dynamic gRPC-based plugin for managing ActionTarget resources.
type ActionTargetProvider struct {
    ops ActionTargetProviderOps
}

// NewActionTargetProvider allocates a resource provider that delegates to a ops instance.
func NewActionTargetProvider(ops ActionTargetProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &ActionTargetProvider{ops: ops}
}

func (p *ActionTargetProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(ActionTargetToken))
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

func (p *ActionTargetProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(ActionTargetToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        if req.Unknowns[ActionTarget_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: obj.Name}, nil
}

func (p *ActionTargetProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(ActionTargetToken))
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

func (p *ActionTargetProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(ActionTargetToken))
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

func (p *ActionTargetProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(ActionTargetToken))
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
        if diff.Changed("topicName") {
            replaces = append(replaces, "topicName")
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

func (p *ActionTargetProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(ActionTargetToken))
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

func (p *ActionTargetProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(ActionTargetToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *ActionTargetProvider) Unmarshal(
    v *pbstruct.Struct) (*ActionTarget, resource.PropertyMap, mapper.DecodeError) {
    var obj ActionTarget
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable ActionTarget structure(s) */

// ActionTarget is a marshalable representation of its corresponding IDL type.
type ActionTarget struct {
    Name string `json:"name"`
    TopicName *string `json:"topicName,omitempty"`
    DisplayName *string `json:"displayName,omitempty"`
    Subscription *[]__sns.TopicSubscription `json:"subscription,omitempty"`
}

// ActionTarget's properties have constants to make dealing with diffs and property bags easier.
const (
    ActionTarget_Name = "name"
    ActionTarget_TopicName = "topicName"
    ActionTarget_DisplayName = "displayName"
    ActionTarget_Subscription = "subscription"
)

/* RPC stubs for Alarm resource provider */

// AlarmToken is the type token corresponding to the Alarm package type.
const AlarmToken = tokens.Type("aws:cloudwatch/alarm:Alarm")

// AlarmProviderOps is a pluggable interface for Alarm-related management functionality.
type AlarmProviderOps interface {
    Check(ctx context.Context, obj *Alarm) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *Alarm) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Alarm, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Alarm, new *Alarm, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Alarm, new *Alarm, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// AlarmProvider is a dynamic gRPC-based plugin for managing Alarm resources.
type AlarmProvider struct {
    ops AlarmProviderOps
}

// NewAlarmProvider allocates a resource provider that delegates to a ops instance.
func NewAlarmProvider(ops AlarmProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &AlarmProvider{ops: ops}
}

func (p *AlarmProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(AlarmToken))
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

func (p *AlarmProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(AlarmToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        if req.Unknowns[Alarm_Name] {
            return nil, errors.New("Name property cannot be computed from unknown outputs")
        }
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: obj.Name}, nil
}

func (p *AlarmProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(AlarmToken))
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

func (p *AlarmProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(AlarmToken))
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

func (p *AlarmProvider) InspectChange(
    ctx context.Context, req *lumirpc.InspectChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(AlarmToken))
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
        if diff.Changed("alarmName") {
            replaces = append(replaces, "alarmName")
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

func (p *AlarmProvider) Update(
    ctx context.Context, req *lumirpc.UpdateRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(AlarmToken))
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

func (p *AlarmProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(AlarmToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *AlarmProvider) Unmarshal(
    v *pbstruct.Struct) (*Alarm, resource.PropertyMap, mapper.DecodeError) {
    var obj Alarm
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable Alarm structure(s) */

// Alarm is a marshalable representation of its corresponding IDL type.
type Alarm struct {
    Name string `json:"name"`
    ComparisonOperator AlarmComparisonOperator `json:"comparisonOperator"`
    EvaluationPeriods float64 `json:"evaluationPerids"`
    MetricName string `json:"metricName"`
    Namespace string `json:"namespace"`
    Period float64 `json:"period"`
    Statistic AlarmStatistic `json:"statistic"`
    Threshold float64 `json:"threshold"`
    ActionsEnabled *bool `json:"actionsEnabled,omitempty"`
    AlarmActions *[]resource.ID `json:"alarmActions,omitempty"`
    AlarmDescription *string `json:"alarmDescription,omitempty"`
    AlarmName *string `json:"alarmName,omitempty"`
    Dimensions *[]AlarmDimension `json:"dimensions,omitempty"`
    InsufficientDataActions *[]resource.ID `json:"insufficientDataActions,omitempty"`
    OKActions *[]resource.ID `json:"okActions,omitempty"`
    Unit *AlarmMetric `json:"unit,omitempty"`
}

// Alarm's properties have constants to make dealing with diffs and property bags easier.
const (
    Alarm_Name = "name"
    Alarm_ComparisonOperator = "comparisonOperator"
    Alarm_EvaluationPeriods = "evaluationPerids"
    Alarm_MetricName = "metricName"
    Alarm_Namespace = "namespace"
    Alarm_Period = "period"
    Alarm_Statistic = "statistic"
    Alarm_Threshold = "threshold"
    Alarm_ActionsEnabled = "actionsEnabled"
    Alarm_AlarmActions = "alarmActions"
    Alarm_AlarmDescription = "alarmDescription"
    Alarm_AlarmName = "alarmName"
    Alarm_Dimensions = "dimensions"
    Alarm_InsufficientDataActions = "insufficientDataActions"
    Alarm_OKActions = "okActions"
    Alarm_Unit = "unit"
)

/* Marshalable AlarmDimension structure(s) */

// AlarmDimension is a marshalable representation of its corresponding IDL type.
type AlarmDimension struct {
    Name string `json:"name"`
    Value interface{} `json:"value"`
}

// AlarmDimension's properties have constants to make dealing with diffs and property bags easier.
const (
    AlarmDimension_Name = "name"
    AlarmDimension_Value = "value"
)

/* Typedefs */

type (
    AlarmComparisonOperator string
    AlarmMetric string
    AlarmStatistic string
)

/* Constants */

const (
    AverageStatistic AlarmStatistic = "Average"
    BitsMetric AlarmMetric = "Bits"
    BitsPerSecondMetric AlarmMetric = "Bits/Second"
    BytesMetric AlarmMetric = "Bytes"
    BytesPerSecondMetric AlarmMetric = "Bytes/Second"
    CountMetric AlarmMetric = "Count"
    CountPerSecondMetric AlarmMetric = "Count/Second"
    GigabitsMetric AlarmMetric = "Gigabits"
    GigabitsPerSecondMetric AlarmMetric = "Gigabits/Second"
    GigabytesMetric AlarmMetric = "Gigabytes"
    GigabytesPerSecondMetric AlarmMetric = "Gigabytes/Second"
    KilobitsMetric AlarmMetric = "Kilobits"
    KilobitsPerSecondMetric AlarmMetric = "Kilobits/Second"
    KilobytesMetric AlarmMetric = "Kilobytes"
    KilobytesPerSecondMetric AlarmMetric = "Kilobytes/Second"
    MaximumStatistic AlarmStatistic = "Maximum"
    MegabitsMetric AlarmMetric = "Megabits"
    MegabitsPerSecondMetric AlarmMetric = "Megabits/Second"
    MegabytesMetric AlarmMetric = "Megabytes"
    MegabytesPerSecondMetric AlarmMetric = "Megabytes/Second"
    MicrosecondsMetric AlarmMetric = "Microseconds"
    MillisecondsMetric AlarmMetric = "Milliseconds"
    MinimumStatistic AlarmStatistic = "Minimum"
    NoMetric AlarmMetric = "None"
    PercentMetric AlarmMetric = "Percent"
    SampleCountStatistic AlarmStatistic = "SampleCount"
    SecondsMetric AlarmMetric = "Seconds"
    SumStatistic AlarmStatistic = "Sum"
    TerabitsMetric AlarmMetric = "Terabits"
    TerabitsPerSecondMetric AlarmMetric = "Terabits/Second"
    TerabytesMetric AlarmMetric = "Terabytes"
    TerabytesPerSecondMetric AlarmMetric = "Terabytes/Second"
    ThresholdGreaterThan AlarmComparisonOperator = "GreaterThanThreshold"
    ThresholdGreaterThanOrEqualTo AlarmComparisonOperator = "GreaterThanOrEqualToThreshold"
    ThresholdLessThan AlarmComparisonOperator = "LessThanThreshold"
    ThresholdLessThanOrEqualTo AlarmComparisonOperator = "LessThanOrEqualToThreshold"
)


