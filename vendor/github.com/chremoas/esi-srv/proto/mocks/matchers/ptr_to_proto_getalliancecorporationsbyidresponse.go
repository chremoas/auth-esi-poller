package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetAllianceCorporationsByIdResponse() *proto.GetAllianceCorporationsByIdResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetAllianceCorporationsByIdResponse))(nil)).Elem()))
	var nullValue *proto.GetAllianceCorporationsByIdResponse
	return nullValue
}

func EqPtrToProtoGetAllianceCorporationsByIdResponse(value *proto.GetAllianceCorporationsByIdResponse) *proto.GetAllianceCorporationsByIdResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetAllianceCorporationsByIdResponse
	return nullValue
}
