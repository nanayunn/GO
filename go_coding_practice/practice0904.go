package main

import (
	"fmt"
	"os"
	"strconv"
)

func inputNumber() int {
	var n string

	fmt.Scan(&n)

	parsednum, err := strconv.Atoi(n)
	if err != nil {
		os.Exit(1)
	}

	num := parsednum

	return num
}

func main() {

	A := inputNumber()
	B := inputNumber()

	rest := A - B

	fmt.Println(rest)
}
