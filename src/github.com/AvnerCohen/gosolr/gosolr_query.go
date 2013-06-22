package gosolr

import "fmt"
import "net/http"
import "net/url"


func ExecuteQuery(options map[string]string)  (*http.Response, error)  {
  var query string

  for key, value := range options {
    query += key + ":" + url.QueryEscape(value) +" "
  }

  conn := "http://localhost:8983/solr/collection1/select?q=%s&wt=json&indent=true"
  url :=  fmt.Sprintf(conn, query)
  fmt.Println(url)
  return  http.Get(url)
}


//Set the values for filtering, q=
func QueryFields(options map[string]string) map[string]string {

  return options;
}

//Set query filter values qf=
func QueryFilter(options map[string]string) map[string]string {

  return options;
}

//Set filter for query using special solr syntax field:[value TO other_value]
func RangeFilter(options map[string]string) map[string]string {
  return options;
}


//set sort oder , sort=
func SortBy(options map[string]string) map[string]string {
  return options;
}


//set offset=
func Offset(options map[string]string) map[string]string {
  return options;
}

//set rows=
func Limit(options map[string]string) map[string]string {
  return options;
}

//Group by  functionallity
func GroupBy(options map[string]string) map[string]string {
  return options;
}

//set field list fl=
func FieldsList(options map[string]string) map[string]string {
  return options;
}


//Set timeout for query , timeout=
func Timeout(options map[string]string) map[string]string {
  return options;
}

