package context

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/julienschmidt/httprouter"
	"github.com/wardonne/gopi/binding"
	"github.com/wardonne/gopi/context/formdata"
	"github.com/wardonne/gopi/support/container/maps"
	"github.com/wardonne/gopi/validation"
)

type Request struct {
	mu      *sync.RWMutex
	Request *http.Request
	Params  httprouter.Params
	Values  *maps.HashMap[string, any]
	form    validation.IValidateForm
}

func NewRequest(r *http.Request, p httprouter.Params) *Request {
	return &Request{
		mu:      new(sync.RWMutex),
		Request: r,
		Params:  p,
		Values:  maps.NewHashMap[string, any](),
	}
}

func (request *Request) Clone() *Request {
	return &Request{
		mu:      new(sync.RWMutex),
		Request: request.Request,
		Params:  request.Params,
		Values:  request.Values,
	}
}

func (request *Request) Set(key string, value any) {
	request.Values.Set(key, value)
}

func (request *Request) Get(key string) (any, bool) {
	request.mu.RLock()
	defer request.mu.RUnlock()
	if ok := request.Values.ContainsKey(key); ok {
		return request.Values.Get(key), true
	}
	return nil, false
}

func (request *Request) MustGet(key string) any {
	if value, exists := request.Get(key); !exists {
		panic(fmt.Errorf("Key \"%s\" does not exists in context Values", key))
	} else {
		return value
	}
}

func (request *Request) GetString(key string, defaultValue ...string) string {
	if value, exists := request.Get(key); exists {
		return value.(string)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return ""
	}
}

func (request *Request) GetInt(key string, defaultValue ...int) int {
	if value, exists := request.Get(key); exists {
		return value.(int)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return 0
	}
}

func (request *Request) GetUint(key string, defaultValue ...uint) uint {
	if value, exists := request.Get(key); exists {
		return value.(uint)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return 0
	}
}

func (request *Request) GetInt8(key string, defaultValue ...int8) int8 {
	if value, exists := request.Get(key); exists {
		return value.(int8)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return 0
	}
}

func (request *Request) GetUint8(key string, defaultValue ...uint8) uint8 {
	if value, exists := request.Get(key); exists {
		return value.(uint8)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return 0
	}
}

func (request *Request) GetInt16(key string, defaultValue ...int16) int16 {
	if value, exists := request.Get(key); exists {
		return value.(int16)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return 0
	}
}

func (request *Request) GetUint16(key string, defaultValue ...uint16) uint16 {
	if value, exists := request.Get(key); exists {
		return value.(uint16)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return 0
	}
}

func (request *Request) GetInt32(key string, defaultValue ...int32) int32 {
	if value, exists := request.Get(key); exists {
		return value.(int32)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return 0
	}
}

func (request *Request) GetUint32(key string, defaultValue ...uint32) uint32 {
	if value, exists := request.Get(key); exists {
		return value.(uint32)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return 0
	}
}

func (request *Request) GetInt64(key string, defaultValue ...int64) int64 {
	if value, exists := request.Get(key); exists {
		return value.(int64)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return 0
	}
}

func (request *Request) GetUint64(key string, defaultValue ...uint64) uint64 {
	if value, exists := request.Get(key); exists {
		return value.(uint64)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return 0
	}
}

func (request *Request) GetFloat32(key string, defaultValue ...float32) float32 {
	if value, exists := request.Get(key); exists {
		return value.(float32)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return 0
	}
}

func (request *Request) GetFloat64(key string, defaultValue ...float64) float64 {
	if value, exists := request.Get(key); exists {
		return value.(float64)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return 0
	}
}

func (request *Request) GetBool(key string, defaultValue ...bool) bool {
	if value, exists := request.Get(key); exists {
		return value.(bool)
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	} else {
		return false
	}
}

func (request *Request) Method() string {
	return request.Request.Method
}

func (request *Request) IsGet() bool {
	return request.Method() == http.MethodGet
}

func (request *Request) IsPost() bool {
	return request.Method() == http.MethodPost
}

func (request *Request) IsPut() bool {
	return request.Method() == http.MethodPut
}

func (request *Request) IsPatch() bool {
	return request.Method() == http.MethodPatch
}

func (request *Request) IsDelete() bool {
	return request.Method() == http.MethodDelete
}

func (request *Request) IsHead() bool {
	return request.Method() == http.MethodHead
}

func (request *Request) IsConnect() bool {
	return request.Method() == http.MethodConnect
}

func (request *Request) IsOptions() bool {
	return request.Method() == http.MethodOptions
}

func (request *Request) IsTrace() bool {
	return request.Method() == http.MethodTrace
}

func (request *Request) Host() string {
	return request.Request.Host
}

func (request *Request) RequestURI() string {
	return request.Request.RequestURI
}

func (request *Request) Path() string {
	return request.Request.URL.Path
}

func (request *Request) Query(key string, defaultValue ...string) formdata.Value {
	if values, exists := request.QueryArray(key); exists {
		return values[0]
	} else if len(defaultValue) > 0 {
		return formdata.NewValue(defaultValue[0])
	} else {
		return ""
	}
}

func (request *Request) QueryArray(key string) (formdata.Values, bool) {
	if request.Request.URL.Query().Has(key) {
		return formdata.NewValues(request.Request.URL.Query()[key]), true
	}
	return formdata.NewValues([]string{}), false
}

func (request *Request) QueryMap(key string) (map[string]formdata.Value, bool) {
	queries := request.Request.URL.Query()
	result := make(map[string]formdata.Value)
	exists := false
	for key, values := range queries {
		if i := strings.IndexByte(key, '['); i > 0 && key[:i] == key {
			if j := strings.IndexByte(key[i+1:], ']'); j > 0 {
				exists = true
				result[key[i+1:][:j]] = formdata.NewValue(values[0])
			}
		}
	}
	return result, exists
}

func (request *Request) PostForm(key string, defaultValue ...string) formdata.Value {
	if values, exists := request.PostFormArray(key); exists {
		return values[0]
	} else if len(defaultValue) > 0 {
		return formdata.NewValue(defaultValue[0])
	} else {
		return ""
	}
}

func (request *Request) PostFormArray(key string) (formdata.Values, bool) {
	if request.Request.PostForm.Has(key) {
		return formdata.NewValues(request.Request.PostForm[key]), true
	} else {
		return formdata.NewValues([]string{}), false
	}
}

func (request *Request) PostFormMap(key string) (map[string]formdata.Value, bool) {
	posts := request.Request.PostForm
	result := make(map[string]formdata.Value)
	exists := false
	for key, values := range posts {
		if i := strings.IndexByte(key, '['); i > 0 && key[:i] == key {
			if j := strings.IndexByte(key[i+1:], ']'); j > 0 {
				exists = true
				result[key[i+1:][:j]] = formdata.NewValue(values[0])
			}
		}
	}
	return result, exists
}

func (request *Request) Param(key string) formdata.Value {
	return formdata.NewValue(request.Params.ByName(key))
}

func (request *Request) File(name string) (*formdata.UploadedFile, error) {
	file, fileHeader, err := request.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	return formdata.NewUploadedFile(file, fileHeader)
}

func (request *Request) Files(name string) formdata.UploadedFiles {
	if err := request.Request.ParseMultipartForm(32 << 20); err != nil {
		panic(err)
	}
	fileHeaders := request.Request.MultipartForm.File[name]
	return formdata.NewUploadedFiles(fileHeaders)
}

func (request *Request) Header(key string, defaultValue ...string) formdata.Value {
	if headers, exists := request.HeaderArray(key); exists {
		return headers[0]
	} else if len(defaultValue) > 0 {
		return formdata.NewValue(defaultValue[0])
	} else {
		return ""
	}
}

func (request *Request) HeaderArray(key string) (formdata.Values, bool) {
	if headers, exists := request.Request.Header[key]; exists {
		return formdata.NewValues(headers), true
	} else {
		return formdata.NewValues([]string{}), false
	}
}

func (request *Request) ClientIP() string {
	xForwardedFor := request.Header("X-Forwarded-For", "").ToString()
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(request.Header("X-Real-IP", "").ToString())
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(request.Request.RemoteAddr); err == nil {
		return ip
	}
	return ""
}

func (request *Request) Bind(form validation.IValidateForm) error {
	var parsers = []binding.IParser{
		new(binding.URIParser),
	}
	if request.IsGet() {
		parsers = append(parsers, new(binding.FormParser))
	} else {
		parsers = append(parsers, new(binding.FormParser))
		contentType := request.Header("Content-Type")
		if contentType == MIMEJSON {
			parsers = append(parsers, new(binding.JSONParser))
		} else if contentType == MIMEXML {
			parsers = append(parsers, new(binding.XMLParser))
		} else if contentType == MIMEYAML {
			parsers = append(parsers, new(binding.YAMLParser))
		} else if contentType == MIMETOML {
			parsers = append(parsers, new(binding.TOMLParser))
		}
	}
	for _, parser := range parsers {
		if err := parser.Parse(request.Request, form); err != nil {
			return err
		}
	}
	request.form = form
	return nil
}

func (request *Request) Validated() validation.IValidateForm {
	return request.form
}
