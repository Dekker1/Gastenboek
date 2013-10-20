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
	ascii string = "                                                            -hMMMMMMMMMMMd:\n                                                          -hMMMMMMMMMMMd:  \n-:::::::::::: :::::::::::::`.::::::::::::::::::::::::::.-hMMMMMMMMMMMd:    \n`oNMMMMMMMMMM.MMMMMMMMMMMMM.ssmMMMMMMMMMMMMMMMMMMMMMMdshMMMMMMMMMMMd:      \n  `oNMMMMMMMM.MMMMMMMMMMMMM.NNysmMMMMMMMMMMMMMMMMMMmshMMMMMMMMMMMd:        \n    `oNMMMMMM.MMMMMMMMMMMMM.NMMNysmMMMMMMMMMMMMMMmshMMMMMMMMMMMd/          \n      `oNMMMM.MMMMMMMMMMMMM.NMMMMNysmMMMMMMMMMMm+ohddddddddddh/`           \n        `oNMM.MMMMMMMMMMMMM.NMMMMMMNysmMMMMMMm/` ````````````              \n          .oN.MMMMMMMMMMMMM.NMMMMMMMMNysmMMm/`                             \n            . :::::::::::::`NMMMMMMMMMMNys+`                               \n                            NMMMMMMMMMMMMNo.                               \n                            NMMMMMMMMMMMMMMmo`                             \n                            NMMMMMMMMMMMMMMMMmo`                           \n                            NMMMMMMMMMMMMMMMMMy.                           \n                            NMMMMMMMMMMMMMMMh-                             \n                            NMMMMMMMMMMMMMh:                               \n                            NMMMMMMMMMMMh:                                 \n                            NMMMMMMMMMh:                                   \n                            NMMMMMMMdyo                                    \n                            NMMMMMdydMy                                    \n                            NMMMdsdMMMy                                    \n                            NMdsdMMMMMy                                    \n                            dsdMMMMMMMy                                    \n                           :hMMMMMMMMMy                                    \n                         :hMMMMMMMMMMMy                                    \n                       -hMMMMMMMMMMMMMy                                    \n                     -hMMMMMMMMMMMMMMMy                                    \n                   -yMMMMMMMMMMMMMMMMMy                                    \n                   :dMMMMMMMMMMMMMMMMMy                                    \n                     :dMMMMMMMMMMMMMMMy                                    \n                       /dMMMMMMMMMMMMMy                                    \n                         /dMMMMMMMMMMMy-.                                  \n                           /dMMMMMMMMMy+Ns.                                \n                             /dMMMMMMMy+MMNs.                              \n                               /mMMMMMy+MMMMNs.                            \n                                 /mMMMy+MMMMMMMs.                          \n                                   /mMy+MMMMMMMMMs.                        \n                                     /++MMMMMMMMMMMy.                      "
	tekst string = "Thalia Constitutieborrel Gastenboek\n1) Schrijf een nieuw bericht.\n2) Lees oude berichten."
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
			}
		case 2:
			Message = "Not yet implemented"
		case 88888888:
			os.Exit(0)
		}
	}
}
