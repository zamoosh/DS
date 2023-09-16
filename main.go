package main

import "ds/array"

func main() {
	items := array.NewArrayTypeAny[string](5)
	items.Insert("a")
	items.Insert("c")
	items.Insert("c")
	items.Insert("b")
	items.Insert("b")
	println(items.IndexOf("b"))
}
