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
	return 2
}
func (m *sampleModel) ObjectStr(o interface{}) string {
	if o == (*sampleModel)(nil) {
		return "**"
	}
	return o.(*sampleModel).name
}
func (m *sampleModel) Root() interface{} {
	return m
}
func (m *sampleModel) ChildCount(parent interface{}) int {
	if parent==(*sampleModel)(nil) {return 0}
	return len(parent.(*sampleModel).children)
}
func (m *sampleModel) Child(parent interface{}, idx int) interface{} {
	return parent.(*sampleModel).children[idx]
}

// Create a Sample model

func NewSampleModel(lvl int) Model {
	if lvl > 9 {lvl = 9}
	sm := &sampleModel{name: "RT", children: nil}
	var build func(*sampleModel,rune,int)
	build=func(parent *sampleModel,prefix rune,level int){
		if level == 0 { return }
		cnt:=rand.Int() % 5
		if cnt < 0 { cnt = -cnt}
		cnt ++
		parent.children =make([]*sampleModel,cnt)
		fmt.Printf("count:%d\n",cnt)
		for i,_ := range parent.children {
			r := rand.Int() % 20
			fmt.Println(r)
			if r < 19 {
				parent.children[i]=&sampleModel{name:fmt.Sprintf("%c%d",prefix,i),children:nil}
				build(parent.children[i],prefix + 1,level -1)
			}
		}
	}
	build(sm,'a',lvl)

	return sm
}
