package src

import (
    "io"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type WishService struct{}

func (wishService *WishService) Start() {
    router := mux.NewRouter()
    router.HandleFunc("/", wishService.IndexRouteHandler)
    http.Handle("/", router)

    server := &http.Server{Addr: ":12345", Handler: router}
    server.ListenAndServe()
}

func (wishService *WishService) IndexRouteHandler(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")

    if request.URL.Path != "/" {
        writer.WriteHeader(http.StatusNotFound)
        jsonResponse, _ := json.Marshal(Message{"Url not found"})
        io.WriteString(writer, string(jsonResponse))
        return
    }

    writer.WriteHeader(http.StatusOK)
    io.WriteString(writer, "{\"success\":true}")
}
