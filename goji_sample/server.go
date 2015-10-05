package main

import (
    "encoding/json"
    "net/http"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

type Result struct {
    Response string
}

func result(c web.C, w http.ResponseWriter, r *http.Request) {
    result := &Result{
        Response:  "OK",
    }

    encoder := json.NewEncoder(w)
    encoder.Encode(result)
}

func main() {
    goji.Get("/", result)
    goji.Serve()
}
