package main

import (
	"fmt"
	"github.com/james-bowman/nlp"
	"github.com/james-bowman/nlp/measures/pairwise"
	"gonum.org/v1/gonum/mat"
)

func main() {
	testCorpus := []string{"The quick brown fox jumped over the lazy dog",
		"hey diddle diddle, the cat and the fiddle",
		"the fast cunning brown fox liked the slow canine dog ",
		"the little dog laughed to see such fun",
		"and the dish ran away with the spoon"}
	query := "the cunning creature ran around the canine"
	vectoriser := nlp.NewCountVectoriser("the")
	transformer := nlp.NewTfidfTransformer()            // set k (the number of dimensions following truncation) to 2
	reducer := nlp.NewTruncatedSVD(2)                   // Fit and Transform the corpus into a term document matrix fitting the // model to the documents in the process
	matrix, _ := vectoriser.FitTransform(testCorpus...) // transform the query into the same dimensional space - any terms in
	// the query not in the original training data the model was fitted to
	// will be ignored
	queryMat, _ := vectoriser.Transform(query)
	calcCosine(queryMat, matrix, testCorpus, "Raw TF")

	tfidfmat, _ := transformer.FitTransform(matrix)
	tfidfquery, _ := transformer.Transform(queryMat)
	calcCosine(tfidfquery, tfidfmat, testCorpus, "TF-IDF")

	lsi, _ := reducer.FitTransform(tfidfmat)
	queryVector, _ := reducer.Transform(tfidfquery)
	calcCosine(queryVector, lsi, testCorpus, "LSA")
}
func calcCosine(query mat.Matrix, tdmat mat.Matrix, corpus []string, name string) {
	// iterate over document feature vectors (columns) in the LSI and
	// compare with the query vector for similarity. Similarity is determined
	// by the difference between the angles of the vectors known as the cosine
	// similarity
	_, docs := tdmat.Dims()
	fmt.Printf("Comparing based on %s\n", name)
	for i := 0; i < docs; i++ {
		// similarity := pairwise.CosineSimilarity(queryVector.(mat.ColViewer).ColView(0), lsi.(mat.ColViewer).ColView(i))
		queryVec := query.(mat.ColViewer).ColView(0)
		docVec := tdmat.(mat.ColViewer).ColView(i)
		similarity := pairwise.CosineSimilarity(queryVec, docVec)
		fmt.Printf("Comparing '%s' = %f\n", corpus[i], similarity)
	}
}
