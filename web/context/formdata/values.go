package formdata

import "time"

// Values is a simple superset of [][Value]
type Values []Value

// NewValues creates a new [Values] instance from a string slice
func NewValues(values []string) Values {
	items := make([]Value, 0, len(values))
	for _, value := range values {
		items = append(items, Value(value))
	}
	return items
}

// ToStrings returns values as []string
func (values Values) ToStrings() []string {
	items := make([]string, 0, len(values))
	for _, value := range values {
		items = append(items, value.String())
	}
	return items
}

// ToInts returns values as []int
func (values Values) ToInts() []int {
	items := make([]int, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToInt())
	}
	return items
}

// ToUints returns values as []uint
func (values Values) ToUints() []uint {
	items := make([]uint, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToUint())
	}
	return items
}

// ToInt8s returns values as []int
func (values Values) ToInt8s() []int8 {
	items := make([]int8, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToInt8())
	}
	return items
}

// ToUint8s returns values as []uint8
func (values Values) ToUint8s() []uint8 {
	items := make([]uint8, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToUint8())
	}
	return items
}

// ToInt16s returns values as []int16
func (values Values) ToInt16s() []int16 {
	items := make([]int16, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToInt16())
	}
	return items
}

// ToUint16s returns values as []uint16
func (values Values) ToUint16s() []uint16 {
	items := make([]uint16, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToUint16())
	}
	return items
}

// ToInt32s returns values as []int32
func (values Values) ToInt32s() []int32 {
	items := make([]int32, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToInt32())
	}
	return items
}

// ToUint32s returns values as []uint32
func (values Values) ToUint32s() []uint32 {
	items := make([]uint32, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToUint32())
	}
	return items
}

// ToInt64s returns values as []int64
func (values Values) ToInt64s() []int64 {
	items := make([]int64, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToInt64())
	}
	return items
}

// ToUint64s returns values as []uint64
func (values Values) ToUint64s() []uint64 {
	items := make([]uint64, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToUint64())
	}
	return items
}

// ToFloat32s returns values as []float32
func (values Values) ToFloat32s() []float32 {
	items := make([]float32, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToFloat32())
	}
	return items
}

// ToFloat64s returns values as []float64
func (values Values) ToFloat64s() []float64 {
	items := make([]float64, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToFloat64())
	}
	return items
}

// ToBools returns values as []bool
func (values Values) ToBools() []bool {
	items := make([]bool, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToBool())
	}
	return items
}

// ToDurations returns values as []time.Duration
func (values Values) ToDurations() []time.Duration {
	items := make([]time.Duration, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToDuration())
	}
	return items
}

// ToTimes returns values as []time.Time
func (values Values) ToTimes(layout string) []time.Time {
	items := make([]time.Time, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToTime(layout))
	}
	return items
}
