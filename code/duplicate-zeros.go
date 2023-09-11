package main

import "fmt"

func duplicateZeros(arr []int)  {
	l := len(arr)
	i := 0
	for i<l{
		if arr[i]==0{
			copy(arr[i+1:],arr[i:])
			i+=2
			continue
		}
		i++
	}
}

// 算法3： append复制转移法
func duplicateZeros1(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == 0 {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++
		}
	}
}

func main() {
	arr := []int{1,0,2,3,0,4,5,0}
	duplicateZeros(arr)
	duplicateZeros1(arr)
	fmt.Println("arr:",arr)
}
