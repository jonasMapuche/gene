package files

import (
	"io/ioutil"
)

func Body(file string) {
	expression := "package main\n\nimport (\n	\"fmt\"\n)\n\n"
	expression = expression + "func main() {\n	fmt.Println(\"| Start...\")\n}"
	text := []byte(expression)
	err := ioutil.WriteFile("out/"+file, text, 0644)
	if err != nil {
		panic(err)
	}
}
