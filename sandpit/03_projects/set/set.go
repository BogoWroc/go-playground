package set

type Set interface {
	Add(e ...interface{})
	Contains(e ...interface{}) bool
	Length() int
	Clean()
	Iterator() <-chan interface{}
	ToSlice() []interface{}
}
