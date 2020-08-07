package recursion

var result []interface{}

// divideConquer 分治 模板代码
func divideConquer(problem, params interface{}) interface{} {
	// 1.首先写递归终结条件 recursion terminator
	if problem == nil {
		// process result
		return nil
	}
	// 2.prepare data
	data := prepareData(problem)
	subProblems := spiltProblem(problem, data)
	// 3.conquer subProblems
	subResult1 := divideConquer(subProblems[0], params)
	subResult2 := divideConquer(subProblems[1], params)
	subResult3 := divideConquer(subProblems[2], params)
	// 4.process and generate the final result
	result = processResult(subResult1, subResult2, subResult3)
	// 5.restore current status
	return result
}

func prepareData(problem interface{}) interface{} {
	return nil
}
func spiltProblem(problem, data interface{}) []interface{} {
	return nil
}
func processResult(result ...interface{}) []interface{} {
	return nil
}
