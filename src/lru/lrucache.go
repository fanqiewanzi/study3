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
//map与链表中的元素相对应，提升了查找速度
//capacity是用来表示缓存的容量大小
//first和last分别是指向链表的头指针和尾指针
type LRUCache struct {
	reflect     map[int]*Node
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

//在链表首节点处添加结点,最后一个节点始终为空
func (this *LRUCache) add(node *Node) {
	node.pre = nil
	node.next = this.first
	this.first.pre = node
	this.first = node
}

//移除节点
func (this *LRUCache) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

//将节点移至首节点
func (this *LRUCache) moveToHead(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
	node.pre = nil
	node.next = this.first
	this.first.pre = node
	this.first = node
}

//如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1
func (this *LRUCache) Get(key int) int {
	if node, ok := this.reflect[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

//如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。
//当缓存容量达到上限，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空时间。
func (this *LRUCache) Put(key int, value int) {
	//检查在Map中是否存在要插入的元素，存在就更新值并且在链表总将节点移至首节点
	if node, ok := this.reflect[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		//判断为新节点，判断缓存是否已满，如果已满则删除链表最后一个节点
		if len(this.reflect) >= this.capacity {
			this.remove(this.last.pre)
			delete(this.reflect, this.last.pre.key)
		}
		//添加新节点
		p := &Node{key, value, nil, nil}
		this.reflect[key] = p
		this.add(p)
	}
}
