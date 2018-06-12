package main

import (
    "log"
    "net/http"
    "html"
    "fmt"
)

func main() {
    // your http.Handle calls here
    // http.Handle("/foo", fooHandler)

    http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
