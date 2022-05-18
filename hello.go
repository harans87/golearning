package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

const s string = "constant"

func main() {
	fmt.Println("Hello, Hari!")
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	e := "yashicaa"
	fmt.Println(e)

	fmt.Println(s)

	const n = 500000000

	const da = 3e20 / n
	fmt.Println(da)

	fmt.Println(int64(da))

	fmt.Println(math.Abs(da))

	i := 1
	for i <= 9 {
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	if num := 9; num < 10 {
		fmt.Println(num)
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("boolean value")
		case int:
			fmt.Println("int value")
		default:
			fmt.Println("dont know")
		}

	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("")

	var arr [4]int
	fmt.Println(arr[0])

	arr[0] = 56

	fmt.Println(arr[0])
	fmt.Println(len(arr))

	arr2 := [...]int{21, 22, 23, 24, 25}
	fmt.Println("length of arr2:", len(arr2))

	var twoD [10][10]int
	fmt.Println(len(twoD))

	// fmt.Println(twoD)

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	//slices - power ful interface in go

	slices := make([]string, 3)
	fmt.Println("slices:", slices)

	slices[0] = "first"
	slices[1] = "second"
	slices[2] = "third"

	fmt.Println("length of slices:", len(slices))

	slices = append(slices, "forth")
	slices = append(slices, "fifth", "sixth")

	fmt.Println("slices with append:", slices)

	slice1 := slices[2:]

	fmt.Println("first slice:", slice1)

	slice2 := slices[:5]

	fmt.Println("second slice:", slice2)

	slice3 := slices[3:5]

	fmt.Println("third slice:", slice3)

	twoDSlices := make([][]int, 10)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDSlices[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDSlices[i][j] = i * j
		}
	}
	fmt.Println("twoD slices:", twoDSlices)

	// map ds
	m1 := make(map[string]int)

	fmt.Println("empty map:", m1)

	m1["yashi"] = 1
	m1["tanvi"] = 2

	fmt.Println("map contents:", m1)

	fmt.Println("map value:", m1["tanvi"])

	fmt.Println("map length:", len(m1))

	_, pres := m1["tanvi"]

	fmt.Println("value present:", pres)

	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", m2)

	// range
	nums := []int{23, 21, 34, 56, 77}
	sum := 0
	for _, i := range nums {
		sum += i
	}
	fmt.Println("sum:", sum)

	sum2 := 0
	for k, v := range nums {
		if v == 23 {
			fmt.Println("index:", k)
		}

	}
	fmt.Println("sum2:", sum2)

	kvs := map[string]string{"yashi": "1", "tanvi": "2"}

	for h, vv := range kvs {
		fmt.Printf("%s --> %s\t", h, vv)
	}
	sum3 := add(1,2)
	fmt.Println("\nfunction sum:",sum3)

	a1,b1 := multiFunc(22,22)

	fmt.Println("multi function a1:",a1)
	fmt.Println("multi function b1:",b1)

	_, b2 := multiFunc(1,0)

	fmt.Println("with optional return", b2)

	t1 := variadicFunctions(1,3,4,5)
	fmt.Println("total:",t1)

	slice10 := []int{10,4,5,6}
	fmt.Println("total2:",variadicFunctions(slice10...))

	closureExample := closures() 

	fmt.Println(closureExample())
	fmt.Println(closureExample())

	fmt.Println("recursive function:",fact(7))

	var fib func(m int) int 

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n - 1) + fib(n-2)
	}

	fmt.Println("fib recursive:",fib(7))

	ptr := 1
	pointerDemo(ptr)
	fmt.Println("before changing memory address, value is :",ptr)
	fmt.Println("before changing memory address, address is :",&ptr)
	ipointer(&ptr)
	fmt.Println("after changing memory address, value is:",ptr)
	fmt.Println("after changing memory address, address is:",&ptr)

	const str = "hello"
	fmt.Println("string:",str);
	fmt.Println("string length:",len(str))

	//string loop
	for i :=0;i<len(str);i++ {
		fmt.Println("string loop:",str[i])
	}
	//rune count
	fmt.Println("Rune count:", utf8.RuneCountInString(str))
}

func add (a,b int) int {
	return a + b
}

// multi return values
func multiFunc(a,b int) (int, int) {
	s := a + b
	d := a - b
	return s , d
}

func variadicFunctions(nums...int) int {
	total := 0;
	for _,num := range nums {
		total += num
	}
	return total
}

// closures are anonymous functions
func closures() func() int {
	finalInt := 0
	return func() int {
		finalInt++;
		finalInt += 1 
		return finalInt
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

//pointers
func pointerDemo(val int) {
	val = 0;
}

func ipointer(ptr *int) {
	*ptr = 1;
}
