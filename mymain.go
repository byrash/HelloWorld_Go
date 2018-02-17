package main

import (
	"fmt"

	"github.com/byrash/HelloWorld_Go/MyStuff"
)

const (
	x = iota
)

func main() {
	pointersTest()
}

func pointersTest() {
	x := 10
	fmt.Println(&x)
	changeIt(&x)
	fmt.Println(x)
}

func changeIt(x *int) {
	fmt.Println(&x)
	fmt.Println("Recived x", *x)
	*x = 9999
}

func funcReturnTest() {
	increment := shiv.Wrapper()
	increment1 := shiv.Wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment1())
	fmt.Println(increment1())
}

func scopeTest() {
	fmt.Printf("%v %v \n", shiv.Name, x)
	// fmt.Printf("%v %v \n", shiv.myAliasName, x)
	shiv.PrintVars()
}

func panicTest() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
