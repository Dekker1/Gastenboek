package main

import (
	"fmt"
	"os"
	"os/signal"
)

func signalCatcher() {
	c := make(chan os.Signal)
	signal.Notify(c)
	for sig := range c {
		fmt.Printf("Signal received: %v", sig)
	}
}

func main() {
	go signalCatcher()
	go showMenu()
	return
}
