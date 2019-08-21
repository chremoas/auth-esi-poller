package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetAllianceNamesByIdsRequest() *proto.GetAllianceNamesByIdsRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetAllianceNamesByIdsRequest))(nil)).Elem()))
	var nullValue *proto.GetAllianceNamesByIdsRequest
	return nullValue
}

func EqPtrToProtoGetAllianceNamesByIdsRequest(value *proto.GetAllianceNamesByIdsRequest) *proto.GetAllianceNamesByIdsRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetAllianceNamesByIdsRequest
	return nullValue
}
