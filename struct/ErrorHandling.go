package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cannot accept 42")
	}
	return arg, nil
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &customErr{arg, "cannot work with it"}
	} else {
		return arg, nil
	}
}

type customErr struct {
	msg  int
	prob string
}

func (cst *customErr) Error() string {
	return fmt.Sprintf("%d - %s", cst.msg, cst.prob)
}

func main() {
	for _, i := range []int{72, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e.Error())
		} else {
			fmt.Println("f1 passed:", r)
		}
	}

	for _, j := range []int{72, 42} {
		if r1, e1 := f2(j); e1 != nil {
			fmt.Println("f2 failed:", e1.Error())
		} else {
			fmt.Println("f2 passed:", r1)
		}
	}

	_, e2 := f2(42)
	fmt.Println("f2 execution:", e2.Error())
	ae := e2.(*customErr)
	fmt.Println("custom error msg:", ae.msg)
	fmt.Println("custom error prob:", ae.prob)
}
