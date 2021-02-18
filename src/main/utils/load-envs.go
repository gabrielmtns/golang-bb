package utils

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func LoadEnvs() {
	dir, _ := os.Getwd()
	filePath := path.Join(dir , ".env")
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer f.Close()

	lines := make([]string, 0, 100)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("%s", err)
	}

	for _, l := range lines {
		pair := strings.Split(l, "=")
		os.Setenv(pair[0], pair[1])
	}
}