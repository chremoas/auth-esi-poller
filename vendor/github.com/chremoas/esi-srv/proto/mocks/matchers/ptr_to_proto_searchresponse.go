package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	proto "github.com/chremoas/esi-srv/proto"
)

func AnyPtrToProtoSearchResponse() *proto.SearchResponse {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*proto.SearchResponse))(nil)).Elem()))
	var nullValue *proto.SearchResponse
	return nullValue
}

func EqPtrToProtoSearchResponse(value *proto.SearchResponse) *proto.SearchResponse {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *proto.SearchResponse
	return nullValue
}
