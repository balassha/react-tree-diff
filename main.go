package main

func main() {
	reactTree := NewReactTree(20)
	reactTree.InsertNode("a", 1)
	reactTree.InsertNode("b", 3)
	reactTree.InsertNode("c", 5)
	reactTree.InsertNode("d", 6)
	reactTree.InsertNode("f", 8)

	reactTree2 := NewReactTree(20)
	reactTree2.InsertNode("a", 1)
	reactTree2.InsertNode("b", 3)
	reactTree2.InsertNode("c", 5)
	reactTree2.InsertNode("d", 7)
	reactTree2.InsertNode("h", 9)

	reactTree.DiffTree(reactTree2, INSERT)
}