package main

import "fmt"

func main() {
	sm := NewSkipMap[string, int]()
	sm.Insert("hello", 0)
	sm.Insert("world", 1)
	sm.Insert("foo", 0)
	sm.Insert("bar", 1)
	sm.Insert("baz", 2)

	fmt.Println(sm.Get("hello"))
	fmt.Println(sm.Get("world"))
	fmt.Println(sm.Get("foo"))
	fmt.Println(sm.Get("bar"))
	fmt.Println(sm.Get("baz"))

	sm.Insert("foo", 7)
	fmt.Println(sm.Get("foo"))

	sm.Delete("foo")
	_, err := sm.Get("foo")
	if err != nil {
		panic(err)
	}
}
