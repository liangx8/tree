package treeview_test

import (
	"github.com/liangx8/tree/treeview"
	"os"
)

func ExamplePrintTree() {
	treeview.PrintTree(os.Stdout, treeview.NewSampleModel())
}
