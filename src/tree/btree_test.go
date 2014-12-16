package tree

import (
	"testing"
)
func intcmp(l,r interface{})int {
	return l.(int)-r.(int)
}

func Test_unique_tree(t *testing.T){
	bt:=New(intcmp,false)
	el,err:=bt.Add(1)
	if el != nil || err != nil {
		t.Fatalf("add a new element but ruturn value or error %v",err)
	}
	el,err=bt.Add(0)
	if el != nil || err != nil {
		t.Fatalf("add a new element but ruturn value or error %v",err)
	}
	el,err=bt.Add(-1)
	if el != nil || err != nil {
		t.Fatalf("add a new element but ruturn value or error %v",err)
	}
	el,err=bt.Add(1)
	if el == nil {
		t.Fatalf("add a duplicate element and expect return value but nothing return")
	}
	if el.(int) != 1 {
		t.Fatalf("expect return value is 1 but %d",el)
	}
	if err != Duplicate {
		t.Fatal(err)
	}
	el,err=bt.Add(110)
	if el != nil || err != nil {
		t.Fatalf("add a new element but ruturn value or error %v",err)
	}
	bt.Each(func(e interface{})error{
		t.Logf("%d,",e)
		return nil
	})
	el,err=bt.Get(110)
	if err != nil {
		t.Fatalf("Can't found a exists element")
	}
	el,err=bt.Get(100)
	if err != NoFound {
		t.Fatalf("try to find a not exists element but return is not %v",NoFound)
	}

	el,err = bt.Remove(1)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_duplicate(t *testing.T){
	bt:=New(intcmp,true)
	el,err:=bt.Add(1)
	if el != nil || err != nil {
		t.Fatalf("add a new element but ruturn value or error %v",err)
	}
	el,err=bt.Add(1)
	if err != nil {
		t.Fatal(err)
	}
	if el != nil {
		t.Fatalf("duplicate tree expect no return for duplicate element add")
	}
	bt.Each(func(e interface{})error{
		t.Logf("%d,",e)
		return nil
	})
}
