package binding

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"reflect"
	"strings"
	"time"

	"github.com/wardonne/gopi/support/utils"
	"github.com/wardonne/gopi/web/context/formdata"
)

func formtostruct(form *multipart.Form, container any, tag string) error {
	containerType := reflect.TypeOf(container)
	if containerType.Kind() != reflect.Ptr {
		return fmt.Errorf("Parser(non-pointer: (%s)", containerType.String())
	}
	containerType = containerType.Elem()
	containerValue := reflect.ValueOf(container)
	numField := containerType.NumField()
	for i := 0; i < numField; i++ {
		field := containerType.Field(i)
		tagValue, ok := field.Tag.Lookup(tag)
		if !ok {
			continue
		}
		tagValue = strings.TrimSpace(tagValue)
		if tagValue == "-" || tagValue == "" {
			continue
		}
		v := createValue(field.Type, form, tagValue, field.Tag)
		if v.IsValid() && !v.IsZero() {
			containerValue.Elem().Field(i).Set(v)
		}
	}
	return nil
}

func isFileHeader(fieldType reflect.Type) bool {
	return fieldType == reflect.TypeOf(multipart.FileHeader{})
}

func isUploadedFile(fieldType reflect.Type) bool {
	return fieldType == reflect.TypeOf(formdata.UploadedFile{})
}

func isFileHeaderSlice(fieldType reflect.Type) bool {
	if fieldType.Kind() != reflect.Slice {
		return false
	}
	itemType := fieldType.Elem()
	if itemType.Kind() == reflect.Ptr {
		itemType = itemType.Elem()
	}
	return isFileHeader(itemType)
}

func isFileHeaderArray(fieldType reflect.Type) bool {
	if fieldType.Kind() != reflect.Array {
		return false
	}
	itemType := fieldType.Elem()
	if itemType.Kind() == reflect.Ptr {
		itemType = itemType.Elem()
	}
	return isFileHeader(itemType)
}

func isUploadedFileSlice(fieldType reflect.Type) bool {
	if fieldType.Kind() != reflect.Slice {
		return false
	}
	itemType := fieldType.Elem()
	if itemType.Kind() == reflect.Ptr {
		itemType = itemType.Elem()
	}
	return isUploadedFile(itemType)
}

func isUploadedFileArray(fieldType reflect.Type) bool {
	if fieldType.Kind() != reflect.Array {
		return false
	}
	itemType := fieldType.Elem()
	if itemType.Kind() == reflect.Ptr {
		itemType = itemType.Elem()
	}
	return isUploadedFile(itemType)
}

func hasFile(form *multipart.Form, key string) bool {
	if _, ok := form.File[key]; ok {
		return true
	}
	return false
}

func hasValue(form *multipart.Form, key string) bool {
	if _, ok := form.Value[key]; ok {
		return true
	}
	return false
}

func has(form *multipart.Form, key string) (exists bool, hasfile bool, hasvalue bool) {
	hasfile = hasFile(form, key)
	hasvalue = hasValue(form, key)
	exists = hasfile || hasvalue
	return
}

func createValue(fieldType reflect.Type, form *multipart.Form, key string, tag reflect.StructTag) reflect.Value {
	isPtr := fieldType.Kind() == reflect.Ptr
	if isPtr {
		fieldType = fieldType.Elem()
	}
	var v reflect.Value
	if exists, hasfile, _ := has(form, key); exists {
		if hasfile {
			if isFileHeader(fieldType) {
				files := form.File[key]
				if len(files) > 0 {
					v = reflect.ValueOf(files[0]).Elem()
				}
			} else if isUploadedFile(fieldType) {
				files := form.File[key]
				if len(files) > 0 {
					fh := files[0]
					if f, err := fh.Open(); err != nil {
						panic(err)
					} else if file, err := formdata.NewUploadedFile(f, fh); err != nil {
						panic(err)
					} else {
						v = reflect.ValueOf(file).Elem()
					}
				}
			} else if isFileHeaderSlice(fieldType) {
				files := form.File[key]
				slice := reflect.MakeSlice(fieldType, 0, 0)
				if len(files) == 0 {
					v = slice
				} else {
					for _, file := range files {
						item := reflect.ValueOf(file)
						if fieldType.Elem().Kind() != reflect.Ptr {
							item = item.Elem()
						}
						slice = reflect.Append(slice, item)
					}
					v = slice
				}
			} else if isFileHeaderArray(fieldType) {
				files := form.File[key]
				if fieldType.Len() != len(files) {
					panic("file length isn't equal to array length")
				}
				array := reflect.New(fieldType).Elem()
				for index, file := range files {
					item := reflect.ValueOf(file)
					if fieldType.Elem().Kind() != reflect.Ptr {
						item = item.Elem()
					}
					array.Index(index).Set(item)
				}
				v = array
			} else if isUploadedFileSlice(fieldType) {
				files := form.File[key]
				if fieldType.Elem().Kind() == reflect.Ptr {
					v = reflect.ValueOf(formdata.NewUploadedFiles(files))
				} else {
					slice := reflect.MakeSlice(fieldType, 0, 0)
					uploadedFiles := formdata.NewUploadedFiles(files)
					for _, uploadedFile := range uploadedFiles {
						slice = reflect.Append(slice, reflect.ValueOf(uploadedFile).Elem())
					}
					v = slice
				}
			} else if isUploadedFileArray(fieldType) {
				files := form.File[key]
				if fieldType.Len() != len(files) {
					panic("file length isn't equal to array length")
				}
				array := reflect.New(fieldType).Elem()
				uploadedFiles := formdata.NewUploadedFiles(files)
				for index, uploadedFile := range uploadedFiles {
					item := reflect.ValueOf(uploadedFile)
					if fieldType.Kind() != reflect.Ptr {
						item = item.Elem()
					}
					array.Index(index).Set(item)
				}
				v = array
			} else {
				panic("unsupported file struct")
			}
		} else {
			values := form.Value[key]
			if len(values) == 0 {

			} else if fieldType == reflect.TypeOf(time.Duration(0)) {
				v = reflect.ValueOf(utils.StrToDuration(values[0]))
			} else if fieldType == reflect.TypeOf(time.Time{}) {
				dateFormat := time.DateTime
				if tf, ok := tag.Lookup("date_format"); ok && tf != "" {
					dateFormat = tf
				}
				v = reflect.ValueOf(utils.StrToTime(values[0], dateFormat))
			} else if fieldType.Kind() == reflect.String {
				v = reflect.ValueOf(values[0])
			} else if fieldType.Kind() == reflect.Int || fieldType.Kind() == reflect.Int8 || fieldType.Kind() == reflect.Int16 || fieldType.Kind() == reflect.Int32 || fieldType.Kind() == reflect.Int64 {
				v = reflect.ValueOf(utils.StrToInt64(values[0])).Convert(fieldType)
			} else if fieldType.Kind() == reflect.Uint || fieldType.Kind() == reflect.Uint8 || fieldType.Kind() == reflect.Uint16 || fieldType.Kind() == reflect.Uint32 || fieldType.Kind() == reflect.Uint64 {
				v = reflect.ValueOf(utils.StrToUint64(values[0])).Convert(fieldType)
			} else if fieldType.Kind() == reflect.Float32 || fieldType.Kind() == reflect.Float64 {
				v = reflect.ValueOf(utils.StrToFloat64(values[0])).Convert(fieldType)
			} else if fieldType.Kind() == reflect.Bool {
				v = reflect.ValueOf(utils.StrToBool(values[0]))
			} else if fieldType.Kind() == reflect.Slice {
				slice := reflect.MakeSlice(fieldType, 0, 0)
				for _, item := range values {
					slice = reflect.Append(slice, createValue(slice.Type().Elem(), &multipart.Form{
						Value: map[string][]string{key: {item}},
					}, key, tag))
				}
				v = slice
			} else if fieldType.Kind() == reflect.Array {
				if len(values) != fieldType.Len() {
					panic("value length isn't equal to array length")
				}
				array := reflect.New(fieldType).Elem()
				for index, item := range values {
					array.Index(index).Set(createValue(array.Type().Elem(), &multipart.Form{
						Value: map[string][]string{key: {item}},
					}, key, tag))
				}
				v = array
			} else if fieldType.Kind() == reflect.Map {
				fieldValue := reflect.New(fieldType)
				v = reflect.ValueOf(json.Unmarshal(utils.StrToBytes(values[0]), fieldValue.Interface()))
			} else if fieldType.Kind() == reflect.Struct {
				fieldValue := reflect.New(fieldType)
				v = reflect.ValueOf(json.Unmarshal(utils.StrToBytes(values[0]), fieldValue.Interface()))
			} else {
				panic("unsupported field type")
			}
		}
	}
	if isPtr {
		if v.CanAddr() {
			v = v.Addr()
		} else if v.IsValid() && !v.IsZero() {
			valueAddr := reflect.New(fieldType)
			valueAddr.Elem().Set(v)
			v = valueAddr
		}
	}
	return v
}
