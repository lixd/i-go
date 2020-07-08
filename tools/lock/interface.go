package lock

type ILock interface {
	Lock(key int) bool
	UnLock(key int) bool
}
