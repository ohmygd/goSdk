package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("first")
	a := l.PushBack("second")
	l.PushFront("67")
	l.PushFront("68")
	l.InsertAfter(89, a)

	for i:=l.Front(); i!=nil; i=i.Next() {
		fmt.Println(i.Value, "===============")
	}
}