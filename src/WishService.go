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
    router.HandleFunc("/wishes/{userId}", wishService.WishesHandler)
    router.HandleFunc("/", wishService.IndexHandler)
    http.Handle("/", router)

    server := &http.Server{Addr: ":12345", Handler: router}
    server.ListenAndServe()
}

func (wishService *WishService) IndexHandler(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "application/json")


    if request.URL.Path != "/" {
        response.WriteHeader(http.StatusNotFound)
        jsonResponse, _ := json.Marshal(Message{"Url not found"})
        io.WriteString(response, string(jsonResponse))
        return
    }

    response.WriteHeader(http.StatusOK)
    jsonResponse, _ := json.Marshal(Message{"Url found"})
    io.WriteString(response, string(jsonResponse))
}

func (wishService *WishService) WishesHandler(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(request)

    if len(vars["userId"]) == 0 {
        response.WriteHeader(http.StatusNotFound)
        jsonResponse, _ := json.Marshal(Message{"Url found"})
        io.WriteString(response, string(jsonResponse))
        return
    }
}
