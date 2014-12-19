package main
import (
	"fmt"
	"os"
//	"time"
	"math/rand"

	"tree"
	"treeview"
)
func comp(l,r interface{}) int {
	return l.(int)-r.(int)
}
const MAX=200
func main(){
	test3()
}
func test3(){
	bi:=make([]int,MAX)
	for i:=0;i<MAX;i++ {
		bi[i]=i
	}
	for i:= range bi {
		n:=rand.Int() % MAX
		bi[i],bi[n]=bi[n],bi[i]
	}

	bt:=tree.New(comp,false)
	for _,i := range bi {
		bt.Add(i)
	}
	fn:=func(v interface{})string{
		if v == nil { return "" }

		return fmt.Sprintf("%03d",v)
	}
	model,err:=tree.CreateModel(bt,fn,3)
	if err != nil {
		fmt.Println(err)
	}
	treeview.PrintTree(os.Stdout,model)
}
func test2(){
	m:=treeview.NewSampleModel()
	treeview.PrintTree(os.Stdout,m)
}
func test1(){
	bi:=make([]int,MAX)
	for i:=0;i<MAX;i++ {
		bi[i]=i
	}
	for i:= range bi {
		n:=rand.Int() % MAX
		bi[i],bi[n]=bi[n],bi[i]
	}

	bt:=tree.New(comp,false)
	for _,i := range bi {
		bt.Add(i)
	}
	fn:=func(v interface{})string{
		return fmt.Sprintf("%03d",v)
	}
	tree.PrintTree(os.Stdout,bt,3,fn)
}
