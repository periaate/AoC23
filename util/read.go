package util

import (
	"bufio"
	"os"
)

func ReadLines(src string) (lines []string, err error) {
	file, err := os.OpenFile(src, os.O_RDONLY, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		data := sc.Text()
		lines = append(lines, data)
	}

	return lines, nil
}
