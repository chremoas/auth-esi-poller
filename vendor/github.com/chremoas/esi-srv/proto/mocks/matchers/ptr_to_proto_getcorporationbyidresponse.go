package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetCorporationByIdResponse() *proto.GetCorporationByIdResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetCorporationByIdResponse))(nil)).Elem()))
	var nullValue *proto.GetCorporationByIdResponse
	return nullValue
}

func EqPtrToProtoGetCorporationByIdResponse(value *proto.GetCorporationByIdResponse) *proto.GetCorporationByIdResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetCorporationByIdResponse
	return nullValue
}
