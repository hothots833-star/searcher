package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
)

var writer = bufio.NewWriter(os.Stdout)

func write(s string) {
	_, err := writer.WriteString(s)
	if err != nil {
		writer.WriteString("ошибка выведения в консоль")
		writer.Flush()
		return
	}
}
func main() {
	defer writer.Flush()
	scannerConsole := bufio.NewScanner(os.Stdin)
	files, err := os.ReadDir(".")
	if err != nil {
		write("ошибка поиска файлов в директории\n")
		return
	}
	var fileNames []string
	for _, entry := range files {
		if !entry.IsDir() {
			fileNames = append(fileNames, entry.Name())
		}
	}
	prompt := promptui.Select{Label: "выбери файл:", Items: fileNames}
	_, fileName, err := prompt.Run()
	if err != nil {
		write("ошибка выбора файла")
		return
	}
	file, err := os.Open(fileName)
	if err != nil {
		write("ошибка при открытии файла\n")
		return
	}
	defer file.Close()
	scannerFile := bufio.NewScanner(file)
	write("введите слово для поиска: ")
	writer.Flush()
	if !scannerConsole.Scan() {
		write("ошибка чтения консоли\n")
		return
	}
	input := scannerConsole.Text()
	lineNum := 0
	found := false
	for scannerFile.Scan() {
		lineNum++
		line := strings.Fields(scannerFile.Text())
		for wordIdx, word := range line {
			if word == input {
				write("найдено!\nномер строки: " + strconv.Itoa(lineNum) + "\nномер слова: " + strconv.Itoa(wordIdx+1) + "\n\n")
				found = true
			}
		}
	}
	if err = scannerFile.Err(); err != nil {
		write("ошибка чтения файла\n")
		return
	}
	if !found {
		write("не найдено\n")
	}
}
