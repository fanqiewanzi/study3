package list

import (
	"errors"
	"fmt"
)

/*
从链表左边取出一个节点:首节点指针右移一位，首节点为nil,长度减一；
给链表末尾新增一个节点:以传入的数据new一个节点，
  1，链表为空时，该node作为链表的首节点、尾节点；
  2，不为空时，该节点赋值到旧的末节点的next，该节点的前置指针为旧的末节点，该节点赋值为链表的末节点，该节点的next为nil，链表长度+1；
反转链表：思路较多，此处随便写写，详见https://lan6193.blog.csdn.net/article/details/104660223
*/
type Node struct {
	pre  *Node
	next *Node
	data interface{}
}

type DoubleList struct {
	first *Node
	last  *Node
	size  int
	List
}

// 给链表末尾新增一个节点
func (list *DoubleList) Add(obj ...interface{}) error {
	for _, elem := range obj {
		p := new(Node)
		p.data = elem
		list.last.next = p
		p.pre = list.last
		list.last = p
		list.size++
	}
	return nil
}

//向指定位置加入元素
func (list *DoubleList) Insert(location int, obj interface{}) error {
	if location <= 0 || location > list.size {
		return errors.New("位置超出")
	}
	//找相应位置元素
	p := list.first
	count := 0
	for p.next != nil && count != location {
		p = p.next
		count++
	}
	//在相应位置元素进行赋值
	q := new(Node)
	q.next = p.next
	q.pre = p
	p.next = q
	q.data = obj
	list.size++
	return nil
}

//向指定位置修改元素
func (list *DoubleList) Set(location int, obj interface{}) error {
	if location <= 0 || location > list.size {
		return errors.New("位置超出")
	}
	//找相应位置元素
	p := list.first
	count := 0
	for p.next != nil && count != location {
		p = p.next
		count++
	}
	//赋值
	p.data = obj
	return errors.New("no error")
}

//是否存在某元素
func (list *DoubleList) Contain(obj interface{}) bool {
	//找相应位置元素
	p := list.first
	for p.next != nil && p.data != obj {
		p = p.next
	}
	//若已经走到尽头该值不等于要找的值则返回false否则返回false
	if p.next == nil || p.data != obj {
		return false
	}
	return true
}

//是否为空
func (list *DoubleList) IsEmpty() bool {
	if list.last == list.first {
		return true
	}
	return false
}

//查看某一位置上的元素
func (list *DoubleList) Get(location int) (interface{}, error) {
	if location <= 0 || location > list.size {
		return nil, errors.New("位置超出")
	}
	//找相应位置元素
	p := list.first
	count := 0
	for p.next != nil && count != location {
		p = p.next
		count++
	}
	return p.data, nil
}

//判断是否相等
func (list *DoubleList) Equals(list1 List) bool {
	//对每个元素一一进行比较，如果循环完后两个指针都为空说明是相等的
	p := list.first.next
	q := list1
	var value interface{}
	i := 1
	for value, _ = q.Get(i); p != nil && q != nil && value == p.data; {
		i++
		value, _ = q.Get(i)
		p = p.next
	}
	if p == nil && value == nil {
		return true
	}
	return false
}

//转换为Slice类型
func (list *DoubleList) ToSlice() []interface{} {
	//定义一个切片进入循环将值一一赋给切片你即可
	p := list.first.next
	sli := make([]interface{}, list.size)
	i := 0
	//如果为空链表则直接返回一个空切片
	if list.size == 0 {
		return sli
	}
	for p != nil {
		sli[i] = p.data
		p = p.next
		i++
	}
	return sli
}

//输出当前list的长度
func (list *DoubleList) Size() int {
	return list.size
}

// 打印(遍历)链表
func (list *DoubleList) Print() {
	fmt.Printf("\n链表输出为:")
	p := list.first.next
	for p != nil {
		fmt.Print(p.data)
		fmt.Print(" ")
		p = p.next
	}

}

// 创建一个空的双链表
func NewDoubleList() (list *DoubleList) {
	p := &DoubleList{nil, nil, 0, nil}
	p.last = new(Node)
	p.first = p.last
	p.List = p
	return p
}

func Doubletest() {
	list1 := NewDoubleList()
	list2 := NewDoubleList()
	list3 := NewDoubleList()
	list3.Add(1, 2, 3, 4)
	list1.Add(1, 2, 3, 4)
	list2.Add(1, 2, 3, 4)
	list1.Print()
	fmt.Println(list1.Equals(list2))
	list1.Insert(2, 5)
	list1.Print()
	fmt.Println(list1.Contain(5))
	list1.Set(3, 1)
	list1.Print()
	fmt.Println(list1.IsEmpty())
	fmt.Println(list1.Get(4))
	fmt.Println(list1.Equals(list2))
	fmt.Println(list1.ToSlice())
	fmt.Println(list1.Size())
}
