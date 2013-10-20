package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	maxContent int = 10000
	maxSign    int = 1000
)

func msgContent() (string, bool) {
	var msg string
	reader := bufio.NewReader(os.Stdin)

S:
	for {
		switch line, err := reader.ReadString('\n'); {
		case err != nil:
			return msg, false
		case strings.TrimSpace(line) == "Aldus ons bericht.":
			msg += "Aldus ons bericht.\n\n"
			break S
		default:
			msg += line
			if len(msg) > maxContent {
				fmt.Println("Uw bericht werd te lang en is om deze reden hier beëindigd.")
				break S
			}
		}
	}

	return msg, true
}

func msgSign() (string, bool) {
	var sign string
	var end int = -1
	reader := bufio.NewReader(os.Stdin)

S:
	for {
		switch line, err := reader.ReadString('\n'); {
		case err != nil:
			return sign, false
		case strings.TrimSpace(line) == "":
			if end > 0 {
				break S
			}
			sign += line
			end++
		default:
			sign += line
			end = 0
			if len(sign) > maxSign {
				fmt.Println("Uw bericht werd te lang en is om deze reden hier beëindigd.")
				break S
			}
		}
	}

	return sign, true
}

func writeToFile(content string) bool {

}

func makeEntry() bool {
	clear()

	fmt.Printf("Schrijf hieronder uw bericht aan het nieuwe bestuur van Thalia.\nU kunt alle tekens binnen de unicode gebruiken, u kunt maximaal %d bytes gebruiken.\nU Schrijft per regel, dus nadat een regel getypt is, is deze definitief.\nSluit uw bericht af met een regel \"Aldus ons bericht.\", u kunt het bericht daarna nog persoonlijk ondertekenen.\n", maxContent)
	var width, _ int = terminalSize()
	fmt.Print(strings.Repeat("-", width-1) + "\n")
	content, ok := msgContent()
	if !ok {
		return false
	}

	fmt.Print("\n\nHieronder heeft u de kans om uw bericht persoonlijk te ondertekenen, sluit dit af met twee witregels/enters'.\n")
	fmt.Print(strings.Repeat("-", width-1) + "\n")
	sign, ok := msgSign()
	if !ok {
		return false
	}

	ok = writeToFile(content + "\n\n" + sign)

	return true
}
