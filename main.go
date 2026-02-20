package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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
	extensions := []string{".txt", ".md", ".log", ".json", ".csv", ".xml", ".yaml", ".yml", ".conf", ".ini", ".html", ".css", ".go", ".py", ".js", ".sql", ".sh"}
	scannerConsole := bufio.NewScanner(os.Stdin)
	write("введи название файла: ")
	writer.Flush()
	if !scannerConsole.Scan() {
		write("ошибка чтения с консоли\n")
		return
	}
	fileName := strings.TrimSpace(scannerConsole.Text())
	write("выбери расширение файла из списка:\n0. .txt     4. .csv     8. .ini     12. .html\n1. .md      5. .yaml    9. .go      13. .css\n2. .log     6. .xml     10. .py     14. .sql\n3. .json    7. .conf    11. .js     15. .sh\n")
	writer.Flush()
	scannerConsole.Scan()
	extension, err := strconv.Atoi(scannerConsole.Text())
	if err != nil {
		write("Введите число, а не текст\n")
		return
	}
	if extension < 0 || extension >= len(extensions) {
		write("Такого номера расширения нет (введите от 0 до " + strconv.Itoa(len(extensions)-1) + ")\n")
		return
	}
	fileName += extensions[extension]
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
				write("найдено!\nномер строки : " + strconv.Itoa(lineNum) + "\nномер слова: " + strconv.Itoa(wordIdx+1) + "\n")
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
