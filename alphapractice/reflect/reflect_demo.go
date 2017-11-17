package reflect

import (
	"fmt"
	"reflect"
)

type Human struct {
	Name string
	Age  int
	sex  int
}

func hello(params ...int) string {
	for _, i := range params {
		fmt.Println(i, " ")
	}
	return "method return"
}

func test() {
	hl := hello
	hl(1, 2, 3)

	v := reflect.ValueOf(hl)

	fmt.Println(v.Type())
	fmt.Println(v.Kind())
	params := make([]reflect.Value, 2) //定义两个参数
	params[0] = reflect.ValueOf(20)
	params[1] = reflect.ValueOf(10)
	values := v.Call(params)
	fmt.Println(values[0].Interface())

	human := &Human{
		Name: "jia",
		Age:  18,
		sex:  0,
	}

	rv := reflect.ValueOf(human)
	fmt.Println(rv.Elem())
}
