package search

import (
	"context"
	"encoding/json"
	"io"
	"reflect"

	"github.com/olivere/elastic/v7"
	"i-go/core/db/elasticsearch"
)

// 网址相关
const (
	Keywords  = "keywords"
	SiteIndex = "sites"
)

// Site 网址
type Site struct {
	SearchIndex Index
	DataType    ESite // 具体数据类型 查询时用于解析文档
}

// SiteClient 网址
var SiteClient = &Site{
	SearchIndex: Index{
		Index: SiteIndex,
	},
}

// Search 根据关键字搜索相关网站
func (s *Site) Search(keywords string) (int64, []ESite, error) {
	var (
		list = make([]ESite, 0)
	)
	// 只根据内容进行搜索
	query := elastic.NewBoolQuery().Should(elastic.NewMatchPhrasePrefixQuery(Keywords, keywords))

	svc := elasticsearch.ESClient.Search(s.SearchIndex.Index)
	res, err := svc.
		Query(query).
		From(0).
		Size(20).
		Do(context.TODO())
	if err != nil {
		return -1, list, err
	}
	for _, item := range res.Each(reflect.TypeOf(s.DataType)) {
		t := item.(ESite)
		list = append(list, t)
	}

	return res.TotalHits(), list, err
}

// SearchAll 使用scroll获取所有满足条件的文档
func (s *Site) SearchAll(keywords string) (<-chan ESite, error) {
	var (
		ch = make(chan ESite)
	)
	// 根据关键字搜索
	query := elastic.NewBoolQuery().Should(elastic.NewMatchPhrasePrefixQuery(Keywords, keywords))

	svc := elasticsearch.ESClient.
		Scroll(s.SearchIndex.Index).
		Query(query).
		Size(20)

	go func() {
		defer func() {
			_ = svc.Clear(context.TODO())
			close(ch)
		}()
		for {
			res, err := svc.Do(context.TODO())
			if err != nil {
				if err == io.EOF {
					break
				}
				continue
			}
			for _, item := range res.Each(reflect.TypeOf(s.DataType)) {
				t := item.(ESite)
				ch <- t
			}
		}
	}()
	return ch, nil
}

// Create 写入
func (s *Site) Create(e ESite) error {
	_, err := elasticsearch.ESClient.Update().
		Index(s.SearchIndex.Index).
		Id(e.ID).
		Doc(&e).
		DetectNoop(true).
		DocAsUpsert(true).
		Do(context.TODO())
	return err
}

// Read 根据id获取文档
func (s *Site) Read(id string) (ESite, error) {
	var e ESite
	res, err := elasticsearch.ESClient.Get().
		Index(s.SearchIndex.Index).
		Id(id).
		Do(context.TODO())
	if err != nil {
		return e, err
	}
	err = json.Unmarshal(res.Source, &e)
	return e, err
}

// Update 部分字段更新
func (s *Site) Update(aid string, m map[string]interface{}) error {
	_, err := elasticsearch.ESClient.Update().
		Index(s.SearchIndex.Index).
		Id(aid).
		Doc(m).
		Do(context.TODO())
	return err
}

// Delete 根据ID删除文档
func (s *Site) Delete(id string) error {
	_, err := elasticsearch.ESClient.Delete().
		Index(s.SearchIndex.Index).
		Id(id).
		Do(context.TODO())
	return err
}
