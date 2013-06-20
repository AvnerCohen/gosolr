package gosolr

import "fmt"
import "net/http"
import "net/url"

func DeleteById(field string, value string)  (*http.Response, error) {
    conn := "http://localhost:8983/solr/collection1/select?q=%s:'%s'&wt=json&indent=true"
    url := fmt.Sprintf(conn, field, url.QueryEscape(value))
    fmt.Println(url)
    return  http.Get(url)
}
