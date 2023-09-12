package utils

import (
	"strconv"
	"time"
	"unsafe"
)

func StrToInt(s string) int64 {
	if value, err := strconv.ParseInt(s, 10, 64); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToUint(s string) uint64 {
	if value, err := strconv.ParseUint(s, 10, 64); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToFloat(s string) float64 {
	if value, err := strconv.ParseFloat(s, 64); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToBool(s string) bool {
	if value, err := strconv.ParseBool(s); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToDuration(s string) time.Duration {
	if value, err := time.ParseDuration(s); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToTime(s string, layout string) time.Time {
	if value, err := time.Parse(layout, s); err != nil {
		panic(err)
	} else {
		return value
	}
}

func StrToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
