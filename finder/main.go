package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// readLine reads line from the buffer
func readLine(r *bufio.Reader) (string, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func main() {
	args := os.Args[1:]
	fileName := args[0]
	query := args[1]

	if len(query) > 2 {
		secondChar := strings.ToLower(string(query[1]))
		regex, _ := regexp.Compile("^" + query)

		file, _ := os.Open(fileName)
		defer file.Close()
		reader := bufio.NewReader(file)

		isFound := false
		line, e := readLine(reader)
		for e == nil {
			if len(line) > 2 {
				if regex.Match([]byte(line)) {
					if !isFound {
						isFound = true
					}
					fmt.Println(line)
				} else if isFound {
					lineSecondChar := strings.ToLower(string(line[1]))
					if lineSecondChar != secondChar {
						break
					}
				}
			}
			line, e = readLine(reader)
		}
	}
}
