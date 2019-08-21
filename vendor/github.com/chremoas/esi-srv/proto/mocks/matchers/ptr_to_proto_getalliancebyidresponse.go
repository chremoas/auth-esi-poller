package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetAllianceByIdResponse() *proto.GetAllianceByIdResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetAllianceByIdResponse))(nil)).Elem()))
	var nullValue *proto.GetAllianceByIdResponse
	return nullValue
}

func EqPtrToProtoGetAllianceByIdResponse(value *proto.GetAllianceByIdResponse) *proto.GetAllianceByIdResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetAllianceByIdResponse
	return nullValue
}
