package main

import (
  "net/http"
  cpprof "github.com/sunmyinf/cpprof/libs"
)

func main()  {
  http.HandleFunc("/debug/pprof/cheap", cpprof.HeapProfile)
}
