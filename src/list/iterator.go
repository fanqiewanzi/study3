package list

//迭代器接口
type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
	//HasPrevious() bool
	//Previous() (interface{}, error)

}
