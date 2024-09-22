package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please enter a word")
		return
	}
	file, err := os.Open("program/standard.txt")
	if err != nil {
		fmt.Println("Erreur")
		return
	}
	defer file.Close()
	sentence := os.Args[1]

	scanner := bufio.NewScanner(file)

	count := 0

	symbole := []string{}

	symboles := [][]string{}

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), "\n")

		symbole = append(symbole, line)

		count++

		if count == 9 {

			symboles = append(symboles, symbole)

			symbole = []string{}

			count = 0
		}

	}

	result := Split(sentence)

	fmt.Println(PrintChar(result, symboles))
}

func Split(sen string) []string {
	sl := []string{}

	a := ""

	for i := 0; i < len(sen); i++ {
		if i < len(sen)-1 && sen[i] == '\\' && sen[i+1] == 'n' {

			if a != "" {
				sl = append(sl, a)
				a = ""
			}
			sl = append(sl, "")
			i++
		} else if sen[i] != '\t' {
			a += string(sen[i])
		} else {
			if a != "" {
				sl = append(sl, a)
				a = ""
			}
		}
	}
	if a != "" {
		sl = append(sl, a)
	}
	return sl
}

func PrintChar(words []string, sl [][]string) string {
	char := ""
	for _, w := range words {
		for i := 1; i <= 8; i++ {
			if len(w) == 0 {
				char += "\n"
				break
			}
			for _, e := range w {
				if 0 <= int(e)-32 && int(e)-32 <= len(sl)-1 {
					char += sl[int(e)-32][i]
				} else {
					log.Fatal("This error")
				}
			}
			if i < 8 {
				char += "\n"
			}
		}
	}
	if char[len(char)-1] == '\n' && len(words) == 1 {
		return strings.Trim(char, "\n")
	} else if char[len(char)-1] == '\n' && char[len(char)-2] == '\n' {
		char = char[:len(char)-1]
	}
	return char
}
