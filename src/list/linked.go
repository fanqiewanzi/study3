package list

import (
	"errors"
	"fmt"
)

const defaultCount = 0

//单向链表数据结构
type LinkList struct {
	data interface{}
	next *LinkList
	size int
	List
}

func NewLinkList() *LinkList {
	return &LinkList{nil, nil, defaultCount, nil}
}

func (linkList *LinkList) Add(obj ...interface{}) error {
	//找最后一个元素
	p := linkList
	for p.next != nil {
		p = p.next
	}
	//在最后一个元素进行赋值，因为不知道有多少个值所以用循环
	for _, elem := range obj {
		q := new(LinkList)
		p.next = q
		q.data = elem
		p = q
		linkList.size++
	}
	return errors.New("no error")
}

func (linkList *LinkList) Insert(location int, obj interface{}) error {
	if location <= 0 || location > linkList.size {
		return errors.New("位置超出")
	}
	//找相应位置元素
	p := linkList
	count := 0
	for p.next != nil && count != location {
		p = p.next
		count++
	}
	//在相应位置元素进行赋值
	q := new(LinkList)
	q.next = p.next
	p.next = q
	q.data = obj
	linkList.size++
	return errors.New("no error")
}

//修改某元素
func (linkList *LinkList) Set(location int, obj interface{}) error {
	if location <= 0 || location > linkList.size {
		return errors.New("位置超出")
	}
	//找相应位置元素
	p := linkList
	count := 0
	for p.next != nil && count != location {
		p = p.next
		count++
	}
	//赋值
	p.data = obj
	return errors.New("no error")
}

func (linkList *LinkList) Contain(obj interface{}) bool {
	//找相应位置元素
	p := linkList
	for p.next != nil && p.data != obj {
		p = p.next
	}
	//若已经走到尽头该值不等于要找的值则返回false否则返回false
	if p.next == nil || p.data != obj {
		return false
	}
	return true
}

//单链表第一个位置为空，它next指针为空说明此链表为空
func (linkList *LinkList) IsEmpty() bool {
	if linkList.next == nil {
		return true
	} else {
		return false
	}
}

func (linkList *LinkList) Get(location int) (interface{}, error) {
	if location <= 0 || location > linkList.size {
		return nil, errors.New("位置超出")
	}
	//找相应位置元素
	p := linkList
	count := 0
	for p.next != nil && count != location {
		p = p.next
		count++
	}
	return p.data, errors.New("no error")
}

//判断是否相等
func (linkList *LinkList) Equals(linkList1 *LinkList) bool {
	//对每个元素一一进行比较，如果循环完后两个指针都为空说明是相等的
	p := linkList
	q := linkList1
	for p != nil && q != nil && p.data == q.data {
		p = p.next
		q = q.next
	}
	if p == nil && q == nil {
		return true
	}
	return false
}

//转换为Slice类型
func (linkList *LinkList) ToSlice() []interface{} {
	//定义一个切片进入循环将值一一赋给切片你即可
	p := linkList.next
	sli := make([]interface{}, linkList.size)
	i := 0
	//如果为空链表则直接返回一个空切片
	if linkList.size == 0 {
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
func (linkList *LinkList) Size() int {
	return linkList.size
}

//打印链表，头指针不打印
func (linkList *LinkList) Print() {
	fmt.Printf("\n链表输出为:")
	p := linkList.next
	for p != nil {
		fmt.Print(p.data)
		fmt.Print(" ")
		p = p.next
	}
}
