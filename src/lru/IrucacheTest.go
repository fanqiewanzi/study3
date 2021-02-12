package lru

import "fmt"

func IrucacheTest() {
	fmt.Println()
	ir := Constructor(2)
	ir.Put(1, 1)
	ir.Put(2, 2)
	ir.Put(3, 3)
	ir.Put(4, 4)
	ir.Put(5, 5)
	ir.Put(6, 6)
	fmt.Println(ir.Get(5))
	fmt.Println(ir.last.pre.value)
	fmt.Println(ir.Get(4))
	fmt.Println(len(ir.reflect))

}
