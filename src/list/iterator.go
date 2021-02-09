package list

//迭代器接口
type Iterator interface {
	HashNext() bool
	Next() (interface{}, error)
}

//构造函数
type Container interface {
	ArrayIterator() Iterator
	LinkIterator() Iterator
}
