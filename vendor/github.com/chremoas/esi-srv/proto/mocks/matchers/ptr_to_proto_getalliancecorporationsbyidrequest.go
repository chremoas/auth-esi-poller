package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetAllianceCorporationsByIdRequest() *proto.GetAllianceCorporationsByIdRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetAllianceCorporationsByIdRequest))(nil)).Elem()))
	var nullValue *proto.GetAllianceCorporationsByIdRequest
	return nullValue
}

func EqPtrToProtoGetAllianceCorporationsByIdRequest(value *proto.GetAllianceCorporationsByIdRequest) *proto.GetAllianceCorporationsByIdRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetAllianceCorporationsByIdRequest
	return nullValue
}
