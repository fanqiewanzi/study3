package lru

type LRUCache struct {
}

//构造函数
func Constructor(capacity int) LRUCache {
}

//如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1
func (this *LRUCache) Get(key int) int {
}

//如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
//当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
func (this *LRUCache) Put(key int, value int) {
}
