// In ~/go/src/crawler/crawl.go
package main

import (
  "fmt"
  "flag"
  "os"
  // all dependencies are specified as directory names, so the slash here is important
  "net/http"
  "io/ioutil"
)

func main() {
  flag.Parse()

  args := flag.Args()

  if len(args) < 1 {
    fmt.Println("A starting web page is necessary to crawl the internet")
    os.Exit(1)
  }

  resp, err := http.Get(args[0])

  fmt.Println("Error is:", err)

  body, err := ioutil.ReadAll(resp.Body)
  fmt.Println(string(body))
}
