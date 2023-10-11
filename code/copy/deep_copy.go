package main

// 定义一个Robot结构体
type Robot struct {
	Name  string
	Color string
	Model string
}

//func main() {
//	fmt.Println("深拷贝 内容一样，改变其中一个对象的值时，另一个不会变化。")
//	robot1 := Robot{
//		Name:  "小白-X型-V1.0",
//		Color: "白色",
//		Model: "小型",
//	}
//	robot2 := robot1
//	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, &robot1)
//	fmt.Printf("Robot 2：%s\t内存地址：%p \n", robot2, &robot2)
//
//	fmt.Println("修改Robot1的Name属性值")
//	robot1.Name = "小白-X型-V1.1"
//
//	fmt.Printf("Robot 1：%s\t内存地址：%p \n", robot1, &robot1)
//	fmt.Printf("Robot 2：%s\t内存地址：%p \n", robot2, &robot2)
//}
