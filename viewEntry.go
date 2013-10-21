package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func listFiles(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		fmt.Println(info.Name() + " - " +  path)
	}
	return nil
}

func viewEntry() {
	clear()
	filepath.Walk(".", filepath.WalkFunc(listFiles))

	fmt.Print("\nSelecteer hier welk bericht u graag zou willen lezen: ")
	var i int
	fmt.Scan(&i)
}
