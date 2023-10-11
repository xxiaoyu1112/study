package main

import "fmt"

func calculate(s string) int {
	nums := []int{} // 栈。解析过的数字入栈
	curNum := 0     // 当前遇到的数字
	prevOp := '+'   // 初始化为加号，保证第一个数正常入栈（上一个遇到的运算符）
	s = s + "+"     // 追加一个运算符，不然会少做最后一次运算

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' { // 如果是数字
			curNum = curNum*10 + int(s[i]-'0') // 更新curNum
		} else if s[i] == ' ' { // 如果遇到空格，跳过
			continue
		} else { // 如果遇到加减乘除
			if prevOp == '+' { // 如果上一个遇到的运算符是加号
				nums = append(nums, curNum) // 将curNum入栈
			} else if prevOp == '-' {
				nums = append(nums, -1*curNum) // 将负的curNum入栈
			} else if prevOp == '*' {
				nums[len(nums)-1] = nums[len(nums)-1] * curNum // 更新栈顶为乘过的值
			} else if prevOp == '/' {
				nums[len(nums)-1] = nums[len(nums)-1] / curNum // 更新栈顶为除过的值
			}
			// 更换当前运算符，s[i] 为byte类型，通过rune转化为int32类型
			prevOp = rune(s[i]) // 记录当前遇到的运算符
			curNum = 0          // curNum 重置
		}
	}

	sum := 0
	for _, v := range nums { // 将nums栈中的数字逐个累加
		sum += v
	}
	return sum
}

func calculate1(s string) (ans int) {
	stack := []int{}
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}

func main() {
	s := "3+2*2-2+2/1"
	res := calculate1(s)
	fmt.Println("res:", res)
}
