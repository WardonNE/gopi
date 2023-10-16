package formdata

import (
	"time"

	"github.com/wardonne/gopi/utils"
)

// Value is a simple superset of string that provides type convertion methods
type Value string

// String returns the string value
func (value Value) String() string {
	return string(value)
}

// ToInt converts the value to type
func (value Value) ToInt() int {
	return utils.StrToInt(value)
}

// ToUint converts the value to uint
func (value Value) ToUint() uint {
	return utils.StrToUint(value)
}

// ToInt8 converts the value to uint8
func (value Value) ToInt8() int8 {
	return utils.StrToInt8(value)
}

// ToUint8 converts the value to uint8
func (value Value) ToUint8() uint8 {
	return utils.StrToUint8(value)
}

// ToInt16 converts the value to int16
func (value Value) ToInt16() int16 {
	return utils.StrToInt16(value)
}

// ToUint16 converts the value to uint16
func (value Value) ToUint16() uint16 {
	return utils.StrToUint16(value)
}

// ToInt32 converts the value to int32
func (value Value) ToInt32() int32 {
	return utils.StrToInt32(value)
}

// ToUint32 converts the value to uint32
func (value Value) ToUint32() uint32 {
	return utils.StrToUint32(value)
}

// ToInt64 converts the value to int64
func (value Value) ToInt64() int64 {
	return utils.StrToInt64(value)
}

// ToUint64 converts the value to uint64
func (value Value) ToUint64() uint64 {
	return utils.StrToUint64(value)
}

// ToFloat32 converts the value to float32
func (value Value) ToFloat32() float32 {
	return utils.StrToFloat32(value)
}

// ToFloat64 converts the value to float64
func (value Value) ToFloat64() float64 {
	return utils.StrToFloat64(value)
}

// ToBool converts the value to bool
func (value Value) ToBool() bool {
	return utils.StrToBool(value)
}

// ToDuration converts the value to time.Duration
func (value Value) ToDuration() time.Duration {
	return utils.StrToDuration(value)
}

// ToTime converts the value to time.Time with specific layout
func (value Value) ToTime(layout string) time.Time {
	return utils.StrToTime(value, layout)
}
