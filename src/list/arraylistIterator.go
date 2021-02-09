package list

import "errors"

//ArrayList迭代器
//array为原数组
//cursor为一个移动的下标
//end为最近已经输出元素的下标
type ArraylistIterator struct {
	array  *Array
	cursor int
	end    int
}

//判断是否存在下一个元素
func (arrayIterator *ArraylistIterator) HashNext() bool {
	//如果当前下标等于它的大小说明没有下一个元素了
	return arrayIterator.cursor != arrayIterator.array.size+1
}

//返回下一个元素
func (arrayIterator *ArraylistIterator) Next() (interface{}, error) {
	//首先获取当前下标的位置
	i := arrayIterator.cursor
	if i >= arrayIterator.array.size+1 {
		return nil, errors.New("没有这样的索引")
	}
	//下标位置往后移
	arrayIterator.cursor = arrayIterator.cursor + 1
	arrayIterator.end = i
	return arrayIterator.array.data[arrayIterator.end], nil
}

func (array *Array) ArrayIterator() Iterator {
	it := new(ArraylistIterator)
	it.array = array
	it.cursor = 0
	it.end = -1
	return it
}
