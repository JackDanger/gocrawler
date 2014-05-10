package main

import (
  "fmt"
  "flag"        // 'flag' helps you parse command line arguments"
  "os"          // 'os' gives you access to system calls
)

func main() {
  flag.Parse()         // Convert the command line arguments into a usable form
  args := flag.Args()  // and assign the arguments to a new variable named 'args'
                       // (the ':=' form means "this is a brand new variable")
  if len(args) < 1 {   // if a starting page wasn't provided as an argument
    fmt.Println("Please specify start page")  // show a usage banner and exit
    os.Exit(1)
  }
}

