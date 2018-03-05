package tree


import (
	"testing"
	"math/rand"
)
func intcmp(l,r interface{})int {
	return l.(int)-r.(int)
}

func TestTree(t *testing.T){
	bt:=New(intcmp)
	for i:=0;i<100000;i++{
		bt.Add(rand.Int() % 1000)
	}
	check_depth(bt.(*btree).top,t)
}

func check_depth(n *node, t *testing.T) int{
	if n == nil {
		return 0
	}
	if check_depth(n.l,t) != n.ln {
		t.Fatal("depth is incorrected")
	}
	if check_depth(n.r,t) != n.rn {
		t.Fatal("depth is incorrected")
	}
	switch(n.rn-n.ln){
	case -1,0,1:
	default:
		t.Fatal("Lost balance")
	}
	
	if n.rn > n.ln {
		return n.rn+1
	}
	return n.ln+1
}
