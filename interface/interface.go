package main

import "fmt"

type infoer interface {
	info() string
}

type promotion interface {
	discount() int
}

type presenter interface {
	// info()
	// discount() int
	promotion // can send type
	infoer
}

type course struct {
}

func (course) info() string {
	return "for test"
}

func (course) discount() int {
	return 555
}

func summarypromotion(val presenter) {
	fmt.Printf("val.discount(): %v\n", val.discount()) // pass parameter
	fmt.Printf("val.info(): %v\n", val.info())
}

func show(val interface{}) {
	i, ok := val.(int)
	if ok {
		i = i + 2
		fmt.Println("int:", i)
	}

	j, ok := val.(string)
	if ok {
		j = j + " hello"
		fmt.Println("string:", j)
	}

	// use switch case easily (for this any case)
	switch v := val.(type) {
	case int:
		i := v + 2
		fmt.Println("int in switch:", i)
	case string:
		s := v + " hello"
		fmt.Println("string in switch:", s)
	default:
		fmt.Println("not int and string")

	}
}

func main() {
	// var v interface{}
	var v any
	v = 13
	show(v)

	va := course{}
	summarypromotion(va)

}
