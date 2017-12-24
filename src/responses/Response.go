package responses

import "net/http"

type Response interface {
    processResponseWriter(responseWriter http.ResponseWriter) Response
    getBody() string
}
