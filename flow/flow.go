package flow

import (
	"github.com/juju/errors"

	"math"

	"github.com/golang/protobuf/ptypes"
	google_protobuf "github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/struct"
)

// https://medium.com/@pokstad/sending-any-any-thing-in-golang-with-protobuf-3-95f84838028d
func BuildStructReply(flowName string, data map[string]interface{}) (*Reply, error) {
	pData := &structpb.Struct{
		Fields: make(map[string]*structpb.Value),
	}

	for key, val := range data {

		pbVal, err := ValToPBVal(val)
		if err != nil {
			return nil, errors.Trace(err)
		}
		pData.Fields[key] = pbVal

	}

	any, err := ptypes.MarshalAny(pData)
	return &Reply{
		Payload: any,
	}, errors.Trace(err)
}

// TODO(waigani) remove these funcs once ReveiwRequest have been refactored out.
func ReviewRequestToRequestOptions(req *ReviewRequest) map[string]*Value {

	return map[string]*Value{
		// TODO(waigani) support slice of strings
		"Patches":       &Value{Value: &Value_StringProp{StringProp: req.Patches[0]}},
		"Host":          &Value{Value: &Value_StringProp{StringProp: req.Host}},
		"Hostname":      &Value{Value: &Value_StringProp{StringProp: req.Hostname}},
		"Repo":          &Value{Value: &Value_StringProp{StringProp: req.Repo}},
		"Sha":           &Value{Value: &Value_StringProp{StringProp: req.Sha}},
		"IsPullRequest": &Value{Value: &Value_BoolProp{BoolProp: req.IsPullRequest}},
		"PullRequestID": &Value{Value: &Value_Int64Prop{Int64Prop: req.PullRequestID}},
		"Vcs":           &Value{Value: &Value_StringProp{StringProp: req.Vcs}},
		"Dotlingo":      &Value{Value: &Value_StringProp{StringProp: req.Dotlingo}},
		"Dir":           &Value{Value: &Value_StringProp{StringProp: req.Dir}},
		"OwnerOrDepot":  &Value{Value: &Value_StringProp{StringProp: req.OwnerOrDepot.(*ReviewRequest_Owner).Owner}},
	}

}

// TODO(waigani) remove these funcs once ReveiwRequest have been refactored out.
func RequestToReviewRequest(req *Request) *ReviewRequest {

	opts := req.Options
	return &ReviewRequest{
		Host:     opts["Host"].GetStringProp(),
		Hostname: opts["Hostname"].GetStringProp(),
		Repo:     opts["Repo"].GetStringProp(),
		Sha:      opts["Sha"].GetStringProp(),
		// TODO(waigani) support slice of strings
		Patches:       []string{opts["Patch"].GetStringProp()},
		IsPullRequest: opts["IsPullRequest"].GetBoolProp(),
		PullRequestID: opts["PullRequestID"].GetInt64Prop(),
		Vcs:           opts["Vcs"].GetStringProp(),
		Dotlingo:      opts["Dotlingo"].GetStringProp(),
		Dir:           opts["Dir"].GetStringProp(),
		OwnerOrDepot:  &ReviewRequest_Owner{Owner: opts["Owner"].GetStringProp()},
	}
}

func stringSlice(val *structpb.Value) []string {

	vals := val.GetListValue().GetValues()
	result := make([]string, len(vals))

	for i := 0; i < len(vals); i++ {
		result[i] = vals[i].GetStringValue()
	}
	return result
}

func ValToIssueRange(val *structpb.Value) *IssueRange {

	fields := val.GetStructValue().GetFields()

	startFields := fields["Start"].GetStructValue().GetFields()
	endFields := fields["End"].GetStructValue().GetFields()

	return &IssueRange{
		Start: FieldsToPosition(startFields),
		End:   FieldsToPosition(endFields),
	}
}

func ValToStrMap(fields map[string]*structpb.Value) map[string]string {
	strMap := make(map[string]string, len(fields))
	for key, val := range fields {

		strMap[key] = val.GetStringValue()

	}
	return strMap
}

func FieldsToPosition(fields map[string]*structpb.Value) *Position {
	return &Position{
		Filename: fields["Filename"].GetStringValue(),
		Offset:   int64(fields["Filename"].GetNumberValue()),
		Line:     int64(fields["Filename"].GetNumberValue()),
		Column:   int64(fields["Filename"].GetNumberValue()),
	}
}

// TODO(waigani) remove these funcs once Issue has been refactored out.
func ReplyToIssue(reply *Reply) (*Issue, error) {

	pData := structpb.Struct{
		Fields: make(map[string]*structpb.Value),
	}

	if err := ptypes.UnmarshalAny(reply.Payload, &pData); err != nil {
		return nil, errors.Annotate(err, "could not unmarshal fields from payload")
	}

	fields := pData.Fields
	return &Issue{
		Name:          fields["Name"].String(),
		Position:      ValToIssueRange(fields["Position"]),
		Comment:       fields["Comment"].String(),
		CtxBefore:     fields["CtxBefore"].String(),
		LineText:      fields["LineText"].String(),
		CtxAfter:      fields["CtxAfter"].String(),
		Metrics:       ValToStrMap(fields["Metrics"].GetStructValue().GetFields()),
		Tags:          stringSlice(fields["Tags"]),
		Link:          fields["Link"].String(),
		NewCode:       fields["NewCode"].GetBoolValue(),
		Patch:         fields["Patch"].String(),
		Err:           fields["Err"].String(),
		Discard:       fields["Discard"].GetBoolValue(),
		DiscardReason: fields["DiscardReason"].String(),
	}, nil
}

func AnyStructToMap(any *google_protobuf.Any) (map[string]interface{}, error) {

	mp := make(map[string]interface{})
	var stpb structpb.Struct
	if err := ptypes.UnmarshalAny(any, &stpb); err != nil {
		return nil, errors.Annotate(err, "Could not unmarshal struct from Payload field")
	}

	for key, val := range stpb.Fields {
		kind := val.GetKind()
		switch kind.(type) {
		case *structpb.Value_NullValue:
			mp[key] = val.GetNullValue()
		case *structpb.Value_NumberValue:
			mp[key] = val.GetNumberValue()
		case *structpb.Value_StringValue:
			mp[key] = val.GetStringValue()
		case *structpb.Value_BoolValue:
			mp[key] = val.GetBoolValue()
		case *structpb.Value_StructValue:
			mp[key] = val.GetStructValue()
		case *structpb.Value_ListValue:
			mp[key] = val.GetListValue()
		default:
			return nil, errors.Errorf("unknown type %T", val)
		}
	}
	return mp, nil
}

func SliceToListValue(vals []interface{}) (*structpb.ListValue, error) {
	valLen := len(vals)

	pbVals := make([]*structpb.Value, valLen)

	for i := 0; i < valLen; i++ {
		pbVal, err := ValToPBVal(vals[i])
		if err != nil {
			return nil, errors.Trace(err)
		}
		pbVals[i] = pbVal
	}

	return &structpb.ListValue{Values: pbVals}, nil
}

func ValToPBVal(val interface{}) (pbVal *structpb.Value, err error) {

	switch v := val.(type) {
	case int:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{math.Float64frombits(uint64(v))}}, nil
	case uint64:
		return &structpb.Value{Kind: &structpb.Value_NumberValue{math.Float64frombits(v)}}, nil
	case string:
		return &structpb.Value{Kind: &structpb.Value_StringValue{v}}, nil
	case bool:
		return &structpb.Value{Kind: &structpb.Value_BoolValue{v}}, nil
	case structpb.Value_NullValue:
		return &structpb.Value{Kind: &v}, nil
	case structpb.Value_StructValue:
		return &structpb.Value{Kind: &v}, nil
	case structpb.Value_ListValue:
		return &structpb.Value{Kind: &v}, nil
	}
	return nil, errors.Errorf("unknown type %T", val)
}
