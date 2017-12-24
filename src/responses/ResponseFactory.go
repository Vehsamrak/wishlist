package responses

import "net/http"

type ResponseFactory struct{}

const (
    ERROR   = iota
    SUCCESS
)

func (factory ResponseFactory) create(responseWriter http.ResponseWriter, responseType uint64) Response {
    var response Response
    switch responseType {
    case ERROR:
        response = ErrorResponse{}.processResponseWriter(responseWriter)
    case SUCCESS:
        fallthrough
    default:
        response = SuccessResponse{}.processResponseWriter(responseWriter)

    }

    return response
}
