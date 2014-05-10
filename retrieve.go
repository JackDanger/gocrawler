package main
import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main() {
  resp, err := http.Get("http://motherfuckingwebsite.com")

  fmt.Println("http transport error is:", err)

  body, err := ioutil.ReadAll(resp.Body)

  fmt.Println("read error is:", err)

  fmt.Println(string(body))
}

