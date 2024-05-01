type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	newArray := make([]int, len(nums)+1)

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		newArray[i+1] = sum
	}

	return NumArray{
		nums: newArray,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.nums[right+1] - this.nums[left]
}