package gomerkeltreestream

import (
	flat "github.com/SirRujak/goflattree"
)

type TreeLeaf struct {
	Index  uint
	Parent uint
	Hash   []byte
	Size   uint
	Data   []byte
}

type Hash []byte

type NodeParts struct {
	Node PartialNode
	Hash []byte
}

type MerkelTreeStream struct {
	Handler HashMethods
	Roots   []DefaultNode
	Blocks  uint
}

type HashMethods struct {
	Leaf   StreamLeaf
	Parent StreamParent
}

type StreamLeaf func(leaf PartialNode, roots []DefaultNode) []byte

type StreamParent func(a, b DefaultNode) []byte

func (mTS *MerkelTreeStream) New(handler HashMethods, roots []DefaultNode) {
	mTS.Handler = handler
	mTS.Roots = roots
	mTS.Blocks = 0
}

// Note: nodes is a pointer to a list and is edited in this function
// so nodes is not returned.
func (mTS *MerkelTreeStream) Next(data *[]uint8, nodes *[]DefaultNode) {
	var index uint
	index = 2 * mTS.Blocks
	mTS.Blocks++

	var tempLeaf Leaf
	tempLeaf = Leaf{}
	tempLeaf.Data = *data

	var tempNodeKind NodeKind
	tempNodeKind = NodeKind{}
	tempNodeKind.Leaf = &tempLeaf

	var leaf PartialNode
	leaf = PartialNode{
		Index:  index,
		Parent: flat.Parent(index),
		Length: uint(len(*data)),
		Data:   tempNodeKind,
	}

	var hash []uint8
	hash = mTS.Handler.Leaf(leaf, mTS.Roots)
	var parts NodeParts
	parts = NodeParts{
		Node: leaf,
		Hash: hash,
	}

	var node DefaultNode
	node = NodeFromParts(parts)

	mTS.Roots = append(mTS.Roots, node)
	*nodes = append(*nodes, node)

	for len(mTS.Roots) > 1 {
		var left, right DefaultNode
		left = mTS.Roots[len(mTS.Roots)-2]
		right = mTS.Roots[len(mTS.Roots)-1]

		if left.Parent != right.Parent {
			break
		}
		mTS.Roots = mTS.Roots[:len(mTS.Roots)-1]

		var hash []uint8
		hash = mTS.Handler.Parent(left, right)

		var tempBool bool
		tempBool = true

		var partial PartialNode
		partial = PartialNode{
			Index:  left.Parent,
			Parent: flat.Parent(left.Parent),
			Length: left.Len() + right.Len(),
			Data:   NodeKind{Parent: &tempBool},
		}

		var tempNode DefaultNode
		tempNode = DefaultNode{}
		tempNode.FromPartial(partial, hash)

		mTS.Roots[len(mTS.Roots)-1] = tempNode
		*nodes = append(*nodes, tempNode)
	}
}

/*

type TreeOpts struct {
	// TODO: This is probably wrong.
	Leaf   func(leaf PartialNode, roots []Node) Hash
	Parent uint
}

func MerkelGenerator(opts uint, roots uint) {

}
*/
