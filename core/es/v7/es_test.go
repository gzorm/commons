package v7

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"log"
	"testing"
	"time"
)

func TestHttpAdd(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
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
func TestSearch2(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})

	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}

	index := "win_betslips"
	// Define query conditions
	conditions := []QueryCondition{
		{Field: "created_at", Operator: GreaterThanOrEqual, Value: 1730044800},
		{Field: "created_at", Operator: LessThanOrEqual, Value: 1730649599},
	}

	sources, total, err := esClientHTTP.Search(index, conditions, []string{"*"})
	if err != nil {
		fmt.Println("Error querying:", err)
		return
	}
	fmt.Println("total:", total)
	for _, source := range sources {
		fmt.Println(string(source))
	}

}

// 查询带分页
func TestSearchPaging(t *testing.T) {
	// 使用HTTP
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}

	// 分页查询
	from := 2
	size := 10

	queryConditions := []QueryCondition{
		{Field: "created_at", Operator: GreaterThanOrEqual, Value: 1730044800},
		{Field: "created_at", Operator: LessThanOrEqual, Value: 1730649599},
		//{Field: "username", Operator: Wildcard, Value: "bin"},
		//{Field: "id", Operator: Equal, Value: 9},
		//{
		//	Field:    "id",
		//	Operator: In,
		//	Value:    []interface{}{3, 11, 16},
		//},
	}

	sources, total, err := esClientHTTP.SearchWithPagination("win_betslips", queryConditions, from, size, nil)
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
	queryConditions := []QueryCondition{
		//{Field: "username", Operator: Wildcard, Value: "bin"},
		//{Field: "id", Operator: Equal, Value: 9},
		{
			Field:    "id",
			Operator: In,
			Value:    []interface{}{3, 11, 16},
		},
	}

	results, total, err := esClientHTTP.SearchWithScroll("win_user", queryConditions, scrollTime, size)
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
		//{Field: "username", Operator: Wildcard, Value: "bin"},
		//{Field: "id", Operator: Equal, Value: 9},
		{
			Field:    "id",
			Operator: In,
			Value:    []interface{}{3, 11, 16},
		},
	}
	sources, nextAfter, total, err := esClientHTTP.SearchWithAfter("win_user", []interface{}{}, 20, queryConditions)
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
	esClientHTTP, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"http://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (HTTP): %s", err)
	}
	conditions := []QueryCondition{
		//{Field: "username", Operator: Wildcard, Value: "bin"},
		//{Field: "age", Operator: GreaterThan, Value: 25},
		{Field: "id", Operator: NotIn,
			Value: []interface{}{3, 11, 16},
		},
	}
	fields := []string{}
	results, total, err := esClientHTTP.Search("win_user", conditions, fields)
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
func TestESSQL(t *testing.T) {
	esClientTLS, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"https://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (TLS): %s", err)
	}

	// Replace "your_index" with the actual index name
	index := "win_betslips"
	// Replace with the actual condition based on your requirement
	var reqCategory = 0
	var conditions []QueryCondition
	var aggs map[string]interface{}

	if reqCategory == 1 {
		startTime := time.Now().Truncate(24 * time.Hour).Unix()

		conditions = []QueryCondition{
			{
				Field:    "created_at",
				Operator: GreaterThanOrEqual,
				Value:    startTime,
			},
		}

		aggs = map[string]interface{}{
			"group_by_username_and_gameTypeId": map[string]interface{}{
				"composite": map[string]interface{}{
					"sources": []map[string]interface{}{
						{"xb_username.keyword": map[string]interface{}{"terms": map[string]interface{}{"field": "xb_username.keyword"}}},
						{"game_provider_subtype_id": map[string]interface{}{"terms": map[string]interface{}{"field": "game_provider_subtype_id"}}},
					},
				},
				"aggs": map[string]interface{}{
					"total_profit": map[string]interface{}{
						"sum": map[string]interface{}{
							"field": "xb_profit",
						},
					},
					"total_validStake": map[string]interface{}{
						"sum": map[string]interface{}{
							"field": "valid_stake",
						},
					},
				},
			},
		}
	} else {
		aggs = map[string]interface{}{
			"group_by_username": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "xb_username.keyword",
					"order": map[string]interface{}{
						"total_profit": "desc",
					},
					"size": 10,
				},
				"aggs": map[string]interface{}{
					"total_profit": map[string]interface{}{
						"sum": map[string]interface{}{
							"field": "xb_profit",
						},
					},
					"total_validStake": map[string]interface{}{
						"sum": map[string]interface{}{
							"field": "valid_stake",
						},
					},
				},
			},
		}
	}

	result, err := esClientTLS.SearchSQL(index, conditions, nil, aggs, nil)
	if err != nil {
		log.Fatalf("Error performing search: %s", err)
	}

	// Process the results
	if reqCategory == 1 {
		aggregations := result["aggregations"].(map[string]interface{})
		groupBy := aggregations["group_by_username_and_gameTypeId"].(map[string]interface{})
		buckets := groupBy["buckets"].([]interface{})

		for _, bucket := range buckets {
			b := bucket.(map[string]interface{})
			username := b["key"].(map[string]interface{})["xb_username.keyword"].(string)
			gameTypeId := b["key"].(map[string]interface{})["game_provider_subtype_id"].(string)
			totalProfit := b["total_profit"].(map[string]interface{})["value"].(float64)
			totalValidStake := b["total_validStake"].(map[string]interface{})["value"].(float64)
			fmt.Printf("Username: %s, GameTypeId: %s, Total Profit: %f, Total Valid Stake: %f\n", username, gameTypeId, totalProfit, totalValidStake)
		}
	} else {
		aggregations := result["aggregations"].(map[string]interface{})
		groupBy := aggregations["group_by_username"].(map[string]interface{})
		buckets := groupBy["buckets"].([]interface{})

		for _, bucket := range buckets {
			b := bucket.(map[string]interface{})
			username := b["key"].(string)
			totalProfit := b["total_profit"].(map[string]interface{})["value"].(float64)
			totalValidStake := b["total_validStake"].(map[string]interface{})["value"].(float64)
			fmt.Printf("Username: %s, Total Profit: %f, Total Valid Stake: %f\n", username, totalProfit, totalValidStake)
		}
	}

}
func TestSQL(t *testing.T) {
	esClientTLS, err := NewElasticsearchClient(false, "", "elastic", "123456", []string{"https://192.168.114.133:9200"})
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client (TLS): %s", err)
	}
	// 获取当前时间的时间戳（以秒为单位）
	//currentTimestamp := time.Now().AddDate(time.Now().Year(), cast.ToInt(time.Now().Month()), -30).Unix()
	/*
				  POST _opendistro/_sql?format=json
		{
		  "query": "SELECT xb_uid , game_type_id AS gameTypeId, SUM(xb_profit) as profit, SUM(valid_stake)  AS validStake  from win_betslips  WHERE created_at >=0     GROUP BY xb_uid, game_type_id ORDER BY profit DESC LIMIT 2"
		}

			    POST _opendistro/_sql
				{
				  "query": "SELECT 1"
				}


	*/
	//currentTimestamp = 0
	//query := fmt.Sprintf(`
	//   SELECT
	//       xb_uid as uid,
	//       game_type_id AS gameTypeId,
	//       SUM(xb_profit) AS profit,
	//       SUM(valid_stake) AS validStake
	//   FROM win_betslips
	//   WHERE created_at >= %d
	//   GROUP BY uid, gameTypeId
	//   ORDER BY profit DESC LIMIT 10
	//`, currentTimestamp)
	query := "select sum(xb_profit) as totalProfit from win_betslips"
	result, err := esClientTLS.QueryByOpenDistroSQL(query, "json") // csv
	if err != nil {
		log.Fatalf("error performing SQL query: %s", err)
	}
	hits, ok := result["hits"].(map[string]interface{})
	if !ok {
		return
	}

	total, ok := hits["total"].(map[string]interface{})
	if !ok {
		return
	}
	value, ok := total["value"].(float64)
	if !ok {
		return
	}
	fmt.Println(value)
	//response, count, err := extractFields(result)
	//fmt.Println(response)
	//fmt.Println(count)
	//// 打印结果
	//fmt.Println(result)
}

type AggregationResult struct {
	Uid        int64   `json:"uid"`
	GameTypeId int64   `json:"gameTypeId"`
	Profit     float64 `json:"profit"`
	ValidStake float64 `json:"validStake"`
}

func extractFields(result map[string]interface{}) ([]AggregationResult, int, error) {
	// Extract total record count
	hits, ok := result["hits"].(map[string]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("invalid hits format")
	}

	total, ok := hits["total"].(map[string]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("invalid total format")
	}

	recordCount, ok := total["value"].(float64)
	if !ok {
		return nil, 0, fmt.Errorf("invalid record count format")
	}

	// Extract aggregations
	aggregations, ok := result["aggregations"].(map[string]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("invalid aggregations format")
	}

	var aggregationResults []AggregationResult

	uidBuckets, ok := aggregations["uid"].(map[string]interface{})["buckets"].([]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("invalid uid buckets format")
	}

	for _, uidBucket := range uidBuckets {
		uidBucketMap, ok := uidBucket.(map[string]interface{})
		if !ok {
			return nil, 0, fmt.Errorf("invalid uid bucket format")
		}

		uid, ok := uidBucketMap["key"].(float64)
		if !ok {
			return nil, 0, fmt.Errorf("invalid uid key format")
		}

		gameTypeBuckets, ok := uidBucketMap["gameTypeId"].(map[string]interface{})["buckets"].([]interface{})
		if !ok {
			return nil, 0, fmt.Errorf("invalid gameTypeId buckets format")
		}

		for _, gameTypeBucket := range gameTypeBuckets {
			gameTypeMap, ok := gameTypeBucket.(map[string]interface{})
			if !ok {
				return nil, 0, fmt.Errorf("invalid gameType bucket format")
			}

			gameTypeID, ok := gameTypeMap["key"].(float64)
			if !ok {
				return nil, 0, fmt.Errorf("invalid gameTypeId key format")
			}

			profitMap, ok := gameTypeMap["profit"].(map[string]interface{})
			if !ok {
				return nil, 0, fmt.Errorf("invalid profit format")
			}
			profit, ok := profitMap["value"].(float64)
			if !ok {
				return nil, 0, fmt.Errorf("invalid profit value format")
			}

			validStakeMap, ok := gameTypeMap["validStake"].(map[string]interface{})
			if !ok {
				return nil, 0, fmt.Errorf("invalid validStake format")
			}
			validStake, ok := validStakeMap["value"].(float64)
			if !ok {
				return nil, 0, fmt.Errorf("invalid validStake value format")
			}

			aggregationResults = append(aggregationResults, AggregationResult{
				Uid:        int64(uid),        // Convert UID from float64 to int64
				GameTypeId: int64(gameTypeID), // Convert GameTypeId from float64 to int64
				Profit:     profit,
				ValidStake: validStake,
			})
		}
	}

	return aggregationResults, int(recordCount), nil
}

func extractFields2(result map[string]interface{}) ([]AggregationResult, int, error) {
	// Extract total record count
	hits, ok := result["hits"].(map[string]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("invalid hits format")
	}

	total, ok := hits["total"].(map[string]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("invalid total format")
	}

	recordCount, ok := total["value"].(float64)
	if !ok {
		return nil, 0, fmt.Errorf("invalid record count format")
	}

	// Extract aggregations
	aggregations, ok := result["aggregations"].(map[string]interface{})
	if !ok {
		return nil, 0, fmt.Errorf("invalid aggregations format")
	}

	var aggregationResults []AggregationResult

	for _, item := range aggregations {
		buckets, ok := item.(map[string]interface{})["buckets"].([]interface{})
		if !ok {
			return nil, 0, fmt.Errorf("invalid buckets format")
		}

		for _, bucket := range buckets {
			bucketMap, ok := bucket.(map[string]interface{})
			if !ok {
				return nil, 0, fmt.Errorf("invalid bucket format")
			}

			// Extract UID
			uid, ok := bucketMap["key"].(string)
			if !ok {
				return nil, 0, fmt.Errorf("invalid key format in bucket")
			}

			// Extract gameTypeId buckets
			gameTypeBuckets, ok := bucketMap["gameTypeId"].(map[string]interface{})["buckets"].([]interface{})
			if !ok {
				return nil, 0, fmt.Errorf("invalid gameTypeId buckets format")
			}

			for _, gameTypeBucket := range gameTypeBuckets {
				gameTypeMap, ok := gameTypeBucket.(map[string]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("invalid gameType bucket format")
				}

				// Extract GameTypeId
				gameTypeID, ok := gameTypeMap["key"].(float64)
				if !ok {
					return nil, 0, fmt.Errorf("invalid gameTypeId key format")
				}

				// Extract Profit
				profitMap, ok := gameTypeMap["profit"].(map[string]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("invalid profit format")
				}
				profit, ok := profitMap["value"].(float64)
				if !ok {
					return nil, 0, fmt.Errorf("invalid profit value format")
				}

				// Extract ValidStake
				validStakeMap, ok := gameTypeMap["validStake"].(map[string]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("invalid validStake format")
				}
				validStake, ok := validStakeMap["value"].(float64)
				if !ok {
					return nil, 0, fmt.Errorf("invalid validStake value format")
				}

				aggregationResults = append(aggregationResults, AggregationResult{
					Uid:        cast.ToInt64(uid), // Convert UID from string to int64
					GameTypeId: int64(gameTypeID), // Convert GameTypeId from float64 to int64
					Profit:     profit,
					ValidStake: validStake,
				})
			}
		}
	}

	return aggregationResults, int(recordCount), nil
}
func TestBBB(t *testing.T) {

}

/*remark
 ================SearchWithAfterAndSort=========================
queryBody["sort"] = []map[string]string{
    {"_doc": "asc"},
}

=================SearchWithScrollAndSort===========================
results, totalHits, err := es.SearchWithScrollAndSort("my_index", conditions, 1*time.Minute, 10, "timestamp", "desc")
if err != nil {
    log.Fatalf("Search failed: %v", err)
}
fmt.Printf("Found %d documents\n", totalHits)

==============================SearchWithPaginationAndSort=====================================
sortFields := []SortField{
    {Field: "timestamp", Order: "desc"},
    {Field: "sort_order", Order: "asc"},
}

sources, totalCount, err := es.SearchWithPaginationAndSort("my_index", conditions, 0, 10, sortFields)
if err != nil {
    log.Fatalf("Search failed: %v", err)
}
==============================SearchSQLAndSort==============================================
sortFields := []SortField{
    {Field: "timestamp", Order: "desc"},
    {Field: "sort_order", Order: "asc"},
}

results, err := es.SearchSQLAndSort("my_index", conditions, fields, aggs, sortFields)
if err != nil {
    log.Fatalf("Search failed: %v", err)
}
*/
