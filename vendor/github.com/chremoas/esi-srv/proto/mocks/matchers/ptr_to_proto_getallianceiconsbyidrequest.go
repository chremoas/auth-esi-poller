package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetAllianceIconsByIdRequest() *proto.GetAllianceIconsByIdRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetAllianceIconsByIdRequest))(nil)).Elem()))
	var nullValue *proto.GetAllianceIconsByIdRequest
	return nullValue
}

func EqPtrToProtoGetAllianceIconsByIdRequest(value *proto.GetAllianceIconsByIdRequest) *proto.GetAllianceIconsByIdRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetAllianceIconsByIdRequest
	return nullValue
}
