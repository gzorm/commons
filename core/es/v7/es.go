package v7

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

type ElasticsearchClient struct {
	client  *elasticsearch7.Client
	baseURL string
}

// NewElasticsearchClient 初始化Elasticsearch客户端
func NewElasticsearchClient(useTLS bool, caCertPath, username, password string, addresses []string) (*ElasticsearchClient, error) {
	var transport *http.Transport
	if useTLS {
		// 读取本地证书文件
		cert, err := os.ReadFile(caCertPath)
		if err != nil {
			return nil, fmt.Errorf("Error reading CA certificate: %s", err)
		}

		// 创建一个新的证书池，并添加CA证书
		caCertPool := x509.NewCertPool()
		if ok := caCertPool.AppendCertsFromPEM(cert); !ok {
			return nil, fmt.Errorf("Failed to append CA certificate")
		}

		// 配置TLS客户端，使用CA证书池
		tlsConfig := &tls.Config{
			RootCAs: caCertPool,
		}

		transport = &http.Transport{
			TLSClientConfig:       tlsConfig,
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).DialContext,
		}
	} else {
		transport = &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DialContext: (&net.Dialer{
				Timeout: 10 * time.Second,
			}).DialContext,
		}
	}

	client, err := elasticsearch7.NewClient(elasticsearch7.Config{
		Addresses: addresses,
		Username:  username,
		Password:  password,
		Transport: transport,
	})
	if err != nil {
		return nil, fmt.Errorf("Error creating the client: %s", err)
	}

	return &ElasticsearchClient{client: client, baseURL: addresses[0]}, nil
}

// CreateIndex 创建索引
func (es *ElasticsearchClient) CreateIndex(index string, mapping string) error {
	res, err := es.client.Indices.Create(index, es.client.Indices.Create.WithBody(strings.NewReader(mapping)))
	if err != nil {
		return fmt.Errorf("Error creating index: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}
	return nil
}

// AddDocument 添加文档
func (es *ElasticsearchClient) AddDocument(index string, id string, doc interface{}) error {
	docJson, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("Error marshalling document: %s", err)
	}

	res, err := es.client.Index(index, strings.NewReader(string(docJson)), es.client.Index.WithDocumentID(id))
	if err != nil {
		return fmt.Errorf("Error indexing document: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}
	return nil
}

// GetDocument 获取文档
func (es *ElasticsearchClient) GetDocument(index string, id string) (map[string]interface{}, error) {
	res, err := es.client.Get(index, id)
	if err != nil {
		return nil, fmt.Errorf("Error getting document: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return nil, fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}

	var doc map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&doc); err != nil {
		return nil, fmt.Errorf("Error parsing the response body: %s", err)
	}
	return doc, nil
}

// Search 查询(不能查询超过1W条的记录)
func (es *ElasticsearchClient) Search(index string, conditions []QueryCondition, fields []string) ([][]byte, int, error) {
	// 构建查询体
	queryBody := make(map[string]interface{})

	// 指定返回的字段
	if len(fields) > 0 {
		queryBody["_source"] = fields
	}

	// 根据查询条件构建查询语句
	if len(conditions) > 0 {
		boolQuery := make(map[string]interface{})
		mustClauses := make([]map[string]interface{}, 0)
		mustNotClauses := make([]map[string]interface{}, 0)
		for _, condition := range conditions {
			switch condition.Operator {
			case GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual:
				// 数值类型范围查询
				rangeQuery := map[string]interface{}{
					string(condition.Operator): condition.Value,
				}
				mustClauses = append(mustClauses, map[string]interface{}{"range": map[string]interface{}{condition.Field: rangeQuery}})
			case Equal:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			case Match:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"match": map[string]interface{}{condition.Field: condition.Value}})
			case WildcardLeft:
				// 左通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v", condition.Value)}})
			case WildcardRight:
				// 右通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("%v*", condition.Value)}})
			case Wildcard:
				// 双通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v*", condition.Value)}})
			case In:
				// In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("Value for 'in' operator must be a slice of interfaces")
				}
				mustClauses = append(mustClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotIn:
				// Not In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("Value for 'not in' operator must be a slice of interfaces")
				}
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotEqual:
				// 不等于查询
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			default:
				return nil, 0, fmt.Errorf("Unsupported operator: %s", condition.Operator)
			}
		}
		boolQuery["must"] = mustClauses
		if len(mustNotClauses) > 0 {
			boolQuery["must_not"] = mustNotClauses
		}
		queryBody["query"] = map[string]interface{}{"bool": boolQuery}
	}

	// 设置 size 为一个较大的值以获取所有记录
	queryBody["size"] = 10000 // 或其他适当的大值

	// 添加聚合以获取总记录数
	queryBody["aggs"] = map[string]interface{}{
		"total_count": map[string]interface{}{
			"cardinality": map[string]interface{}{
				"field": "_id",
			},
		},
	}

	// 将查询体转换为 JSON 字符串
	queryJson, err := json.Marshal(queryBody)
	if err != nil {
		return nil, 0, fmt.Errorf("Error marshalling query: %s", err)
	}

	// 创建搜索请求
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader(queryJson),
	}

	// 执行搜索
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	res, err := req.Do(ctx, es.client.Transport)
	if err != nil {
		return nil, 0, fmt.Errorf("Error executing search document: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, 0, fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}

	// 解析响应
	var result struct {
		Hits struct {
			Hits []struct {
				Source json.RawMessage `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
		Aggregations struct {
			TotalCount struct {
				Value int `json:"value"`
			} `json:"total_count"`
		} `json:"aggregations"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, fmt.Errorf("Error parsing the response body: %s", err)
	}

	// 准备返回的结果
	var results [][]byte
	for _, hit := range result.Hits.Hits {
		results = append(results, hit.Source)
	}

	// 返回结果和总记录数
	return results, result.Aggregations.TotalCount.Value, nil
}

// Search 查询并排序
func (es *ElasticsearchClient) SearchAndSort(index string, conditions []QueryCondition, fields []string, sortFields []string) ([][]byte, int, error) {
	// 构建查询体
	queryBody := make(map[string]interface{})

	// 指定返回的字段
	if len(fields) > 0 {
		queryBody["_source"] = fields
	}

	// 根据查询条件构建查询语句
	if len(conditions) > 0 {
		boolQuery := make(map[string]interface{})
		mustClauses := make([]map[string]interface{}, 0)
		mustNotClauses := make([]map[string]interface{}, 0)
		for _, condition := range conditions {
			switch condition.Operator {
			case GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual:
				// 数值类型范围查询
				rangeQuery := map[string]interface{}{
					string(condition.Operator): condition.Value,
				}
				mustClauses = append(mustClauses, map[string]interface{}{"range": map[string]interface{}{condition.Field: rangeQuery}})
			case Equal:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			case Match:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"match": map[string]interface{}{condition.Field: condition.Value}})
			case WildcardLeft:
				// 左通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v", condition.Value)}})
			case WildcardRight:
				// 右通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("%v*", condition.Value)}})
			case Wildcard:
				// 双通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v*", condition.Value)}})
			case In:
				// In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("Value for 'in' operator must be a slice of interfaces")
				}
				mustClauses = append(mustClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotIn:
				// Not In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("Value for 'not in' operator must be a slice of interfaces")
				}
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotEqual:
				// 不等于查询
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			default:
				return nil, 0, fmt.Errorf("Unsupported operator: %s", condition.Operator)
			}
		}
		boolQuery["must"] = mustClauses
		if len(mustNotClauses) > 0 {
			boolQuery["must_not"] = mustNotClauses
		}
		queryBody["query"] = map[string]interface{}{"bool": boolQuery}
	}

	// 设置 size 为一个较大的值以获取所有记录
	queryBody["size"] = 10000 // 或其他适当的大值

	// 添加聚合以获取总记录数
	queryBody["aggs"] = map[string]interface{}{
		"total_count": map[string]interface{}{
			"cardinality": map[string]interface{}{
				"field": "_id",
			},
		},
	}
	// 添加排序字段
	if len(sortFields) > 0 {
		sort := make([]map[string]interface{}, len(sortFields))
		for i, field := range sortFields {
			sort[i] = map[string]interface{}{
				field: map[string]interface{}{
					"order": "asc", // 默认升序，可根据需要修改为 "desc"
				},
			}
		}
		queryBody["sort"] = sort
	}
	// 将查询体转换为 JSON 字符串
	queryJson, err := json.Marshal(queryBody)
	if err != nil {
		return nil, 0, fmt.Errorf("Error marshalling query: %s", err)
	}

	// 创建搜索请求
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader(queryJson),
	}

	// 执行搜索
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	res, err := req.Do(ctx, es.client.Transport)
	if err != nil {
		return nil, 0, fmt.Errorf("Error executing search document: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, 0, fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}

	// 解析响应
	var result struct {
		Hits struct {
			Hits []struct {
				Source json.RawMessage `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
		Aggregations struct {
			TotalCount struct {
				Value int `json:"value"`
			} `json:"total_count"`
		} `json:"aggregations"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, fmt.Errorf("Error parsing the response body: %s", err)
	}

	// 准备返回的结果
	var results [][]byte
	for _, hit := range result.Hits.Hits {
		results = append(results, hit.Source)
	}

	// 返回结果和总记录数
	return results, result.Aggregations.TotalCount.Value, nil
}

// SearchSQL 查询
func (es *ElasticsearchClient) SearchSQL(index string, conditions []QueryCondition, fields []string, aggs map[string]interface{}, sortFields []SortField) (map[string]interface{}, error) {
	queryBody := make(map[string]interface{})

	// 指定返回的字段
	if len(fields) > 0 {
		queryBody["_source"] = fields
	}

	// 根据查询条件构建查询语句
	if len(conditions) > 0 {
		boolQuery := make(map[string]interface{})
		mustClauses := make([]map[string]interface{}, 0)
		mustNotClauses := make([]map[string]interface{}, 0)
		for _, condition := range conditions {
			switch condition.Operator {
			case GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual:
				// 数值类型范围查询
				rangeQuery := map[string]interface{}{
					string(condition.Operator): condition.Value,
				}
				mustClauses = append(mustClauses, map[string]interface{}{"range": map[string]interface{}{condition.Field: rangeQuery}})
			case Equal:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			case Match:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"match": map[string]interface{}{condition.Field: condition.Value}})
			case WildcardLeft:
				// 左通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v", condition.Value)}})
			case WildcardRight:
				// 右通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("%v*", condition.Value)}})
			case Wildcard:
				// 双通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v*", condition.Value)}})
			case In:
				// In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, fmt.Errorf("Value for 'in' operator must be a slice of interfaces")
				}
				mustClauses = append(mustClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotIn:
				// Not In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, fmt.Errorf("Value for 'not in' operator must be a slice of interfaces")
				}
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotEqual:
				// 不等于查询
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			default:
				return nil, fmt.Errorf("Unsupported operator: %s", condition.Operator)
			}
		}

		boolQuery["must"] = mustClauses
		if len(mustNotClauses) > 0 {
			boolQuery["must_not"] = mustNotClauses
		}
		queryBody["query"] = map[string]interface{}{"bool": boolQuery}
	}

	// 添加聚合功能
	if len(aggs) > 0 {
		queryBody["aggs"] = aggs
	}

	// 添加排序字段
	if len(sortFields) > 0 {
		sortArray := make([]map[string]interface{}, 0)
		for _, sortField := range sortFields {
			sortOrder := "asc" // 默认排序方式是升序
			if strings.ToLower(sortField.Order) == "desc" {
				sortOrder = "desc"
			}
			sortArray = append(sortArray, map[string]interface{}{
				sortField.Field: map[string]interface{}{
					"order": sortOrder,
				},
			})
		}
		queryBody["sort"] = sortArray
	}
	// 显式指定返回的记录数，假设需要不超过10000条记录
	queryBody["size"] = 10000

	// 将查询体转换为 JSON 字符串
	queryJson, err := json.Marshal(queryBody)
	if err != nil {
		return nil, fmt.Errorf("Error marshalling query: %s", err)
	}

	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader(queryJson),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// 执行查询请求
	res, err := req.Do(ctx, es.client.Transport)
	if err != nil {
		return nil, fmt.Errorf("Error executing search document: %s", err)
	}
	defer res.Body.Close()

	// 检查是否出错
	if res.IsError() {
		return nil, fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}

	// 解析响应
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Error parsing the response body: %s", err)
	}

	return result, nil
}

// QueryByOpenDistroSQLHaWei 原始SQL查询 华为es sql 查询
func (es *ElasticsearchClient) QueryByOpenDistroSQLHaWei(query string, formatType string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"query": query,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling query body: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/_opendistro/_sql?format=%v", es.baseURL, formatType), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := es.client.Perform(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("error performing request: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("error response from Elasticsearch: %s", res.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error parsing response body: %s", err)
	}

	return result, nil
}

// QueryByOpenDistroSQL 原始SQL查询 google es sql 查询
func (es *ElasticsearchClient) QueryByOpenDistroSQL(query string, formatType string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"query": query,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling query body: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/_sql?format=%v", es.baseURL, formatType), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := es.client.Perform(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("error performing request: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("error response from Elasticsearch: %s", res.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error parsing response body: %s", err)
	}

	return result, nil
}

func (es *ElasticsearchClient) QueryByXPackSQL(query string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"query": query,
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling query body: %s", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/_sql?format=json", es.baseURL), bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err)
	}

	req.Header.Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	res, err := es.client.Perform(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("error performing request: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, fmt.Errorf("error response from Elasticsearch: %s", res.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error parsing response body: %s", err)
	}

	return result, nil
}

// SearchWithPagination 查询带分页和排序功能
func (es *ElasticsearchClient) SearchWithPagination(index string, conditions []QueryCondition, pageIndex, size int, sortFields []SortField) ([][]byte, int, error) {
	// 构建查询体
	queryBody := make(map[string]interface{})
	queryBody["from"] = (pageIndex - 1) * size
	queryBody["size"] = size

	// 根据查询条件构建查询语句
	if len(conditions) > 0 {
		boolQuery := make(map[string]interface{})
		mustClauses := make([]map[string]interface{}, 0)
		mustNotClauses := make([]map[string]interface{}, 0)
		for _, condition := range conditions {
			switch condition.Operator {
			case GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual:
				// 数值类型范围查询
				rangeQuery := map[string]interface{}{
					string(condition.Operator): condition.Value,
				}
				mustClauses = append(mustClauses, map[string]interface{}{"range": map[string]interface{}{condition.Field: rangeQuery}})
			case Equal:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			case Match:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"match": map[string]interface{}{condition.Field: condition.Value}})
			case WildcardLeft:
				// 左通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v", condition.Value)}})
			case WildcardRight:
				// 右通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("%v*", condition.Value)}})
			case Wildcard:
				// 双通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v*", condition.Value)}})
			case In:
				// In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("Value for 'in' operator must be a slice of interfaces")
				}
				mustClauses = append(mustClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotIn:
				// Not In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("Value for 'not in' operator must be a slice of interfaces")
				}
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotEqual:
				// 不等于查询
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			default:
				return nil, 0, fmt.Errorf("Unsupported operator: %s", condition.Operator)
			}
		}

		boolQuery["must"] = mustClauses
		if len(mustNotClauses) > 0 {
			boolQuery["must_not"] = mustNotClauses
		}
		queryBody["query"] = map[string]interface{}{"bool": boolQuery}
	}

	// 添加聚合以获取总记录数
	queryBody["aggs"] = map[string]interface{}{
		"total_count": map[string]interface{}{
			"cardinality": map[string]interface{}{
				"field": "_id", // 这里根据实际需要选择合适的字段
			},
		},
	}

	// 添加排序字段
	if len(sortFields) > 0 {
		sortArray := make([]map[string]interface{}, 0)
		for _, sortField := range sortFields {
			sortOrder := "asc" // 默认排序方式是升序
			if strings.ToLower(sortField.Order) == "desc" {
				sortOrder = "desc"
			}
			sortArray = append(sortArray, map[string]interface{}{
				sortField.Field: map[string]interface{}{
					"order": sortOrder,
				},
			})
		}
		queryBody["sort"] = sortArray
	}

	// 将查询体转换为 JSON 字符串
	body, err := json.Marshal(queryBody)
	if err != nil {
		return nil, 0, fmt.Errorf("Error marshalling query body: %w", err)
	}

	// 创建搜索请求
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader(body),
	}

	// 执行搜索
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	res, err := req.Do(ctx, es.client.Transport)
	if err != nil {
		return nil, 0, fmt.Errorf("Error executing search request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, 0, fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}

	// 解析响应
	var result struct {
		Hits struct {
			Hits []struct {
				Source json.RawMessage `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
		Aggregations struct {
			TotalCount struct {
				Value int `json:"value"`
			} `json:"total_count"`
		} `json:"aggregations"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, fmt.Errorf("Error parsing the response body: %w", err)
	}

	// 提取 _source 字段作为字节数组
	sources := make([][]byte, len(result.Hits.Hits))
	for i, hit := range result.Hits.Hits {
		sources[i] = hit.Source
	}

	// 返回结果和总记录数
	return sources, result.Aggregations.TotalCount.Value, nil
}

// SearchWithScroll 使用Scroll API分页查询文档(适用于需要遍历大量数据的场景，它通过保持一个快照来避免每次查询重新计算的开销)
func (es *ElasticsearchClient) SearchWithScroll(index string, conditions []QueryCondition, scrollTime time.Duration, size int) ([][]byte, int, error) {
	// 构建查询体
	queryBody := make(map[string]interface{})
	queryBody["size"] = size

	// 根据查询条件构建查询语句
	if len(conditions) > 0 {
		boolQuery := make(map[string]interface{})
		mustClauses := make([]map[string]interface{}, 0)
		mustNotClauses := make([]map[string]interface{}, 0)
		for _, condition := range conditions {
			switch condition.Operator {
			case GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual:
				// 数值类型范围查询
				rangeQuery := map[string]interface{}{
					string(condition.Operator): condition.Value,
				}
				mustClauses = append(mustClauses, map[string]interface{}{"range": map[string]interface{}{condition.Field: rangeQuery}})
			case Equal:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			case Match:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"match": map[string]interface{}{condition.Field: condition.Value}})
			case WildcardLeft:
				// 左通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v", condition.Value)}})
			case WildcardRight:
				// 右通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("%v*", condition.Value)}})
			case Wildcard:
				// 双通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v*", condition.Value)}})
			case In:
				// In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("Value for 'in' operator must be a slice of interfaces")
				}
				mustClauses = append(mustClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotIn:
				// Not In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("Value for 'not in' operator must be a slice of interfaces")
				}
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotEqual:
				// 不等于查询
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})

			default:
				return nil, 0, fmt.Errorf("Unsupported operator: %s", condition.Operator)
			}
		}

		boolQuery["must"] = mustClauses
		if len(mustNotClauses) > 0 {
			boolQuery["must_not"] = mustNotClauses
		}
		queryBody["query"] = map[string]interface{}{"bool": boolQuery}
	}

	// 将查询体转换为 JSON 字符串
	body, err := json.Marshal(queryBody)
	if err != nil {
		return nil, 0, fmt.Errorf("Error marshalling query body: %w", err)
	}

	var results [][]byte
	var totalHits int

	res, err := es.client.Search(
		es.client.Search.WithIndex(index),
		es.client.Search.WithBody(bytes.NewReader(body)),
		es.client.Search.WithScroll(scrollTime),
		es.client.Search.WithSize(size),
	)
	if err != nil {
		return nil, 0, fmt.Errorf("Error executing search query: %s", err)
	}
	defer res.Body.Close()

	for {
		var result struct {
			ScrollID string `json:"_scroll_id"`
			Hits     struct {
				Total struct {
					Value int `json:"value"`
				} `json:"total"`
				Hits []struct {
					Source json.RawMessage `json:"_source"`
				} `json:"hits"`
			} `json:"hits"`
		}

		if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
			return nil, 0, fmt.Errorf("Error parsing the response body: %s", err)
		}

		// 更新总命中数
		totalHits = result.Hits.Total.Value //result.Total.Value

		// 收集搜索结果
		for _, hit := range result.Hits.Hits {
			results = append(results, hit.Source)
		}

		// 如果没有更多结果，退出循环
		if len(result.Hits.Hits) == 0 {
			break
		}

		// 获取 scroll_id
		scrollID := result.ScrollID
		res, err = es.client.Scroll(
			es.client.Scroll.WithScrollID(scrollID),
			es.client.Scroll.WithScroll(scrollTime),
		)
		if err != nil {
			return nil, 0, fmt.Errorf("Error executing scroll query: %s", err)
		}
		defer res.Body.Close()
	}

	return results, totalHits, nil
}

// SearchWithScroll 使用Scroll API分页查询文档(适用于需要遍历大量数据的场景，它通过保持一个快照来避免每次查询重新计算的开销)
func (es *ElasticsearchClient) SearchWithScrollAndSort(index string, conditions []QueryCondition, scrollTime time.Duration, size int, sortField string, sortOrder string) ([][]byte, int, error) {
	// 构建查询体
	queryBody := make(map[string]interface{})
	queryBody["size"] = size

	// 根据查询条件构建查询语句
	if len(conditions) > 0 {
		boolQuery := make(map[string]interface{})
		mustClauses := make([]map[string]interface{}, 0)
		mustNotClauses := make([]map[string]interface{}, 0)
		for _, condition := range conditions {
			switch condition.Operator {
			case GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual:
				// 数值类型范围查询
				rangeQuery := map[string]interface{}{
					string(condition.Operator): condition.Value,
				}
				mustClauses = append(mustClauses, map[string]interface{}{"range": map[string]interface{}{condition.Field: rangeQuery}})
			case Equal:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			case Match:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"match": map[string]interface{}{condition.Field: condition.Value}})
			case WildcardLeft:
				// 左通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v", condition.Value)}})
			case WildcardRight:
				// 右通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("%v*", condition.Value)}})
			case Wildcard:
				// 双通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v*", condition.Value)}})
			case In:
				// In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("Value for 'in' operator must be a slice of interfaces")
				}
				mustClauses = append(mustClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotIn:
				// Not In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, 0, fmt.Errorf("Value for 'not in' operator must be a slice of interfaces")
				}
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotEqual:
				// 不等于查询
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			default:
				return nil, 0, fmt.Errorf("Unsupported operator: %s", condition.Operator)
			}
		}

		boolQuery["must"] = mustClauses
		if len(mustNotClauses) > 0 {
			boolQuery["must_not"] = mustNotClauses
		}
		queryBody["query"] = map[string]interface{}{"bool": boolQuery}
	}

	// 添加排序功能
	if sortField != "" {
		// Elasticsearch 需要的排序字段格式
		sortOrderStr := "asc" // 默认升序
		if strings.ToLower(sortOrder) == "desc" {
			sortOrderStr = "desc"
		}
		queryBody["sort"] = []map[string]interface{}{
			{
				sortField: map[string]interface{}{
					"order": sortOrderStr,
				},
			},
		}
	}

	// 将查询体转换为 JSON 字符串
	body, err := json.Marshal(queryBody)
	if err != nil {
		return nil, 0, fmt.Errorf("Error marshalling query body: %w", err)
	}

	var results [][]byte
	var totalHits int

	res, err := es.client.Search(
		es.client.Search.WithIndex(index),
		es.client.Search.WithBody(bytes.NewReader(body)),
		es.client.Search.WithScroll(scrollTime),
		es.client.Search.WithSize(size),
	)
	if err != nil {
		return nil, 0, fmt.Errorf("Error executing search query: %s", err)
	}
	defer res.Body.Close()

	for {
		var result struct {
			ScrollID string `json:"_scroll_id"`
			Hits     struct {
				Total struct {
					Value int `json:"value"`
				} `json:"total"`
				Hits []struct {
					Source json.RawMessage `json:"_source"`
				} `json:"hits"`
			} `json:"hits"`
		}

		if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
			return nil, 0, fmt.Errorf("Error parsing the response body: %s", err)
		}

		// 更新总命中数
		totalHits = result.Hits.Total.Value

		// 收集搜索结果
		for _, hit := range result.Hits.Hits {
			results = append(results, hit.Source)
		}

		// 如果没有更多结果，退出循环
		if len(result.Hits.Hits) == 0 {
			break
		}

		// 获取 scroll_id
		scrollID := result.ScrollID
		res, err = es.client.Scroll(
			es.client.Scroll.WithScrollID(scrollID),
			es.client.Scroll.WithScroll(scrollTime),
		)
		if err != nil {
			return nil, 0, fmt.Errorf("Error executing scroll query: %s", err)
		}
		defer res.Body.Close()
	}

	return results, totalHits, nil
}
func (es *ElasticsearchClient) SearchWithAfter(index string, after []interface{}, size int, conditions []QueryCondition) ([][]byte, []interface{}, int, error) {
	var res *esapi.Response
	var err error

	// 构建查询体
	queryBody := make(map[string]interface{})
	queryBody["size"] = size
	if len(after) > 0 {
		queryBody["search_after"] = after
	}
	queryBody["sort"] = []map[string]string{
		{"_doc": "asc"},
	}

	// 根据查询条件构建查询语句
	if len(conditions) > 0 {
		boolQuery := make(map[string]interface{})
		mustClauses := make([]map[string]interface{}, 0)
		mustNotClauses := make([]map[string]interface{}, 0)
		for _, condition := range conditions {
			switch condition.Operator {
			case GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual:
				// 数值类型范围查询
				rangeQuery := map[string]interface{}{
					string(condition.Operator): condition.Value,
				}
				mustClauses = append(mustClauses, map[string]interface{}{"range": map[string]interface{}{condition.Field: rangeQuery}})
			case Equal:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			case Match:
				// 精确匹配
				mustClauses = append(mustClauses, map[string]interface{}{"match": map[string]interface{}{condition.Field: condition.Value}})
			case WildcardLeft:
				// 左通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v", condition.Value)}})
			case WildcardRight:
				// 右通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("%v*", condition.Value)}})
			case Wildcard:
				// 双通配符匹配
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v*", condition.Value)}})
			case In:
				// In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, nil, 0, fmt.Errorf("Value for 'in' operator must be a slice of interfaces")
				}
				mustClauses = append(mustClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotIn:
				// Not In 查询
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, nil, 0, fmt.Errorf("Value for 'not in' operator must be a slice of interfaces")
				}
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotEqual:
				// 不等于查询
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			default:
				return nil, nil, 0, fmt.Errorf("Unsupported operator: %s", condition.Operator)
			}
		}

		boolQuery["must"] = mustClauses
		if len(mustNotClauses) > 0 {
			boolQuery["must_not"] = mustNotClauses
		}
		queryBody["query"] = map[string]interface{}{"bool": boolQuery}
	}

	// 将查询体转换为 JSON 字符串
	body, err := json.Marshal(queryBody)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("Error marshalling query body: %w", err)
	}

	// 创建搜索请求
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader(body),
	}

	// 执行搜索
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	res, err = req.Do(ctx, es.client.Transport)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("Error executing search request: %w", err)
	}
	defer res.Body.Close()

	// 解析响应
	var result struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
			Hits []struct {
				Sort   []interface{}   `json:"sort"`
				Source json.RawMessage `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, nil, 0, fmt.Errorf("Error parsing response: %w", err)
	}

	// 准备下一个搜索的 after 参数
	var nextAfter []interface{}
	if len(result.Hits.Hits) > 0 {
		lastHitSort := result.Hits.Hits[len(result.Hits.Hits)-1].Sort
		for _, sortVal := range lastHitSort {
			nextAfter = append(nextAfter, sortVal)
		}
	}

	// 返回结果和下一个搜索的 after 参数
	sources := make([][]byte, len(result.Hits.Hits))
	for i, hit := range result.Hits.Hits {
		sources[i] = hit.Source
	}

	return sources, nextAfter, result.Hits.Total.Value, nil
}
func (es *ElasticsearchClient) SearchWithAfterAndSort(index string, after []interface{}, size int, conditions []QueryCondition, sortFields []map[string]interface{}) ([][]byte, []interface{}, int, error) {
	var res *esapi.Response
	var err error

	// Build the query body
	queryBody := make(map[string]interface{})
	queryBody["size"] = size
	if len(after) > 0 {
		queryBody["search_after"] = after
	}

	// Add sort functionality
	if len(sortFields) > 0 {
		queryBody["sort"] = sortFields
	} else {
		// Default sort, if no sort fields provided
		queryBody["sort"] = []map[string]string{
			{"_doc": "asc"},
		}
	}

	// Build the query conditions
	if len(conditions) > 0 {
		boolQuery := make(map[string]interface{})
		mustClauses := make([]map[string]interface{}, 0)
		mustNotClauses := make([]map[string]interface{}, 0)

		for _, condition := range conditions {
			switch condition.Operator {
			case GreaterThan, GreaterThanOrEqual, LessThan, LessThanOrEqual:
				// Range query for numeric values
				rangeQuery := map[string]interface{}{
					string(condition.Operator): condition.Value,
				}
				mustClauses = append(mustClauses, map[string]interface{}{"range": map[string]interface{}{condition.Field: rangeQuery}})
			case Equal:
				// Exact match
				mustClauses = append(mustClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			case Match:
				// Full-text match
				mustClauses = append(mustClauses, map[string]interface{}{"match": map[string]interface{}{condition.Field: condition.Value}})
			case WildcardLeft:
				// Left wildcard match
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v", condition.Value)}})
			case WildcardRight:
				// Right wildcard match
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("%v*", condition.Value)}})
			case Wildcard:
				// Full wildcard match
				mustClauses = append(mustClauses, map[string]interface{}{"wildcard": map[string]interface{}{condition.Field: fmt.Sprintf("*%v*", condition.Value)}})
			case In:
				// In query
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, nil, 0, fmt.Errorf("Value for 'in' operator must be a slice of interfaces")
				}
				mustClauses = append(mustClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotIn:
				// Not In query
				values, ok := condition.Value.([]interface{})
				if !ok {
					return nil, nil, 0, fmt.Errorf("Value for 'not in' operator must be a slice of interfaces")
				}
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"terms": map[string]interface{}{condition.Field: values}})
			case NotEqual:
				// Not equal query
				mustNotClauses = append(mustNotClauses, map[string]interface{}{"term": map[string]interface{}{condition.Field: condition.Value}})
			default:
				return nil, nil, 0, fmt.Errorf("Unsupported operator: %s", condition.Operator)
			}
		}

		boolQuery["must"] = mustClauses
		if len(mustNotClauses) > 0 {
			boolQuery["must_not"] = mustNotClauses
		}
		queryBody["query"] = map[string]interface{}{"bool": boolQuery}
	}

	// Convert query body to JSON
	body, err := json.Marshal(queryBody)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("Error marshalling query body: %w", err)
	}

	// Create search request
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader(body),
	}

	// Execute search
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	res, err = req.Do(ctx, es.client.Transport)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("Error executing search request: %w", err)
	}
	defer res.Body.Close()

	// Parse the response
	var result struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
			Hits []struct {
				Sort   []interface{}   `json:"sort"`
				Source json.RawMessage `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, nil, 0, fmt.Errorf("Error parsing response: %w", err)
	}

	// Prepare next search's "after" parameter
	var nextAfter []interface{}
	if len(result.Hits.Hits) > 0 {
		lastHitSort := result.Hits.Hits[len(result.Hits.Hits)-1].Sort
		for _, sortVal := range lastHitSort {
			nextAfter = append(nextAfter, sortVal)
		}
	}

	// Return results and "after" parameter for the next search
	sources := make([][]byte, len(result.Hits.Hits))
	for i, hit := range result.Hits.Hits {
		sources[i] = hit.Source
	}

	return sources, nextAfter, result.Hits.Total.Value, nil
}

// UpdateDocument 更新文档
func (es *ElasticsearchClient) UpdateDocument(index string, id string, doc interface{}) error {
	updateDoc := map[string]interface{}{
		"doc": doc,
	}
	docJson, err := json.Marshal(updateDoc)
	if err != nil {
		return fmt.Errorf("Error marshalling document: %s", err)
	}

	res, err := es.client.Update(index, id, strings.NewReader(string(docJson)))
	if err != nil {
		return fmt.Errorf("Error updating document: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}
	return nil
}

// DeleteDocument 删除文档
func (es *ElasticsearchClient) DeleteDocument(index string, id string) error {
	res, err := es.client.Delete(index, id)
	if err != nil {
		return fmt.Errorf("Error deleting document: %s", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("Error response from Elasticsearch: %s", res.String())
	}
	return nil
}
