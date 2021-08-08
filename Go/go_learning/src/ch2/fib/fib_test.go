package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	{
		// type 1
		//var a int = 1
		//var b int = 1

		// type 2
		//var (
		//	a int = 1
		//	b     = 1
		//)
	}
	// type 3
	a := 1
	b := 1
	// fmt.Print(a)
	t.Log(a)
	for i := 0; i < 5; i++ {
		//fmt.Print(" ", b)
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	// fmt.Println()
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)

}
