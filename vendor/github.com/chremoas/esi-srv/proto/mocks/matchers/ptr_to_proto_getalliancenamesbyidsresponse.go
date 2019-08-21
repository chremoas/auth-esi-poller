package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetAllianceNamesByIdsResponse() *proto.GetAllianceNamesByIdsResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetAllianceNamesByIdsResponse))(nil)).Elem()))
	var nullValue *proto.GetAllianceNamesByIdsResponse
	return nullValue
}

func EqPtrToProtoGetAllianceNamesByIdsResponse(value *proto.GetAllianceNamesByIdsResponse) *proto.GetAllianceNamesByIdsResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetAllianceNamesByIdsResponse
	return nullValue
}
