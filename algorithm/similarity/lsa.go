package main

import (
	"fmt"
	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"gonum.org/v1/gonum/mat"
)

/*
1. 分析文档集合，建立Term-Document矩阵。
2. 对Term-Document矩阵进行奇异值分解。
3. 对SVD分解后的矩阵进行降维，也就是奇异值分解一节所提到的低阶近似。
4. 使用降维后的矩阵构建潜在语义空间，或重建Term-Document矩阵。
*/
func main() {
	// testCorpus := []string{
	// 	"常见 文本 相似度 计算 方法",
	// 	"文本 相似度 计算",
	// 	"常见 计算 方法",
	// 	"常见 方法",
	// }

	testCorpus := []string{
		"我 不喜欢 吃 苹果",
		"我 喜欢 看 电影 和 吃 苹果",
		"我 不喜欢 看 电影吃 苹果",
		"我 不喜欢 看 电影",
	}

	var stopWords = []string{"常见"}
	query := "我 喜欢 吃 苹果"
	vectoriser := nlp.NewCountVectoriser(stopWords...)
	transformer := nlp.NewTfidfTransformer()

	// set k (the number of dimensions following truncation) to 4
	// SVD 降维
	reducer := nlp.NewTruncatedSVD(4)
	lsiPipeline := nlp.NewPipeline(vectoriser, transformer, reducer)

	// Transform the corpus into an LSI fitting the model to the documents in the process
	// LSI
	lsi, err := lsiPipeline.FitTransform(testCorpus...)
	if err != nil {
		fmt.Printf("Failed to process documents because %v", err)
		return
	}

	// run the query through the same pipeline that was fitted to the corpus and
	// to project it into the same dimensional space
	queryVector, err := lsiPipeline.Transform(query)
	if err != nil {
		fmt.Printf("Failed to process documents because %v", err)
		return
	}

	// iterate over document feature vectors (columns) in the LSI matrix and compare
	// with the query vector for similarity.  Similarity is determined by the difference
	// between the angles of the vectors known as the cosine similarity
	highestSimilarity := -1.0
	var matched int
	_, docs := lsi.Dims()
	for i := 0; i < docs; i++ {
		// 余弦相似度
		similarity := pairwise.CosineSimilarity(queryVector.(mat.ColViewer).ColView(0), lsi.(mat.ColViewer).ColView(i))
		fmt.Printf("文本:'%s' 相似度:%v \n", testCorpus[i], similarity)
		if similarity > highestSimilarity {
			matched = i
			highestSimilarity = similarity
		}
	}
	fmt.Printf("Matched '%s'", testCorpus[matched])
	// Output: Matched 'The quick brown fox jumped over the lazy dog'
}
