package builder

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// StringBuilder is used to build string
type StringBuilder struct {
	*strings.Builder
}

// NewStringBuilder creates a new instance of StringBuilder with init data.
//
// The type of init data can be string, rune and []byte
//
// example:
//
//	builderWithString := NewStringBuilder[string]("string")
//	builderWithRune := NewStringBuilder[rune]('s')
//	builderWithBytes := NewStringBuilder[[]byte]([]byte("string"))
func NewStringBuilder[S ~string | ~rune | ~[]byte](str S) *StringBuilder {
	var builder = new(strings.Builder)
	builder.WriteString(string(str))
	return &StringBuilder{
		Builder: builder,
	}
}

// WriteInt writes an int value to the builder
func (sb *StringBuilder) WriteInt(value int) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

// WriteUint writes an uint value to the builder
func (sb *StringBuilder) WriteUint(value uint) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

// WriteInt8 writes an int8 value to the builder
func (sb *StringBuilder) WriteInt8(value int8) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

// WriteUint8 writes an uint8 value to the builder
func (sb *StringBuilder) WriteUint8(value uint8) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

// WriteInt16 writes an int16 value to the builder
func (sb *StringBuilder) WriteInt16(value int16) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

// WriteUint16 writes an uint16 value to the builder
func (sb *StringBuilder) WriteUint16(value uint16) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

// WriteUint32 writes an int32 value to the builder
func (sb *StringBuilder) WriteInt32(value int32) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

// WriteUint32 writes an uint32 value to the builder
func (sb *StringBuilder) WriteUint32(value uint32) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

// WriteInt64 writes an int64 value to the builder
func (sb *StringBuilder) WriteInt64(value int64) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

// WriteUint64 writes an uint64 value to the builder
func (sb *StringBuilder) WriteUint64(value uint64) (int, error) {
	return sb.WriteString(fmt.Sprintf("%d", value))
}

// WriteStringBuilder writes another builder to current builder
func (sb *StringBuilder) WriteStringBuilder(builder *StringBuilder) (int, error) {
	return sb.WriteString(builder.String())
}

// WriteBuffer writes a bytes buffer to the builder
func (sb *StringBuilder) WriteBuffer(buf *bytes.Buffer) (int, error) {
	return sb.Write(buf.Bytes())
}

// WriteBool writes a bool value to the builder
func (sb *StringBuilder) WriteBool(value bool) (int, error) {
	return sb.WriteString(strconv.FormatBool(value))
}

// TrimSpace cuts space from both side of the builder
func (sb *StringBuilder) TrimSpace() {
	str := strings.TrimSpace(sb.String())
	sb.Reset()
	sb.WriteString(str)
}
