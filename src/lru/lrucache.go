package lru

//双向链表结构
//key指的是map中用来查找节点的值
//value用来保存每个节点下的值
//pre和next用来指向上一个或者下一个节点的指针
type Node struct {
	key, value int
	pre, next  *Node
}

//LRUCache的数据结构
//map用来筛选进入缓存，在缓存满了之后map中的数据个数比链表中的数据个数要多
//capacity是用来表示缓存的容量大小
//first和last分别是指向链表的头指针和尾指针
type LRUCache struct {
	m           map[int]*Node
	capacity    int
	first, last *Node
}

//构造函数
func Constructor(capacity int) LRUCache {
	first := &Node{-1, -1, nil, nil}
	last := &Node{-1, -1, nil, nil}
	first.next = last
	last.pre = first
	return LRUCache{make(map[int]*Node), capacity, first, last}
}
func (this *LRUCache) Add(node *Node) {
	node.pre = this.first.pre
	node.next = this.first.next
	this.first.pre = node
	this.first = node
}
func (this *LRUCache) Remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

//将节点移至首节点
func (this *LRUCache) Move(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
	this.first.pre = node
	node.next = this.first
	node.pre = nil
}

//如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1
func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		return node.value
	}
	return -1
}

//如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
//当缓存容量达到上限，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空时间。
func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		node.value++
	} else {
		if len(this.m) > this.capacity {
			this.Remove(this.last)
		}
	}
	p := &Node{key, value, nil, nil}
	this.Add(p)
	this.m[key] = p
}
