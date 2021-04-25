package q1206

import "math/rand"

const MaxLayers = 16

type SkipList struct {
	level  int
	length int
	head   *SkipListNode
}

type SkipListNode struct {
	val   int
	layer []*SkipListNode
}

func NewSkipListNode(val, layers int) *SkipListNode {
	return &SkipListNode{
		val:   val,
		layer: make([]*SkipListNode, layers),
	}
}

func Constructor() SkipList {
	head := NewSkipListNode(-1<<31, MaxLayers)
	return SkipList{
		level:  1,
		length: 0,
		head:   head,
	}
}

func (s *SkipList) Search(target int) bool {
	if s.length < 1 {
		return false
	}
	cur := s.head
	for i := MaxLayers - 1; i >= 0; i-- {
		for cur.layer[i] != nil {
			val := cur.layer[i].val
			if val > target {
				break
			} else if val == target {
				return true
			}
			cur = cur.layer[i]
		}

	}
	return false
}

func (s *SkipList) Add(num int) {
	cur := s.head
	path := make([]*SkipListNode, MaxLayers)

	for i := MaxLayers - 1; i >= 0; i-- {
		for cur.layer[i] != nil {
			if cur.layer[i].val >= num {
				path[i] = cur
				break
			}
			cur = cur.layer[i]
		}
		if cur.layer[i] == nil {
			path[i] = cur
		}
	}

	level := 1
	for i := level; i < MaxLayers; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}
	node := NewSkipListNode(num, level)
	for i := 0; i < level; i++ {
		node.layer[i], path[i].layer[i] = path[i].layer[i], node
	}
	s.length++
	if level > s.level {
		s.level = level
	}
}

func (s *SkipList) Erase(num int) bool {
	if s.length < 1 {
		return false
	}
	cur := s.head
	path := make([]*SkipListNode, MaxLayers)
	for i := MaxLayers - 1; i >= 0; i-- {
		for cur.layer[i] != nil && cur.layer[i].val < num {
			cur = cur.layer[i]
		}
		path[i] = cur
	}

	if cur.layer[0] != nil && cur.layer[0].val == num {
		for i := MaxLayers - 1; i >= 0; i-- {
			if path[i].layer[i] != nil && path[i].layer[i].val == num {
				path[i].layer[i] = path[i].layer[i].layer[i]
			}
		}
		return true
	}
	return false
}

/**
 * Your SkipList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
