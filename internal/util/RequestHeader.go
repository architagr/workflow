package util

import (
	"net/http"
)

const (
	ApiKeyHeaderKey = 1
)

var RequestHeaderKey = map[int]string{
	ApiKeyHeaderKey: "Auth",
}

func GetHeaderKeyName(code int) string {
	return RequestHeaderKey[code]
}

func GetHeader(request *http.Request, key string) string {
	return request.Header.Get(key)
}

func GetApiKey(request *http.Request) string {
	return GetHeader(request, GetHeaderKeyName(ApiKeyHeaderKey))
}
