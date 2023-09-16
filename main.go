package main

import (
	"ds/stack"
)

func main() {
	str := "{s<s(s)s>sdf(we)s{d}w[e<f>w(e)f[s]d]ewf}"
	items := stack.NewStack()
	println(items.IsStable(str))
}
