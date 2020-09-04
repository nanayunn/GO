package main

import (
	"fmt"
	"os"
	"strconv"
)

func inputNumber(num int) int {
	var n string

	fmt.Scan(&n)

	parsednum, err := strconv.Atoi(n)
	if err != nil {
		os.Exit(1)
	}

	num = parsednum

	return num
}

func main() {
	var a int
	var b int

	A := inputNumber(a)
	B := inputNumber(b)

	rest := A - B

	fmt.Println(rest)
}
