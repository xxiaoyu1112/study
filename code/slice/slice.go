package main

import "fmt"

func main() {
	//slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s1 := slice[2:5]
	//s2 := s1[2:6:7]
	//
	//s2 = append(s2, 100)
	//s2 = append(s2, 200)
	//
	//s1[2] = 20
	//
	//fmt.Println(s1)
	//fmt.Println(len(s1), cap(s1))
	//fmt.Println(s2)
	//fmt.Println(len(s2), cap(s2))
	//fmt.Println(slice)

	//arr1 := []int{1, 2, 3, 4, 5, 6}
	//arr2 := arr1
	//arr2 = append(arr2[:2], arr2[3:]...)
	//fmt.Println(arr1)
	//fmt.Println(arr2)

	//e := []int32{1, 2, 3}
	//fmt.Println("cap of e before:", cap(e))
	//e = append(e, 4, 5, 6, 7)
	//fmt.Println("cap of e after:", cap(e))

	//arr := [3]int{1, 3, 4}
	//x := arr[0]
	//arr = arr[1:]
	//fmt.Println(x, arr)

	//a1 := []int{1, 2, 3}
	//a2 := []int{1, 2, 3}
	////fmt.Println(a1 == a2)  // 报错
	//fmt.Println(reflect.DeepEqual(a1, a2))

	//type A struct {
	//	name string
	//}
	//fmt.Println([...]int{1, 2, 3, 4} == [...]int{1, 2, 3, 4})
	////fmt.Println([]int{1, 2, 3, 4} == []int{1, 2, 3, 4})
	//fmt.Println(&A{name: "a"} == &A{name: "a"})
	//fmt.Println(A{name: "a"} == A{name: "a"})

	// 观察 slice 的容量变化
	//s := []int{1, 2, 3, 4}
	//s1 := s[:4]
	//s1 = append(s1, 5)
	//fmt.Printf("len of s: %d, cap of s: %d\n", len(s), cap(s))
	//fmt.Printf("len of s1: %d, cap of s1: %d\n", len(s1), cap(s1))

	// 批量扩容
	//s := []int{1}
	//s1 := []int{1, 2, 3, 4}
	//s = append(s, append([]int{}, s1...)...)
	//fmt.Printf("len of s: %d, cap of s: %d\n", len(s), cap(s))

	//s := make([]int, 10, 12)
	//s1 := s[8:]
	//fmt.Printf("s: %v, len of s: %d, cap of s: %d\n", s, len(s), cap(s))
	//fmt.Printf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))

	//s := make([]int, 10, 12)
	//s1 := s[8:]
	//s1 = append(s1, 2)
	//s2 := []int{1}
	//changeSlice(s1)
	//changeSlice(s2)
	//fmt.Printf("s: %v, len of s: %d, cap of s: %d\n", s, len(s), cap(s))
	//fmt.Printf("s1: %v, len of s1: %d, cap of s1: %d\n", s1, len(s1), cap(s1))
	//fmt.Printf("s2: %v, len of s2: %d, cap of s2: %d\n", s2, len(s2), cap(s2))
	//i := []int{5, 6, 7}
	//hello(i...)

	var n, k int
	fmt.Scan(&n, &k)
	// 如果 n == 1，则数组中只有一个元素，直接输出该元素
	if n == 1 {
		var x int
		fmt.Scan(&x)
		fmt.Println(x)
		return
	}

	// 如果 k == 1，则数组中的元素必须互质
	// 可以构造出一个和最小的等差数列
	if k == 1 {
		sum := n * (n + 1) / 2
		fmt.Println(sum)
		return
	}

	// 如果 n == 2，则两个元素必须互质，直接输出和
	if n == 2 {
		var x, y int
		fmt.Scan(&x, &y)
		if gcd(x, y) == 1 {
			fmt.Println(x + y)
		} else {
			fmt.Println(-1)
		}
		return
	}

	// 如果 n >= 3 且 k > 1，则构造一个和最小的等差数列
	// 数列首项为 k，公差为 k
	sum := k * n * (n - 1) / 2
	fmt.Println(sum)
}

//func main() {
//	var n, k int
//	fmt.Scan(&n, &k)
//
//	// 构造最小的满足条件的数组
//	a := make([]int, n)
//	for i := 0; i < n; i++ {
//		a[i] = (i + 1) * k
//	}
//
//	// 调整数组使得所有元素两两不相等
//	for i := 1; i < n; i++ {
//		for j := i - 1; j >= 0; j-- {
//			if gcd(a[i], a[j]) == k {
//				a[i], a[j] = a[j], a[i]
//			}
//		}
//	}
//
//	// 输出数组元素之和
//	sum := 0
//	for i := 0; i < n; i++ {
//		sum += a[i]
//	}
//	fmt.Println(sum)
//}

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func findMinSum(n, k int) int {
	lcm := k
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcd(i, j) == 1 {
				lcm = lcm * j / gcd(lcm, j)
			}
		}
	}
	ans := lcm * (n / 2)
	if n%2 == 1 {
		ans += lcm / 2
	}
	return ans
}

func recursiveFunc(i int) int {
	if i <= 1 {
		return i
	}
	return recursiveFunc(i-2) + i
}

func changeSlice(s1 []int) {
	s1[0] = -1
	s1 = append(s1, 3)
}

func hello(num ...int) {
	num[0] = 18
}
