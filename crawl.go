package main

import (
  "flag"
  "fmt"
  "os"
)

func usage() {
  fmt.Fprintf(os.Stderr, "usage: crawl http://example.com/path/file.html\n")
  flag.PrintDefaults()
  os.Exit(2)
}

func main() {
  flag.Usage = usage
  flag.Parse()

  args := flag.Args()
  if len(args) < 1 {
    usage()
    fmt.Println("A starting web page is necessary to crawl the internet")
    os.Exit(1)
  }
  Start(args[0])
}
