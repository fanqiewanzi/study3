package list

import "fmt"

func IteratorTest() {
	list := NewArray(100)

	for i := 0; i < 10; i++ {
		list.Add(1)
		list.Add(2)
	}

	//fmt.Println(list)
	//list.Set(0, 100)
	it := list.Iterator()

	for it.HashNext() {
		i, _ := it.Next()
		fmt.Print(i)
	}

	list1 := NewLinkList()

	for i := 0; i < 10; i++ {
		list1.Add(1)
		list1.Add(2)
	}

	//fmt.Println(list)
	//list.Set(0, 100)
	it1 := list.Iterator()

	for it1.HashNext() {
		i, _ := it1.Next()
		fmt.Print(i)
	}

}
