package models

const (
	ErrorTypeFatal      = 1
	ErrorTypeError      = 2
	ErrorTypeValidation = 3
	ErrorTypeInfo       = 4
	ErrorTypeDebug      = 5
)

var ErrorTypes = map[int]string{
	ErrorTypeFatal:      "Fatal",
	ErrorTypeError:      "Error",
	ErrorTypeValidation: "Validation Error",
	ErrorTypeInfo:       "Info",
	ErrorTypeDebug:      "Debug",
}

func GetErrorTypeName(errorCode int) string {
	if errorCode > 0 && errorCode <= 5 {
		return ErrorTypes[errorCode]
	}
	return ""
}

type ErrorDetail struct {
	ErrorType    string
	ErrorMessage string
}
type Response struct {
	Data    interface{}
	Status  int
	Error   []ErrorDetail
	Message string
}

type ListResponse struct {
	Data           interface{}
	TotalCount     int
	StartPageIndex int
	EndPageIndex   int
	Status         int
	Error          []ErrorDetail
	Message        string
}
