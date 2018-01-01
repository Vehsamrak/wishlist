package responses

import "net/http"

type Response interface {
    ProcessResponseWriter(responseWriter http.ResponseWriter) Response
    GetBody() string
    GetWriter() http.ResponseWriter
}
