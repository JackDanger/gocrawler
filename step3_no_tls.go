package main

import (
  "crypto/tls"   // we'll import this package to get access to some
  "flag"         // low-level transport customizations
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
                 InsecureSkipVerify: true,  // new(thing(a: b)) in other languages.
               }                            // It gives you a new 'thing' object (in this
                                            // case a new 'tls.Config' object) and sets the
                                            // 'a' attribute to a value of 'b'.

  transport := &http.Transport{    // And we take that tlsConfig object we instantiated
    TLSClientConfig: tlsConfig,    // and use it as the value for another new object's
  }                                // 'TLSClientConfig' attribute.

  client := http.Client{Transport: transport}  // Go typicaly gives you sane defaults (like 'http.Get')
                                               // and also provides a way to override them.

  resp, err := client.Get(args[0])  // this line is basically the same as before, only
  if err != nil {                   // we're calling 'Get' on a customized client rather
    return                          // than the 'http' package directly.
  }
  defer resp.Body.Close()
  
  links := collectlinks.All(resp.Body)

  for _, link := range(links) {
    fmt.Println(link)
  }
}
