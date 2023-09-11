package main

import (
	"fmt"
)

func main() {
	//fmt.Println('A')
	//s := "abAB"
	//for i, v := range s {
	//	fmt.Println(i, v)
	//}
	//for i := 0; i < len(s); i++ {
	//	fmt.Println(s[i])
	//}
	//b := []byte(s)
	//for i, v := range b {
	//	fmt.Println(i, v)
	//	if v >= 97 {
	//		b[i] -= 32
	//	} else {
	//		b[i] += 32
	//	}
	//}
	//fmt.Println(string(b))

	//s := strings.Builder{}
	//s.WriteString("a")
	//s.WriteString("b")
	//s.WriteByte('c')
	//s.WriteRune('1')
	//s.Write([]byte{'a', 'c', '2', '1'})
	//fmt.Println(s.String())
	//
	//s1 := bytes.Buffer{}
	//s1.Write([]byte{'a', 'c', '2', '1'})
	//fmt.Println(s1.String())

	var a *int = nil
	var b interface{} = nil
	fmt.Println(a == b)
}
