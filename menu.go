package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	ascii = "                                                            -hMMMMMMMMMMMd:\n                                                          -hMMMMMMMMMMMd:  \n-:::::::::::: :::::::::::::`.::::::::::::::::::::::::::.-hMMMMMMMMMMMd:    \n`oNMMMMMMMMMM.MMMMMMMMMMMMM.ssmMMMMMMMMMMMMMMMMMMMMMMdshMMMMMMMMMMMd:      \n  `oNMMMMMMMM.MMMMMMMMMMMMM.NNysmMMMMMMMMMMMMMMMMMMmshMMMMMMMMMMMd:        \n    `oNMMMMMM.MMMMMMMMMMMMM.NMMNysmMMMMMMMMMMMMMMmshMMMMMMMMMMMd/          \n      `oNMMMM.MMMMMMMMMMMMM.NMMMMNysmMMMMMMMMMMm+ohddddddddddh/`           \n        `oNMM.MMMMMMMMMMMMM.NMMMMMMNysmMMMMMMm/` ````````````              \n          .oN.MMMMMMMMMMMMM.NMMMMMMMMNysmMMm/`                             \n            . :::::::::::::`NMMMMMMMMMMNys+`                               \n                            NMMMMMMMMMMMMNo.                               \n                            NMMMMMMMMMMMMMMmo`                             \n                            NMMMMMMMMMMMMMMMMmo`                           \n                            NMMMMMMMMMMMMMMMMMy.                           \n                            NMMMMMMMMMMMMMMMh-                             \n                            NMMMMMMMMMMMMMh:                               \n                            NMMMMMMMMMMMh:                                 \n                            NMMMMMMMMMh:                                   \n                            NMMMMMMMdyo                                    \n                            NMMMMMdydMy                                    \n                            NMMMdsdMMMy                                    \n                            NMdsdMMMMMy                                    \n                            dsdMMMMMMMy                                    \n                           :hMMMMMMMMMy                                    \n                         :hMMMMMMMMMMMy                                    \n                       -hMMMMMMMMMMMMMy                                    \n                     -hMMMMMMMMMMMMMMMy                                    \n                   -yMMMMMMMMMMMMMMMMMy                                    \n                   :dMMMMMMMMMMMMMMMMMy                                    \n                     :dMMMMMMMMMMMMMMMy                                    \n                       /dMMMMMMMMMMMMMy                                    \n                         /dMMMMMMMMMMMy-.                                  \n                           /dMMMMMMMMMy+Ns.                                \n                             /dMMMMMMMy+MMNs.                              \n                               /mMMMMMy+MMMMNs.                            \n                                 /mMMMy+MMMMMMMs.                          \n                                   /mMy+MMMMMMMMMs.                        \n                                     /++MMMMMMMMMMMy.                      "
)

func terminalSize() (width, height int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	sizes := strings.Split(string(out), " ")
	width, _ = strconv.Atoi(sizes[1])
	height, _ = strconv.Atoi(sizes[0])
	return width, height
}

func showMenu() {
	fmt.Printf("%v\nThalia Constitutieborrel Gastenboek\n Optie 1 1: Optie 2 2:", ascii)
}
