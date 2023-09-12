package formdata

import (
	"bufio"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

type UploadedFile struct {
	fileHeader *multipart.FileHeader
	file       multipart.File
	mime       *mimetype.MIME
}

func NewUploadedFile(file multipart.File, fileHeader *multipart.FileHeader) (*UploadedFile, error) {
	mime, err := mimetype.DetectReader(file)
	if _, err := file.Seek(0, 0); err != nil {
		return nil, err
	}
	uploadedFile := &UploadedFile{
		fileHeader: fileHeader,
		file:       file,
		mime:       mime,
	}
	return uploadedFile, err
}

func (uploadedFile *UploadedFile) Name() string {
	return uploadedFile.fileHeader.Filename
}

func (uploadedFile *UploadedFile) ClientExtension() string {
	return filepath.Ext(uploadedFile.fileHeader.Filename)
}

func (uploadedFile *UploadedFile) ClientMimeType() string {
	return uploadedFile.fileHeader.Header.Get("Content-type")
}

func (uploadedFile *UploadedFile) MimeType() string {
	return uploadedFile.mime.String()
}

func (uploadedFile *UploadedFile) Extension() string {
	return uploadedFile.mime.Extension()
}

func (uploadedFile *UploadedFile) Content() ([]byte, error) {
	content := make([]byte, 0, uploadedFile.fileHeader.Size)
	scanner := bufio.NewScanner(uploadedFile.file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		content = append(content, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		return content, err
	}
	return content, nil
}

func (uploadedFile *UploadedFile) SaveAs(filename string) error {
	dst, err := os.Create(filename)
	if err != nil {
		return err
	}
	if _, err = io.Copy(dst, uploadedFile.file); err != nil {
		return err
	}
	return nil
}

func (uploadedFile *UploadedFile) Close() error {
	return uploadedFile.file.Close()
}
