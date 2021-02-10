package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := "Kairi Hojo"

	tpl := `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Kairi</title>
</head>
<body>
<h1>` + name + `</h1>
</html>
`

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	io.Copy(file, strings.NewReader(tpl))
}
