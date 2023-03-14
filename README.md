# Govecs for HTTP

This package is a work in progress. It is not yet ready for production use.

## Installation

```bash
go get github.com/cantoramann/govecs-http
```

## Usage

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/cantoramann/govecs-http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        p := govecshttp.NewHttpPoint(0, "name", 1, 1, 1,)
        random := govecshttp.NewRandomHttpPoint()

        // ...

        fmt.Fprintf(w, "Hello, %q", r.URL.Path)
    })

    govecs_http.ListenAndServe(":8080", nil)
}




```
