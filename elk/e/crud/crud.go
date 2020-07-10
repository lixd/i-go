package crud

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"i-go/utils"
	"reflect"
)

type IHello interface {
}

type hello struct {
	Index string `json:"index"`
	Type  string `json:"type"` // 7.0 后只有一个 type _doc
	ESCli *elastic.Client
}
type HelloItem struct {
	Name string `json:"name"`
}

// MatchPrefix 前缀匹配分类
func (s *hello) MatchPrefix(keyword string, count int) ([]string, error) {
	query := elastic.NewMatchPhrasePrefixQuery("name", keyword)
	res, err := s.ESCli.Search().
		Index(s.Index).
		Type(s.Type).
		Query(query).
		Size(count).
		Do(context.Background())
	if err != nil {
		logrus.WithFields(logrus.Fields{"caller": utils.Caller()}).Error(err)
	}
	var list = make([]string, 0)
	var item HelloItem
	resp := res.Each(reflect.TypeOf(item))
	for _, item := range resp {
		t := item.(HelloItem)
		list = append(list, t.Name)
	}
	return list, err
}

// Insert 添加到 ES，已经存在就不添加了 创建、更新问题时记得更新
func (s *hello) Insert(keyword string) {
	// 查询是否存在 完全匹配
	where := elastic.NewTermQuery("name.keyword", keyword)
	docId, err := s.ESCli.Search().
		Index(s.Index).
		Type(s.Type).
		Query(where).
		Do(context.Background())
	if err != nil {
		logrus.WithFields(logrus.Fields{"CallerName": utils.Caller(), "Scenes": "ES 标签查询"}).Error(err)
		return
	}
	// 数量为不为 0 就直接返回
	total := docId.Hits.TotalHits
	if total != 0 {
		return
	}
	// 否则写入
	body := HelloItem{Name: keyword}
	_, err = s.ESCli.Index().
		Index(s.Index).
		Type(s.Type).
		BodyJson(body).
		Do(context.Background())
	if err != nil {
		logrus.WithFields(logrus.Fields{"CallerName": utils.Caller(), "Scenes": "标签写入 ES"}).Error(err)
	}
	return
}

// Delete
func (s *hello) Delete(keyword string) error {
	where := elastic.NewTermQuery("name.keyword", keyword)
	_, err := s.ESCli.DeleteByQuery().
		Index(s.Index).
		Type(s.Type).
		Query(where).
		Do(context.Background())
	return err
}
