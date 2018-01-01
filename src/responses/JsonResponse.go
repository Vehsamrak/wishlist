package responses

import "net/http"

type JsonResponse struct {
    Code           int
    Body           string
    writer http.ResponseWriter
}

func (response JsonResponse) ProcessResponseWriter(responseWriter http.ResponseWriter) Response {
    response.writer = responseWriter
    response.writer.Header().Set("Content-Type", "application/json")
    response.writer.WriteHeader(response.Code)

    return response
}

func (response JsonResponse) GetBody() string {
    return response.Body
}

func (response JsonResponse) GetWriter() http.ResponseWriter {
    return response.writer
}

func (response JsonResponse) GetWriter() http.ResponseWriter {
    return response.writer
}
