package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	timeLayout = "3:04pm (MST)"
)

var (
	entries []string
)

// Get the information from entries and format it.
func listFilenames() {
	for i, val := range entries {
		val = strings.Trim(val, ".txt")
		val = strings.Trim(val, "\n")
		parts := strings.Split(val, " ")
		unix, _ := strconv.ParseInt(parts[0], 10, 64)
		t := time.Unix(unix, 0)
		fmt.Printf("%d) %v - %v \n", i+1, t.Format(timeLayout), parts[1])
	}
}

// Adds a file to entries if it's an actual file
func getFilenames(path string, info os.FileInfo, err error) error {
	if !info.IsDir() && info.Name() == path && filepath.Ext(path) == ".txt" {
		entries = append(entries, path)
	}
	return nil
}

// A function that works like "less"
func showFile(i int) {
	f, openErr := os.Open(entries[i])
	if openErr != nil {
		fmt.Printf("\n\tCouldn't open file: %v", openErr)
		return
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	user := bufio.NewReader(os.Stdin)

	clear()
	fmt.Print("Druk op enter om volgende regels te zien en uiteindelijk terug te keren naar het keuzemenu.\n")
	var width, height int = terminalSize()
	fmt.Print(strings.Repeat("-", width-1) + "\n")

	var lines int
	var done = false
	for !done {
		line, err := reader.ReadString('\n')
		done = err != nil
		fmt.Print(line)
		if lines >= height {
			user.ReadString('\n')
		}
		lines++
	}
	fmt.Print(strings.Repeat("-", width-1) + "\n")
	user.ReadString('\n')
}

// The coordinating function to show the different entries.
func viewEntry() {
	entries = make([]string, 0)
	err := filepath.Walk(".", filepath.WalkFunc(getFilenames))
	if err != nil {
		return
	}

	for {
		clear()
		listFilenames()
		fmt.Printf("\nKies hier welk bericht u graag zou willen lezen, toets 0 om terug te keren naar het menu: ")
		var i int
		fmt.Scan(&i)
		if i > 0 && i <= len(entries) {
			showFile(i - 1)
		} else {
			return
		}
	}
}
