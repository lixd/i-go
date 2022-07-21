package t

import (
	"fmt"
	"i-go/algorithm/similarity/tokenizer"
	"math"
	"runtime"
)

func Similar() {
	// 1.加载查询语句和语料库
	query, corpus := loadCorpus()
	// 2.分词
	token := tokenizer.NewParticiple()
	queryWords := token.Cut(query)
	fmt.Println("queryWords:", queryWords)
	// 3.计算余弦相似度
	for _, v := range corpus {
		corpusWords := token.Cut(v)
		similar := CosineSimilar(queryWords, corpusWords)
		fmt.Println("相似度:", similar, "corpusWords:", corpusWords)
	}
}

func getCurrentFilePath() string {
	_, filePath, _, _ := runtime.Caller(1)
	return filePath
}

/*
1.找到所有词
2.找出每个文档的单词的词频
3.计算余弦相似度
*/
func CosineSimilar(srcWords, dstWords []string) float64 {
	// get all words
	allWordsMap := make(map[string]int, 0)
	for _, word := range srcWords {
		if _, found := allWordsMap[word]; !found {
			allWordsMap[word] = 1
		} else {
			allWordsMap[word] += 1
		}
	}
	for _, word := range dstWords {
		if _, found := allWordsMap[word]; !found {
			allWordsMap[word] = 1
		} else {
			allWordsMap[word] += 1
		}
	}

	// stable the sort
	allWordsSlice := make([]string, 0)
	for word := range allWordsMap {
		allWordsSlice = append(allWordsSlice, word)
	}
	// assemble vector
	srcVector := make([]int, len(allWordsSlice))
	dstVector := make([]int, len(allWordsSlice))

	for _, word := range srcWords {
		if index := indexOfArray(allWordsSlice, word); index != -1 {
			srcVector[index] += 1
		}
	}
	for _, word := range dstWords {
		if index := indexOfArray(allWordsSlice, word); index != -1 {
			dstVector[index] += 1
		}
	}
	// calc cos
	numerator := float64(0)
	srcSq := 0
	dstSq := 0
	for i, srcCount := range srcVector {
		dstCount := dstVector[i]
		numerator += float64(srcCount * dstCount)
		srcSq += srcCount * srcCount
		dstSq += dstCount * dstCount
	}
	denominator := math.Sqrt(float64(srcSq * dstSq))
	similarity := numerator / denominator
	// 保留4位小数
	similarity = math.Floor(similarity*10000) / 10000
	return similarity
}

func indexOfArray(data []string, target string) int {
	for i, v := range data {
		if v == target {
			return i
		}
	}
	return -1
}

// loadCroups 加载目标文档和文档库
func loadCorpus() (query string, corpus []string) {
	query = query2
	corpus = corpus2
	return
}

var (
	query1  = "你好，我想问一下，我想离婚，他不想离，孩子他说不要，是六个月就自动生效离婚吗？"
	corpus1 = []string{
		"无偿居间介绍买卖毒品的行为应如何定性",
		"吸毒男动态持有大量毒品的行为该如何认定",
		"如何区分是非法种植毒品原植物罪还是非法制造毒品罪",
		"为毒贩贩卖毒品提供帮助构成贩卖毒品罪",
		"将自己吸食的毒品原价转让给朋友吸食的行为该如何认定",
		"为获报酬帮人购买毒品的行为该如何认定",
		"毒贩出狱后再次够买毒品途中被抓的行为认定",
		"虚夸毒品功效劝人吸食毒品的行为该如何认定",
		"妻子下落不明丈夫又与他人登记结婚是否为无效婚姻",
		"一方未签字办理的结婚登记是否有效",
		"夫妻双方1990年按农村习俗举办婚礼没有结婚证 一方可否起诉离婚",
		"结婚前对方父母出资购买的住房写我们二人的名字有效吗",
		"身份证被别人冒用无法登记结婚怎么办？",
		"同居后又与他人登记结婚是否构成重婚罪",
		"未办登记只举办结婚仪式可起诉离婚吗",
		"同居多年未办理结婚登记，是否可以向法院起诉要求离婚",
	}
	query2 = `四川成都，武术爱好者张先生仅用一部手机和同伴拍出了节奏精彩的武术视频短片，视频中的武术动作设计看起来很专业，网友看过之后也直呼精彩，
称其不亚于大片。据主创者张先生介绍，他们从小就练习武术，拍视频纯粹是属于个人爱好，就是想把自己练习多年的功夫展示出来，作为一个生活的记录，
也没有什么专业的设备，没有人投资，都是他们自己拍着玩儿，自娱自乐，因拍摄的多是运动镜头，手机已经摔坏了好几部。`
	corpus2 = []string{
		`据今日俄罗斯电视台（RT）援引当地媒体报道，以色列情报和特殊使命局（又称摩萨德）已获得中国新冠肺炎疫苗，并带回开始进行“研究”。
以色列媒体12频道（Channel 12）周一报道称，最近几周，摩萨德将中国疫苗带到了以色列。目前还不清楚该情报机构究竟是如何获得这种疫苗的。`,
		`多个政府消息人士向当地媒体“间接”证实了这一报道。据说，采购该疫苗的目的是研究其配方，并进一步探索疫苗接种方案。除此之外，以色列卫生部一位高级官员证实，
以色列正在寻求从其他国家购买新冠疫苗。据报道，以色列已经研制出了自己的新冠病毒疫苗，但其测试过程仍处于早期阶段。`,
		`上周，位于耐斯茨奥纳的以色列生物研究所（IIBR）宣布了其名为“Brilife”的新冠疫苗。该疫苗已经获得了所有必要的批准，预计将于下周开始进行第一阶段临床试验。
首先，这种疫苗将在大约100名志愿者身上进行试验，如果一切顺利，明年春天将有另外1000人参加第二阶段的试验`,
		`最近几周，以色列新冠病毒感染病例急剧增加，目前病例总数达到31万。虽然这个数字与受影响最严重的国家相比相对较小，但对于这个890万人口的国家来说，
这个数字相当可观，约占总人口的3.4%。美国约翰·霍普金斯大学的最新统计数据显示，截至10月26日，以色列新冠确诊病例累计超31万例，2397人死亡。`,
		`新疆维吾尔自治区卫生健康委最新通报，10月27日0时至24时，新疆维吾尔自治区（含新疆生产建设兵团）报告新增新冠肺炎确诊病例22例（均为无症状感染者转确诊），
新增无症状感染者19例，均为喀什地区疏附县报告。截至10月27日24时，新疆（含兵团）现有确诊病例22例，无症状感染者161例，均为喀什地区疏附县报告。`,
		`一经上映后就大受欢迎，观看量飙升，投资人表示意料之外，`,
	}
)
