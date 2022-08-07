type MyStack struct {
    values []int
}


func Constructor() MyStack {
    return MyStack{[]int{}}
}


func (this *MyStack) Push(x int)  {
    this.values = append(this.values, x)
}


func (this *MyStack) Pop() int {
    elem := this.Top()
    this.values = this.values[:len(this.values) - 1]
    return elem
}


func (this *MyStack) Top() int {
    return this.values[len(this.values) - 1]
}


func (this *MyStack) Empty() bool {
    return len(this.values) == -0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
