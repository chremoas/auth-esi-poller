package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetAllianceIconsByIdResponse() *proto.GetAllianceIconsByIdResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetAllianceIconsByIdResponse))(nil)).Elem()))
	var nullValue *proto.GetAllianceIconsByIdResponse
	return nullValue
}

func EqPtrToProtoGetAllianceIconsByIdResponse(value *proto.GetAllianceIconsByIdResponse) *proto.GetAllianceIconsByIdResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetAllianceIconsByIdResponse
	return nullValue
}
