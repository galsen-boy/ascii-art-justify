package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"main/utils"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 2 || len(os.Args) > 4 {
		if strings.HasPrefix(os.Args[1], "--color") {
			utility.PrintErrColor()
			return
		}
		if strings.HasPrefix(os.Args[1], "--output") {
			utility.PrinterrOutput()
			return
		}
		if strings.HasPrefix(os.Args[1], "--align") {
			utility.PrintErrjustify()
			return
		} else {
			utility.AsciiArt()
			return
		}
	}
	if len(os.Args) == 3 {
		if strings.HasPrefix(os.Args[1], "--color=") {
			AsciiArtColor()
			return
		}
		utility.AsciiArtFs()
		return
	}
	if len(os.Args) == 4 && strings.HasPrefix(os.Args[1], "--output=") {
		if (os.Args[3] != "standard") && (os.Args[3] != "shadow") && (os.Args[3] != "thinkertoy") {
			utility.PrinterrOutput()
			return
		}
		utility.ConvertStr(os.Args[2])
		file, err := os.Create(os.Args[1][9:])
		if err == nil {
			file.WriteString(utility.ConvertStr(os.Args[2]))
		}
		return
	}
	if len(os.Args) == 4 && strings.HasPrefix(os.Args[1], "--color=") {
		AsciiArtColor()
		return
	}
	if len(os.Args) == 4 && strings.HasPrefix(os.Args[1], "--align=") {
		AsciiArtJustify()
		return
	} else {
		utility.PrintErrjustify()
		return
	}
}

var Colors = map[string]string{
	"Init":   "\033[0m",
	"Red":    "\033[31m",
	"Green":  "\033[32m",
	"Yellow": "\033[33m",
	"Blue":   "\033[34m",
	"Purple": "\033[35m",
	"Cyan":   "\033[36m",
	"Gray":   "\033[37m",
	"Orange": "\033[38;2;255;140;0m",
	"Pink":   "\033[38;2;255;105;108m",
	"Indigo": "\033[38;2;138;43;226m",
	"Brown":  "\033[38;2;165;42;42m",
}

func ColorCode(color string) string {
	switch color {
	case "red":
		return Colors["Red"]
	case "green":
		return Colors["Green"]
	case "yellow":
		return Colors["Yellow"]
	case "blue":
		return Colors["Blue"]
	case "purple":
		return Colors["Purple"]
	case "cyan":
		return Colors["Cyan"]
	case "gray":
		return Colors["Gray"]
	case "orange":
		return Colors["Orange"]
	case "pink":
		return Colors["Pink"]
	case "indigo":
		return Colors["Indigo"]
	case "brown":
		return Colors["Brown"]
	default:
		log.Fatalf("Invalid Color !!")
		return Colors["Init"]
	}
}

var tabA = map[int][]string{}
var taillel = 9
var color string

func AsciiArtColor() {

	args := os.Args
	if len(args[1:]) < 2 || len(args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	}
	ReadBanner()
	color = ColorCode(strings.ToLower(os.Args[1][8:]))
	input := os.Args[2]
	letterToColor := input
	if len(os.Args)-1 > 2 {
		input = os.Args[3]
		letterToColor = os.Args[2]
	}
	AsciiConvert(input, letterToColor)
}
func ReadBanner() {
	banner := "standard.txt"
	var tabB []string
	scanner, err := ioutil.ReadFile("./banner/"+banner)
	if err != nil {
		fmt.Println("Invalid banner")
		return
	}
	data := bufio.NewScanner(strings.NewReader(string(scanner)))
	for data.Scan() {
		lines := data.Text()
		tabB = append(tabB, lines)
	}
	elems := len(tabB) / taillel
	for i := 0; i < elems; i++ {
		charLines := tabB[i*taillel : (i+1)*taillel]
		tabA[int(i)] = charLines
	}
}
func AsciiConvert(input, letterToColor string) {
	var output string
	divInput := strings.Split(input, "\\n")
	for _, InputPart := range divInput {
		for i := 1; i < taillel; i++ {
			for _, runeValue := range InputPart {
				if int(runeValue) >= 32 && int(runeValue) <= 126 {
					NumCharsInTab := int(runeValue) - 32
					outputLine := ""
					if strings.ContainsRune(letterToColor, runeValue) {
						outputLine = (color + tabA[NumCharsInTab][i] + Colors["Init"])
					} else {
						outputLine = tabA[NumCharsInTab][i]

					}
					if outputLine != "" {
						output += outputLine
					}
				} else {
					fmt.Println("Invalid Input")
					return
				}
			}
			output += "\n"
		}
		if InputPart == "" {
			fmt.Println()
		} else {
			fmt.Print(output)
		}
		output = ""
	}
}
