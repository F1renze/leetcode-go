package q703


type KthLargest struct {
	heap []int
	k int
}


func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		heap: []int{-1},
		k: k + 1,
	}
	for i := range nums {
		kth.insert(nums[i])
	}
	return kth
}


func (h *KthLargest) Add(val int) int {
	h.insert(val)
	return h.heap[1]
}

func (h *KthLargest) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *KthLargest) parent(pos int) int {
	return pos >> 1
}

func (h *KthLargest) left(pos int) int{
	return pos << 1
}

func (h *KthLargest) right(pos int) int {
	return (pos << 1) + 1
}

func (h *KthLargest) insert(val int) {
	size := len(h.heap)
	if size >= h.k && val < h.heap[1] {
		return
	}

	if size >= h.k {
		h.heap[1] = val
		h.sink(1)
	} else {
		h.heap = append(h.heap, val)
		h.swim(len(h.heap) -1)
	}
}

// 上浮
func (h *KthLargest) swim(pos int) {
	if pos < 2 {
		return
	}
	par := h.parent(pos)
	for ; h.heap[par] > h.heap[pos]; {
		h.swap(par, pos)
		pos, par = par, h.parent(par)
		if par == 0 {
			break
		}
	}
}

// 下沉
func (h *KthLargest) sink(pos int) {
	left, right := h.left(pos), h.right(pos)
	smallest := pos

	if left < len(h.heap) && h.heap[left] < h.heap[smallest] {
		smallest = left
	}

	if right < len(h.heap) && h.heap[right] < h.heap[smallest] {
		smallest = right
	}

	if smallest != pos {
		h.swap(pos, smallest)
		h.sink(smallest)
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

