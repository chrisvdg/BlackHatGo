package main

import (
	"html/template"
	"log"
	"os"
)

var templateStr = `
<html>
	<body>
		Hello {{ . }}
	</body>
</body>
`

func main() {
	t, err := template.New("Hello").Parse(templateStr)
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(os.Stdout, "<script>alert('world')</script>")
}
