package main

type BTreeDictionary interface {
	Get(key int) BTreeDictionaryEntry
	Insert(key int, value interface{})
	Delete(key int)
	Traverse()
}

type BTreeDictionaryImpl struct {
	Root *BTreeDictionaryEntryImpl
}

func BTreeDictionaryConstructor() BTreeDictionary {
	return &BTreeDictionaryImpl{}
}

func (dictionary *BTreeDictionaryImpl) Traverse() {
	Traverse(dictionary.Root)
}

func (dictionary *BTreeDictionaryImpl) Get(key int) BTreeDictionaryEntry {
	target, _ := Find(dictionary.Root, key)

	if target == nil {
		panic("target not found")
	}

	return target
}

func (dictionary *BTreeDictionaryImpl) Insert(key int, value interface{}) {
	if dictionary.Root == nil {
		dictionary.Root = BTreeDictionaryEntryImplConstructor(key, value)
		return
	}
	Insert(dictionary.Root, key, value)
}

func (dictionary *BTreeDictionaryImpl) Delete(key int) {
	dictionary.Root = Delete(dictionary.Root, key)
}