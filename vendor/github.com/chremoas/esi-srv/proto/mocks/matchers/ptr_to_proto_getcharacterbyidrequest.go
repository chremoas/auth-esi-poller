package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoGetCharacterByIdRequest() *proto.GetCharacterByIdRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.GetCharacterByIdRequest))(nil)).Elem()))
	var nullValue *proto.GetCharacterByIdRequest
	return nullValue
}

func EqPtrToProtoGetCharacterByIdRequest(value *proto.GetCharacterByIdRequest) *proto.GetCharacterByIdRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.GetCharacterByIdRequest
	return nullValue
}
