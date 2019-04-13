package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intList list.List
	intList.PushBack(1)
	intList.PushBack(2)
	intList.PushBack(3)
	for el := intList.Front(); el != nil; el = el.Next() {
		fmt.Println(el.Value.(int))
	}
}
