package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"unicode/utf8"

	"code.sajari.com/docconv/v2"
)

const ENGmuiltiplier = 1.2

func main() {

	folderPath := "./"

	files, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Ошибка при чтении папки:", err)
		return
	}

	var sum int

	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			if strings.HasSuffix(fileName, ".docx") {
				sum += countContents(fileName)
			}
		}
	}
	fmt.Println("Cумма страниц:", sum)

	fmt.Println("Программа завершена. Нажмите Enter, чтобы закрыть консоль.")
	bufio.NewReader(os.Stdin).ReadString('\n')

}

func countContents(name string) int {
	folderPath := "./"
	f, err := os.Open(folderPath + name)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	defer f.Close()

	var r io.Reader
	r = f

	tmpl, _, err := docconv.ConvertDocx(r)
	if err != nil {
		fmt.Print(tmpl)
		panic(err)
	}

	letterCountTrimmed := (strings.Join(strings.Fields(tmpl), " "))
	letterCount := utf8.RuneCountInString(letterCountTrimmed)
	pages := math.Floor((float64(letterCount)*ENGmuiltiplier)/1800) + 1

	fmt.Println("Название файла:", name)

	fmt.Println("Количество символов (utf8) RuneCountInString:", letterCount)
	// fmt.Println("Количество cлов:", wordCount)
	fmt.Println("Количество страниц:", pages)
	fmt.Println("")
	return int(pages)
}
