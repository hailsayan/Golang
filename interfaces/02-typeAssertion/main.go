package main

import "fmt"

func myfunc(a interface{}) {

	val := a.(string)

	fmt.Println("Value: ", val)

}

func main() {

	var val interface{} = "VALUE"

	myfunc(val)
}
