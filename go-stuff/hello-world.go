package main

import "fmt"
import "math"

const s string = "constant string"

func main() {
	// hello world
	fmt.Println("hello world")

	// values
	fmt.Println("go" + "lang")

	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	// variables
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e string
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

	//constants
	fmt.Println(s)

	const n = 50000

	const dn = 3e20 / n
	fmt.Println(dn)

	fmt.Println(int64(dn))

	fmt.Println(math.Sin(n))

	// for
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
}
