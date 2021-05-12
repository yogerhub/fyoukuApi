package es

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
)

var esUrl string

func init() {
	esUrl = "http://localhost:9200/"
}

// EsSearch 搜索功能
func EsSearch(indexName string, query map[string]interface{}, from, size int, sort []map[string]string) HitsData {
	searchQuery := map[string]interface{}{
		"query": query,
		"from":  from,
		"size":  size,
		"sort":  sort,
	}
	req := httplib.Post(esUrl + indexName + "/_search")
	req.JSONBody(searchQuery)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	var stb ReqSearchData
	err = json.Unmarshal([]byte(str), &stb)
	return stb.Hits
}

type ReqSearchData struct {
	Hits HitsData `json:"hits"`
}

type HitsData struct {
	Total TotalData     `json:"total"`
	Hits  []HitsTwoData `json:"hits"`
}
type HitsTwoData struct {
	Source json.RawMessage `json:"_source"`
}
type TotalData struct {
	Value    int
	Relation string
}

// EsAdd 添加ES
func EsAdd(indexName, id string, body map[string]interface{}) bool {
	req := httplib.Post(esUrl + indexName + "/_doc/" + id)
	req.JSONBody(body)
	_, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	return true
}

// EsEdit 修改ES
func EsEdit(indexName, id string, body map[string]interface{}) bool {
	req := httplib.Post(esUrl + indexName + "/_doc/" + id)
	req.JSONBody(body)
	_, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	return true
}

// EsDelete 删除Es
func EsDelete(indexName, id string) bool {
	req := httplib.Delete(esUrl + indexName + "/_doc/" + id)
	_, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	return true
}
