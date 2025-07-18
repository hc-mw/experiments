package main

type Node struct {
	keys [][]byte
	vals [][]byte
	kids []*Node
}

func (n *Node) Encode() []byte {

}

func (n *Node) Decode() *Node {

}
