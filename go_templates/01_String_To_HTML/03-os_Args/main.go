package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := strings.Join(os.Args[1:], " ")
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

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

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))

}
