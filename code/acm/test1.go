package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1、输入: 包括两个正整数a,b(1 <= a, b <= 10^9)，输入数据包括多组。
	// 输出: a+b的结果
	//var a, b int
	//for {
	//	n, _ := fmt.Scan(&a, &b)
	//	if n == 0 {
	//		break
	//	} else {
	//		fmt.Println(a + b)
	//	}
	//}

	// 2、第一行包括一个数据组数t(1 <= t <= 100)  接下来每行包括两个正整数a,b(1 <= a, b <= 10^9)
	// 输出: a+b的结果
	//var t, a, b int
	//_, _ = fmt.Scanln(&t)
	//for i := 0; i < t; i++ {
	//	_, _ = fmt.Scan(&a, &b)
	//	fmt.Println(a + b)
	//}

	// 3、包括两个正整数a,b(1 <= a, b <= 10^9),输入数据有多组, 如果输入为0 0则结束输入
	// 输出: a+b的结果
	//var a, b int
	//for {
	//	fmt.Scan(&a, &b)
	//	if a == 0 && b == 0 {
	//		break
	//	} else {
	//		fmt.Println(a + b)
	//	}
	//}

	// 4、输入数据包括多组。 每组数据一行,包含两个字符串形式的非负整数
	// 输出: 对于每组测试数据，计算它们的和，输出字符串。
	//var s1, s2 string
	//for {
	//	n, _ := fmt.Scan(&s1, &s2)
	//	if n == 0 {
	//		break
	//	} else {
	//		fmt.Println(bigNumberAdd(s1, s2))
	//	}
	//}

	// 5、输入:
	// 输入数据包括多组。
	// 每组数据一行,每行的第一个整数为整数的个数n(1 <= n <= 100), n为0的时候结束输入。
	// 接下来n个正整数,即需要求和的每个正整数。
	// 输出: 每组数据输出求和的结果
	//var n int
	//for {
	//	var sum int
	//	fmt.Scan(&n)
	//	if n == 0 {
	//		break
	//	} else {
	//		a := make([]int, n)
	//		for i := 0; i < n; i++ {
	//			fmt.Scan(&a[i])
	//			sum += a[i]
	//		}
	//	}
	//	fmt.Println(sum)
	//}

	// 6、输入:
	//第一行包括一个正整数t(1 <= t <= 100), 表示数据组数。
	//接下来t行, 每行一组数据。
	//每行的第一个整数为整数的个数n(1 <= n <= 100)。
	//接下来n个正整数, 即需要求和的每个正整数。
	//
	//输出: 每组数据输出求和的结果
	//var t int
	//fmt.Scan(&t)
	//for i := 0; i < t; i++ {
	//	var num, sum, a int
	//	fmt.Scan(&num)
	//	for i := 0; i < num; i++ {
	//		fmt.Scan(&a)
	//		sum += a
	//	}
	//	fmt.Println(sum)
	//}

	// 7、输入:
	//输入数据有多组, 每行表示一组输入数据。
	//每行的第一个整数为整数的个数n(1 <= n <= 100)。
	//接下来n个正整数, 即需要求和的每个正整数。
	//
	//输出:每组数据输出求和的结果
	//for {
	//	var n int
	//	fmt.Scan(&n)
	//	var sum, num int
	//	for i := 0; i < n; i++ {
	//		fmt.Scan(&num)
	//		sum += num
	//	}
	//	fmt.Println(sum)
	//}

	// 8、输入:
	//输入数据有多组, 每行表示一组输入数据。
	//每行不定有n个整数，空格隔开。(1 <= n <= 100)。
	//
	//输出: 每组数据输出求和的结果
	//sc := bufio.NewScanner(os.Stdin)
	//for sc.Scan() { //每次读入一行
	//	strs := strings.Split(sc.Text(), " ") //通过空格将他们分割，并存入一个字符串切片
	//	var sum int
	//	for i := range strs {
	//		val, _ := strconv.Atoi(strs[i]) //将字符串转换为int
	//		sum += val
	//	}
	//	fmt.Println(sum)
	//}

	// 9、输入:
	//输入有两行，第一行n
	//第二行是n个空格隔开的字符串
	//
	//输出:
	//输出一行排序后的字符串，空格隔开，无结尾空格
	//sc := bufio.NewScanner(os.Stdin)
	//fmt.Println(sc.Scan())
	//for sc.Scan() {
	//	strs := strings.Split(sc.Text(), " ")
	//	sort.Strings(strs)                   //排序
	//	fmt.Println(strings.Join(strs, " ")) //将切片连接成字符串
	//}

	// 10、输入:
	//多个测试用例，每个测试用例一行。
	//每行通过,隔开，有n个字符串，n＜100
	//
	//输出:
	//对于每组用例输出一行排序后的字符串，用','隔开，无结尾空格
	//sc := bufio.NewScanner(os.Stdin)
	//for sc.Scan() {
	//	strs := strings.Split(sc.Text(), ",")
	//	sort.Strings(strs)
	//	fmt.Println(strings.Join(strs, ","))
	//}
	n := 0
	fmt.Scan(&n)
	input := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}
	fmt.Println(input)
}

func bigNumberAdd(s1 string, s2 string) string {
	add := 0
	res := ""
	for i, j := len(s1)-1, len(s2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(s1[i] - '0')
		}
		if j >= 0 {
			y = int(s2[i] - '0')
		}
		result := x + y + add
		res += strconv.Itoa(result % 10)
		add /= 10
	}
	return res
}
