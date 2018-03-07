package treeview

import (
	"fmt"
	"io"
	"strings"
)

// w output stream
//
// ml model
//
func PrintTree(w io.Writer, ml Model) {
	printWithPrefix(w,ml,"")
}

func printWithPrefix(w io.Writer,tm Model, prefix string){
	if tm == nil {
		fmt.Fprintln(w,"=")
		return
	}
	if tm == nil {panic("impossible")}
	fmt.Fprint(w,tm)
	cnt := tm.ChildCount()
	if cnt == 0 {
		fmt.Fprintln(w,"$")
		return
	} else {
		fmt.Fprint(w,"-+-")
	}
	sb1 := prefix + strings.Repeat(" ",tm.ObjectWidth() + 1)
	sb2 := sb1 + "  "
	ssb := sb1 + "+-"
	sb1 = sb1 + "| "
	for i := 0; i < cnt;i++ {
		if i == cnt -1 {
			printWithPrefix(w,tm.ChildAt(i),sb2)
		} else {
			printWithPrefix(w,tm.ChildAt(i),sb1)
			fmt.Fprint(w,ssb)
		}
	}
}
