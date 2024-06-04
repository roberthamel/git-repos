package helpers

import (
	"bufio"
	"fmt"
	"os"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func IsFileEmpty(filename string) (bool, error) {
	if !FileExists(filename) {
		return false, fmt.Errorf("file %s does not exist", filename)
	}
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false, err
	}
	return fileInfo.Size() == 0, err
}

func WriteFile(filename, content string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func TruncateFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func ReadLastLineOfFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var lastLine string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}
	err = scanner.Err()
	if err != nil {
		return "", err
	}
	return lastLine, nil
}
