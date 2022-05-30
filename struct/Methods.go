package main

import "fmt"

type dimensions struct {
	height, weight int
	age            int
}

func (d *dimensions) bmi() int {
	return d.height * d.weight / d.age
}

func main() {
	dimen := dimensions{height: 180, weight: 210, age: 33}

	fmt.Println("bmi -->", dimen.bmi())

	dimen2 := &dimen

	fmt.Println("bmi 2 -->", dimen2.bmi())
}
