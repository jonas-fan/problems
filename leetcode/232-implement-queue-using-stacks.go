// #stack

type MyQueue struct {
    pushStack []int
    popStack  []int
}


func Constructor() MyQueue {
    return MyQueue{
        pushStack: []int{},
        popStack:  []int{},
    }
}


func (this *MyQueue) Push(x int)  {
    for i := len(this.popStack) - 1; i >= 0; i-- {
        this.pushStack = append(this.pushStack, this.popStack[i])
    }

    this.popStack = this.popStack[:0]
    this.pushStack = append(this.pushStack, x)
}


func (this *MyQueue) Pop() int {
    top := this.Peek()

    this.popStack = this.popStack[:len(this.popStack)-1]

    return top
}


func (this *MyQueue) Peek() int {
    for i := len(this.pushStack) - 1; i >= 0; i-- {
        this.popStack = append(this.popStack, this.pushStack[i])
    }

    this.pushStack = this.pushStack[:0]

    return this.popStack[len(this.popStack)-1]
}


func (this *MyQueue) Empty() bool {
    if len(this.pushStack) != 0 {
        return false
    } else if len(this.popStack) != 0 {
        return false
    }

    return true
}
