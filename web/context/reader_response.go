package context

import (
	"io"
	"net/http"

	"github.com/gabriel-vasile/mimetype"
)

type ReaderResponse struct {
	*Response
	contentType string
	reader      io.Reader
}

func (readerResponse *ReaderResponse) SetReader(reader io.Reader) *ReaderResponse {
	readerResponse.reader = reader
	return readerResponse
}

func (readerResponse *ReaderResponse) SetContentType(contentType string) *ReaderResponse {
	readerResponse.contentType = contentType
	return readerResponse
}

func (readerResponse *ReaderResponse) Send(w http.ResponseWriter, r *http.Request) {
	// set cookies
	for _, cookie := range readerResponse.cookies {
		http.SetCookie(w, cookie)
	}
	// set headers
	for key, value := range readerResponse.headers {
		w.Header()[key] = value
	}
	// set content type
	if readerResponse.contentType != "" {
		w.Header().Set("Content-Type", readerResponse.contentType)
	} else {
		mime, _ := mimetype.DetectReader(readerResponse.reader)
		w.Header().Set("Content-Type", mime.String())
		// rewind reader
		if _, err := readerResponse.reader.(io.ReadSeeker).Seek(0, 0); err != nil {
			panic(err)
		}
	}
	// set http status code
	w.WriteHeader(readerResponse.statusCode)
	if _, err := io.Copy(w, readerResponse.reader); err != nil {
		panic(err)
	}
}
