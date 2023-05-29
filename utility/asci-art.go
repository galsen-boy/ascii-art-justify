package utility

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt() {

	if len(os.Args) < 2 || len(os.Args) > 4 || len(os.Args) == 1{
		PrintErr()
		return
	}
	//Recuperation du ficher testing
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	//Traitement des arguments des parametre
	arguments := strings.Split(os.Args[1], "\\n")
	dataSplit := strings.Split(string(data), "\n")

	for j, args := range arguments {
		if args != "" {
			var result []string
			for _, char := range args {
				if char < 32 || char > 127 {
					fmt.Println("There is an unprintable caracter !!!")
					return
				}
				for i := 1; i <= 8; i++ {
					result = append(result, dataSplit[((char-32)*9)+rune(i)])
				}
			}
			//tableau a deux dimension pour aligner le sortie
			var tab [8][]string
			for i, val := range result {
				tab[i%8] = append(tab[i%8], val)
			}
			for _, ligne := range tab {
				for _, part := range ligne {
					fmt.Print(part)
				}
				fmt.Println()
			}
		} else if j != len(arguments)-1 && args != "\n" {
			fmt.Println()
		}
	}

}
