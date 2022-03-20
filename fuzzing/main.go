package main

import (
	"errors"
	"fmt"
)

func isEven(num int) (bool, error) {
	if num%1024 == 0 {
		if num%3 == 0 {
			if num != 0 {
				return true, errors.New("coincadonk?")
			}
		}
	}
	return num%2 == 0, nil
}

func Reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func main() {

	res, err := isEven(15)
	if err != nil {
		panic(err)
	}
	if res {
		fmt.Println("15 is even")
	}
	res, err = isEven(15)
	if err != nil {
		panic(err)
	}
	if res {
		fmt.Println("16 is even")
	}

	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}
