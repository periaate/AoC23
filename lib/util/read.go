package util

import (
	"bufio"
	"log"
	"os"
)

func ReadLines(src string) (lines []string) {
	file, err := os.OpenFile(src, os.O_RDONLY, 0777)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		data := sc.Text()
		lines = append(lines, data)
	}

	return lines
}
