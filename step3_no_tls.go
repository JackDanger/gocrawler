package main

import (
  "crypto/tls"
  "flag"
  "fmt"
  "github.com/jackdanger/collectlinks"
  "net/http"
  "os"
)

func main() {
  flag.Parse()

  args := flag.Args()
  fmt.Println(args)
  if len(args) < 1 {
    fmt.Println("Please specify start page")
    os.Exit(1)
  }

  tlsConfig := &tls.Config{                 // The &thing{a: b} syntax is equivalent to
                 InsecureSkipVerify: true,  // new(thing(a: b)) in other languages
               }
  transport := &http.Transport{    // We don't have to define custom http transport
    TLSClientConfig: tlsConfig,    // or TLS config but you may find it's really handy
  }                                // when you're working with SSL
  client := http.Client{Transport: transport}

  resp, err := client.Get(args[0])  // this line is basically the same, only
  if err != nil {                   // we're calling 'Get' on the client rather
    return                          // than the 'http' package directly.
  }
  defer resp.Body.Close()
  
  links := collectlinks.All(resp.Body)

  for _, link := range(links) {
    fmt.Println(link)
  }
}

