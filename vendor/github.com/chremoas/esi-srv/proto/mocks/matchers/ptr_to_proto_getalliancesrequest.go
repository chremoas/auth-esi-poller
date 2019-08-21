package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetAlliancesRequest() *proto.GetAlliancesRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetAlliancesRequest))(nil)).Elem()))
	var nullValue *proto.GetAlliancesRequest
	return nullValue
}

func EqPtrToProtoGetAlliancesRequest(value *proto.GetAlliancesRequest) *proto.GetAlliancesRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetAlliancesRequest
	return nullValue
}
