package main

import "fmt"
import "io/ioutil"
import "github.com/AvnerCohen/gosolr"

func main() {

  var query_values,query_filter map[string]string
  query_values = make(map[string]string)
  query_filter = make(map[string]string)

  query_values["title_t"] ="*Warbler*"
  query_filter["body_t"] ="*Israel*"

  fmt.Println(query_filter)
  //query = gosolr.QueryFields(query_values)
  resp, error := gosolr.ExecuteQuery(query_values)

  //resp, error := gosolr.ExactMatch("title_t", "Alpine Accentor")
  if error != nil {
    fmt.Println("Some Fatal error occured:", error)
  }
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("Some Fatal error occured:", error)
  }

  resp.Body.Close()
  fmt.Println( fmt.Printf("%s", data))
}


