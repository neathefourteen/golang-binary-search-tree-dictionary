package main

func Delete(root *BTreeDictionaryEntryImpl, key int) *BTreeDictionaryEntryImpl {
	target, parent := Find(root, key)
	if target == nil {
		panic("target entry not found")
	}

	if target.Left == nil && target.Right == nil {
		if parent == nil {
			return nil
		}
		parent.Replace(target, nil)
	}

	if target.Left != nil && target.Right == nil {
		if parent == nil {
			root = target.Left
			return root
		}
		parent.Replace(target, target.Left)
	}

	if target.Right != nil && target.Left == nil {
		if parent == nil {
			root = target.Right
			return root
		}
		parent.Replace(target, target.Right)
	}

	if target.Left != nil && target.Right != nil {
		iop, iopParent := target.FindIop()

		if iopParent != nil {
			iopParent.Right = iop.Left
		}

		if target.Left != iop {
			iop.Left = target.Left
		}
		iop.Right = target.Right

		if parent != nil {
			parent.Replace(target, iop)
		} else {
			root = iop
		}

	}

	return root
}
