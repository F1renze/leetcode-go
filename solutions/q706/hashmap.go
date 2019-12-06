package q706

import (
	"strconv"
)

/*
Base on Hash Table
*/

const (
	defaultSize      int     = 10
	growLoadFactor   float32 = 0.75
	shrinkLoadFactor float32 = 0.25
)

type Node struct {
	Kv  *KeyVal
	Next *Node
	Prev *Node
}

type KeyVal struct {
	Key int
	Val int
}

func newNode(key, val int) *Node{
	return &Node{
		Kv: &KeyVal{
			Key: key,
			Val: val,
		},
	}
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		bucketArr: make([]*Node, defaultSize),
	}
}

func djb2Hash(key int) uint {
	bs := []byte(strconv.Itoa(key))

	h := uint(5381)
	for _, r := range bs {
		h = (h << 5) + h + uint(r)
	}
	return h
}

func getSlot(key, size int) int {
	return int(djb2Hash(key)) % size
}

type MyHashMap struct {
	bucketArr []*Node
	size      int
}

func (t *MyHashMap) loadFactor() float32 {
	return float32(t.size) / float32(len(t.bucketArr))
}

func (t *MyHashMap) resize(size int) {
	newBucket := make([]*Node, size)
	for i := range t.bucketArr {
		node := t.bucketArr[i]
		var keysInBucket []*KeyVal
		for {
			if node == nil {
				break
			}
			keysInBucket = append(keysInBucket, node.Kv)
			node = node.Next
		}

		for j := range keysInBucket {
			kv := keysInBucket[j]
			newSlot := getSlot(kv.Key, size)
			node = newBucket[newSlot]
			head := newNode(kv.Key, kv.Val)
			head.Next = node
			if node != nil {
				node.Prev = head
			}
			newBucket[newSlot] = head
		}
	}
	t.bucketArr = newBucket
}

func (t *MyHashMap) grow() {
	newSize := len(t.bucketArr) * 2
	t.resize(newSize)
}

func (t *MyHashMap) shrink() {
	newSize := len(t.bucketArr) / 2
	if defaultSize >= newSize {
		newSize = defaultSize
	}
	t.resize(newSize)
}

func  (t *MyHashMap) searchNode(key int)(*Node, bool) {
	slot := getSlot(key, len(t.bucketArr))
	node :=  t.bucketArr[slot]

	for {
		switch {
		case node == nil:
			fallthrough
		case node.Next == nil && node.Kv.Key != key:
			return node, false
		case node.Kv.Key == key:
			return node, true
		default:
			node = node.Next
		}
	}
}

/** value will always be non-negative. */
func (t *MyHashMap) Put(key , val int) {
	node, ok := t.searchNode(key)
	if ok {
		node.Kv.Val = val
		return
	}
	newNode := newNode(key, val)
	if node != nil {
		node.Next = newNode
		newNode.Prev = node
	} else {
		slot := getSlot(key, len(t.bucketArr))
		t.bucketArr[slot] = newNode
	}

	t.size++

	if t.loadFactor() >= growLoadFactor {
		t.grow()
	}
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (t *MyHashMap) Remove(key int) {
	node, ok := t.searchNode(key)
	if !ok {
		return
	}

	slot := getSlot(key, len(t.bucketArr))
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		t.bucketArr[slot] = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	t.size--

	if t.loadFactor() <= shrinkLoadFactor {
		t.shrink()
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (t *MyHashMap) Get(key int) int {
	node, ok := t.searchNode(key)
	if ok {
		return node.Kv.Val
	} else {
		return -1
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */