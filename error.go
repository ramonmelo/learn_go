package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Value Error")
	}

	return arg, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Can't work with it"}
	}

	return arg, nil
}

func main() {

	for _, i := range []int{7, 42, 30} {
		if v, e := f1(i); e != nil {
			println("Error", e.Error)
		} else {
			println(v)
		}
	}

	for _, i := range []int{7, 42, 30} {
		if v, e := f2(i); e != nil {
			println("Error", e.Error)
		} else {
			println(v)
		}
	}

	_, e := f2(42)

	if ae, ok := e.(*argError); ok {
		println(ae.arg)
		println(ae.prob)
	}
}
