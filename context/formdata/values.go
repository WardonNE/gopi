package formdata

import "time"

type Values []Value

func NewValues(values []string) Values {
	items := make([]Value, 0, len(values))
	for _, value := range values {
		items = append(items, NewValue(value))
	}
	return items
}

func (values Values) ToStrings() []string {
	items := make([]string, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToString())
	}
	return items
}

func (values Values) ToInts() []int {
	items := make([]int, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToInt())
	}
	return items
}

func (values Values) ToUints() []uint {
	items := make([]uint, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToUint())
	}
	return items
}

func (values Values) ToInt16s() []int16 {
	items := make([]int16, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToInt16())
	}
	return items
}

func (values Values) ToUint16s() []uint16 {
	items := make([]uint16, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToUint16())
	}
	return items
}

func (values Values) ToInt32s() []int32 {
	items := make([]int32, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToInt32())
	}
	return items
}

func (values Values) ToUint32s() []uint32 {
	items := make([]uint32, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToUint32())
	}
	return items
}

func (values Values) ToInt64s() []int64 {
	items := make([]int64, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToInt64())
	}
	return items
}

func (values Values) ToUint64s() []uint64 {
	items := make([]uint64, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToUint64())
	}
	return items
}

func (values Values) ToFloat32s() []float32 {
	items := make([]float32, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToFloat32())
	}
	return items
}

func (values Values) ToFloat64s() []float64 {
	items := make([]float64, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToFloat64())
	}
	return items
}

func (values Values) ToBools() []bool {
	items := make([]bool, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToBool())
	}
	return items
}

func (values Values) ToDuration() []time.Duration {
	items := make([]time.Duration, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToDuration())
	}
	return items
}

func (values Values) ToTimes(layout string) []time.Time {
	items := make([]time.Time, 0, len(values))
	for _, value := range values {
		items = append(items, value.ToTime(layout))
	}
	return items
}
