package main
import (
	"fmt"
	"os"
//	"time"
	"math/rand"

	"tree"
)
func comp(l,r interface{}) int {
	return l.(int)-r.(int)
}
const MAX=10
func main(){
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
