package main

import "fmt"

type work struct { // assume you crate slice my youself
	name  string
	price float64
	wfh   bool
}

// var testman struct {
// 	name: "testwork",
// 	price: 5000,
// }

func main() {
	// c.name, c.price
	c := work{"A work", 3000, true} // fix position

	fmt.Printf("c: %#v\n", c)

	c_work := work{name: "ss"}
	fmt.Printf("c_work: %#v\n", c_work)

	c_wfh := work{wfh: true, name: "dd"} // need tag
	fmt.Printf("c_wfh: %#v\n", c_wfh)
}
