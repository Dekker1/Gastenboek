package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
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
	for i := 0; i < 100000; i++ {
		fmt.Printf("At %v\n", i)
	}
	fmt.Printf("Done!")
	os.Exit(0)
}

func main() {
	go showMenu()

	c := make(chan os.Signal)
	signal.Notify(c)
	for sig := range c {
		fmt.Println(sig.String())
	}
	return
}
