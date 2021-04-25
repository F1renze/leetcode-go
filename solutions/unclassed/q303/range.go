package q303

type NumArray struct {
	data []int
}

func Constructor(nums []int) NumArray {
	var obj NumArray
	if len(nums) < 1 {
		return obj
	}
	obj.data = make([]int, len(nums)+1)
	for i := range nums {
		obj.data[i+1] = obj.data[i] + nums[i]
	}
	return obj
}

func (n *NumArray) SumRange(i int, j int) int {
	return n.data[j+1] - n.data[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
