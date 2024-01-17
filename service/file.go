package service

import (
	"filelineserve/entity"
)

type FileService struct {
	fileRepository *entity.FileRepository
}

func NewFileService(fileRepo *entity.FileRepository) *FileService {
	return &FileService{
		fileRepository: fileRepo,
	}
}

func (s FileService) GetLine(lineNum int) (string, error) {
	return s.fileRepository.GetLine(lineNum)
}
