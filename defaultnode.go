package gomerkeltreestream

type DefaultNode struct {
	Parent uint
	Data   *[]byte
	Hash   []byte
	Length uint
	Index  uint
}

func (dN *DefaultNode) FromPartial(partial PartialNode, hash []byte) {
	dN.Index = partial.Index
	dN.Length = partial.Length
	dN.Hash = hash
	dN.Parent = partial.Parent

	var data NodeKind
	data = partial.GetData()
	if data.Leaf != nil {
		dN.Data = &data.Leaf.Data
	} else {
		dN.Data = nil
	}
}

func (dN *DefaultNode) GetHash() []byte {
	return dN.Hash
}

func (dN *DefaultNode) Len() uint {
	return dN.Length
}

func (dN *DefaultNode) IsEmpty() bool {
	return dN.Length == 0
}

func (dN *DefaultNode) GetIndex() uint {
	return dN.Index
}

func (dN *DefaultNode) GetParent() uint {
	return dN.Parent
}

func NodeFromParts(nodeParts NodeParts) DefaultNode {
	var tempDefaultNode DefaultNode
	tempDefaultNode = DefaultNode{}
	tempDefaultNode.FromPartial(nodeParts.Node, nodeParts.Hash)
	return tempDefaultNode
}
