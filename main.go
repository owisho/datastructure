package main

import (
	"datastructure/tree"
	"fmt"
)

func main() {
	t := &tree.RedBlackTree{}
	t.Insert(1)
	t.Insert(6)
	t.Insert(8)
	t.Insert(11)
	t.Insert(13)
	t.Insert(17)
	t.Insert(15)
	t.Insert(25)
	t.Insert(22)
	t.Insert(27)

	t.ToString()
	fmt.Println()
	t.LayerPrint()
	fmt.Println()

	t.Delete(1)
	t.ToString()
	fmt.Println()
	t.LayerPrint()
	fmt.Println()

	t.Insert(4)
	t.Insert(19)
	t.Insert(21)
	t.Insert(23)
	t.Insert(14)

	t.ToString()
	fmt.Println()
	t.LayerPrint()
	fmt.Println()

}
