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
	ascii string = "-hMMMMMMMMMMMd:\n-hMMMMMMMMMMMd:\n-:::::::::::::::::::::::::`.::::::::::::::::::::::::::.-hMMMMMMMMMMMd:\n`oNMMMMMMMMMM.MMMMMMMMMMMMM.ssmMMMMMMMMMMMMMMMMMMMMMMdshMMMMMMMMMMMd:\n`oNMMMMMMMM.MMMMMMMMMMMMM.NNysmMMMMMMMMMMMMMMMMMMmshMMMMMMMMMMMd:\n`oNMMMMMM.MMMMMMMMMMMMM.NMMNysmMMMMMMMMMMMMMMmshMMMMMMMMMMMd/\n`oNMMMM.MMMMMMMMMMMMM.NMMMMNysmMMMMMMMMMMm+ohddddddddddh/`\n`oNMM.MMMMMMMMMMMMM.NMMMMMMNysmMMMMMMm/`````````````\n.oN.MMMMMMMMMMMMM.NMMMMMMMMNysmMMm/`\n.:::::::::::::`NMMMMMMMMMMNys+`\nNMMMMMMMMMMMMNo.\nNMMMMMMMMMMMMMMmo`\nNMMMMMMMMMMMMMMMMmo`\nNMMMMMMMMMMMMMMMMMy.\nNMMMMMMMMMMMMMMMh-\nNMMMMMMMMMMMMMh:\nNMMMMMMMMMMMh:\nNMMMMMMMMMh:\nNMMMMMMMdyo\nNMMMMMdydMy\nNMMMdsdMMMy\nNMdsdMMMMMy\ndsdMMMMMMMy\n:hMMMMMMMMMy\n:hMMMMMMMMMMMy\n-hMMMMMMMMMMMMMy\n-hMMMMMMMMMMMMMMMy\n-yMMMMMMMMMMMMMMMMMy\n:dMMMMMMMMMMMMMMMMMy\n:dMMMMMMMMMMMMMMMy\n/dMMMMMMMMMMMMMy\n/dMMMMMMMMMMMy-.\n/dMMMMMMMMMy+Ns.\n/dMMMMMMMy+MMNs.\n/mMMMMMy+MMMMNs.\n/mMMMy+MMMMMMMs.\n/mMy+MMMMMMMMMs.\n/++MMMMMMMMMMMy."
	tekst string = "Thalia Constitutieborrel Gastenboek\n1) Schrijf een nieuw bericht.\n2) Lees oude berichten."
)

var (
	menu string = ""
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

func generateMenu() string {
	if menu == "" {
		var width, height int = terminalSize()
		var content []string = strings.Split(ascii, "\n")
		content = append(content, strings.Split(tekst, "\n")...)
		menu += strings.Repeat("-", width)
		for _, value := range content {
			var length int = len(value)
			menu += "|"
			menu += strings.Repeat(" ", math.Floor((width-length-2)/2.0))
			menu += value
			menu += strings.Repeat(" ", math.Ceil((width-length-2)/2.0))
			menu += "|"
		}
		menu += strings.Repeat("-", width)
	}
	return menu
}

func showMenu() {
	fmt.Println(generateMenu())
}
