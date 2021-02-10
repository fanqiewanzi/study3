package list

import "fmt"

func IteratorTest() {
	list := NewArray(100)
	for i := 0; i < 2; i++ {
		list.Add(1)
		list.Add(3)
	}
	it := list.Iterator()

	for it.HasNext() {
		i, _ := it.Next()
		fmt.Print(i)
	}

	list1 := NewDoubleList()

	for i := 0; i < 2; i++ {
		list1.Add(1)
		list1.Add(2)
	}
	fmt.Print(list1.Equals(list1))
	fmt.Print(list.Equals(list))
	it1 := list1.Iterator()

	for it1.HasNext() {
		i, _ := it1.Next()
		fmt.Print(i)
	}
	for it1.HasNext() {
		i, _ := it1.Next()
		fmt.Print(i)
	}
}
