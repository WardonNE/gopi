package formdata

import (
	"mime/multipart"
	"path/filepath"
)

// UploadedFiles is a slice of [UploadedFile] instance
type UploadedFiles []*UploadedFile

// NewUploadedFiles creates an instance of [UploadedFiles]
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

// Save saves all uploaded files under the given directory
func (uploadedFiles UploadedFiles) Save(dirpath string) {
	for _, uploadedFile := range uploadedFiles {
		if err := uploadedFile.SaveAs(filepath.Join(dirpath, uploadedFile.Name())); err != nil {
			panic(err)
		}
	}
}

// Close closes all uploaded files
func (uploadedFiles UploadedFiles) Close() error {
	for _, uploadedFile := range uploadedFiles {
		if err := uploadedFile.Close(); err != nil {
			return err
		}
	}
	return nil
}
