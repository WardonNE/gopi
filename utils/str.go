package utils

import (
	"strconv"
	"time"
	"unsafe"
)

// StrToInt converts string to int
func StrToInt[T ~string](s T) int {
	if value, err := strconv.Atoi(string(s)); err != nil {
		panic(err)
	} else {
		return value
	}
}

// StrToUint converts string to uint
func StrToUint[T ~string](s T) uint {
	return uint(StrToUint64[T](s))
}

// StrToInt8 converts string to int8
func StrToInt8[T ~string](s T) int8 {
	return int8(StrToInt64[T](s))
}

// StrToUint8 converts string to uint8
func StrToUint8[T ~string](s T) uint8 {
	return uint8(StrToUint64[T](s))
}

// StrToInt16 converts string to int16
func StrToInt16[T ~string](s T) int16 {
	return int16(StrToInt16[T](s))
}

// StrToUint16 converts string to uint16
func StrToUint16[T ~string](s T) uint16 {
	return uint16(StrToUint64[T](s))
}

// StrToInt32 converts string to int32
func StrToInt32[T ~string](s T) int32 {
	return int32(StrToInt64[T](s))
}

// StrToUint32 converts string to uint32
func StrToUint32[T ~string](s T) uint32 {
	return uint32(StrToUint64[T](s))
}

// StrToInt64 converts string to int64
func StrToInt64[T ~string](s T) int64 {
	if value, err := strconv.ParseInt(string(s), 10, 64); err != nil {
		panic(err)
	} else {
		return value
	}
}

// StrToUint64 converts string to uint64
func StrToUint64[T ~string](s T) uint64 {
	if value, err := strconv.ParseUint(string(s), 10, 64); err != nil {
		panic(err)
	} else {
		return value
	}
}

// StrToFloat32 converts string to float32
func StrToFloat32[T ~string](s T) float32 {
	return float32(StrToFloat64[T](s))
}

// StrToFloat64 converts string to float64
func StrToFloat64[T ~string](s T) float64 {
	if value, err := strconv.ParseFloat(string(s), 64); err != nil {
		panic(err)
	} else {
		return value
	}
}

// StrToBool converts string to bool
func StrToBool[T ~string](s T) bool {
	if value, err := strconv.ParseBool(string(s)); err != nil {
		panic(err)
	} else {
		return value
	}
}

// StrToDuration converts string to [time.Duration]
func StrToDuration[T ~string](s T) time.Duration {
	if value, err := time.ParseDuration(string(s)); err != nil {
		panic(err)
	} else {
		return value
	}
}

// StrToTime converts string to [time.Time]
func StrToTime[T ~string](s T, layout string) time.Time {
	if value, err := time.Parse(layout, string(s)); err != nil {
		panic(err)
	} else {
		return value
	}
}

// StrToBytes converts string to byte slice
func StrToBytes[T ~string](s T) []byte {
	return unsafe.Slice(unsafe.StringData(string(s)), len(s))
}

// StrToRunes converts string to rune slice
func StrToRunes[T ~string](s T) []rune {
	return []rune(string(s))
}
