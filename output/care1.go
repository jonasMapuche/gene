package main

import (
	"fmt"
	"time"

	"gene.go/trees"
)

func main() {
	fmt.Println(" -------------------------------------------------")
	fmt.Println("| Start...")
	fmt.Println(" -------------------------------------------------")
	fmt.Println("| Data...")
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| Tree...")
	var article Notice
	article.Spawn = "raiz"
	article.Date = time.Date(2023, 1, 17, 0, 0, 0, 0, time.Local)
	article.Quantity = 4
	var tree *trees.Tree = trees.Add(trees.Notice(article))
	fmt.Println("| Tree: ", tree.String())
	fmt.Println("|-------------------------------------------------")
}
