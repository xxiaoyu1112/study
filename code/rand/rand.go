package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2006-01-02 00:00:05", time.Local)
	//x := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	th := strconv.Itoa(t.Hour())
	tm := strconv.Itoa(t.Minute())
	if len(tm) < 2 {
		tm = "0" + tm
	}
	if len(th) < 2 {
		th = "0" + th
	}
	ShowEnterpriseID := "1" + th + tm + fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	fmt.Println(ShowEnterpriseID)
}

//var nil = new(int)
//
//func main() {
//	var p *int
//	fmt.Printf("p:%#v\n", p)
//	if p == nil {
//		fmt.Println("p is nil")
//	} else {
//		fmt.Printf("nil:%#v\n", nil)
//		fmt.Println("p is not nil")
//	}
//}

//func A() int {
//	time.Sleep(100 * time.Millisecond)
//	return 1
//}
//
//func B() int {
//	time.Sleep(1000 * time.Millisecond)
//	return 2
//}
//
//func main() {
//	ch := make(chan int, 1)
//	go func() {
//		select {
//		case ch <- A():
//		case ch <- B():
//		default:
//			ch <- 3
//		}
//	}()
//	fmt.Println(<-ch)
//}
