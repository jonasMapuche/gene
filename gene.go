package main

import (
	"fmt"
	"net/http"
	"time"

	"gene.go/files"
	"gene.go/router"
	"gene.go/trees"
)

type Notice struct {
	Spawn    string
	Date     time.Time
	Quantity float32
}

func main() {
	fmt.Println(" -------------------------------------------------")
	fmt.Println("| Start...")
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| Data...")
	/*
		var article Notice
		article.Spawn = ""
		article.Date = time.Date(2023, 1, 17, 0, 0, 0, 0, time.Local)
		article.Quantity = 4
		data.Init()
		data.Add(data.Notice(article))
	*/
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| Tree...")
	var article Notice
	article.Spawn = "raiz"
	article.Date = time.Date(2023, 1, 17, 0, 0, 0, 0, time.Local)
	article.Quantity = 4
	var tree *trees.Tree = trees.Add(trees.Notice(article))
	fmt.Println("| Tree: ", tree.String())
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| Body...")
	file := "care1.go"
	files.Body("output/" + file)
	fmt.Println("| Body:", file)
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| Http...")
	router.Controller()
	http.ListenAndServe(":8000", nil)
	fmt.Println("| End...")
	fmt.Println(" -------------------------------------------------")
}
