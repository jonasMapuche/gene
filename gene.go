package main

import (
	"fmt"
	"time"

	"gene.go/files"
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
	file := "care1.go"
	files.Body(file)
	/*
		var article Notice
		article.Spawn = ""
		article.Date = time.Date(2023, 1, 17, 0, 0, 0, 0, time.Local)
		article.Quantity = 4
		data.Init()
		data.Add(data.Notice(article))
	*/
	fmt.Println("| Body:", file)
	fmt.Println("|-------------------------------------------------")
	fmt.Println("| End...")
	fmt.Println(" -------------------------------------------------")
}
