package context

import (
	"net/http"
	"os"
)

type FileResponse struct {
	*ReaderResponse
	filename string
}

func (fileResponse *FileResponse) SetFile(filename string) *FileResponse {
	fileResponse.filename = filename
	return fileResponse
}

func (fileResponse *FileResponse) Send(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(fileResponse.filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fileResponse.SetContent(f)
	fileResponse.ReaderResponse.Send(w, r)
}
