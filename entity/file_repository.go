package entity

import (
	"bufio"
	"io"
	"log"
	"os"
)

type FileRepository struct {
	FilePath string
}

func NewFileRepository(filePath string) *FileRepository {
	return &FileRepository{
		FilePath: filePath,
	}
}

func (r *FileRepository) GetLine(lineNum int) (string, error) {
	file, err := os.Open(r.FilePath)
	if err != nil {
		log.Fatal("File not found or inaccessible")
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	var currentLine int
	for sc.Scan() {
		currentLine++
		if currentLine == lineNum {
			return sc.Text(), sc.Err()
		}
	}
	return "", io.EOF
}
