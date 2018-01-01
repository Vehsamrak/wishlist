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
    router.HandleFunc("/wishes", wishService.WishesHandler)
    router.HandleFunc("/", wishService.IndexHandler)
    http.Handle("/", router)

    server := &http.Server{Addr: ":12345", Handler: router}
    server.ListenAndServe()
}

func (wishService *WishService) IndexHandler(responseWriter http.ResponseWriter, request *http.Request) {
    responseWriter.WriteHeader(http.StatusNotFound)
    jsonResponse, _ := json.Marshal(Message{"Url not found"})
    io.WriteString(responseWriter, string(jsonResponse))
}

func (wishService *WishService) WishesHandler(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "application/json")

    userParameter := request.URL.Query()["userId"]

    if len(userParameter) < 1 {
        response.WriteHeader(http.StatusBadRequest)
        jsonResponse, _ := json.Marshal(Message{"Parameter \"userId\" is missing"})
        io.WriteString(response, string(jsonResponse))
        return
    }

    userId := userParameter[0]

    response.WriteHeader(http.StatusOK)
    jsonResponse, _ := json.Marshal(Message{"userId:" + userId})
    io.WriteString(response, string(jsonResponse))
}
