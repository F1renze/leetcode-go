package q347

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)

	for _, n := range nums {
		countMap[n]++
	}
	heap := &Heap{
		cntMap: countMap,
		data: []int{-1},
	}

	for k := range countMap {
		heap.insert(k)
	}

	var result []int
	for i := 0; i < k; i++ {
		result = append(result, heap.pop())
	}
	return result
}

// max heap
type Heap struct {
	cntMap map[int]int
	data []int
}

func (h *Heap) insert(val int) {
	h.data = append(h.data, val)
	h.swim(h.size() -1)
}

func (h *Heap) pop() int {
	val := h.data[1]

	h.swap(1, h.size()-1)
	h.data = h.data[:h.size()-1]
	h.sink(1)
	return val
}


func (h *Heap) size() int {
	return len(h.data)
}

func (h *Heap) parent(i int) int {
	return i >> 1
}

func (h *Heap) left(i int) int {
	return i << 1
}

func (h *Heap) right(i int) int {
	return (i << 1) + 1
}

func (h *Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) getValCnt(i int) int {
	return h.cntMap[h.data[i]]
}

func (h *Heap) swim(i int) {
	if i < 2 {
		return
	}

	par := h.parent(i)
	for ;h.getValCnt(par) < h.getValCnt(i); {
		h.swap(par, i)
		i, par = par, h.parent(par)
		if par == 0 {
			break
		}
	}
}

func (h *Heap) sink(i int) {
	l, r, b := h.left(i), h.right(i), i

	if l < h.size() && h.getValCnt(l) > h.getValCnt(b) {
		b = l
	}

	if r < h.size() && h.getValCnt(r) > h.getValCnt(b) {
		b = r
	}

	if b != i {
		h.swap(i, b)
		h.sink(b)
	}
}
