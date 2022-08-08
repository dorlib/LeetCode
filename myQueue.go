type MyQueue struct {
    values []int
}


func Constructor() MyQueue {
    return MyQueue{values: []int{}}
}


func (this *MyQueue) Push(x int)  {
    elem := []int{x}
    this.values = append(elem, this.values...)
}


func (this *MyQueue) Pop() int {
    elem := this.Peek()
    this.values = this.values[:len(this.values) -1]
    return elem
}


func (this *MyQueue) Peek() int {
    return this.values[len(this.values) - 1]
}


func (this *MyQueue) Empty() bool {
    return len(this.values) == 0 
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
