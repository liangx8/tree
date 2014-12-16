tree
====

a simple balance tree implement

tree package is super simple balance tree implementation.

	
	bt := tree.New(func(l,r interface{})int { return l.(int)-r.(int) },false)
	for i:=0;i<10;i++ {
		bt.Add(rand.Int()%1000)
	}
	tree.PrintTree(os.Stdout,bt,3,func(e interface{}) string{
		return fmt.Sprintf("%03d",e)
	}
