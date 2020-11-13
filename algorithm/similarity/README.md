
Semantic analysis of webpages with machine learning in Go
http://www.jamesbowman.me/post/semantic-ana lysis-of-webpages-with-machine-learning-in-go/

文档相似度对比算法
1.基于词向量
余弦相似度
曼哈顿距离
欧几里得距离
明式距离（是前两种距离测度的推广），在极限情况下的距离是切比雪夫距离
2.基于字符的
编辑距离
simhash
共有字符数（有点类似 onehot 编码 ，直接统计两个文本的共有字符数，最 naive 的相似度算法了）
3.基于概率统计的
杰卡德相似系数
4.基于词嵌入模型的
word2vec/doc2vec

博客 https://www.cnblogs.com/daniel-D/p/3244718.html
闵可夫斯基距离
欧几里得距离
曼哈顿距离
切比雪夫距离
马氏距离
余弦相似度
皮尔逊相关系数
汉明距离
杰卡德相似系数
编辑距离
DTW 距离
KL 散度


（1）使用TF-IDF算法，找出两篇文章的关键词；
（2）每篇文章各取出若干个关键词（比如20个），合并成一个集合，计算每篇文章对于这个集合中的词的词频（为了避免文章长度的差异，可以使用相对词频）；
（3）生成两篇文章各自的词频向量；
（4）计算两个向量的余弦相似度，值越大就表示越相似。


Latent Semantic Analysis(LSA) 潜在语义分析
https://en.wikipedia.org/wiki/Latent_semantic_analysis
Eigenvalues and Eigenvectors 特征值与特征向量
Singular Value Decomposition(SVD) 奇异值分解
Latent Semantic Indexing 潜在语义索引
LDI
n维矩阵向量
文档*文档
文档*terms
terms*terms

svd https://cs.fit.edu/~dmitra/SciComp/Resources/singular-value-decomposition-fast-track-tutorial.pdf
svb0 https://arxiv.org/pdf/1305.2452.pdf
lsa http://webhome.cs.uvic.ca/~thomo/svd.pdf
lda  隐含狄利克雷分布（Latent Dirichlet Allocation，简称LDA），一种处理文档的主题模型