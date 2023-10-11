package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name-field"`
	Age  int    `json:"age"`
}

type List struct {
	name string `conditions:"name"`
	age  int    `conditions:"age"`
}

func main() {
	// 通过反射获取 tag
	user := &User{"John Doe The Fourth", 20}
	field, ok := reflect.TypeOf(user).Elem().FieldByName("Name")
	for i := 0; i < reflect.TypeOf(user).Elem().NumField(); i++ {
		fmt.Println(reflect.TypeOf(user).Elem().Field(i).Tag)
	}
	if !ok {
		fmt.Println("Field not found")
	}
	fmt.Println(getStructTag(field))

	fmt.Println("------------------")
	// 循环遍历 结构体拿到 tag
	list := &List{"xiao yu", 24}
	conditionType := reflect.TypeOf(list).Elem()
	conditionValue := reflect.ValueOf(list).Elem()
	for i := 0; i < conditionType.NumField(); i++ {
		field := conditionType.Field(i)
		fmt.Println("field:", field)
		fmt.Println("field.Tag:", field.Tag)
		tag := field.Tag.Get("conditions")
		value := conditionValue.FieldByName(field.Name)
		fmt.Println("tag:", tag)
		fmt.Println("value:", value)
	}

	// 反射 DeepEqual
	type MyInt int
	type YourInt int
	m := MyInt(1)
	y := YourInt(1)
	fmt.Println(reflect.DeepEqual(m, y)) // false
}

func getStructTag(f reflect.StructField) string {
	return string(f.Tag)
}
