package context

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/julienschmidt/httprouter"
	"github.com/wardonne/gopi/support/maps"
	"github.com/wardonne/gopi/utils"
	"github.com/wardonne/gopi/validation"
	"github.com/wardonne/gopi/web/binding"
	"github.com/wardonne/gopi/web/context/formdata"
)

// Request http request with custom attributes and methods
type Request struct {
	mu      *sync.RWMutex
	Request *http.Request
	Params  httprouter.Params
	Values  *maps.HashMap[string, any]
	form    validation.IValidateForm
}

// NewRequest creates a new [Request] instance with http.Request and httprouter.Params
func NewRequest(r *http.Request, p httprouter.Params) *Request {
	return &Request{
		mu:      new(sync.RWMutex),
		Request: r,
		Params:  p,
		Values:  maps.NewHashMap[string, any](),
	}
}

// Clone clones a new [Request] instance from current one
func (request *Request) Clone() *Request {
	return &Request{
		mu:      new(sync.RWMutex),
		Request: request.Request,
		Params:  request.Params,
		Values:  request.Values,
	}
}

// Set sets a value with specific key to current request
func (request *Request) Set(key string, value any) {
	request.Values.Set(key, value)
}

// Get gets a value from current request by specific key
//
// if the specific key is not exist, it returns nil and false
func (request *Request) Get(key string) (any, bool) {
	request.mu.RLock()
	defer request.mu.RUnlock()
	if ok := request.Values.ContainsKey(key); ok {
		return request.Values.Get(key), true
	}
	return nil, false
}

// MustGet gets a value from current request by specific key
//
// if the specific key is not exist, it will panic
func (request *Request) MustGet(key string) any {
	if value, exists := request.Get(key); !exists {
		panic(fmt.Errorf("Key \"%s\" does not exists in context Values", key))
	} else {
		return value
	}
}

// GetString gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetString(key string, defaultValue ...string) *string {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(string))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetInt gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetInt(key string, defaultValue ...int) *int {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(int))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetUint gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetUint(key string, defaultValue ...uint) *uint {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(uint))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetInt8 gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetInt8(key string, defaultValue ...int8) *int8 {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(int8))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetUint8 gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetUint8(key string, defaultValue ...uint8) *uint8 {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(uint8))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetInt16 gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetInt16(key string, defaultValue ...int16) *int16 {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(int16))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetUint16 gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetUint16(key string, defaultValue ...uint16) *uint16 {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(uint16))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetInt32 gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetInt32(key string, defaultValue ...int32) *int32 {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(int32))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetUint32 gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetUint32(key string, defaultValue ...uint32) *uint32 {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(uint32))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetInt64 gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetInt64(key string, defaultValue ...int64) *int64 {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(int64))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetUint64 gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetUint64(key string, defaultValue ...uint64) *uint64 {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(uint64))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetFloat32 gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetFloat32(key string, defaultValue ...float32) *float32 {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(float32))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetFloat64 gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetFloat64(key string, defaultValue ...float64) *float64 {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(float64))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// GetBool gets a value from current request by specific key
//
// if the specific key is not exist, it will return the default value
//
// if the default value is not provided, it will return nil
func (request *Request) GetBool(key string, defaultValue ...bool) *bool {
	if value, exists := request.Get(key); exists {
		return utils.Ptr(value.(bool))
	} else if len(defaultValue) > 0 {
		return &defaultValue[0]
	} else {
		return nil
	}
}

// Method returns the http request method
func (request *Request) Method() string {
	return request.Request.Method
}

// IsGet returns if the http request method is GET
func (request *Request) IsGet() bool {
	return request.Method() == http.MethodGet
}

// IsPost returns if the http request method is POST
func (request *Request) IsPost() bool {
	return request.Method() == http.MethodPost
}

// IsPut returns if the http request method is PUT
func (request *Request) IsPut() bool {
	return request.Method() == http.MethodPut
}

// IsPatch returns if the http request method is PATCH
func (request *Request) IsPatch() bool {
	return request.Method() == http.MethodPatch
}

// IsDelete returns if the http request method is DELETE
func (request *Request) IsDelete() bool {
	return request.Method() == http.MethodDelete
}

// IsHead returns if the http request method is HEAD
func (request *Request) IsHead() bool {
	return request.Method() == http.MethodHead
}

// IsConnect returns if the http request method is CONNECT
func (request *Request) IsConnect() bool {
	return request.Method() == http.MethodConnect
}

// IsOptions returns if the http request method is OPTIONS
func (request *Request) IsOptions() bool {
	return request.Method() == http.MethodOptions
}

// IsTrace returns if the http request method is GET
func (request *Request) IsTrace() bool {
	return request.Method() == http.MethodTrace
}

// Host returns the request's host
func (request *Request) Host() string {
	return request.Request.Host
}

// RequestURI returns the request's uri
func (request *Request) RequestURI() string {
	return request.Request.RequestURI
}

// Path returns the request's url path
func (request *Request) Path() string {
	return request.Request.URL.Path
}

// Query returns the query value by specific key
//
// if the specific key is not exist, default value will be returned
//
// if default value is not provided, nil will be returned
func (request *Request) Query(key string, defaultValue ...string) *formdata.Value {
	if values, exists := request.QueryArray(key); exists {
		return &values[0]
	} else if len(defaultValue) > 0 {
		return utils.Ptr(formdata.Value(defaultValue[0]))
	} else {
		return nil
	}
}

// QueryArray returns the query values by specific key
//
// if the specific key is not exist, the second return value will be false
func (request *Request) QueryArray(key string) (formdata.Values, bool) {
	if request.Request.URL.Query().Has(key) {
		return formdata.NewValues(request.Request.URL.Query()[key]), true
	}
	return formdata.NewValues([]string{}), false
}

// QueryMap returns the query values by specific key as a map
//
// if the specific key is not exist, the second return value will be false
func (request *Request) QueryMap(key string) (map[string]formdata.Value, bool) {
	queries := request.Request.URL.Query()
	result := make(map[string]formdata.Value)
	exists := false
	for key, values := range queries {
		if i := strings.IndexByte(key, '['); i > 0 && key[:i] == key {
			if j := strings.IndexByte(key[i+1:], ']'); j > 0 {
				exists = true
				result[key[i+1:][:j]] = formdata.Value(values[0])
			}
		}
	}
	return result, exists
}

// PostForm returns the form value by specific key
//
// if the specific key is not exist, default value will be returned
//
// if default value is not provided, nil will be returned
func (request *Request) PostForm(key string, defaultValue ...string) *formdata.Value {
	if values, exists := request.PostFormArray(key); exists {
		return &values[0]
	} else if len(defaultValue) > 0 {
		return utils.Ptr(formdata.Value(defaultValue[0]))
	} else {
		return nil
	}
}

// PostFormArray returns the form values by specific key
//
// if the specific key is not exist, the second return value will be false
func (request *Request) PostFormArray(key string) (formdata.Values, bool) {
	if request.Request.PostForm.Has(key) {
		return formdata.NewValues(request.Request.PostForm[key]), true
	}
	return formdata.NewValues([]string{}), false
}

// PostFormMap returns the form values by specific key as a map
//
// if the specific key is not exist, the second return value will be false
func (request *Request) PostFormMap(key string) (map[string]formdata.Value, bool) {
	posts := request.Request.PostForm
	result := make(map[string]formdata.Value)
	exists := false
	for key, values := range posts {
		if i := strings.IndexByte(key, '['); i > 0 && key[:i] == key {
			if j := strings.IndexByte(key[i+1:], ']'); j > 0 {
				exists = true
				result[key[i+1:][:j]] = formdata.Value(values[0])
			}
		}
	}
	return result, exists
}

// Param returns path param value of the specific key
func (request *Request) Param(key string) formdata.Value {
	return formdata.Value(request.Params.ByName(key))
}

// File returns an instance of [formdata.UploadedFile] of the specific name
func (request *Request) File(name string) (*formdata.UploadedFile, error) {
	file, fileHeader, err := request.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	return formdata.NewUploadedFile(file, fileHeader)
}

// Files returns an instance of [formdata.UploadedFiles] of the specific name
func (request *Request) Files(name string) formdata.UploadedFiles {
	if err := request.Request.ParseMultipartForm(32 << 20); err != nil {
		panic(err)
	}
	fileHeaders := request.Request.MultipartForm.File[name]
	return formdata.NewUploadedFiles(fileHeaders)
}

// Header returns the header value by specific key
//
// if the specific key is not exist, default value will be returned
//
// if default value is not provided, nil will be returned
func (request *Request) Header(key string, defaultValue ...string) *formdata.Value {
	if headers, exists := request.HeaderArray(key); exists {
		return &headers[0]
	} else if len(defaultValue) > 0 {
		return utils.Ptr(formdata.Value(defaultValue[0]))
	} else {
		return nil
	}
}

// HeaderArray returns the header values by specific key
//
// if the specific key is not exist, the second return value will be false
func (request *Request) HeaderArray(key string) (formdata.Values, bool) {
	if headers, exists := request.Request.Header[key]; exists {
		return formdata.NewValues(headers), true
	}
	return formdata.NewValues([]string{}), false
}

// ClientIP returns the ip address of client
func (request *Request) ClientIP() string {
	xForwardedFor := request.Header("X-Forwarded-For", "").String()
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(request.Header("X-Real-IP", "").String())
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(request.Request.RemoteAddr); err == nil {
		return ip
	}
	return ""
}

// Bind parses request into an instance of [validation.IValidateForm]
//
// if bindings is not provided, it uses [binding.Form] and if the request method is POST
// it will use binding implement according to Content-Type header
func (request *Request) Bind(form validation.IValidateForm, bindings ...binding.Binding) error {
	if len(bindings) == 0 {
		bindings = append(bindings, binding.Form)
		if h := request.Header("Content-Type"); h != nil {
			contentType := h.String()
			if contentType == MIMEJSON {
				bindings = append(bindings, binding.JSON)
			} else if contentType == MIMEXML {
				bindings = append(bindings, binding.XML)
			} else if contentType == MIMEYAML {
				bindings = append(bindings, binding.YAML)
			} else if contentType == MIMETOML {
				bindings = append(bindings, binding.TOML)
			}
		}
	}
	for _, binding := range bindings {
		if err := binding(request.Request, form); err != nil {
			return err
		}
	}
	request.form = form
	return nil
}

// Validated returns the validated instance of [validation.IValidateForm]
func (request *Request) Validated() validation.IValidateForm {
	return request.form
}
