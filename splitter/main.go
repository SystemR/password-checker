package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

const passwordRegex = "(?i)[a-z0-9_@#%&=" + `\(\)\$\*\^\-\+` + "]"

func main() {
	startTime := time.Now()

	outputFolder := "pass"

	args := os.Args[1:]
	if len(args) == 0 {
		panic("Please specify the dictionary file")
	} else if len(args) == 2 {
		outputFolder = args[1]
	}

	// Handle password file
	passwordFile, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	defer passwordFile.Close()
	passwordFileReader := bufio.NewReader(passwordFile)

	// Create output folder
	createFolderIfNotExist(outputFolder)

	// Result summary file
	resultFile, _ := os.Create(outputFolder + "/result.json")
	defer resultFile.Close()
	resultWriter := bufio.NewWriter(resultFile)

	// Create password files
	fileWriters, filesArray := createPasswordFiles(outputFolder)
	defer closeFiles(filesArray)

	// Result set
	resultSet := make(map[string]int)

	// Matchers
	passwordRegex, _ := regexp.Compile(passwordRegex)
	alphanumericRegex, _ := regexp.Compile("[a-z0-9]")

	// Process passwords
	line, err := readLine(passwordFileReader)
	for err == nil {
		if len(line) > 0 {
			var char = strings.ToLower(string(line[0]))
			if passwordRegex.Match([]byte(char)) {
				incrementResult(resultSet, char)
				writeLineToPasswordFile(fileWriters, alphanumericRegex, char, line)
			}
		}
		line, err = readLine(passwordFileReader)
	}

	timeElapsed := time.Since(startTime)
	fmt.Printf("Time spent %s", timeElapsed)

	jsonString, _ := json.Marshal(resultSet)
	resultWriter.WriteString(string(jsonString))
	resultWriter.Flush()

	fmt.Println(string(jsonString))
}

func incrementResult(resultSet map[string]int, char string) {
	if resultSet[char] > 0 {
		resultSet[char]++
	} else {
		resultSet[char] = 1
	}
}

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

// createFolderIfNotExist checks if folder exists and creates the folder
func createFolderIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func createPasswordFiles(outputFolder string) (map[string]bufio.Writer, []os.File) {
	passwordFiles := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o",
		"p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "other"}

	numOfFiles := len(passwordFiles)
	filesArray := make([]os.File, numOfFiles)
	fileWriters := make(map[string]bufio.Writer)

	for i := 0; i < numOfFiles; i++ {
		char := passwordFiles[i]
		fileName := outputFolder + "/" + passwordFiles[i] + ".txt"

		// Create file
		outFile, err := os.Create(fileName)
		if err != nil {
			panic("Create file error")
		}
		filesArray[i] = *outFile

		// Create writer
		outFileWriter := bufio.NewWriter(outFile)
		fileWriters[char] = *outFileWriter
	}

	return fileWriters, filesArray
}

func closeFiles(filesArray []os.File) {
	length := len(filesArray)
	for i := 0; i < length; i++ {
		filesArray[i].Close()
	}
}

// writeLine writes line to buffer
func writeLineToPasswordFile(fileWriters map[string]bufio.Writer, alphanumericRegex *regexp.Regexp, char string, line string) {
	var writer bufio.Writer
	if alphanumericRegex.Match([]byte(char)) {
		writer = fileWriters[char]
	} else {
		writer = fileWriters["other"]
	}

	writer.WriteString(line + "\n")
	writer.Flush()
}
