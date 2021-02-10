package main

import "fmt"

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

	fmt.Println(tpl)
}
