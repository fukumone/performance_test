package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
)

func index(c web.C, w http.ResponseWriter, r *http.Request) {
    j, _ := json.Marshal("OK")
    fmt.Fprintf(w, string(j))
}

func main() {
    goji.Get("/", index)
    goji.Serve()
}
