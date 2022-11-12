// Simple balance tree implements
package tree

import (
	"fmt"

	"github.com/liangx8/tree/treeview"
)

// compare two elements
//
// return negative indicate l < r, 0 indicate l == r, positive indicate l > r

type (
	Compare func(l, r interface{}) int
	node    struct {
		e      interface{}
		l, r   *node
		ln, rn int
	}
	Btree interface {
		// add element
		Add(e interface{}) error
		// get element
		Get(e interface{}) (interface{}, error)
		// iterate each element
		// it will be stop and return with error immediately if callback return a error
		Each(callback func(int, interface{}) error) error
		// remove element and return it if found element in btree
		Remove(e interface{}) (interface{}, error)
	}
	btree struct {
		c   Compare
		top *node
	}
	printModel struct {
		n         *node
		asString  func(interface{}) string
		viewWidth func() int
	}
)

func (pm *printModel) ChildCount() int {
	if pm.n.l == nil && pm.n.r == nil {
		return 0
	}
	return 2
}
func (pm *printModel) ChildAt(idx int) treeview.Model {
	if idx == 0 {
		if pm.n.l == nil {
			return nil
		}
		return &printModel{n: pm.n.l, asString: pm.asString, viewWidth: pm.viewWidth}
	}
	if pm.n.r == nil {
		return nil
	}
	return &printModel{n: pm.n.r, asString: pm.asString, viewWidth: pm.viewWidth}
}
func (pm *printModel) ObjectWidth() int {
	return pm.viewWidth()
}
func (pm *printModel) String() string {
	return pm.asString(pm.n.e)
}
func ToModel(bt Btree, fn func(interface{}) string, width int) treeview.Model {
	return &printModel{
		n:         bt.(*btree).top,
		asString:  fn,
		viewWidth: func() int { return width },
	}
}
func (t *btree) Each(cb func(int, interface{}) error) error {
	idx := 0
	if t.top == nil {
		return nil
	}

	return walk(t.top, cb, &idx)
}
func (t *btree) Add(e interface{}) error {
	ntop, err := recurs_add(t.top, e, t.c)
	if err != nil {
		return err
	}
	t.top = ntop
	return nil
}
func (t *btree) Remove(e interface{}) (interface{}, error) {
	ntop, old, err := find_for_remove(t.top, e, t.c)
	t.top = ntop
	return old, err
}
func (t *btree) Get(e interface{}) (interface{}, error) {
	n, err := find(t.top, e, t.c)
	if err != nil {
		return nil, err
	}
	return n.e, nil
}

// cp compare two elements
func New(cp Compare) Btree {
	return &btree{c: cp}
}
func walk(top *node, cb func(int, interface{}) error, idx *int) error {
	if top.l != nil {
		if err := walk(top.l, cb, idx); err != nil {
			return err
		}
	}

	if err := cb(*idx, top.e); err != nil {
		return err
	}
	*idx++
	if top.r != nil {
		if err := walk(top.r, cb, idx); err != nil {
			return err
		}
	}

	return nil
}

func find(top *node, e interface{}, c Compare) (*node, error) {
	if top == nil {
		return nil, NoFound
	}
	cp := c(top.e, e)
	if cp == 0 {
		return top, nil
	}
	if cp > 0 {
		return find(top.l, e, c)
	}
	return find(top.r, e, c)
}

/*
func max(x,y int) int {
	if x>y {return x} else { return y}
}
*/
func calc_depth(n *node) int {
	if n == nil {
		return 0
	}
	if n.ln > n.rn {
		return n.ln + 1
	}
	return n.rn + 1
}
func recurs_add(top *node, e interface{}, cp Compare) (*node, error) {
	if top == nil {
		return &node{e: e}, nil
	}
	cpv := cp(top.e, e)
	if cpv > 0 {
		newl, err := recurs_add(top.l, e, cp)
		if err != nil {
			return nil, err
		}
		top.l = newl
		top.ln = calc_depth(top.l)
	} else {
		// right site
		newr, err := recurs_add(top.r, e, cp)
		if err != nil {
			return nil, err
		}
		top.r = newr
		top.rn = calc_depth(top.r)
	}
	return doBalance(top), nil
}
func doBalance(top *node) *node {
	bal := top.ln - top.rn
	switch bal {
	case -2:
		// left < right
		rcp := top.r.ln - top.r.rn
		switch rcp {
		case 1:
			ntop := rotate_rl(top)
			return ntop
		case 0, -1:
			ntop := rotate_left(top)
			return ntop
		default:
			panic("left < right rotate error(unexpect error)")
		}
	case 2:
		// left > right
		lcp := top.l.ln - top.l.rn
		switch lcp {
		case -1:
			ntop := rotate_lr(top)
			return ntop
		case 0, 1:
			ntop := rotate_right(top)
			return ntop
		default:
			panic("left > right rotate error(unexpect error)")
		}
	case 0, 1, -1:
	default:
		panic("Not a balance tree(Unexpect error)")
	}
	return top
}

/*
     b                 d
    / \               / \
   a   d             b   e
   |  / \    ====>  / \  |
   ? c   e         a   c ?
     |   |         |   |
     ?   ?         ?   ?
*/

func rotate_left(top *node) *node {
	if top == nil || top.r == nil {
		panic("right node should not nil for left rotation!")
	}
	b := top
	c := top.r.l
	d := top.r

	b.r = c
	b.rn = calc_depth(b.r)
	d.l = b
	d.ln = calc_depth(d.l)
	return d
}

/*
      d               b
     / \             / \
    b   e           a   d
   / \  |  ====>    |  / \
  a   c ?           ? c   e
  |   |               |   |
  ?   ?               ?   ?
*/

func rotate_right(top *node) *node {
	if top == nil || top.l == nil {
		panic("left node should not nil for right rotation")
	}
	d := top
	b := top.l
	c := top.l.r

	d.l = c
	d.ln = calc_depth(d.l)
	b.r = d
	b.rn = calc_depth(b.r)
	return b
}
func rotate_lr(top *node) *node {
	if top == nil || top.l == nil || top.l.r == nil {
		panic("not fulfill left than right rotation")
	}
	top.l = rotate_left(top.l)
	top.ln = calc_depth(top.l)
	return rotate_right(top)
}
func rotate_rl(top *node) *node {
	if top == nil || top.r == nil || top.r.l == nil {
		panic("not fulfill right than left rotation")
	}
	top.r = rotate_right(top.r)
	top.rn = calc_depth(top.r)
	return rotate_left(top)
}

var (
	NoFound = fmt.Errorf("Element no found")
)
