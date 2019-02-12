package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Week day")
	}

	t := time.Now()

	switch {
	case t.Hour() < 19:
		fmt.Println(t)
	default:
		fmt.Println("other time")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It's a bool")
		case int:
			fmt.Println("It's a int")
		default:
			fmt.Printf("Other: %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI("asd")
	whatAmI(10)
}
