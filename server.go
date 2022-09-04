package main

import (
  "fmt"
  "io"
  "log"
  "net/http"
  "os"
  "path"
  "strconv"
)

func main() {

  errorHandler := func(w http.ResponseWriter, req *http.Request) {
    errNumStr := path.Base(req.URL.Path)

    errNum, err := strconv.Atoi(errNumStr)
    if err != nil {
      w.WriteHeader(500)
      fmt.Fprintf(w, "'%s' is not a valid HTTP code\n", errNumStr)
    } else {
      w.WriteHeader(errNum)
      fmt.Fprintf(w, "HTTP Code %d\n", errNum)
    }
  }

  healthHandler := func(w http.ResponseWriter, _ *http.Request) {
    io.WriteString(w, "OK\n")
  }
  indexHandler := func(w http.ResponseWriter, _ *http.Request) {
    io.WriteString(w, "<!DOCTYPE html><html><head><title>Cloud87</title></head><body><h1>Hello World! - Cloud87 Dummy App</h1></body></html>\n")
  }
  headersHandler := func(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
      for _, h := range headers {
        fmt.Fprintf(w, "%v: %v\n", name, h)
      }
    }
  }

  panicHandler := func(w http.ResponseWriter, _ *http.Request) {
    panic("oh noes")
  }

  exitHandler := func(w http.ResponseWriter, _ *http.Request) {
    os.Exit(1)
  }
  
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/health", healthHandler)
  http.HandleFunc("/headers", headersHandler)
  http.HandleFunc("/panic", panicHandler)
  http.HandleFunc("/exit", exitHandler)
  http.HandleFunc("/error/", errorHandler)

  listenAddr := fmt.Sprintf(":%s", os.Getenv("PORT"))


  fmt.Printf("Starting server at port %s\n", listenAddr)
  if err := http.ListenAndServe(listenAddr, nil); err != nil {
      log.Fatal(err)
  }
}