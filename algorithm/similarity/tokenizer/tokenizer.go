package tokenizer

import (
	"github.com/go-ego/gse"
	"path"
	"runtime"
	"sync"
)

type tokenizer struct {
	Seg *gse.Segmenter
}

var (
	Tokenizer *tokenizer
	once      sync.Once
)

func NewParticiple() *tokenizer {
	once.Do(func() {
		Tokenizer = &tokenizer{Seg: &gse.Segmenter{}}
		// 初始化并加载词典和停用词
		dictDir := path.Join(path.Dir(getCurrentFilePath()), "dictionary.txt")
		err := Tokenizer.Seg.LoadDict(dictDir)
		if err != nil {
			panic(err)
		}
		stopDir := path.Join(path.Dir(getCurrentFilePath()), "stopwords.txt")
		err = Tokenizer.Seg.LoadStop(stopDir)
		if err != nil {
			panic(err)
		}
	})
	return Tokenizer
}

func getCurrentFilePath() string {
	_, filePath, _, _ := runtime.Caller(1)
	return filePath
}

// Cut 分词
func (t *tokenizer) Cut(target string) []string {
	tokens := t.Seg.Slice(target, true)
	trimStop := t.Seg.Trim(tokens)
	return trimStop
}
