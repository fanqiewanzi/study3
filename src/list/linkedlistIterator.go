package list

//LinkedList迭代器
//array为原数组
//cursor为一个移动的下标
//end为最近已经输出元素的下标
type LinkedlistIterator struct {
	linklist *LinkList
	cursor   *LinkList
	end      *LinkList
}

//判断是否存在下一个元素
func (linkedlistIterator *LinkedlistIterator) HashNext() bool {
	//如果当前下标等于它的大小说明没有下一个元素了
	return linkedlistIterator.cursor != nil
}

//返回下一个元素
func (linkedlistIterator *LinkedlistIterator) Next() (interface{}, error) {
	//首先获取当前下标的位置
	i := linkedlistIterator.cursor
	//if i >= arrayIterator.array.size+1 {
	//	return nil, errors.New("没有这样的索引")
	//}
	//下标位置往后移
	linkedlistIterator.cursor = linkedlistIterator.cursor.next
	linkedlistIterator.end = i
	return linkedlistIterator.linklist.data, nil
}

func (linkedList *LinkList) Iterator() Iterator {
	it := new(LinkedlistIterator)
	it.linklist = linkedList
	it.cursor = linkedList
	it.end = linkedList
	return it
}
