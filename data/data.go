package data

import "time"

type Notice struct {
	Spawn    string
	Date     time.Time
	Quantity float32
}

var note []Notice

func Init() {
	var article Notice
	article.Spawn = "pepper"
	article.Date = time.Date(2023, 1, 17, 0, 0, 0, 0, time.Local)
	article.Quantity = 6
	note = append(note, article)
	article.Spawn = "pepper"
	article.Date = time.Date(2023, 1, 18, 0, 0, 0, 0, time.Local)
	article.Quantity = 4
	note = append(note, article)
	article.Spawn = "pepper"
	article.Date = time.Date(2023, 1, 19, 0, 0, 0, 0, time.Local)
	article.Quantity = 4
	note = append(note, article)
	article.Spawn = "pepper"
	article.Date = time.Date(2023, 1, 20, 0, 0, 0, 0, time.Local)
	article.Quantity = 4
	note = append(note, article)
	article.Spawn = "pepper"
	article.Date = time.Date(2023, 1, 21, 0, 0, 0, 0, time.Local)
	article.Quantity = 0
	note = append(note, article)
	article.Spawn = "pepper"
	article.Date = time.Date(2023, 1, 22, 0, 0, 0, 0, time.Local)
	article.Quantity = 0
	note = append(note, article)
	article.Spawn = "pepper"
	article.Date = time.Date(2023, 1, 23, 0, 0, 0, 0, time.Local)
	article.Quantity = 6
	note = append(note, article)
	article.Spawn = "pepper"
	article.Date = time.Date(2023, 1, 24, 0, 0, 0, 0, time.Local)
	article.Quantity = 4
	note = append(note, article)
	article.Spawn = "pepper"
	article.Date = time.Date(2023, 1, 25, 0, 0, 0, 0, time.Local)
	article.Quantity = 4
	note = append(note, article)
	article.Spawn = "pepper"
	article.Date = time.Date(2023, 1, 26, 0, 0, 0, 0, time.Local)
	article.Quantity = 4
	note = append(note, article)
}

func Add(item Notice) {
	var article Notice
	article.Spawn = item.Spawn
	article.Date = item.Date
	article.Quantity = item.Quantity
	note = append(note, article)
}
