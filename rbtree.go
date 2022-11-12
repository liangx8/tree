package tree

import "github.com/liangx8/gl/stack"

func otherWalk(top *node, cb func(any) error) error {
	ss := stack.New[*node]()
	for {
		for top.l != nil {
			ss.Push(top)
			top = top.l
		}
		if err := cb(top.e); err != nil {
			return err
		}
		for {
			if top.r == nil {
				if err := ss.Pop(&top); err != nil {
					return nil
				}
				if err := cb(top.e); err != nil {
					return err
				}

			} else {
				top = top.r
				break
			}
		}
	}
}
