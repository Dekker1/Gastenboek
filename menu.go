package gastenboek

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func terminalSize() (width, height int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	fmt.Printf("out: %#v\n", string(out))
	fmt.Printf("err: %#v\n", err)
	if err != nil {
		log.Fatal(err)
	}
	return 0, 0
}

func showMenu() {
	cmd := exec.Command("echo", "Thalia Constitutieborrel Gastenboek\n Optie 1 1: Optie 2 2:")
	err := cmd.Run()
	if err != nil {
		fmt.Printf("There has been an error: %s ", err)
	}

}
