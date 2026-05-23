package search

type Node struct {
	Point  Record
	Radius float32

	Left  *Node
	Right *Node
}
