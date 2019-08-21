package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetAlliancesResponse() *proto.GetAlliancesResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetAlliancesResponse))(nil)).Elem()))
	var nullValue *proto.GetAlliancesResponse
	return nullValue
}

func EqPtrToProtoGetAlliancesResponse(value *proto.GetAlliancesResponse) *proto.GetAlliancesResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetAlliancesResponse
	return nullValue
}
