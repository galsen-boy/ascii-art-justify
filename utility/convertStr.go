package utility

import (
	"fmt"
	"os"
	"strings"
)

func ConvertStr(str string) string {
	data, err := os.ReadFile(os.Args[3] + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	dataSplit := strings.Split(string(data), "\n")
	strTab := strings.Split(str, "\\n")
	var newStr string
	for j, args := range strTab {
		if args != "" {
			var result []string
			for _, char := range args {
				for i := 1; i <= 8; i++ {
					result = append(result, dataSplit[((char-32)*9)+rune(i)])
				}
			}
			var tab [8][]string
			for i, val := range result {
				tab[i%8] = append(tab[i%8], val)
			}
			for _, ligne := range tab {
				for _, part := range ligne {
					newStr += part
				}
				newStr += "\n"
			}
		} else if j != len(str)-1 && args != "\n" {
			fmt.Println()
		}
	}
	return newStr
}

// --------------------------------

func TraiteFile(s []string) []string {
	data, err := os.ReadFile(os.Args[2] + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	dataSplit := strings.Split(string(data), "\n")

	for j, args := range s {
		if args != "" {
			var result []string
			for _, char := range args {
				if char < 32 || char > 127 {
					fmt.Println("There is an unprintable caracter !!!")
					os.Exit(0)
				}
				for i := 1; i <= 8; i++ {
					result = append(result, dataSplit[((char-32)*9)+rune(i)])
				}
			}
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
		} else if j != len(s)-1 && args != "\n" {
			fmt.Println()
		}
	}
	return s
}
