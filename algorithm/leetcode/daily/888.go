package daily

// https://leetcode-cn.com/problems/fair-candy-swap/
/*
记爱丽丝的糖果棒的总大小为sumA，鲍勃的糖果棒的总大小为 sumB。设答案为{x,y}，即爱丽丝的大小为 x 的糖果棒与鲍勃的大小为 y 的糖果棒交换，则有如下等式：
sumA-x+y=sumB-y+x -->x=y+(sumA+sumB)/2
所以只需要遍历B，对于B中的每一个y都去寻找看A中是否存在对于的x即可
*/
func fairCandySwap(A []int, B []int) []int {
	var (
		sumA, sumB int
		x, y       int
		set        = make(map[int]struct{})
	)
	for _, v := range A {
		sumA += v
		set[v] = struct{}{}
	}
	for _, v := range B {
		sumB += v
	}
	delta := (sumA - sumB) / 2
	for _, v := range B {
		y = v
		x = y + delta
		_, ok := set[x]
		if ok {
			return []int{x, y}
		}
	}
	return nil
}

/*
git log --author="lixd" --pretty=tformat: --numstat |\
gawk '{ add += $1 ; subs += $2 ; loc += $1 - $2 } END { printf "增加的行数:%s 删除的行数:%s 总行数: %s\n",add,subs,loc }'
*/
