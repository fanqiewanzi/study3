package list

//迭代器接口
type Iterator interface {
	HashNext() bool
	Next() (interface{}, error)
}

//
type Container interface {
	Iterator() Iterator
}
