package binding

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/wardonne/gopi/utils"
)

type FormParser struct{}

func (formParser *FormParser) Parse(request *http.Request, container any) error {
	if err := request.ParseForm(); err != nil {
		return err
	}
	if err := request.ParseMultipartForm(32 << 20); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}
	form := &multipart.Form{
		Value: request.Form,
	}
	if request.MultipartForm != nil {
		form.File = request.MultipartForm.File
	}
	return utils.FormDataToStruct(form, container, "form")
}
