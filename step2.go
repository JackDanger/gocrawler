package main

import (
  "flag"          // 'flag', 'fmt' and 'os' we'll keep around
  "fmt"
  "net/http"      // 'http' will retrieve pages for us
  "io/ioutil"     // 'ioutil' will help us print pages to the screen
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
  retrieve(args[0])  // The reason we can call 'retrieve' is because
                     // it's defined in the same package as the calling function.
}

func retrieve(uri string) {  // This func(tion) takes a parameter and the
                             // format for a function parameter definition is
                             // to say what the name of the parameter is and then
                             // the type.
                             // So here we're expecting to be given a
                             // string that we'll refer to as 'uri'
  resp, err := http.Get(uri)
  if err != nil {            // This is the way error handling typically works in Go.
    return                   // It's a bit verbose but it works.
  }
  defer resp.Body.Close()  // Important: we need to close the resource we opened
                           // (the TCP connection to some web server and our reference
                           // to the stream of data it sends us).
                           // `defer` delays an operation until the function ends.
                           // It's basically the same as if you'd moved the code
                           // you're deferring to the very last line of the func.

  body, _ := ioutil.ReadAll(resp.Body)  // I'm assigning the err to _ 'cause
                                        // I don't care about it but Go will whine
  fmt.Println(string(body))             // if I name it and don't use it
}
