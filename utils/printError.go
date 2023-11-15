package utility

import (
	"fmt"
)

func PrintErr() {
	fmt.Println("Usage: go run . [STRING]")
	fmt.Println()
	fmt.Print("EX: go run . something")
	fmt.Println()
}
func PrinterrOutput() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Print("EX: go run . --output=<fileName.txt> something standard")
	fmt.Println()
}
func PrintErrColor() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Print("EX: go run . --color=<color> <letters to be colored> something")
	fmt.Println()
}
func PrintErrjustify() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Print("EX: go run . --align=justify something standard")
	fmt.Println()
}