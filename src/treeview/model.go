//  treeview is package for displaying tree in text console
package treeview
// Model is tree model for display tree
type Model interface{
	// root node
	Root() interface{}
	// how many childrens under parent node
	ChildCount(parent interface{}) int
	// child of parent at idx
	Child(parent interface{},idx int) interface{}
	// width of object
	ObjectWidth() int
	// how to convert object to string
	ObjectStr(interface{}) string
}


type sampleModel struct{
	name string
	childs []*sampleModel
}
func (m *sampleModel)ObjectWidth() int{
	return 4
}
func (m *sampleModel)ObjectStr(o interface{})string {
	return o.(*sampleModel).name
}
func (m *sampleModel)Root() interface{}{
	return m
}
func (m *sampleModel)ChildCount(parent interface{}) int{
	return len(parent.(*sampleModel).childs)
}
func (m *sampleModel)Child(parent interface{},idx int) interface{} {
	return parent.(*sampleModel).childs[idx]
}
// Create a Sample model
func NewSampleModel() Model {
	sm:=&sampleModel{name:"root",childs:nil}
	sm.childs=make([]*sampleModel,4)
	sm.childs[0]=&sampleModel{name:"[c1]",childs:make([]*sampleModel,3)}
	sm.childs[1]=&sampleModel{name:"[c2]",childs:nil}
	sm.childs[2]=&sampleModel{name:"[c3]",childs:make([]*sampleModel,2)}
	sm.childs[3]=&sampleModel{name:"[c4]",childs:make([]*sampleModel,2)}

	sm.childs[0].childs[0]=&sampleModel{name:"[k1]",childs:nil}
	sm.childs[0].childs[1]=&sampleModel{name:"[k2]",childs:nil}
	sm.childs[0].childs[2]=&sampleModel{name:"[k3]",childs:nil}

	sm.childs[0].childs[0].childs=make([]*sampleModel,2)
	sm.childs[0].childs[0].childs[0]=&sampleModel{name:"[x1]",childs:nil}
	sm.childs[0].childs[0].childs[1]=&sampleModel{name:"[x2]",childs:nil}

	sm.childs[0].childs[2].childs=make([]*sampleModel,4)
	sm.childs[0].childs[2].childs[0]=&sampleModel{name:"[x3]",childs:nil}
	sm.childs[0].childs[2].childs[1]=&sampleModel{name:"[x4]",childs:nil}
	sm.childs[0].childs[2].childs[2]=&sampleModel{name:"[x5]",childs:nil}
	sm.childs[0].childs[2].childs[3]=&sampleModel{name:"[x6]",childs:nil}

	sm.childs[2].childs[0]=&sampleModel{name:"[k4]",childs:nil}
	sm.childs[2].childs[1]=&sampleModel{name:"[k5]",childs:nil}
	sm.childs[3].childs[0]=&sampleModel{name:"[k6]",childs:nil}
	sm.childs[3].childs[1]=&sampleModel{name:"[k7]",childs:nil}

	return sm
}
