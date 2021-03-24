package main

import "fmt"

func Traverse(nodes ...*BTreeDictionaryEntryImpl) {
	for _, node := range nodes {
		fmt.Printf("%d ", node.Key())
	}
	fmt.Printf("\n")

	nds := Aggregate(nodes)

	if nds == nil {
		return
	}

	Traverse(nds...)
}

func Aggregate(nodes []*BTreeDictionaryEntryImpl) []*BTreeDictionaryEntryImpl {
	var aggregation []*BTreeDictionaryEntryImpl

	for _, node := range nodes {
		if node.Left != nil {
			aggregation = append(aggregation, node.Left)
		}
		if node.Right != nil {
			aggregation = append(aggregation, node.Right)
		}
	}

	return aggregation
}