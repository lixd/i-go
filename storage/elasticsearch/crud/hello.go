package main

import (
	"context"
	"reflect"
	"sync"

	"i-go/utils"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type IHello interface {
	MatchPrefix(keyword string, count int) ([]string, error)
	Upsert(id, keyword string) error
	Delete(id string) error
}

type hello struct {
	Index string `json:"index"`
	Cli   *elastic.Client
}

type HelloItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var (
	once     sync.Once
	HelloCli *hello
)

func NewHelloCli(index, _type string, cli *elastic.Client) IHello {
	once.Do(func() {
		HelloCli = &hello{
			Index: index,
			Cli:   cli,
		}
	})
	return HelloCli
}

// MatchPrefix 前缀匹配分类
func (s *hello) MatchPrefix(keyword string, count int) ([]string, error) {
	query := elastic.NewMatchPhrasePrefixQuery("name", keyword)
	res, err := s.Cli.Search().
		Index(s.Index).
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

// Upsert 根据 id 更新文档
func (s *hello) Upsert(id, keyword string) error {
	body := HelloItem{Name: keyword}
	_, err := s.Cli.Update().
		Index(s.Index).
		Id(id).
		Doc(&body).
		// ES 检测到 文档无变化则不执行任何操作
		DetectNoop(true).
		// 文档不存在则创建
		DocAsUpsert(true).
		Do(context.Background())
	if err != nil {
		logrus.WithFields(logrus.Fields{"CallerName": utils.Caller(), "Scenes": "ES:Index"}).Error(err)
	}
	return err
}

// Delete 根据 id 删除文档
func (s *hello) Delete(id string) error {
	_, err := s.Cli.Delete().
		Index(s.Index).
		Id(id).
		Do(context.Background())
	return err
}
