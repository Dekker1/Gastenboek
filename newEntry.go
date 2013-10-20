package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	maxContent int = 10000
	maxSign    int = 1000
	maxName    int = 25
)

var (
	reader *bufio.Reader = bufio.NewReader(os.Stdin)
)

func msgContent() (string, bool) {
	var msg string

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

func writeToFile(filename, content string) bool {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer f.Close()

	buf := bufio.NewWriter(f)

	bytes, err := buf.WriteString(content)
	if err != nil {
		return false
	}

	buf.Flush()
	fmt.Printf("\nUw bericht van %v bytes is opgeslagen in het gastenboek. Druk op enter om terug te keren in het menu.\n", bytes)

	return true
}

func makeEntry() bool {
	clear()

	fmt.Print("Vul de naam van uw vereniging of commissie in: ")
	name, nameOk := reader.ReadString('\n')
	if nameOk != nil {
		return false
	} else if len(name) > maxName {
		name = name[:maxName-1]
	}

	fmt.Printf("\n\nSchrijf hieronder uw bericht aan het nieuwe bestuur van Thalia.\nU kunt alle tekens binnen de unicode gebruiken, u kunt maximaal %d bytes gebruiken.\nU Schrijft per regel, dus nadat een regel getypt is, is deze definitief.\nSluit uw bericht af met een regel \"Aldus ons bericht.\", u kunt het bericht daarna nog persoonlijk ondertekenen.\n", maxContent)
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

	ok = writeToFile(strconv.FormatInt(time.Now().Unix(), 10)+" "+name, content+"\n\n"+sign)

	reader.ReadString('\n')

	return ok
}
