//  treeview is package for displaying tree in text console
package treeview

import (
	"math/rand"
	"fmt"
)

// Model is tree model for display tree
type Model interface {
	// root node
	Root() interface{}
	// how many childrens under parent node
	ChildCount(parent interface{}) int
	// child of parent at idx
	Child(parent interface{}, idx int) interface{}
	// width of object
	ObjectWidth() int
	// how to convert object to string
	ObjectStr(interface{}) string
}

type sampleModel struct {
	name   string
	children []*sampleModel
}

func (m *sampleModel) ObjectWidth() int {
	return 4
}
func (m *sampleModel) ObjectStr(o interface{}) string {
	return o.(*sampleModel).name
}
func (m *sampleModel) Root() interface{} {
	return m
}
func (m *sampleModel) ChildCount(parent interface{}) int {
	return len(parent.(*sampleModel).children)
}
func (m *sampleModel) Child(parent interface{}, idx int) interface{} {
	return parent.(*sampleModel).children[idx]
}

// Create a Sample model
func NewSampleModel() Model {
	sm := &sampleModel{name: "root", children: nil}
	sm.children = make([]*sampleModel, 4)
	sm.children[0] = &sampleModel{name: "[c1]", children: make([]*sampleModel, 3)}
	sm.children[1] = &sampleModel{name: "[c2]", children: nil}
	sm.children[2] = &sampleModel{name: "[c3]", children: make([]*sampleModel, 2)}
	sm.children[3] = &sampleModel{name: "[c4]", children: make([]*sampleModel, 2)}

	sm.children[0].children[0] = &sampleModel{name: "[k1]", children: nil}
	sm.children[0].children[1] = &sampleModel{name: "[k2]", children: nil}
	sm.children[0].children[2] = &sampleModel{name: "[k3]", children: nil}

	sm.children[0].children[0].children = make([]*sampleModel, 2)
	sm.children[0].children[0].children[0] = &sampleModel{name: "[x1]", children: nil}
	sm.children[0].children[0].children[1] = &sampleModel{name: "[x2]", children: nil}

	sm.children[0].children[2].children = make([]*sampleModel, 4)
	sm.children[0].children[2].children[0] = &sampleModel{name: "[x3]", children: nil}
	sm.children[0].children[2].children[1] = &sampleModel{name: "[x4]", children: nil}
	sm.children[0].children[2].children[2] = &sampleModel{name: "[x5]", children: nil}
	sm.children[0].children[2].children[3] = &sampleModel{name: "[x6]", children: nil}

	sm.children[2].children[0] = &sampleModel{name: "[k4]", children: nil}
	sm.children[2].children[1] = &sampleModel{name: "[k5]", children: nil}
	sm.children[3].children[0] = &sampleModel{name: "[k6]", children: nil}
	sm.children[3].children[1] = &sampleModel{name: "[k7]", children: nil}

	return sm
}
func RandSampleModel() Model {
	sm := &sampleModel{name: "root", children: nil}
	var build func(*sampleModel,rune,int,int)
	build=func(parent *sampleModel,prefix rune,cnt,level int){
		if level == 0 { return }
		parent.children =make([]*sampleModel,cnt)
		fmt.Printf("%d,",len(parent.children))

		for i,_ := range sm.children {
			sm.children[i]=&sampleModel{name:fmt.Sprintf("[%c%d]",prefix,i),children:nil}
			build(sm.children[i],prefix + 1,rand.Int() %6,level -1)
		}
	}
	build(sm,'a',rand.Int() % 6+4,4)
	
	return sm
}
