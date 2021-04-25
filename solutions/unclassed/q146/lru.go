package q146

type (
	LRUCache struct {
		kvs  map[int]*Node
		head *Node
		tail *Node
		cap  int
		size int
	}
	Node struct {
		key  int
		val  int
		next *Node
		prev *Node
	}
)

func Constructor(capacity int) LRUCache {
	return LRUCache{
		kvs: make(map[int]*Node),
		cap: capacity,
	}
}

func (l *LRUCache) removeNode(node *Node) *Node {
	if node == nil {
		return node
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}

	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}
	node.prev, node.next = nil, nil

	return node
}

func (l *LRUCache) insertFront(node *Node) *Node {
	if l.head == nil {
		l.head, l.tail = node, node
	} else {
		l.head.prev, node.next = node, l.head
		l.head = node
	}

	return node
}

func (l *LRUCache) moveToHead(node *Node) *Node {
	l.removeNode(node)
	return l.insertFront(node)
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.kvs[key]
	if !ok {
		return -1
	}
	l.moveToHead(node)
	return node.val
}

func (l *LRUCache) Put(key int, value int) {
	node, ok := l.kvs[key]
	if ok {
		node.val = value
		l.kvs[key] = l.moveToHead(node)
		return
	}

	if l.size >= l.cap {
		delete(l.kvs, l.tail.key)
		l.tail = l.tail.prev
		if l.tail != nil {
			l.tail.next = nil
		}
	} else {
		l.size++
	}

	node = &Node{
		key: key, val: value,
	}
	l.kvs[key] = l.insertFront(node)
}

func (l *LRUCache) MoveToHead(node *Node) *Node {
	if node == nil || node.prev == nil {
		return node
	}

	// delete node
	if node == l.tail {
		l.tail = l.tail.prev
	} else {
		node.next.prev = node.prev
	}

	// move to head
	l.head.prev, node.next = node, l.head
	l.head = node

	if l.head.next != nil {
		l.head.next.prev, node.prev = node, l.head
		l.head.next, node.next = node, l.head.next
	} else {
		l.head.next = node
		node.prev = l.head
	}
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
