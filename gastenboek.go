package main

import (
	"os"
	"os/signal"
)

// SUPER SECRET CODE
const code int = 14101992

// IT'S A TRAP! (This catches all signals so the program can't be interrupted)
func signalCatcher() {
	c := make(chan os.Signal)
	signal.Notify(c)
	for range c {
		//fmt.Printf("Signal received: %v", sig)
	}
}

// This is where it all starts
func main() {
	go signalCatcher()
	showMenu()
}
