package factory

import (
	"fmt"
	"list"
)

type Factory interface {
	ListFactory(string) list.List
}
type ListFactory struct {
	Factory
}

func (listfactory *ListFactory) GetListType(str string) list.List {
	if str == "ArrayList" {
		return list.NewArrayWithoutNoCap()
	} else if str == "LinkList" {
		return list.NewDoubleList()
	}
	return nil
}

func FactoryTest() {
	factory := new(ListFactory)
	list1 := factory.GetListType("ArrayList")
	list1.Add(1, 2, 3, 4, 5)
	it := list1.Iterator()
	for it.HasNext() {
		i, _ := it.Next()
		fmt.Print(i)
	}
}
