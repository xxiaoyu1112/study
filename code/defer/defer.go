package main

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func test() {
	x := 2
	defer func() {
		fmt.Printf("first %d\n", x)
	}()
	x = 1
	defer func() {
		fmt.Printf("second %d\n", x)
	}()
	defer func() {
		fmt.Printf("third %d\n", x)
	}()
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}

func DeferTest1(i int) (r int) {
	r = i
	defer func() {
		r += 3
	}()
	return r
}

func DeferTest2(i int) (r int) {
	defer func() {
		r += i
	}()
	return 2
}

func main() {
	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	test()
	//fmt.Println("return:", *(c()))
	//println(DeferTest1(1))
	//println(DeferTest2(1))

	//defer func() {
	//	recover()
	//}()
	//defer fmt.Println("A")
	//defer fmt.Println("B")
	//defer panic("a")\

	//fmt.Println("return:", *(deferPoint()))
}

func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i
}

func deferPoint() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}
