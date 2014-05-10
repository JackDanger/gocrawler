package main

import (
  "flag"
  "fmt"
  "github.com/jackdanger/collectlinks"  // This is the little library I made for 
  "net/http"                            // parsing links. Go natively allows sourcing
  "os"                                  // Github projects as dependencies. They'll be
)                                       // downloaded to $GOPATH/src/github.com/... on your
                                        // filesystem but you don't have to worry about that.
func main() {
  flag.Parse()

  args := flag.Args()
  fmt.Println(args)
  if len(args) < 1 {
    fmt.Println("Please specify start page")
    os.Exit(1)
  }

  resp, err := http.Get(args[0])
  if err != nil {
    return
  }
  defer resp.Body.Close()
  
  links := collectlinks.All(resp.Body)  // Here we use the collectlinks package

  for _, link := range(links) {  // 'for' + 'range' in Go is like .each in Ruby or
    fmt.Println(link)            // an iterator in many other languages.
  }                              // When we call range() on a list each iteration of the
}                                // contents will set two variables: the index and the value.
                                 // Here we don't care about the index so we set it to '_'
                                 // because if we write 'for index, link := range(links)'
                                 // Go would point out that we left 'index' unused.
