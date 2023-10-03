package test

import (
	"fmt"
	"math"
)

const s string = "constant"

func privateFunction() int {
	return 0
}

func PublicFunction() bool {

	fmt.Println("Worked")

	return true
}

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {

	fmt.Println(s)

	const n = 5000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	if n > 20 {
		fmt.Println("Weird")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Print(Monday)

}
