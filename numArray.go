type NumArray struct {
    values []int
}


func Constructor(nums []int) NumArray {
    return NumArray{values: nums}        
}


func (this *NumArray) SumRange(left int, right int) int {
    var sum int
    for i := left; i <= right; i++ {
        sum += this.values[i]
    }
    return sum 
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

