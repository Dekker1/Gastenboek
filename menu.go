package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	ascii    string = "                                       `sMMMMMMMh:\n.-------..-------- -----------------.`sMMMMMMMd:  \n`sNMMMMMhyMMMMMMMM/hdMMMMMMMMMMMMMMhhMMMMMMMd:    \n  `sNMMMhyMMMMMMMM+MNhhMMMMMMMMMMhhMMMMMMMd:      \n    `sNMhyMMMMMMMM+MMMNhhMMMMMMy-+ooooooo:        \n      `ohyMMMMMMMM+MMMMMNhhMMs.                   \n         `........:MMMMMMMNy-                     \n                  :MMMMMMMMMNo`                   \n                  :MMMMMMMMMMMN:                  \n                  :MMMMMMMMMMd:                   \n                  :MMMMMMMMd:                     \n                  :MMMMMMd:                       \n                  :MMMMdhy                        \n                  :MMdhNMd                        \n                  :dhNMMMd                        \n                 `oNMMMMMd                        \n               `+NMMMMMMMd                        \n             `+NMMMMMMMMMd                        \n            `dMMMMMMMMMMMd                        \n              /mMMMMMMMMMd                        \n                /mMMMMMMMd.                       \n                  /mMMMMMdyd:                     \n                    /mMMMhyMMd:                   \n                      /mMhyMMMMd:                 \n                        /syMMMMMMd/               "
	asciiLen int    = 50
	tekst    string = "Thalia Constitutieborrel Gastenboek\n1) Schrijf een nieuw bericht.\n2) Lees oude berichten."
)

var (
	menu      string = ""
	menuwidth int    = 0
	Message   string = ""
)

func terminalSize() (width, height int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	sizes := strings.Trim(string(out), "\n")
	size := strings.Split(sizes, " ")
	width, _ = strconv.Atoi(size[1])
	height, _ = strconv.Atoi(size[0])
	return width, height
}

func clear() {
	_, height := terminalSize()
	fmt.Printf("%v", strings.Repeat("\n", height+1))
}

func generateMenu() string {
	var width, height int = terminalSize()
	if menu == "" && width != menuwidth {
		menuwidth = width
		if width-3 < asciiLen {
			var content []string = strings.Split(ascii, "\n")
			var options []string = strings.Split(tekst, "\n")

			content = append(content, make([]string, ((height-4)-len(content)-len(options))/2)...)
			content = append(content, options...)
			content = append(content, make([]string, ((height-4)-len(content)-len(options))/2)...)
			menu += strings.Repeat("-", width-1) + "\n"

			for _, value := range content {
				var length int = len(value)
				menu += "|"
				menu += strings.Repeat(" ", int(math.Floor(float64((width-length-3)/2.0))))
				menu += value
				menu += strings.Repeat(" ", int(math.Ceil(float64((width-length-2)/2.0))))
				menu += "|\n"
			}

			menu += strings.Repeat("-", width-1) + "\n"
		} else {
			menu = tekst
		}
	}
	return menu
}

func showMenu() {
	for {
		clear()
		fmt.Print(generateMenu())
		fmt.Println(Message)

		fmt.Print("Maak uw keuze: ")
		var i int
		fmt.Scan(&i)

		switch i {
		default:
			Message = "Was dat een keuze? Ik dacht van niet, probeer maar opnieuw!"
		case 1:
			if makeEntry() {
				Message = ""
			} else {
				Message = "Something went terribly wrong! Wat heb je nou weer gedaan??"
			}
		case 2:
			viewEntry()
		case 88888888:
			os.Exit(0)
		}
	}
}
