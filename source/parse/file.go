package parse

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func RunProgramWithFile(path string, showContent bool) {
	err := WriteResultByFilePath(os.Stdout, path, showContent)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		os.Exit(1)
	}
}

func WriteResultByFilePath(w io.Writer, path string, writeContent bool) error {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("WriteResultByFilePath: %w", err)
	} else if fInfo, _ := file.Stat(); fInfo.IsDir() {
		err = fmt.Errorf("%v is directory", fInfo.Name())
		return fmt.Errorf("WriteResultByFilePath: %w", err)
	}

	scanner := bufio.NewScanner(file)
	result, err := GetResult(scanner)
	if err != nil {
		return fmt.Errorf("WriteResultByFilePath: %w", err)
	}
	if writeContent {
		file.Seek(0, io.SeekStart)
		_, err = io.Copy(w, file)
		if err != nil {
			return fmt.Errorf("WriteResultByFilePath: %w", err)
		}
		fmt.Fprint(w, "\n\n# result\n")
	}
	result.WriteResult(w)

	return nil
}
