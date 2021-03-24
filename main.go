package main

import "fmt"

func main() {
	dictionary := BTreeDictionaryConstructor()
	dictionary.Insert(6, "")
	dictionary.Insert(3, "")
	dictionary.Insert(8, "test")
	dictionary.Delete(6)

	dictionary.Traverse()

	fmt.Printf("%s", dictionary.Get(8).Value())

}
