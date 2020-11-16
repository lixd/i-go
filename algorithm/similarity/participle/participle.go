package participle

import (
	"github.com/yanyiwu/gojieba"
	"strings"
	"sync"
)

type participle struct {
	JieBa *gojieba.Jieba
}

var (
	Participle *participle
	once       sync.Once
)

func NewParticiple() *participle {
	once.Do(func() {
		Participle = &participle{JieBa: gojieba.NewJieba()}
	})
	return Participle
}

// Release 释放内存
func (p *participle) Release() {
	p.JieBa.Free()
}

// Cut 分词
func (p *participle) Cut(target ...string) []string {
	var (
		cuts = make([][]string, 0, len(target))
		list = make([]string, 0, len(target))
	)
	// 1.分词
	for _, v := range target {
		cut := p.JieBa.CutForSearch(v, true)
		cuts = append(cuts, cut)
	}
	// 2.拼接
	for _, v := range cuts {
		list = append(list, strings.Join(v, " "))
	}
	return list
}
