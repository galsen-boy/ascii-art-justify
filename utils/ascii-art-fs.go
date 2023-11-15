package utility

import (
	"os"
	"strings"
)

func AsciiArtFs() {
	if len(os.Args) == 2 {
		AsciiArt()
		return
	}
	if (len(os.Args) < 2) || (len(os.Args) > 3) || ((os.Args[2] != "standard") && (os.Args[2] != "shadow") && (os.Args[2] != "thinkertoy")) {
		PrintErr()
		return
	}
	arguments := strings.Split(os.Args[1], "\\n")
	TraiteFile(arguments)
}

// -------------------------------------------------------
