package main
import (
	"fmt"
	"errors"
	"strconv"
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
	// 42 : *
	// 43 : + 
	// 45 : - 
	// 47 : / 
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 编写一个方法，返回某个运算符的优先级（计算器的运算符优先级由自己定义）
// [* / => 1   + - => 0]
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

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
	numStack := &Stack{
		MaxTop : 20,
		Top : -1,
	}
	// 符号栈
	operStack := &Stack{
		MaxTop : 20,
		Top : -1,
	}

	exp := "3+3*6-6"
	// 定义一个index， 帮助扫描exp
	index := 0
	num1 := 0 // 数1
	num2 := 0 // 数2
	oper := 0 // 运算符
	result := 0 // 结果

	for {
		ch := exp[index:index + 1] // 字符串
		temp := int([]byte(ch)[0]) // 就是字符串对应的ASCII码
		if operStack.IsOper(temp) {
			// 说明是符号
			if operStack.Top == -1 {
				// 如果operStack 是一个空栈， 直接入栈
				operStack.Push(temp)
			} else {
				// 如果发现 operStack 栈顶的运算符的优先级大于等于当前准备入栈的优先级
				// 就从符号栈pop出一个符号， 并从数栈中pop 两个数进行运算，运算后的结果
				// 再次放入到数栈，最后当前符号才入符号栈
				if operStack.Priority(operStack.arr[operStack.Top]) >= 
					operStack.Priority(temp) {
						num1, _ = numStack.Pop()
						num2, _ = numStack.Pop()
						oper, _ = operStack.Pop()
						result = operStack.Cal(num1, num2, oper)
						// 计算结果重新入数栈
						numStack.Push(result)
						// 当前的符号在压入符号栈
						operStack.Push(temp)
				} else {
					// 直接入符号栈
					operStack.Push(temp)
				}
			}

		} else {
			// 说明是数
			// 参数1 数字的字符串形式
			// 参数2 数字字符串的进制 比如二进制 八进制 十进制 十六进制
			// 参数3 返回结果的bit大小 也就是int8 int16 int32 int64

			val, _ := strconv.ParseInt(ch, 10, 64)
			numStack.Push(int(val))
		}
		// 继续扫描
		// 先判断 index 是否已经扫描到计算表达式的最后
		if index + 1 == len(exp) {
			break
		}
		index++
	}
	// 如果扫描完毕，依次从符号栈取出符号，然后在数栈中取出两个数
	// 运算后的结果，入数栈，再次从符号栈取出符号，依次内推，直到符号栈为空
	for {
		if operStack.Top == -1 {
			break // 符号栈已取完，退出
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		// 计算结果重新入数栈
		numStack.Push(result)
	}
	res, _ := numStack.Pop()
	fmt.Printf("表达式%s = %v", exp, res)
}