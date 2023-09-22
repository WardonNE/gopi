package context

import (
	"fmt"
	"io"
	"net/http"
)

type IResponse interface {
	SetStatusCode(statusCode int) IResponse
	StatusCode() int
	SetContent(content any) IResponse
	Content() any
	SetHeader(key, header string, replace ...bool) IResponse
	SetHeaders(headers map[string]string) IResponse
	HasHeader(key string) bool
	Header(key string) string
	Headers() http.Header
	SetCookie(cookie *http.Cookie) IResponse
	Cookies() []*http.Cookie
	Send(w http.ResponseWriter, r *http.Request)
}

type Response struct {
	headers    http.Header
	cookies    []*http.Cookie
	statusCode int
	content    any
}

func NewResponse(statusCode int, content ...any) *Response {
	response := &Response{
		headers:    make(http.Header),
		cookies:    make([]*http.Cookie, 0),
		statusCode: statusCode,
	}
	if len(content) > 0 {
		response.content = content[0]
	}
	return response
}

func (response *Response) SetStatusCode(statusCode int) IResponse {
	if statusCode < 100 || statusCode > 600 {
		panic(fmt.Errorf("HTTP status code `%d` is invalid", statusCode))
	}
	response.statusCode = statusCode
	return response
}

func (response *Response) StatusCode() int {
	return response.statusCode
}

func (response *Response) SetContent(content any) IResponse {
	response.content = content
	return response
}

func (response *Response) Content() any {
	return response.content
}

func (response *Response) SetHeaders(headers map[string]string) IResponse {
	for header, value := range headers {
		response.headers.Set(header, value)
	}
	return response
}

func (response *Response) SetHeader(key, value string, replace ...bool) IResponse {
	if len(replace) == 0 || (len(replace) > 0 && replace[0]) {
		response.headers.Set(key, value)
	} else {
		response.headers.Add(key, value)
	}
	return response
}

func (response *Response) HasHeader(key string) bool {
	h, ok := response.headers[key]
	return ok && len(h) > 0
}

func (response *Response) Header(key string) string {
	return response.headers.Get(key)
}

func (response *Response) Headers() http.Header {
	return response.headers
}

func (response *Response) SetCookie(cookie *http.Cookie) IResponse {
	response.cookies = append(response.cookies, cookie)
	return response
}

func (response *Response) Cookies() []*http.Cookie {
	return response.cookies
}

func (response *Response) Send(w http.ResponseWriter, r *http.Request) {
	// set cookies
	for _, cookie := range response.cookies {
		http.SetCookie(w, cookie)
	}
	// set headers
	for key, value := range response.headers {
		w.Header()[key] = value
	}
	// set http status code
	w.WriteHeader(response.statusCode)
	// send content
	if response.content != nil {
		switch v := response.content.(type) {
		case []byte:
			if _, err := w.Write(v); err != nil {
				panic(err)
			}
		default:
			if _, err := w.Write([]byte(fmt.Sprintf("%v", response.content))); err != nil {
				panic(err)
			}
		}
	} else {
		if _, err := w.Write([]byte{}); err != nil {
			panic(err)
		}
	}
}

func (response *Response) JSON(data ...any) *JSONResponse {
	json := &JSONResponse{
		Response: response,
	}
	if len(data) > 0 {
		json.data = data[0]
	} else {
		json.data = response.content
	}
	return json
}

func (response *Response) XML(data ...any) *XMLResponse {
	xml := &XMLResponse{
		Response: response,
	}
	if len(data) > 0 {
		xml.data = data[0]
	} else {
		xml.data = response.content
	}
	return xml
}

func (response *Response) YAML(data ...any) *YAMLResponse {
	yaml := &YAMLResponse{
		Response: response,
	}
	if len(data) > 0 {
		yaml.data = data[0]
	} else {
		yaml.data = response.content
	}
	return yaml
}

func (response *Response) TOML(data ...any) *TOMLResponse {
	toml := &TOMLResponse{
		Response: response,
	}
	if len(data) > 0 {
		toml.data = data[0]
	} else {
		toml.data = response.content
	}
	return toml
}

func (response *Response) Protobuf(data ...any) *ProtobufResponse {
	protobuf := &ProtobufResponse{
		Response: response,
	}
	if len(data) > 0 {
		protobuf.data = data[0]
	} else {
		protobuf.data = response.content
	}
	return protobuf
}

func (response *Response) Reader(reader io.Reader) *ReaderResponse {
	r := &ReaderResponse{
		Response: response,
		reader:   reader,
	}
	return r
}

func (response *Response) Redirect(location string) *RedirectResponse {
	redirect := &RedirectResponse{
		Response: response,
		location: location,
	}
	return redirect
}

func (response *Response) File(file string) *FileResponse {
	f := &FileResponse{
		ReaderResponse: &ReaderResponse{
			Response: response,
		},
		filename: file,
	}
	return f
}

func (response *Response) Stream(step func(io.Writer) bool) *StreamedResponse {
	s := &StreamedResponse{
		Response: response,
		step:     step,
	}
	return s
}
