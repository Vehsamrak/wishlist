package responses

import "net/http"

type JsonResponse struct {
    Code        string
    Body        string
}

func (response JsonResponse) processResponseWriter(responseWriter http.ResponseWriter) Response {
    responseWriter.Header().Set("Content-Type", "application/json")
}

func (response JsonResponse) getBody() string {
    return response.Body
}
