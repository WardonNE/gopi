package context

import (
	"fmt"
	"io"
	"net/http"
)

var _ IResponse = new(Response)

// IResponse the response interface
type IResponse interface {
	// SetStatusCode sets the response http status code
	SetStatusCode(statusCode int)
	// StatusCode gets the response http status code
	StatusCode() int
	// SetContent sets the response body content
	SetContent(content any)
	// Content gets the response body content
	Content() any
	// SetHeader sets the response header, if replace is true, it will replace the existing header,
	// and if replace is false, it appends the new value into existing header
	SetHeader(key, header string, replace ...bool)
	// SetHeaders sets headers map to the response
	SetHeaders(headers map[string]string)
	// HasHeader returns if the specific key is exist
	HasHeader(key string) bool
	// Header returns header value of specific header
	Header(key string) string
	// Headers returns all headers as [http.Header]
	Headers() http.Header
	// SetCookie sets cookie to response
	SetCookie(cookie *http.Cookie)
	// Cookies returns all response cookies
	Cookies() []*http.Cookie
	// Send sends the response
	Send(w http.ResponseWriter, r *http.Request)
}

// Response is the basic response implement
type Response struct {
	headers    http.Header
	cookies    []*http.Cookie
	statusCode int
	content    any
}

// NewResponse creates a new [Response] instance
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

// SetStatusCode sets the response http status code
func (response *Response) SetStatusCode(statusCode int) {
	if statusCode < 100 || statusCode > 600 {
		panic(fmt.Errorf("HTTP status code `%d` is invalid", statusCode))
	}
	response.statusCode = statusCode
}

// StatusCode gets the response http status code
func (response *Response) StatusCode() int {
	return response.statusCode
}

// SetContent sets the response body content
func (response *Response) SetContent(content any) {
	response.content = content
}

// Content gets the response body content
func (response *Response) Content() any {
	return response.content
}

// SetHeader sets the response header, if replace is true, it will replace the existing header,
// and if replace is false, it appends the new value into existing header
func (response *Response) SetHeader(key, value string, replace ...bool) {
	if len(replace) == 0 || (len(replace) > 0 && replace[0]) {
		response.headers.Set(key, value)
	} else {
		response.headers.Add(key, value)
	}
}

// SetHeaders sets headers map to the response
func (response *Response) SetHeaders(headers map[string]string) {
	for header, value := range headers {
		response.headers.Set(header, value)
	}
}

// HasHeader returns if the specific key is exist
func (response *Response) HasHeader(key string) bool {
	h, ok := response.headers[key]
	return ok && len(h) > 0
}

// Header returns header value of specific header
func (response *Response) Header(key string) string {
	return response.headers.Get(key)
}

// Headers returns all headers as [http.Header]
func (response *Response) Headers() http.Header {
	return response.headers
}

// SetCookie sets cookie to response
func (response *Response) SetCookie(cookie *http.Cookie) {
	response.cookies = append(response.cookies, cookie)
}

// Cookies returns all response cookies
func (response *Response) Cookies() []*http.Cookie {
	return response.cookies
}

// Send sends the response
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

// JSON returns a JSON response implement
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

// XML returns a XML response implement
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

// YAML returns a YAML response implement
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

// TOML returns a TOML response implement
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

// Protobuf returns a ProtoBuf response implement
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

// Reader returns a Reader response implement
func (response *Response) Reader(reader io.Reader) *ReaderResponse {
	r := &ReaderResponse{
		Response: response,
		reader:   reader,
	}
	return r
}

// Redirect returns a Redirect response implement
func (response *Response) Redirect(location string) *RedirectResponse {
	redirect := &RedirectResponse{
		Response: response,
		location: location,
	}
	return redirect
}

// File returns a File response implement
func (response *Response) File(file string) *FileResponse {
	f := &FileResponse{
		ReaderResponse: &ReaderResponse{
			Response: response,
		},
		filename: file,
	}
	return f
}

// Stream returns a Stream response implement
func (response *Response) Stream(step func(io.Writer) bool) *StreamedResponse {
	s := &StreamedResponse{
		Response: response,
		step:     step,
	}
	return s
}
