package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetAllianceByIdRequest() *proto.GetAllianceByIdRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetAllianceByIdRequest))(nil)).Elem()))
	var nullValue *proto.GetAllianceByIdRequest
	return nullValue
}

func EqPtrToProtoGetAllianceByIdRequest(value *proto.GetAllianceByIdRequest) *proto.GetAllianceByIdRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetAllianceByIdRequest
	return nullValue
}
