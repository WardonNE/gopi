package formdata

import (
	"strconv"
	"time"
)

type Value string

func NewValue(value string) Value {
	return Value(value)
}

func (value Value) ToString() string {
	return string(value)
}

func (value Value) ToInt() int {
	if v, err := strconv.ParseInt(value.ToString(), 10, 64); err != nil {
		panic(err)
	} else {
		return int(v)
	}
}

func (value Value) ToUint() uint {
	if v, err := strconv.ParseUint(value.ToString(), 10, 64); err != nil {
		panic(err)
	} else {
		return uint(v)
	}
}

func (value Value) ToInt8() int8 {
	if v, err := strconv.ParseInt(value.ToString(), 10, 64); err != nil {
		panic(err)
	} else {
		return int8(v)
	}
}

func (value Value) ToUint8() uint8 {
	if v, err := strconv.ParseUint(value.ToString(), 10, 64); err != nil {
		panic(err)
	} else {
		return uint8(v)
	}
}

func (value Value) ToInt16() int16 {
	if v, err := strconv.ParseInt(value.ToString(), 10, 64); err != nil {
		panic(err)
	} else {
		return int16(v)
	}
}

func (value Value) ToUint16() uint16 {
	if v, err := strconv.ParseUint(value.ToString(), 10, 64); err != nil {
		panic(err)
	} else {
		return uint16(v)
	}
}

func (value Value) ToInt32() int32 {
	if v, err := strconv.ParseInt(value.ToString(), 10, 64); err != nil {
		panic(err)
	} else {
		return int32(v)
	}
}

func (value Value) ToUint32() uint32 {
	if v, err := strconv.ParseUint(value.ToString(), 10, 64); err != nil {
		panic(err)
	} else {
		return uint32(v)
	}
}

func (value Value) ToInt64() int64 {
	if v, err := strconv.ParseInt(value.ToString(), 10, 64); err != nil {
		panic(err)
	} else {
		return v
	}
}

func (value Value) ToUint64() uint64 {
	if v, err := strconv.ParseUint(value.ToString(), 10, 64); err != nil {
		panic(err)
	} else {
		return v
	}
}

func (value Value) ToFloat32() float32 {
	if v, err := strconv.ParseFloat(value.ToString(), 64); err != nil {
		panic(err)
	} else {
		return float32(v)
	}
}

func (value Value) ToFloat64() float64 {
	if v, err := strconv.ParseFloat(value.ToString(), 64); err != nil {
		panic(err)
	} else {
		return v
	}
}

func (value Value) ToBool() bool {
	if v, err := strconv.ParseBool(value.ToString()); err != nil {
		panic(err)
	} else {
		return v
	}
}

func (value Value) ToDuration() time.Duration {
	if v, err := time.ParseDuration(value.ToString()); err != nil {
		panic(err)
	} else {
		return v
	}
}

func (value Value) ToTime(layout string) time.Time {
	if v, err := time.Parse(layout, value.ToString()); err != nil {
		panic(err)
	} else {
		return v
	}
}
