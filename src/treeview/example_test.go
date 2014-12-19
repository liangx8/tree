package treeview_test
import (
	"os"
	"treeview"
)
func ExamplePrintTree() {
	treeview.PrintTree(os.Stdout,treeview.NewSampleModel())
}
