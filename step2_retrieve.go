package main    // Note: this is a new file, not our crawl.go

import (
  "fmt"
  "net/http"    // this is the 'http' package we'll be using to retrieve a page
  "io/ioutil"   // we'll only use 'ioutil' to maek reading and printing
)               // the html page a little easier in this example.

func main() {
  resp, err := http.Get("http://6brand.com.com")  // See how we assign two variables at once here?
                                                  // That destructuring is really common
                                                  // in Go. The way you handle errors in
                                                  // Go is to expect that functions you
                                                  // call will return two things and the
                                                  // second one will be an error. If the
                                                  // error is nil then you can continue
                                                  // but if it's not you need to handle it.
  fmt.Println("http transport error is:", err)

  body, err := ioutil.ReadAll(resp.Body)  // resp.Body isn't a string, it's more like a reference
                                          // to a stream of data. So we use the 'ioutil'
                                          // package to read it into memory for us.
  fmt.Println("read error is:", err)

  fmt.Println(string(body))   // We cast the html body to a string because
}                             // Go hands it to us as a byte array
