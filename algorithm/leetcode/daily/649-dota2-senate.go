package daily

// https://leetcode-cn.com/problems/dota2-senate/
func predictPartyVictory(senate string) string {
	var (
		r, d []int
	)
	for i, v := range senate {
		if v == 'R' {
			r = append(r, i)
		} else {
			d = append(d, i)
		}
	}
	for len(r) > 0 && len(d) > 0 {
		// 根据首元素大小判断谁先投票
		if r[0] < d[0] {
			// 投完票后加到数组尾部 等下一轮又可以投票
			r = append(r, r[0]+len(senate))
		} else {
			d = append(d, d[0]+len(senate))
		}
		// 每次投票都把双方首元素移除（不管哪方投票肯定会选择禁止对方阵营一名参议员的权利）
		// 其中敌方是永久移除（因为按照这个顺序投票，每次这个元素都会被前一个对方单位移除） 友方是暂时移除，下一轮还可以投票（所以前面需要加到队尾去）
		r = r[1:]
		d = d[1:]
	}
	if len(r) > 0 {
		return "Radiant"
	}
	return "Dire"
}
