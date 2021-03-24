package main

func Insert(root *BTreeDictionaryEntryImpl, key int, value interface{}) {
	predecessor := root

	for {
		if key == predecessor.Key() {
			predecessor.SetValue(value)
			break
		}

		if key > predecessor.Key() {
			if predecessor.Right == nil {
				predecessor.Right = BTreeDictionaryEntryImplConstructor(key, value)
				break
			}
			predecessor = predecessor.Right
		}

		if key < predecessor.Key() {
			if predecessor.Left == nil {
				predecessor.Left = BTreeDictionaryEntryImplConstructor(key, value)
				break
			}
			predecessor = predecessor.Left
		}
	}
}