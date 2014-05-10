package main
import (
  "github.com/jackdanger/collectlinks"
  "net/http"
  "fmt"
)

func main() {
  resp, _ := http.Get("http://motherfuckingwebsite.com")
  links := collectlinks.All(resp.Body)
  fmt.Println(links)
}

