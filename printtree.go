package tree
import (
	"fmt"
	"io"

	"github.com/liangx8/tree/treeview"
)
type model struct{
	getRoot func() interface{}
	o2s func(interface{}) string
	width int
}
func (m *model)Root() interface{}{
	return m.getRoot()
}
func (m *model)ChildCount(parent interface{}) int{
	if parent == nil {return 0 }
	n:=parent.(*node)
	if n.l != nil { return 2 }
	if n.r != nil { return 2 }
	return 0
}
func (m *model)Child(parent interface{},idx int) interface{}{
	n:=parent.(*node)
	var rtn *node
	switch idx {
	case 0: rtn = n.l
	case 1: rtn = n.r
	default:
		panic("impossible")
	}
	if rtn == nil {
		return &node{}
	}
	return rtn
}
func (m *model)ObjectWidth() int{
	return m.width
}
func (m *model)ObjectStr(o interface{}) string {
	n:=o.(*node)
	return m.o2s(n.e)
}
// CreateModel creates a model for treeview.PrintTree to used
//
// bt a Btree will be printout
//
// fn how to print object
//
// w width of object
func CreateModel(bt Btree,fn func(interface{})string,w int) (treeview.Model,error){
	btr,ok:=bt.(*btree)
	if !ok {
		return nil,fmt.Errorf("I don't know how to create a model")
	}
	return &model{
	getRoot:func() interface{}{
			if btr.top == nil {
				return &node{}
			}
			return btr.top
		},
	o2s:fn,
	width:w,
	},nil
}
// PrintTree print out btree
//
// w output stream
//
// size size of string for returning of str()
//
// str convert element contains in btree to a string for display
//
// Note: a mvc implementation for this function, refer to 
// package treeview.PrintTree
func PrintTree( w io.Writer, bt Btree,size int, str func(interface{})string){
	var total,height int
	tree,ok := bt.(*btree)
	if !ok {
		fmt.Fprintf(w,"Unknow type %T",bt)
		return
	}
	t := tree.top
	height=calc_depth(t)
	col:=0
	var pt func(io.Writer,*node,func(interface{})string,[]int)
	pt=func(w io.Writer,top *node,str func(interface{})string,noshow []int){
		if top == nil {
			fmt.Fprintln(w,"$")
			return
		}
		val:=str(top.e)
		col++
		fmt.Fprint(w,val)
		total++
		if top.l == nil && top.r == nil {
			fmt.Fprintln(w,"=")
			col --
			return
		}
		fmt.Fprint(w,"-+-")
		pt(w,top.l,str,noshow)
		for i:=0;i<col-1;i++ {
			for j:=0;j<size;j++ {
				fmt.Fprint(w," ")
			}
			if exists(i,noshow) {
				fmt.Fprint(w,"   ")
			} else {
				fmt.Fprint(w," | ")
			}
		}
		for i:=0;i<size;i++ {
			fmt.Fprintf(w," ")
		}
		fmt.Fprint(w," +-")
		pt(w,top.r,str,append(noshow,col-1))
		col--
	}
	pt(w,t,str,make([]int,0))
	fmt.Fprintf(w,"height: %d\n",height)
	fmt.Fprintf(w,"size  : %d\n",total)

}
func exists(v int, b []int) bool {
	for _,i := range b {
		if i==v { return true }
	}
	return false
}
