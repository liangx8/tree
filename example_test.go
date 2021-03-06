package tree_test

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/liangx8/tree"
	"github.com/liangx8/tree/treeview"
)

func ExampleCreateModel() {
	ai := make([]int, 10)
	for i := 0; i < 10; i++ {
		ai[i] = i
	}
	for i := range ai {
		n := rand.Int() % 10
		ai[i], ai[n] = ai[n], ai[i]
	}

	bt := tree.New(func(l, r interface{}) int { return l.(int) - r.(int) })
	for _, i := range ai {
		bt.Add(i)
	}
	fn := func(v interface{}) string {
		if v == nil {
			return ""
		}

		return fmt.Sprintf("%03d", v)
	}
	model := tree.ToModel(bt, fn, 3)
	treeview.PrintTree(os.Stdout, model)
}
