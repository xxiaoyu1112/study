package main

import (
	"code/code/utils"
	"fmt"
)

func trap1(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		// 设置左边最高柱子
		leftMax = utils.Max(leftMax, height[left])
		//设置右边最高柱子
		rightMax = utils.Max(rightMax, height[right])
		if height[left] < height[right] {
			//右边必定有柱子挡水，所以遇到所有值小于等于leftMax的，全部加入水池中
			ans += leftMax - height[left]
			left++
		} else {
			//左边必定有柱子挡水，所以，遇到所有值小于等于rightMax的，全部加入水池
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

func trap(height []int) (ans int) {
	// 存着下标，计算的时候用下标对应的柱子高度
	stack := []int{}
	for i, h := range height {
		// 当前遍历的元素（柱子）高度大于栈顶元素的高度，此时就出现凹槽了
		for len(stack) > 0 && h > height[stack[len(stack)-1]] {
			// 栈顶元素是凹槽的底部，也就是中间位置
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			// 注意减一，只求中间宽度
			curWidth := i - left - 1
			// 两侧的较矮一方的高度 - 凹槽底部高度
			curHeight := utils.Min(height[left], h) - height[top]
			ans += curWidth * curHeight
		}
		stack = append(stack, i)
	}
	return
}

func main() {
	ans := trap1([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Println("ans:", ans)
}
