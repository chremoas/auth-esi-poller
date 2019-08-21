package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetCharacterByIdResponse() *proto.GetCharacterByIdResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetCharacterByIdResponse))(nil)).Elem()))
	var nullValue *proto.GetCharacterByIdResponse
	return nullValue
}

func EqPtrToProtoGetCharacterByIdResponse(value *proto.GetCharacterByIdResponse) *proto.GetCharacterByIdResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetCharacterByIdResponse
	return nullValue
}
