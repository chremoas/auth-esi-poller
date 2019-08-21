package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoSearchRequest() *proto.SearchRequest {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.SearchRequest))(nil)).Elem()))
	var nullValue *proto.SearchRequest
	return nullValue
}

func EqPtrToProtoSearchRequest(value *proto.SearchRequest) *proto.SearchRequest {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.SearchRequest
	return nullValue
}
