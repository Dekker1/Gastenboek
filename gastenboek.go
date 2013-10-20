package main

import (
	"os"
	"os/signal"
)

func signalCatcher() {
	c := make(chan os.Signal)
	signal.Notify(c)
	for _ = range c {
		//fmt.Printf("Signal received: %v", sig)
	}
}

func main() {
	go signalCatcher()
	showMenu()
}
