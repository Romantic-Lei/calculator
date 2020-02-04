package main
import (
	"fmt"
	"errors"
)

// 使用数组来模拟一个栈的使用情况
type Stack struct {
	MaxTop int // 表示栈最大个数
	Top int // 表示栈顶， 因为栈顶固定，因此我们直接使用Top
	arr [20]int // 数组模拟栈
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
// 判断一个字符是不是一个运算符[+, -, *, /]
func (this *Stack) IsOper(val int) bool {
	// 根据基本运算符的 ASCII码 来判断
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 编写一个方法，返回某个运算符的优先级（计算器的运算符优先级由自己定义）
// [* / => 1   + - => 0]

// 运算的方法
func (this *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0;
	switch oper {
		case 42 :
			res = num2 * num1
		case 43 :
			res = num2 + num1
		case 45 :
			res = num2 - num1
		case 47 :
			res = num2 / num1
		default :
			fmt.Println("运算符有误")
	}
	return res
}

func main() {
	// 数栈
	numStack = &Stack{
		MaxTop : 20,
		Top : -1,
	}

	exp := "3 + 2 * 6 - 2"
	// 定义一个index， 帮助扫描exp
	index := 0
	for {
		ch := exp[index:index + 1] // 
	}
}