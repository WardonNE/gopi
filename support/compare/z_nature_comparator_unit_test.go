package compare

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

type sortableSlice[T constraints.Ordered] []T

func (a sortableSlice[T]) Len() int           { return len(a) }
func (a sortableSlice[T]) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortableSlice[T]) Less(i, j int) bool { return a[i] < a[j] }

func createNatureComparator[T constraints.Ordered](t *testing.T, desc bool) *NatureComparator[T] {
	t.Helper()
	return NewNatureComparator[T](desc)
}

func createSortInterface[T constraints.Ordered](t *testing.T, elements []T) sort.Interface {
	t.Helper()
	return sortableSlice[T](elements)
}

func TestNatureComparator_Int(t *testing.T) {
	c1 := createNatureComparator[int](t, false)
	// s := []int{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s1 := []int{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s2 := []int{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Ints(s2)
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[int](t, true)
	s3 := []int{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s4 := []int{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(sort.IntSlice(s4)))
	assert.Equal(t, s4, s3)
}

func TestNatureComparator_Int8(t *testing.T) {
	c1 := createNatureComparator[int8](t, false)
	// s := []int{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s1 := []int8{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s2 := []int8{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Sort(createSortInterface[int8](t, s2))
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[int8](t, true)
	s3 := []int8{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s4 := []int8{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(createSortInterface[int8](t, s4)))
	assert.Equal(t, s4, s3)
}

func TestNatureComparator_Int16(t *testing.T) {
	c1 := createNatureComparator[int16](t, false)
	// s := []int{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s1 := []int16{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s2 := []int16{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Sort(createSortInterface[int16](t, s2))
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[int16](t, true)
	s3 := []int16{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s4 := []int16{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(createSortInterface[int16](t, s4)))
	assert.Equal(t, s4, s3)
}

func TestNatureComparator_Int32(t *testing.T) {
	c1 := createNatureComparator[int32](t, false)
	// s := []int{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s1 := []int32{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s2 := []int32{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Sort(createSortInterface[int32](t, s2))
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[int32](t, true)
	s3 := []int32{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s4 := []int32{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(createSortInterface[int32](t, s4)))
	assert.Equal(t, s4, s3)
}

func TestNatureComparator_Int64(t *testing.T) {
	c1 := createNatureComparator[int64](t, false)
	// s := []int{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s1 := []int64{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s2 := []int64{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Sort(createSortInterface[int64](t, s2))
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[int64](t, true)
	s3 := []int64{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s4 := []int64{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(createSortInterface[int64](t, s4)))
	assert.Equal(t, s4, s3)
}

func TestNatureComparator_Uint(t *testing.T) {
	c1 := createNatureComparator[uint](t, false)
	s1 := []uint{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s2 := []uint{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Sort(createSortInterface[uint](t, s2))
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[uint](t, true)
	s3 := []uint{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s4 := []uint{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(createSortInterface[uint](t, s4)))
	assert.Equal(t, s4, s3)
}

func TestNatureComparator_Uint8(t *testing.T) {
	c1 := createNatureComparator[uint8](t, false)
	// s := []int{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s1 := []uint8{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s2 := []uint8{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Sort(createSortInterface[uint8](t, s2))
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[uint8](t, true)
	s3 := []uint8{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s4 := []uint8{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(createSortInterface[uint8](t, s4)))
	assert.Equal(t, s4, s3)
}

func TestNatureComparator_Uint16(t *testing.T) {
	c1 := createNatureComparator[uint16](t, false)
	// s := []int{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s1 := []uint16{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s2 := []uint16{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Sort(createSortInterface[uint16](t, s2))
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[uint16](t, true)
	s3 := []uint16{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s4 := []uint16{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(createSortInterface[uint16](t, s4)))
	assert.Equal(t, s4, s3)
}

func TestNatureComparator_Uint32(t *testing.T) {
	c1 := createNatureComparator[uint32](t, false)
	// s := []int{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s1 := []uint32{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s2 := []uint32{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Sort(createSortInterface[uint32](t, s2))
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[uint32](t, true)
	s3 := []uint32{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s4 := []uint32{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(createSortInterface[uint32](t, s4)))
	assert.Equal(t, s4, s3)
}

func TestNatureComparator_Uint64(t *testing.T) {
	c1 := createNatureComparator[uint64](t, false)
	// s := []int{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s1 := []uint64{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s2 := []uint64{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Sort(createSortInterface[uint64](t, s2))
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[uint64](t, true)
	s3 := []uint64{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s4 := []uint64{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(createSortInterface[uint64](t, s4)))
	assert.Equal(t, s4, s3)
}

func TestNatureComparator_Float32(t *testing.T) {
	c1 := createNatureComparator[float32](t, false)
	// s := []int{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s1 := []float32{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s2 := []float32{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Sort(createSortInterface[float32](t, s2))
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[float32](t, true)
	s3 := []float32{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s4 := []float32{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(createSortInterface[float32](t, s4)))
	assert.Equal(t, s4, s3)
}

func TestNatureComparator_Float64(t *testing.T) {
	c1 := createNatureComparator[float64](t, false)
	// s := []int{90, 13, 48, 72, 99, 94, 75, 32, 88, 34}
	s1 := []float64{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s2 := []float64{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Float64s(s2)
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[float64](t, true)
	s3 := []float64{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	s4 := []float64{-9, -56, 70, 77, -28, -27, -7, -32, -31, 98}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(sort.Float64Slice(s4)))
	assert.Equal(t, s4, s3)
}

func TestNatureComparator_String(t *testing.T) {
	c1 := createNatureComparator[string](t, false)
	s1 := []string{"qgdolh", "ovqdcp", "auwlku", "mmgovi", "iusjze", "phvdmz", "dbfwpm", "ywccoy", "jktqcv", "bchdzn"}
	s2 := []string{"qgdolh", "ovqdcp", "auwlku", "mmgovi", "iusjze", "phvdmz", "dbfwpm", "ywccoy", "jktqcv", "bchdzn"}
	sort.Slice(s1, func(i, j int) bool {
		return c1.Compare(s1[i], s1[j]) < 0
	})
	sort.Strings(s2)
	assert.Equal(t, s2, s1)

	c2 := createNatureComparator[string](t, true)
	s3 := []string{"qgdolh", "ovqdcp", "ovqdcp", "auwlku", "mmgovi", "iusjze", "phvdmz", "dbfwpm", "ywccoy", "jktqcv", "bchdzn"}
	s4 := []string{"qgdolh", "ovqdcp", "ovqdcp", "auwlku", "mmgovi", "iusjze", "phvdmz", "dbfwpm", "ywccoy", "jktqcv", "bchdzn"}
	sort.Slice(s3, func(i, j int) bool {
		return c2.Compare(s3[i], s3[j]) < 0
	})
	sort.Sort(sort.Reverse(sort.StringSlice(s4)))
	assert.Equal(t, s4, s3)
}
