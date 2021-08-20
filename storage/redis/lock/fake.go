package main

const FirstLock = "firstlock"

// 伪代码
//func main() {
//	if getLock() {
//		defer releaseKey()
//		doSomething()
//	}
//}
//
//func getLock() bool {
//	if existKey(FirstLock) {
//		return false
//	} else {
//		setKey(FirstLock)
//		return true
//	}
//}
//func releaseKey() {
//	if getKey(FirstLock) == randomValue {
//		deleteKey(FirstLock)
//	}
//}
