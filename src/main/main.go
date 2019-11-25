package main

import (
  "fmt"
  "net/http"
  "welcome"
)

func main() {
  http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(writer, welcome.Greeting("Code.education Rocks!"))
  })

  fmt.Println(http.ListenAndServe(":8000", nil))
}
