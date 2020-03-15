package q460

type Node struct {
	key  int
	val  int
	freq int
	next *Node
	prev *Node
}

func (n *Node) insert(node *Node) {
	node.prev = n
	node.next = n.next
	n.next.prev = node
	n.next = node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	list := &LinkedList{
		head: &Node{},
		tail: &Node{},
	}
	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		kvMap:   make(map[int]*Node),
		freqMap: make([]*LinkedList, capacity+1),
		cap:     capacity,
		minFreq: 1,
	}
}

type LFUCache struct {
	kvMap   map[int]*Node
	freqMap []*LinkedList
	minFreq int
	size    int
	cap     int
}

func (c *LFUCache) unlink(node *Node) int {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if node.freq > 0 && node.prev == c.freqMap[node.freq].head && node.next == c.freqMap[node.freq].tail {
		c.freqMap[node.freq] = nil
	}
	node.prev = nil
	node.next = nil
	return node.key
}

func (c *LFUCache) updateFreq(node *Node) {
	c.unlink(node)

	node.freq++
	if node.freq >= len(c.freqMap) {
		c.freqMap = append(c.freqMap, make([]*LinkedList, node.freq*2)...)
	}
	if c.freqMap[node.freq] == nil {
		c.freqMap[node.freq] = NewLinkedList()
	}
	c.freqMap[node.freq].tail.prev.insert(node)

	if node.freq < c.minFreq || c.freqMap[c.minFreq] == nil {
		c.minFreq = node.freq
	}
}

func (c *LFUCache) Get(key int) int {
	node, ok := c.kvMap[key]
	if !ok {
		return -1
	}

	c.updateFreq(node)
	return node.val
}

func (c *LFUCache) Put(key int, value int) {
	if c.cap < 1 {
		return
	}
	node, ok := c.kvMap[key]
	if !ok {
		node = &Node{
			key: key,
			val: value,
		}
		c.kvMap[key] = node
		c.size++
	} else {
		node.val = value
	}

	if c.size > c.cap {
		delete(c.kvMap, c.unlink(c.freqMap[c.minFreq].head.next))
		c.size--
	}

	c.updateFreq(node)
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
