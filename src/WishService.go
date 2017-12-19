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
    //router.HandleFunc("/articles", ArticlesHandler)
    http.Handle("/", router)
}

func (wishService *WishService) IndexRouteHandler(writer http.ResponseWriter, request *http.Request) {
    writer.Header().Set("Content-Type", "application/json")

    if request.URL.Path != "/" {
        writer.WriteHeader(http.StatusNotFound)
        io.WriteString(writer, "{\"error\":\"Page not found\"}")
        return
    }

    writer.WriteHeader(http.StatusOK)
}
