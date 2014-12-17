package tree
//import (
//"fmt"
//)

func find_for_remove(top *node,e interface{},c Compare) (*node,interface{},error){
	var old interface{}
	var err error
	if top == nil {
		return top,nil,NoFound
	}
	cp:=c(top.e,e)
	if cp==0 {
		// found it, and remove
		ntop,err:=rm1(top)
		return ntop, top.e,err
	}
	if cp>0 {
		var lnode *node
		lnode,old,err=find_for_remove(top.l,e,c)
		top.l=lnode
	}
	if cp<0 {
		var rnode *node
		rnode,old,err=find_for_remove(top.r,e,c)
		top.r=rnode
	}
	if top.l==nil {top.ln=0} else {top.ln=max(top.l.ln,top.l.rn)+1}
	if top.r==nil {top.rn=0} else {top.rn=max(top.r.ln,top.r.rn)+1}
	top,err=doBalance(top)
	return top,old,err
}
/*

*/
func rm1(nd *node) (*node,error){
	var ntop *node
	var err error
	if nd.ln>nd.rn {
		var top *node
		ntop,top,err=rm_most_right(nd.l)
		if err != nil { return nd,nil }
		if ntop == nil { return nil,nil } // is terminal node
		ntop.l=top
		ntop.r=nd.r
	} else {
		var top *node
//fmt.Println(nd.r.e,nd.r.l,nd.r.r)
		ntop,top,err=rm_most_left(nd.r)
		if err != nil { return nd,nil }
		if ntop == nil { return nil,nil } // is terminal node
		ntop.r=top
		ntop.l=nd.l
	}
	if ntop.l==nil {ntop.ln=0} else {ntop.ln=max(ntop.l.ln,ntop.l.rn)+1}
	if ntop.r==nil {ntop.rn=0} else {ntop.rn=max(ntop.r.ln,ntop.r.rn)+1}
//fmt.Println(ntop.e,ntop.l,ntop.r)
/*
v1,v2:=doBalance(ntop)
fmt.Println(v1.e,v2)
return v1,v2
*/
	return doBalance(ntop)
}
/*
 查找最右边的节点, 该节点将会被放置在被删除的位置
 most_r, 将是新的root
 npn 是取出节点后。需要重新平衡后的当前树的root
*/
func rm_most_right(pn *node) (most_r,npn *node,err error) {
	if pn== nil {
		return nil,nil,nil // empty tree
	}
	if pn.r == nil {
		return pn,pn.l,nil
	}
	most_r,npn,err=rm_most_right(pn.r)
	if err != nil {
		return nil,nil,err
	}
	if most_r == nil {
		panic("for debug, never reach here (remove.go,rm_most_right)")
		//return nil,pn,nil
	}
	if npn == nil {
		pn.r,pn.rn=nil,0
	}else {
		pn.r,pn.rn=npn,max(npn.ln,npn.rn)+1
	}
	// balances pn, and let npn= new top
	npn,err=doBalance(pn)
	return
}
func rm_most_left(pn *node) (most_l,npn *node,err error) {
	if pn== nil {
		return nil,nil,nil // empty tree
	}
	if pn.l == nil {
		return pn,pn.r,nil
	}
	most_l,npn,err=rm_most_left(pn.l)
	if err != nil {
		return nil,nil,err
	}
	if most_l == nil {
		panic("for debug, never reach here (remove.go,rm_most_left)")
		//return nil,pn,nil
	}
	if npn == nil {
		pn.l,pn.ln=nil,0
	}else {
		pn.l,pn.ln=npn,max(npn.ln,npn.rn)+1
	}
	// balances pn, and let npn= new top
	npn,err=doBalance(pn)
	return
}
