package main

import "fmt"

type address struct {
	addressLine1 string
	zipcode      string
}

func newAddress(addressLine1, zipcode string) *address {
	addr := address{addressLine1: addressLine1, zipcode: zipcode}
	return &addr
}

//struct is mutuable DS
func addressDemo() {
	fmt.Println(address{"addressLine1", "30062"})
	fmt.Println(newAddress("addressLine2", "30080"))
	addr1 := address{addressLine1: "addr2", zipcode: "30090"}
	sp := &addr1
	fmt.Println(sp.zipcode)

	sp.zipcode = "30070"

	fmt.Println(sp.zipcode)
}

func main() {
	addressDemo()
}
