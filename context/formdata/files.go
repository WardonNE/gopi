package formdata

import (
	"mime/multipart"
	"path/filepath"
)

type UploadedFiles []*UploadedFile

func NewUploadedFiles(fileHeaders []*multipart.FileHeader) UploadedFiles {
	uploadedFiles := make([]*UploadedFile, 0, len(fileHeaders))
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			panic(err)
		}
		uploadedFile, _ := NewUploadedFile(file, fileHeader)
		uploadedFiles = append(uploadedFiles, uploadedFile)
	}
	return uploadedFiles
}

func (uploadedFiles UploadedFiles) Save(dirpath string) {
	for _, uploadedFile := range uploadedFiles {
		if err := uploadedFile.SaveAs(filepath.Join(dirpath, uploadedFile.Name())); err != nil {
			panic(err)
		}
	}
}

func (uploadedFiles UploadedFiles) Close() {
	for _, uploadedFile := range uploadedFiles {
		uploadedFile.Close()
	}
}
