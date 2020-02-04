package main
import (
	"fmt"
	"errors"
)

// 使用数组来模拟一个栈的使用情况
type Stack struct {
	MaxTop int // 表示栈最大个数
	Top int // 表示栈顶， 因为栈顶固定，因此我们直接使用Top
	arr [5]int // 数组模拟栈
}

// 入栈
func (this *Stack) Push(val int) (err error) {
	// 判断栈是否已满
	if this.Top == this.MaxTop - 1 {
		fmt.Println("stack is full")
		return errors.New("stack is full")
	}
	this.Top++
	// 放入数据
	this.arr[this.Top] = val
	return 
}

// 出栈
func (this *Stack) Pop() (val int, err error) {
	// 判断栈是否为空
	if this.Top == - 1 {
		fmt.Println("stack is empty")
		return -1, errors.New("stack is empty")
	}
	// 放入数据
	val = this.arr[this.Top]
	this.Top--
	return val, nil
}

// 遍历栈，需要从栈顶开始遍历
func (this *Stack) List() {
	// 判断栈是否为空
	if this.Top == -1 {
		fmt.Println("stack is empty")
		return 
	}
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d] = %d\n", i, this.arr[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop : 5, // 最多存放5个数到栈中
		Top : -1, // 当栈顶为-1， 表示栈为空
	}

	// 入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	// 显示
	stack.List()
	val, _ := stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	fmt.Println("出栈val =", val)
	stack.List()
}