package tree_test

import (
	"testing"
	"math/rand"

	"github.com/liangx8/tree"
)
func intcmp(l,r interface{})int {
	return l.(int)-r.(int)
}

func BenchmarkTree(b *testing.B){
	bt:=tree.New(intcmp)
	for i:=0;i<b.N;i++ {
		bt.Add(rand.Int() % 1000)
	}
}
