package main

import (
	"fmt"
	"model"
)

func main() {
	A = 2
	i := 1
	ptr0 := &i
	fmt.Printf("%x, %d, %x,%x\n", ptr0, *ptr0, &ptr0, &i)
	i = 2
	fmt.Printf("%x, %d, %x,%x", ptr0, *ptr0, &ptr0, &i)
	fmt.Println("hot-fix")
}
