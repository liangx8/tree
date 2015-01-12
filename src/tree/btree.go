// Simple balance tree implements
package tree

import (
	"fmt"
)
// compare two elements
//
// return negative indicate l < r, 0 indicate l == r, positive indicate l > r
type Compare func(l,r interface{}) int
type node struct {
	e interface{}
	l,r *node
	ln,rn int
}
type Btree interface{
	// add element
	Add(e interface{}) (interface{},error)
	// get element
	Get(e interface{}) (interface{},error)
	// iterate each element
	// it will be stop and return with error immediately if callback return a error
	Each(callback func(interface{})error) error
	// remove element and return it if found element in btree
	Remove(e interface{}) (interface{},error)
}
type btree struct{
	c Compare
	allowDup bool
	top *node
}
func (t *btree)Each(cb func(interface{})error)error {
	if t.top == nil {
		return nil
	}
	return walk(t.top,cb)
}
func (t *btree)Add(e interface{}) (interface{},error) {
	ntop,old,err:=recurs_add(t.top,e,t.c,t.allowDup)
	if err != nil {
		return old,err
	}
	t.top=ntop
	return nil,nil
}
func (t *btree)Remove(e interface{}) (interface{},error){
	ntop,old,err:=find_for_remove(t.top,e,t.c)
	t.top=ntop
	return old,err
}
func (t *btree)Get(e interface{}) (interface{},error) {
	n,err:=find(t.top,e,t.c)
	if err != nil {
		return nil,err
	}
	return n.e,nil
}
// allowDuplicate will effect method Add,
// true, add a new element whatever it does or not exists in btree,and return (nil,Duplicate) 
//
// false replace the old one, return old element and warnning Duplicate
//
// c compare two elements
func New(c Compare,allowDuplicate bool) Btree{
	return &btree{c:c,allowDup:allowDuplicate}
}
func walk(top *node,cb func(interface{})error) error {
	if top.l != nil {
		if err:=walk(top.l,cb); err != nil {
			return err
		}
	}
	if err:=cb(top.e); err != nil {
		return err
	}
	if top.r != nil {
		if err:=walk(top.r,cb); err != nil {
			return err
		}
	}
	return nil
}
func find(top *node,e interface{},c Compare) (*node,error){
	if top == nil {
		return nil,NoFound
	}
	cp := c(top.e,e)
	if cp == 0 {
		return top,nil
	}
	if cp > 0 {
		return find(top.l,e,c)
	}
	return find(top.r,e,c)
}
func max(x,y int) int {
	if x>y {return x} else { return y}
}
func recurs_add(top *node,e interface{},c Compare,alwDup bool) (*node,interface{},error){
	var err error
	var old interface{}
	if top == nil {
		return &node{e:e},nil,nil
	}
	cp:=c(top.e,e)
	if cp==0 {
		if !alwDup {
			old=top.e
			top.e=e
			return top,old,Duplicate
		}
		var newl *node
		newl,old,err=recurs_add(top.l,e,c,alwDup)
		top.l=newl
		top.ln=max(top.l.ln,top.l.rn)+1
	}
	if cp>0 {
		var newl *node
		newl,old,err=recurs_add(top.l,e,c,alwDup)
		top.l=newl
		top.ln=max(top.l.ln,top.l.rn)+1
	}
// right site
	if cp<0 {
		var newr *node
		newr,old,err=recurs_add(top.r,e,c,alwDup)
		top.r=newr
		top.rn=max(top.r.ln,top.r.rn)+1
	}
	ttt,nerr:=doBalance(top)
	if nerr != nil {
		return ttt,old,nerr
	} else {
		return ttt,old,err
	}
}
func doBalance(top *node) (*node,error) {
	bal:=top.ln-top.rn
	switch bal {
	case -2:
	// left < right
		rcp:=top.r.ln-top.r.rn
		switch rcp {
		case 1:
			ntop:=rotate_rl(top)
			return ntop,nil
		case 0,-1:
			ntop:=rotate_left(top)
			return ntop,nil
		default:
			return top,fmt.Errorf("left < right rotate error(unexpect error)")
		}
	case 2:
	// left > right
		lcp:=top.l.ln-top.l.rn
		switch lcp{
		case -1:
			ntop:=rotate_lr(top)
			return ntop,nil
		case 0,1:
			ntop:=rotate_right(top)
			return ntop,nil
		default:
			return top,fmt.Errorf("left > right rotate error(unexpect error)")
		}
	case 0,1,-1:
	default:
		return top,fmt.Errorf("Not a balance tree(Unexpect error)")
	}
	return top,nil
}

/*
       b                d                    1                3
      / \              / \                  / \              / \
     a   d      =>    b   e                0   3      =>    1   4
        / \          / \  |                   / \          / \
       c   e        a   c ?                  2   4        0   2
           |
           ?
*/
func rotate_left(top *node) *node {
	if top.r == nil {
		panic("right node should not nil in rotate left")
	}
	b:=top
	d:=top.r
	c:=top.r.l

	d.l=b
	b.r=c
	if c== nil {
		b.rn=0
	} else {
		b.rn=max(c.ln,c.rn)+1
	}
	d.ln=max(b.ln,b.rn)+1
	return d
}
/*
       d                   b
      / \                 / \
     b   e        =>     a   d
    / \                  |  / \
   a   c                 ? c   e
   |
   ?
*/
func rotate_right(top *node) *node {
	if top.l == nil {
		panic("left node should not nil in rotate right")
	}
	d:=top
	b:=top.l
	c:=top.l.r

	b.r=d
	d.l=c
	if c== nil {
		d.ln=0
	} else {
		d.ln=max(c.ln,c.rn)+1
	}
	b.rn=max(d.ln,d.rn)+1
	return b
}
/*

       f              f                 d
      / \            / \              /   \
     b   g     =>   d   g   =>       b     f
    / \            / \              / \   / \
   a   d          b   e            a   c e   g
      / \        / \
     c   e      a   c

*/
func rotate_lr(top *node) *node{

	f:=top
	b:=top.l
	d:=top.l.r
	c:=top.l.r.l
	e:=top.l.r.r

	d.l=b
	d.r=f
	b.r=c
	f.l=e
	if c== nil {
		b.rn=0
	} else {
		b.rn=max(c.ln,c.rn)+1
	}
	if e==nil {
		f.ln=0
	} else {
		f.ln=max(e.ln,e.rn)+1
	}
	d.ln=max(b.ln,b.rn)+1
	d.rn=max(f.ln,f.rn)+1

	return d
}
/*
       b              b                 d
      / \            / \              /   \
     a   f     =>   a   d   =>       b     f
        / \            / \          / \   / \
       d   g          c   f        a   c e   g
      / \                / \
     c   e              e   g

*/
func rotate_rl(top *node) *node{
	b:=top
	f:=top.r
	d:=top.r.l
	c:=top.r.l.l
	e:=top.r.l.r

	d.l=b
	d.r=f
	b.r=c
	f.l=e
	if c== nil {
		b.rn=0
	} else {
		b.rn=max(c.ln,c.rn)+1
	}
	if e==nil {
		f.ln=0
	} else {
		f.ln=max(e.ln,e.rn)+1
	}
	d.ln=max(b.ln,b.rn)+1
	d.rn=max(f.ln,f.rn)+1

	return d
}

var (
	Duplicate = fmt.Errorf("Add a duplicated element")
	NoFound= fmt.Errorf("Element no found")
)

