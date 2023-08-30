package main

import (
	funcs "ascii-art-output/functions"
	"fmt"
	"os"
	"strings"
)

const (
	hashStandard   = "ac85e83127e49ec42487f272d9b9db8b"
	hashShodow     = "a49d5fcb0d5c59b2e77674aa3ab8bbb1"
	hashThinkertoy = "86d9947457f6a41a18cb98427e314ff8"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("ERROR: wrong number of arguments")
		return
	} else if len(args) > 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return

	}
	if args[0] == "" {
		return
	}
	if args[0] == "\\n" {
		return
	}

	if len(args) == 2 {
		for _, alter := range args[0] {
			if (rune(alter) < rune(32) || rune(alter) > rune(127)) && alter != rune(10) {
				fmt.Println("ERROR: non printable character")
				return
			}
		}
	} else if len(args) == 3 {
		for _, alter := range args[1] {
			if (rune(alter) < rune(32) || rune(alter) > rune(127)) && alter != rune(10) {
				fmt.Println("ERROR: non printable character")
				return
			}
		}
	}
	filename := "standard.txt"

	if len(args) == 2 {
		switch args[1] {
		case "shadow":
			filename = "shadow.txt"
		case "thinkertoy":
			filename = "thinkertoy.txt"
		case "standard":
			filename = "standard.txt"
		default:
			fmt.Println("ERROR: wrong name of banner")
			return
		}
	} else if len(args) == 3 {
		switch args[2] {
		case "shadow":
			filename = "shadow.txt"
		case "thinkertoy":
			filename = "thinkertoy.txt"
		case "standard":
			filename = "standard.txt"
		default:
			fmt.Println("ERROR: wrong name of banner")
			return
		}
	}

	if funcs.GetHash(filename) != hashStandard || funcs.GetHash(filename) != hashShodow || funcs.GetHash(filename) != hashThinkertoy {
		asciiLines, err := funcs.GetStrings(filename)
		if err != nil {
			fmt.Println("ERROR: can't read file")
			return
		}
		asciiMap := make(map[rune][]string)
		x := 1
		y := 9
		for key := 32; key < 127; key++ {
			asciiMap[rune(key)] = asciiLines[x:y]
			x = x + 9
			y = y + 9

		}

		res := ""
		var text string
		if len(args) == 1 || len(args) == 2 {
			text = strings.ReplaceAll(args[0], "\n", "\\n")
		} else if len(args) == 3 {
			text = strings.ReplaceAll(args[1], "\n", "\\n")
		}

		arg := strings.Split(text, "\\n")

		for i, v := range arg {
			if v == "" {
				arg[i] = ""
			}
		}
		newline := forNewLines(arg)
		for w := 0; w < len(arg); w++ {
			if newline && w == len(arg)-1 {
				break
			}
			if arg[w] != "" {
				for i := 0; i < 8; i++ {
					for _, ch := range arg[w] {
						res = res + asciiMap[ch][i]
					}
					res = res + string(rune(10))
				}
			} else if arg[w] == "" {
				res = res + string(rune(10))
			}
		}
		if len(args) == 3 {
			if args[0][len(args[0])-4:] != string(".txt") {
				fmt.Println("ERROR: wrong name of file extension")
				return
			}
			for w := range args[0] {
				if args[0][w] > 8 {
					w := 0
					if args[0][w:9] == "--output=" {
						funcs.WriteF(args[0][9:], res)
					} else {
						fmt.Println("ERROR: wrong name of flag")
						return
					}
				}
			}

		}
		fmt.Print(res)
	} else {
		fmt.Println("Error: Wrong hash")
		return
	}
}

func forNewLines(s []string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != "" {
			return false
		}
	}
	return true
}
