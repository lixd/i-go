package main

import (
	"fmt"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
)

var (
	seg    gse.Segmenter
	posSeg pos.Segmenter

	new = gse.New("zh,testdata/test_dict3.txt", "alpha")

	text = "你好世界, Hello world, Helloworld."
)

func main() {
	// cut()

	segCut()
}
func segCut() {
	// 加载默认字典
	seg.LoadDict()
	seg.LoadStop("D:/lillusory/projects/i-go/algorithm/similarity/stopwords/stopwords.txt")
	// 载入词典
	// seg.LoadDict("your gopath"+"/src/github.com/go-ego/gse/data/dict/dictionary.txt")

	// 分词文本
	tb := []byte("山达尔星联邦共和国的联邦政府啊")

	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中 ToString 函数的注释。
	// 搜索模式主要用于给搜索引擎提供尽可能多的关键字
	fmt.Println("输出分词结果, 类型为字符串, 使用搜索模式: ", seg.Slice(string(tb), true))
	tokens := seg.Slice(string(tb), false)
	fmt.Println("输出分词结果, 类型为 slice: ", tokens)
	stop := seg.Trim(tokens)
	fmt.Println("过滤停用词之后: ", stop)
	segments := seg.Segment(tb)
	// 处理分词结果
	fmt.Println(gse.ToString(segments))

	segments1 := seg.Segment([]byte(text))
	fmt.Println(gse.ToString(segments1, true))
}

func cut() {
	hmm := new.Cut(text, true)
	fmt.Println("cut use hmm: ", hmm)

	hmm = new.CutSearch(text, true)
	fmt.Println("cut search use hmm: ", hmm)

	hmm = new.CutAll(text)
	fmt.Println("cut all: ", hmm)
}

func posAndTrim(cut []string) {
	cut = seg.Trim(cut)
	fmt.Println("cut all: ", cut)

	posSeg.WithGse(seg)
	po := posSeg.Cut(text, true)
	fmt.Println("pos: ", po)

	po = posSeg.TrimWithPos(po, "zg")
	fmt.Println("trim pos: ", po)
}

func cutPos() {
	fmt.Println(seg.String(text, true))
	fmt.Println(seg.Slice(text, true))

	po := seg.Pos(text, true)
	fmt.Println("pos: ", po)
	po = seg.TrimPos(po)
	fmt.Println("trim pos: ", po)
}
