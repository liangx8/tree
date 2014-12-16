package tree
import (
	"fmt"
	"io"
)
// PrintTree print out btree
//
// w output stream
//
// size size of string for returning of str()
//
// str convert element contains in btree to a string for display
func PrintTree( w io.Writer, bt Btree,size int, str func(interface{})string){
	tree,ok := bt.(*btree)
	if !ok {
		fmt.Fprintf(w,"Unknow type %T",bt)
		return
	}
	t := tree.top
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

}
func exists(v int, b []int) bool {
	for _,i := range b {
		if i==v { return true }
	}
	return false
}
