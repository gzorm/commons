package v8

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestHttpAdd(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"https://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}

	// 示例：添加文档
	doc := map[string]interface{}{
		"id":      3,
		"title":   "CCCCCCC",
		"content": "CCCCCCC Content",
	}
	err = esClientHTTP.AddDocument("win_notice", "3", doc)
	if err != nil {
		log.Fatalf("Error adding document (HTTP): %s", err)
	}
	log.Println("Document added successfully (HTTP)")
}
func TestHttpGet(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}
	// 示例：获取文档
	docRetrieved, err := esClientHTTP.GetDocument("win_notice", "1")
	if err != nil {
		log.Fatalf("Error getting document (HTTP): %s", err)
	}
	log.Printf("Document retrieved (HTTP): %+v\n", docRetrieved)
}
func TestHttpUpdate(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}
	// 示例：更新文档
	docUpdate := map[string]interface{}{
		"title":   "Updated Document",
		"content": "This is an updated test document.",
	}
	err = esClientHTTP.UpdateDocument("test-index", "1", docUpdate)
	if err != nil {
		log.Fatalf("Error updating document (HTTP): %s", err)
	}
	log.Println("Document updated successfully (HTTP)")
}
func TestHttpDelete(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}
	// 示例：删除文档
	err = esClientHTTP.DeleteDocument("win_notice", "1")
	if err != nil {
		log.Fatalf("Error deleting document (HTTP): %s", err)
	}
	log.Println("Document deleted successfully (HTTP)")
}

// 查询带分页
func TestSearchPaging(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}

	// 分页查询
	from := 0
	size := 10

	conditions := []QueryCondition{
		{Field: "username", Operator: Wildcard, Value: "bin"},
		//{Field: "age", Operator: GreaterThan, Value: 25},
	}

	sources, total, err := esClientHTTP.SearchWithPagination("win_user", conditions, from, size)
	if err != nil {
		fmt.Println("Error querying:", err)
		return
	}
	fmt.Println("total:", total)
	for _, source := range sources {
		fmt.Println(string(source))
	}

}

// 大数据量分页A
func TestSearchWithScroll(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}

	// 分页查询
	scrollTime := 2 * time.Minute
	size := 100
	conditions := []QueryCondition{
		{Field: "username", Operator: Wildcard, Value: "bin"},
		//{Field: "age", Operator: GreaterThan, Value: 25},
	}

	results, total, err := esClientHTTP.SearchWithScroll("win_user", conditions, scrollTime, size, nil)
	if err != nil {
		fmt.Println("Error querying:", err)
		return
	}
	fmt.Println("total:", total)
	for _, result := range results {
		fmt.Println(string(result))
	}

}

type MyDoc struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type WinUser struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
}

func TestHttpSearchAfter(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	//
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}
	// 示例查询条件
	queryConditions := []QueryCondition{
		{Field: "username", Operator: Wildcard, Value: "bin"},
		//{Field: "id", Operator: Equal, Value: 9},
	}
	sources, nextAfter, total, err := esClientHTTP.SearchWithAfter("win_user", []interface{}{}, 20, queryConditions, nil)
	if err != nil {
		fmt.Println("Error querying:", err)
		return
	}

	// 反序列化
	var docs []WinUser
	for _, source := range sources {
		var doc WinUser
		if err := json.Unmarshal(source, &doc); err != nil {
			fmt.Println("Error unmarshalling docment:", err)
			continue
		}
		docs = append(docs, doc)
	}
	fmt.Println(nextAfter)
	fmt.Println("total:", total)
	fmt.Println("Docments found:", docs)

}

func TestSearch(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"https://192.168.114.133:9200"})

	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}
	conditions := []QueryCondition{
		{Field: "title", Operator: Wildcard, Value: "cc"},
		//{Field: "age", Operator: GreaterThan, Value: 25},
	}
	fields := []string{}
	results, total, err := esClientHTTP.Search("win_notice", conditions, fields, nil)
	if err != nil {
		fmt.Println("Error querying:", err)
		return
	}
	fmt.Println("total:", total)
	// 打印搜索结果
	for _, docBytes := range results {
		var doc map[string]interface{}
		if err := json.Unmarshal(docBytes, &doc); err != nil {
			fmt.Printf("Error unmarshalling document: %s\n", err)
			continue
		}
		fmt.Printf("Document: %v\n", doc)
	}

}
func TestES(test *testing.T) {
	// 使用HTTPS和证书
	esClientTLS, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"https://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (TLS): %s", err)
	}

	// 示例：创建索引
	indexMapping := `{
		"mappings": {
			"properties": {
				"title": { "type": "text" },
				"content": { "type": "text" }
			}
		}
	}`
	err = esClientTLS.CreateIndex("test-index2", indexMapping)
	if err != nil {
		log.Fatalf("Error creating index (TLS): %s", err)
	}
	log.Println("Index created successfully (TLS)")

}

/* 排序
===================Search====================================================================
sortFields := []map[string]interface{}{
    {"timestamp": "desc"},
}

results, totalCount, err := es.Search("my_index", conditions, []string{"field1", "field2"}, sortFields)
===============================SearchWithPaginationAndMultiSort================================
conditions := []QueryCondition{
    {Field: "status", Operator: Equal, Value: "active"},
}
from := 0
size := 10
sortFields := []SortField{
    {Field: "date", Order: "desc"},
    {Field: "name", Order: "asc"},
}

results, totalCount, err := esClient.SearchWithPaginationAndMultiSort("my_index", conditions, from, size, sortFields)
if err != nil {
    fmt.Println("Error executing search:", err)
} else {
    fmt.Println("Total count:", totalCount)
    fmt.Println("Search results:", results)
}
======================================SearchWithScroll============================================================
conditions := []QueryCondition{
    {Field: "status", Operator: Equal, Value: "active"},
}
scrollTime := time.Minute * 5 // 设置滚动时间为5分钟
size := 10                    // 每次查询返回10条数据

// 定义排序字段，按照日期降序、名称升序排序
sortFields := []SortField{
    {Field: "date", Order: "desc"},
    {Field: "name", Order: "asc"},
}

results, totalHits, err := esClient.SearchWithScroll("my_index", conditions, scrollTime, size, sortFields)
if err != nil {
    fmt.Println("Error executing search:", err)
} else {
    fmt.Println("Total hits:", totalHits)
    fmt.Println("Results:", results)
}
========================================================================================
*/
