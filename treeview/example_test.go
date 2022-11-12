package treeview_test

import (
	"os"

	"github.com/liangx8/tree/treeview"
)

func ExamplePrintTree() {
	treeview.PrintTree(os.Stdout, treeview.NewSampleModel(5))
}
