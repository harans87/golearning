package main

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	name string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		name: "containerName",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.name)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer -->", d.describe())
}
