package main

import (
    "net/http"
    "io/ioutil"
    "time"
    "fmt"
    "flag"
)

func main() {
    start := time.Now()
    var port string
    flag.StringVar(&port, "port", "3000", "文字列を入力します。")
    flag.Parse()

    url := fmt.Sprintf("http://localhost:%s/", port)
    for i := 0; i < 10000; i++ {
        resp, _ := http.Get(url)
        defer resp.Body.Close()
        byteArray, _ := ioutil.ReadAll(resp.Body)
        fmt.Println(string(byteArray))
    }
    end := time.Now().Sub(start)
    fmt.Printf("%s秒", end)
}
