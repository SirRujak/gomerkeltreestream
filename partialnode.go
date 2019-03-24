package gomerkeltreestream

type Leaf struct {
	Data []byte
}
type NodeKind struct {
	Parent *bool
	Leaf   *Leaf
}
type PartialNode struct {
	Parent uint
	Data   NodeKind
	Length uint
	Index  uint
}

func (pN *PartialNode) IsEmpty() bool {
	return pN.Length == 0
}

func (pN *PartialNode) GetIndex() uint {
	return pN.Index
}

func (pN *PartialNode) GetData() NodeKind {
	return pN.Data
}

// Ignoring Deref and DerefMut because I can't figure out what they are for....
