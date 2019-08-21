package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetCorporationByIdRequest() *proto.GetCorporationByIdRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetCorporationByIdRequest))(nil)).Elem()))
	var nullValue *proto.GetCorporationByIdRequest
	return nullValue
}

func EqPtrToProtoGetCorporationByIdRequest(value *proto.GetCorporationByIdRequest) *proto.GetCorporationByIdRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetCorporationByIdRequest
	return nullValue
}
