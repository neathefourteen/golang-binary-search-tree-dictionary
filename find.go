package main

// Iterative solution
func Find(root *BTreeDictionaryEntryImpl, key int) (*BTreeDictionaryEntryImpl, *BTreeDictionaryEntryImpl) {
	var target, parent *BTreeDictionaryEntryImpl = root, nil

	for {
		if target.Key() == key {
			break
		}

		parent = target
		if key > target.Key() && target.Right != nil {
			target = target.Right
			continue
		} else if key < target.Key() && target.Left != nil {
			target = target.Left
			continue
		}

		return nil, nil
	}

	return target, parent
}

// Recursive solution
func FindRecursive(root *BTreeDictionaryEntryImpl, key int) (*BTreeDictionaryEntryImpl, *BTreeDictionaryEntryImpl) {
	if root == nil || root.Key() == key {
		return root, nil
	}
	if root.Left == nil && root.Right == nil {
		return nil, nil
	}

	var target, parent *BTreeDictionaryEntryImpl
	if key > root.Key() {
		target, parent = FindRecursive(root.Right, key)
	} else if key < root.Key() {
		target, parent = FindRecursive(root.Left, key)
	}

	if target != nil && parent == nil {
		return target, root
	}


	return target, parent
}
