package main

type BTreeDictionaryEntry interface {
	Key() int
	Value() interface{}
}

type BTreeDictionaryEntryImpl struct {
	_Key int
	_Value interface{}
	Left *BTreeDictionaryEntryImpl
	Right *BTreeDictionaryEntryImpl
}

func BTreeDictionaryEntryImplConstructor(key int, value interface{}) *BTreeDictionaryEntryImpl {
	instance := &BTreeDictionaryEntryImpl{}
	instance.SetKey(key)
	instance.SetValue(value)

	return instance
}

func (entry *BTreeDictionaryEntryImpl) Key() int {
	return entry._Key
}

func (entry *BTreeDictionaryEntryImpl) SetKey(key int) {
	entry._Key = key
}

func (entry *BTreeDictionaryEntryImpl) Value() interface{} {
	return entry._Value
}

func (entry *BTreeDictionaryEntryImpl) SetValue(value interface{}) {
	entry._Value = value
}

func (entry *BTreeDictionaryEntryImpl) Replace(source *BTreeDictionaryEntryImpl, target *BTreeDictionaryEntryImpl) {
	if source == nil {
		panic("nil cannot be replaced")
	}

	if entry.Left == source {
		entry.Left = target
	} else if entry.Right == source {
		entry.Right = target
	}
}

func (entry *BTreeDictionaryEntryImpl) FindIop() (*BTreeDictionaryEntryImpl, *BTreeDictionaryEntryImpl) {
	var iop, parent *BTreeDictionaryEntryImpl = entry.Left, nil
	for {
		if iop.Right == nil {
			break
		}
		parent = iop
		iop = iop.Right
	}

	return iop, parent
}