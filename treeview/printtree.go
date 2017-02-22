package treeview
import (
	"io"
	"fmt"
)
// w output stream
//
// ml model
//
func PrintTree(w io.Writer,ml Model){
	root:=ml.Root()
	row:=0
	var pt func(io.Writer,interface{},Model,[]int)
	pt=func(w io.Writer,parent interface{},ml Model,nosh []int){
		fmt.Fprint(w,ml.ObjectStr(parent))
		cols := ml.ChildCount(parent)
		if cols == 0 {
			fmt.Fprintln(w,"$")
			return;
		}
		fmt.Fprint(w,"-+-")
		row++
		ai:=nosh
		for i:=0;i<cols;i++ {
			if i== cols-1 {
				ai=append(ai,row-1)
			}
			pt(w,ml.Child(parent,i),ml,ai)
			if i < cols-1 {
				for x:=0;x<row-1;x++ {
					for j:=0;j<ml.ObjectWidth();j++ {
						fmt.Fprint(w," ")
					}
					if exists(ai,x){
						fmt.Fprint(w,"   ")
					} else {
						fmt.Fprint(w," | ")
					}
				}
				for j:=0;j<ml.ObjectWidth();j++ {
					fmt.Fprint(w," ")
				}
				fmt.Fprint(w," +-")
			}
		}
		row--
	}
	pt(w,root,ml,make([]int,0))
	fmt.Fprintln(w)
}
func exists(a []int, v int) bool{
	for _,i:=range a {
		if i == v { return true}
	}
	return false
}
