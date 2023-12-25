package Slice

import "fmt"

type Person struct {
	Name string
	Age  int32
}

func Test() {
	slc := make([]*Person, 0, 8)
	slc = append(slc, &Person{"Name", 15})

	fmt.Printf("name: %s, age: %d\n", slc[0], slc[0])
}
