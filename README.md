tree
====

a simple balance tree implement

tree package is super simple balance tree implementation.

	
	bt := tree.New(func(l,r interface{})int { return l.(int)-r.(int) })
	for i:=0;i<10;i++ {
		bt.Add(rand.Int()%1000)
	}
	tree.PrintTree(os.Stdout,bt,3,func(e interface{}) string{
		return fmt.Sprintf("%03d",e)
	}

package treeview Sample:


	bi:=make([]int,MAX)
	for i:=0;i<MAX;i++ {
		bi[i]=i
	}
	for i:= range bi {
		n:=rand.Int() % MAX
		bi[i],bi[n]=bi[n],bi[i]
	}

	bt:=tree.New(comp,false)
	for _,i := range bi {
		bt.Add(i)
	}
	fn:=func(v interface{})string{
		if v == nil { return "" }

		return fmt.Sprintf("%03d",v)
	}
	model,err:=tree.CreateModel(bt,fn,3)
	if err != nil {
		fmt.Println(err)
	}
	treeview.PrintTree(os.Stdout,model)
