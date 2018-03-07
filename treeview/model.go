//  treeview is package for displaying tree in text console
package treeview

import (
	"fmt"
	"math/rand"
)

// Model is tree model for display tree
type (
	Model interface {
		// how many childrens under parent node
		ChildCount() int
		// child of parent at idx
		ChildAt(int) Model
		// width of object
		ObjectWidth() int
	}
	sampleModel struct {
		name     string
		children []*sampleModel
	}
)

func (m *sampleModel) ChildCount() int {

	if m.children == nil {
		return 0
	}
	return len(m.children)
}
func (m *sampleModel) ChildAt(idx int) Model {
	v := m.children[idx]
	if v == nil {
		return nil
	}
	return v
}
func (m *sampleModel) ObjectWidth() int {
	return 3
}
func (m *sampleModel) String() string {
	return m.name
}

func createSampleModel(namePrefix rune, value, lvl int) *sampleModel {
	if lvl == 0 {
		return nil
	}
	sm := sampleModel{name: fmt.Sprintf("%c%02d", namePrefix, value)}
	cnt := rand.Int() % 5
	if cnt < 0 {
		cnt = -cnt
	}
	cnt++
	if lvl > 1 {
		sm.children = make([]*sampleModel, cnt)

		for i, _ := range sm.children {
			r := rand.Int() % 12
			if r < 10 {
				sm.children[i] = createSampleModel(namePrefix+1, i, lvl-1)
			}

		}
	}

	return &sm
}
func NewSampleModel(level int) Model {
	return createSampleModel('a', 0, level)
}
