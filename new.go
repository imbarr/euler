package main

import (
	"os"
	"path/filepath"
)

func main() {
	n := os.Args[1]

	os.Mkdir("./"+n, os.ModePerm)
	f, _ := os.Create(filepath.Join(".", n, "main.go"))

	f.WriteString("package main\n\n")
	f.WriteString("func main() {\n\n")
	f.WriteString("}\n")

	f.Close()
}
