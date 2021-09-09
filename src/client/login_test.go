package main

import (
	"fmt"
	"testing"
)

func TestGo(t *testing.T) {
	ints := []int{0, 1, 2, 3, 4}
	fmt.Println(ints[:4])
	fmt.Println(ints[4:])
}
