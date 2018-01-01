package responses

import "net/http"

type ResponseFactory struct{}

const (
    ERROR   = iota
    SUCCESS
)

func (factory ResponseFactory) Create(responseWriter *http.ResponseWriter, responseType uint64) Response {
    var response Response
    switch responseType {
    case ERROR:
        response = ErrorResponse{}.ProcessResponseWriter(*responseWriter)
    case SUCCESS:
        fallthrough
    default:
        response = SuccessResponse{}.ProcessResponseWriter(*responseWriter)

    }

    return response
}
