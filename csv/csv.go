package csv

import (
	"bufio"
	"os"
	"strings"
)

func ParseToCsv(outputFilename string, headers []string, lines [][]string) error {

	file, err := os.Create(outputFilename)

	if err != nil {
		return err
	}

	defer file.Close()
	writer := bufio.NewWriter(file)

	firstLine := strings.Join(headers, ",") + "\n"
	_, err = writer.WriteString(firstLine)

	if err != nil {
		return err
	}

	for _, line := range lines {
		_, err = writer.WriteString(strings.Join(line, ",") + "\n")
		if err != nil {
			return err
		}
	}

	return nil

}
