package main

import (
	"i-go/algorithm/similarity/participle"
	"testing"
)

func BenchmarkLsa(b *testing.B) {

	// 	1.提供目标文档和文档库
	query, corpus := loadCorpus()
	// 	2.对文档进行分词
	p := participle.NewParticiple()
	// defer p.Release()

	queryTerm := p.Cut(query)
	corpusTerm := p.Cut(corpus...)
	// 	3.LSA分析
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 	3.LSA分析
		_, _ = lsa(queryTerm[0], corpusTerm)
	}
}
