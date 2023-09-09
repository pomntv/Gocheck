package main

import "fmt"

func show(a []string) {
	fmt.Printf("length %v -- show %s \n", len(a), a)
}

func main() {
	room := []string{"aa", "bb", "cc", "dd", "ee"}
	// fmt.Printf("room: %v , len: %v \n", room, len(room))

	room2 := append(room, "ff") // append
	// fmt.Printf("room2: %v, len: %v\n", room2, len(room2))

	show(room2)

	roomsl := room[0:3] // 0 1 only [) golang เป็นช่วงเปิด
	show(roomsl)

	//normal
	// show(roomsl[0:3])

	//slice assignd use append
	var ss []int
	fmt.Printf("ss: %v\n", ss)

	ss = append(ss, 5)
	fmt.Printf("ss: %v\n", ss)

	v := ss[0] // <----- don't assign nil
	fmt.Printf("v: %v\n", v)

	sss := make([]int, 6)
	fmt.Printf("len: %v ,type: %T,value: %#v \n", sss, sss, sss)

}
