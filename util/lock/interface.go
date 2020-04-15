package lock

type ILock interface {
	getLock(randomValue string)
	releaseLock(randomValue string)
}
