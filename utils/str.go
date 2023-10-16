package utils

import (
	"strconv"
	"time"
	"unsafe"
)

func StrToInt[T ~string](s T) int {
	if value, err := strconv.Atoi(string(s)); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToUint[T ~string](s T) uint {
	return uint(StrToUint64[T](s))
}

func StrToInt8[T ~string](s T) int8 {
	return int8(StrToInt64[T](s))
}

func StrToUint8[T ~string](s T) uint8 {
	return uint8(StrToUint64[T](s))
}

func StrToInt16[T ~string](s T) int16 {
	return int16(StrToInt16[T](s))
}

func StrToUint16[T ~string](s T) uint16 {
	return uint16(StrToUint64[T](s))
}

func StrToInt32[T ~string](s T) int32 {
	return int32(StrToInt64[T](s))
}

func StrToUint32[T ~string](s T) uint32 {
	return uint32(StrToUint64[T](s))
}

func StrToInt64[T ~string](s T) int64 {
	if value, err := strconv.ParseInt(string(s), 10, 64); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToUint64[T ~string](s T) uint64 {
	if value, err := strconv.ParseUint(string(s), 10, 64); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToFloat32[T ~string](s T) float32 {
	return float32(StrToFloat64[T](s))
}

func StrToFloat64[T ~string](s T) float64 {
	if value, err := strconv.ParseFloat(string(s), 64); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToBool[T ~string](s T) bool {
	if value, err := strconv.ParseBool(string(s)); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToDuration[T ~string](s T) time.Duration {
	if value, err := time.ParseDuration(string(s)); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToTime[T ~string](s T, layout string) time.Time {
	if value, err := time.Parse(layout, string(s)); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToBytes[T ~string](s T) []byte {
	return unsafe.Slice(unsafe.StringData(string(s)), len(s))
}

func StrToRunes[T ~string](s T) []rune {
	return []rune(string(s))
}
