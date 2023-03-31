package main

import "go-design-pattern/composite"

func main() {
	comp := composite.NewComposite()
	leaf1 := composite.NewLeaf(100)
	comp.Add(leaf1)
	leaf2 := composite.NewLeaf(101)
	comp.Add(leaf2)
	leaf3 := composite.NewLeaf(102)
	comp.Add(leaf3)
	comp.Execute()
}
