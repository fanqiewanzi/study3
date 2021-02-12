package main

import (
	"channel"
	"factory"
	"list"
	"lru"
)

func main() {
	channel.PrintTest()
	list.IteratorTest()
	factory.FactoryTest()
	lru.IrucacheTest()
}
