package main

import "ds/stack"

func main() {
	str := "(1 + 2) * [3] <[swe]r([23])>"
	
	items := stack.NewStack[string]()
	println(items.IsStable(str))
}
