package src

import (
    "io"
    "net/http"
    "github.com/gorilla/mux"
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
        io.WriteString(writer, "{\"error\":\"Page not found\"}")
        return
    }

    writer.WriteHeader(http.StatusOK)
    io.WriteString(writer, "{\"success\":true}")
}
